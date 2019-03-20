package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"time"
	"errors"
	"runtime/debug"
)

var POLLING_STATUS = true
var MAP_NODE_VALIDATION_POLLER = map[string]*NodeValidationPoller{}

type NodeValidationPoller struct {
	pollingFlag                     bool
	channelValidateSNAndType        chan ValidationStatus
	channelValidateKernelAndBaseISO chan ValidationStatus
	channelValidateIPLLDP           chan ValidationStatus
	finalChannel                    chan ValidationStatus
}

type ValidationStatus struct {
	IsKernelMatched  bool                                 `json:"isKernelMatched"`
	Kernel           string                               `json:"kernel"`
	IsBaseISOMatched bool                                 `json:"isBaseISOMatched"`
	BaseISO          string                               `json:"baseISO"`
	IsSNMatched      bool                                 `json:"isSNMatched"`
	SerialNumber     string                               `json:"serialNumber"`
	IsTypeMatched    bool                                 `json:"isTypeMatched"`
	Type             string                               `json:"type"`
	InterfacesStatus map[string]InterfaceValidationStatus `json:"interfacesStatus"`
	OverallStatus    bool                                 `json:"overallStatus"`
}

type InterfaceValidationStatus struct {
	IsValidIP              bool `json:"isValidIP"`
	IsLLDPMatched          bool `json:"isLLDPMatched"`
	IsRemoteInvaderMatched bool `json:"isRemoteInvaderMatched"`
	LinkStatus             bool `json:"linkStatus"`
	Status                 bool `json:"status"`
}

const (
	ANSIBLE_PLAYBOOK_CMD   = "ansible-playbook"
	ANSIBLE_HOST_FILE_PATH = "./config/hosts.yml"
)

func (validator *NodeValidationPoller) validateNodeInfoAgainstInvader(updatedNode *Node) (ValidationStatus) {
	validator.channelValidateSNAndType = make(chan ValidationStatus, 1)
	validator.channelValidateKernelAndBaseISO = make(chan ValidationStatus, 1)
	validator.channelValidateIPLLDP = make(chan ValidationStatus, 1)
	validator.finalChannel = make(chan ValidationStatus, 1)

	status := ValidationStatus{}
	var counter uint64
	for {
		counter++
		start := time.Now()
		if !validator.pollingFlag {
			break
		}
		status.InterfacesStatus = map[string]InterfaceValidationStatus{}
		go validator.checkStatus(&status)
		go validator.validateSNAndType(updatedNode.Name, updatedNode.SerialNumber, updatedNode.NodeType, validator.channelValidateSNAndType)
		go validator.validateKernelAndBaseISO(updatedNode.Name, updatedNode.Kernel, updatedNode.LinuxISO, validator.channelValidateKernelAndBaseISO)
		go validator.validateInterfaces(updatedNode, validator.channelValidateIPLLDP)

		switch status.OverallStatus {
		case false:
			updatedNode.NodeStatus = "Mismatch"
		case true:
			updatedNode.NodeStatus = "Registered"
		}

		<-validator.finalChannel
		updatedNode.ValidationStatus = status

		//DOUBLE CHECK
		if validator.pollingFlag {
			updatedNode.UpdateDB(updatedNode)
		}

		end := time.Now()
		elapsed := end.Sub(start)

		fmt.Printf("[ Node %v Counter ID: %v] Polling started. Start Time : %v \n", updatedNode.Name, counter, start)
		fmt.Printf("[ Node %v Counter ID: %v] Polling Completed.Elpased Time : %v \n", updatedNode.Name, counter, elapsed)
		fmt.Printf("[ Node %v Counter ID: %v] Validation Status : %v \n", updatedNode.Name, counter, status)
		time.Sleep(time.Duration(GEnv.NodeValidator.PollingInterval) * time.Second)
	}
	return status
}

func (validator *NodeValidationPoller) checkStatus(status *ValidationStatus) {
	flagSnType, flagKernelIso, flagIpLLDP := false, false, false
	resultSnType, resultKernelISO, resultIpLLDP := false, false, false
	for {
		select {
		case statusChannelValidateSNAndType := <-validator.channelValidateSNAndType:
			status.IsSNMatched = statusChannelValidateSNAndType.IsSNMatched
			status.IsTypeMatched = statusChannelValidateSNAndType.IsTypeMatched
			status.OverallStatus = statusChannelValidateSNAndType.OverallStatus
			status.SerialNumber = statusChannelValidateSNAndType.SerialNumber
			status.Type = statusChannelValidateSNAndType.Type
			resultSnType = status.OverallStatus
			flagSnType = true

		case statusChannelValidateKernelAndBaseISO := <-validator.channelValidateKernelAndBaseISO:
			status.IsKernelMatched = statusChannelValidateKernelAndBaseISO.IsKernelMatched
			status.IsBaseISOMatched = statusChannelValidateKernelAndBaseISO.IsBaseISOMatched
			status.OverallStatus = statusChannelValidateKernelAndBaseISO.OverallStatus
			status.Kernel = statusChannelValidateKernelAndBaseISO.Kernel
			status.BaseISO = statusChannelValidateKernelAndBaseISO.BaseISO

			flagKernelIso = true
			resultKernelISO = status.OverallStatus

		case statusChannelValidateIPLLDP := <-validator.channelValidateIPLLDP:
			status.InterfacesStatus = statusChannelValidateIPLLDP.InterfacesStatus
			status.OverallStatus = statusChannelValidateIPLLDP.OverallStatus
			flagIpLLDP = true
			resultIpLLDP = status.OverallStatus
		}
		if flagSnType && flagKernelIso && flagIpLLDP {
			status.OverallStatus = resultSnType && resultKernelISO && resultIpLLDP
			break
		}
	}
	validator.finalChannel <- *status
}

func (validator *NodeValidationPoller) validateInterfaces(node *Node, channel chan ValidationStatus) {
	defer validator.handlePanic()
	status := ValidationStatus{}
	status.OverallStatus = true
	status.InterfacesStatus = map[string]InterfaceValidationStatus{}
	numOfInterfaces := len(node.Interfaces)

	for i := 0; i < numOfInterfaces; i++ {
		interfaceStatus := InterfaceValidationStatus{}
		if ok, _ := validator.validateIPAndLLDP(node.Name, node.Interfaces[i], &interfaceStatus); !ok {
			status.OverallStatus = false
			node.Interfaces[i].ConnectedTo.LldpMatched = "False"
		} else {
			node.Interfaces[i].ConnectedTo.LldpMatched = "True"
		}
		node.Interfaces[i].ConnectedTo.Link = strconv.FormatBool(interfaceStatus.LinkStatus)
		status.InterfacesStatus[node.Interfaces[i].Port] = interfaceStatus
	}
	channel <- status
}

func (validator *NodeValidationPoller) validateSNAndType(inputHostName, inputSerialNumber, inputType string, channel chan ValidationStatus) (bool, error) {
	defer validator.handlePanic()
	var err error
	result := false
	status := ValidationStatus{}
	args := "host_name=" + inputHostName + " sn=" + inputSerialNumber + " type=" + inputType
	cmd := exec.Command(ANSIBLE_PLAYBOOK_CMD, "-i", ANSIBLE_HOST_FILE_PATH, "./config/validate_type_serial_no_kernel_base_iso.yml",
		"--tags", "validateTypeAndSN", "-e", args)

	out, _ := cmd.CombinedOutput()
	tempMap := validator.buildValidationStatusFromAnsibleResponse(out)
	if len(tempMap) > 0 {
		if inputSerialNumber != "" {
			status.IsSNMatched = tempMap["serialNoMatched"].(bool)
		}
		if inputType != "" {
			status.IsTypeMatched = tempMap["typeMatched"].(bool)
		}
		status.SerialNumber = tempMap["SerialNo"].(string)
		status.Type = tempMap["PartNo"].(string)
		status.OverallStatus = status.IsSNMatched && status.IsTypeMatched
	} else {
		err = errors.New("Unable to validate SN and type through Ansible")
		panic(err)
	}
	channel <- status
	return result, err
}

func (validator *NodeValidationPoller) validateKernelAndBaseISO(inputHostName, inputKernel, inputIso string, channel chan ValidationStatus) (bool, error) {
	defer validator.handlePanic()
	var err error
	result := false
	status := ValidationStatus{}
	args := "host_name=" + inputHostName + " kernel='" + inputKernel + "' iso='" + inputIso +"'"
	cmd := exec.Command(ANSIBLE_PLAYBOOK_CMD, "-i", ANSIBLE_HOST_FILE_PATH, "./config/validate_type_serial_no_kernel_base_iso.yml",
		"--tags", "validateKernelAndISO", "-e", args)
	out, _ := cmd.CombinedOutput()
	tempMap := validator.buildValidationStatusFromAnsibleResponse(out)
	if len(tempMap) > 0 {
		if inputIso != "" {
			status.IsBaseISOMatched = tempMap["isoMatched"].(bool)
		}
		if inputKernel != "" {
			status.IsKernelMatched = tempMap["kernelMatched"].(bool)
		}
		status.Kernel = tempMap["kernel"].(string)
		status.BaseISO = tempMap["baseISO"].(string)
		status.OverallStatus = status.IsBaseISOMatched && status.IsKernelMatched
	} else {
		err = errors.New("Unable to validate kernel and iso through Ansible")
		panic(err)
	}
	channel <- status
	return result, err
}

func (validator *NodeValidationPoller) validateIPAndLLDP(inputHostName string, nodeInterface NodeInterface, interfaceStatus *InterfaceValidationStatus) (bool, error) {
	var err error
	if nodeInterface.Port == "eth0" {
		interfaceStatus.Status = true
		interfaceStatus.LinkStatus = true
		return true, err
	}
	args := "host_name=" + inputHostName + " remote_invader=" + nodeInterface.ConnectedTo.Name + " source_interface=" + nodeInterface.Port + " remote_interface=" + nodeInterface.ConnectedTo.Port + " source_interface_ip=" + nodeInterface.IP
	cmd := exec.Command(ANSIBLE_PLAYBOOK_CMD, "-i", ANSIBLE_HOST_FILE_PATH, "./config/validateIpAndLLDP.yml",
		"-e", args)
	out, _ := cmd.CombinedOutput()
	tempMap := validator.buildValidationStatusFromAnsibleResponse(out)

	if len(tempMap) > 0 {
		interfaceStatus.IsValidIP = tempMap["ipMatched"].(bool)
		interfaceStatus.IsLLDPMatched = tempMap["lldpMatchedSysName"].(bool) && tempMap["lldpMatchedPortDescr"].(bool)
		interfaceStatus.IsRemoteInvaderMatched = tempMap["lldpMatchedPortDescr"].(bool)
		interfaceStatus.LinkStatus = tempMap["linkStatus"].(bool)
		interfaceStatus.Status = interfaceStatus.IsValidIP && interfaceStatus.IsLLDPMatched
	} else {
		err = errors.New("Unable to validate interface information through Ansible")
		panic(err)
	}

	return interfaceStatus.Status, err
}

func (validator *NodeValidationPoller) buildValidationStatusFromAnsibleResponse(response []byte) (map[string]interface{}) {
	var lines []string = strings.Split(string(response), "\n")
	var tempMap map[string]interface{}
	defer func() (map[string]interface{}) {
		return tempMap
	}()
	for _, value := range lines {
		if strings.Contains(value, "Result:") {
			validJsonStringResponse := strings.TrimSpace(strings.SplitAfter(value, "Result:")[1])
			validJsonStringResponse = validJsonStringResponse[:len(validJsonStringResponse)-1]
			var val []byte = []byte("\"" + validJsonStringResponse + "\"")
			s, _ := strconv.Unquote(string(val))
			err := json.Unmarshal([]byte(s), &tempMap)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return tempMap
}

func (validator *NodeValidationPoller) closeChannels() {
	close(validator.channelValidateSNAndType)
	close(validator.channelValidateKernelAndBaseISO)
	close(validator.channelValidateIPLLDP)
	close(validator.finalChannel)
}

func (validator NodeValidationPoller) handlePanic() {
	if r := recover(); r!= nil {
		fmt.Println("Exception Occured . Stop signal sent to Poller . Exception reason : ", r)
		validator.pollingFlag = false
		debug.PrintStack()
	}
}


