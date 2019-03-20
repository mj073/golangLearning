package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
	"strings"
	"os"
	"strconv"
	"path/filepath"
	"log"
)
type lossless_config struct{
	shared_limit int
	headroom_limit int
	offset int
	floor int
}

var lossless_confs = [...]lossless_config{
	//tests completed
	/*
	0: lossless_config{8, 517, 300, 100},
	1: lossless_config{16, 517,300,100},
	2: lossless_config{32, 517, 300, 100},
	3: lossless_config{64, 517, 300, 100},
	4: lossless_config{128, 517, 300, 100},
	5: lossless_config{256, 517, 300, 100},
	6: lossless_config{512, 517, 300, 100},
	7: lossless_config{512, 517, 250, 50},
	8: lossless_config{256, 517, 260, 60},
	9: lossless_config{128, 517, 270, 70},
	10: lossless_config{64, 517, 280, 80},
	11: lossless_config{32, 517, 320, 120},
	12: lossless_config{16, 517, 340, 140},
	*/

	0: lossless_config{512, 517, 0, 0},
	1: lossless_config{512, 517, 10, 0},
	2: lossless_config{512, 517, 20, 10},
	3: lossless_config{512, 517, 50, 20},
	4: lossless_config{512, 517, 100, 50},
	5: lossless_config{1024, 517, 0, 0},
	6: lossless_config{1024, 517, 50, 20},
	7: lossless_config{1024, 517, 100, 50},
	8: lossless_config{1024, 517, 250, 100},
	9: lossless_config{2048, 517, 0, 0},
	10: lossless_config{2048, 517, 50, 20},
	11: lossless_config{2048, 517, 100, 50},
	12: lossless_config{2048, 517, 250, 100},
	13: lossless_config{4096, 517, 0, 0},
	14: lossless_config{4096, 517, 50, 20},
	15: lossless_config{4096, 517, 100, 50},
	16: lossless_config{4096, 517, 250, 100},
	17: lossless_config{8192, 517, 0, 0},
	18: lossless_config{8192, 517, 50, 20},
	19: lossless_config{8192, 517, 100, 50},
	20: lossless_config{8192, 517, 250, 100},
	21: lossless_config{512, 10, 0, 0},
	22: lossless_config{512, 50, 20, 10},
	23: lossless_config{512, 100, 50, 20},
	24: lossless_config{512, 150, 100, 50},
	25: lossless_config{512, 250, 150, 100},
	26: lossless_config{512, 600, 200, 100},
	27: lossless_config{512, 700, 250, 100},
	28: lossless_config{512, 800, 300, 100},
	29: lossless_config{512, 900, 350, 150},
	30: lossless_config{512, 1000, 400, 200},
	//
	//0: lossless_config{8192, 517, 250, 100},
	//1: lossless_config{8192, 517, 300, 150},
	//2: lossless_config{8192, 517, 350, 200},
	//3: lossless_config{8192, 517, 400, 250},
	//4: lossless_config{8192, 517, 500, 350},
	//5: lossless_config{9216, 517, 250, 100},
	//6: lossless_config{9216, 517, 300, 150},
	//7: lossless_config{9216, 517, 350, 200},
	//8: lossless_config{9216, 517, 400, 250},
	//9: lossless_config{9216, 517, 500, 350},
	//10: lossless_config{10240, 517, 250, 100},
	//11: lossless_config{10240, 517, 300, 150},
	//12: lossless_config{10240, 517, 350, 200},
	//13: lossless_config{10240, 517, 400, 250},
	//14: lossless_config{10240, 517, 500, 350},
	//15: lossless_config{11264, 517, 250, 100},
	//16: lossless_config{11264, 517, 300, 150},
	//17: lossless_config{11264, 517, 350, 200},
	//18: lossless_config{11264, 517, 400, 250},
	//19: lossless_config{11264, 517, 500, 350},
	//20: lossless_config{8192, 5000, 250, 100},
	//21: lossless_config{8192, 5500, 250, 100},
	//22: lossless_config{8192, 6000, 250, 100},
	//23: lossless_config{8192, 6500, 250, 100},
	//24: lossless_config{8192, 7000, 250, 100},
	//25: lossless_config{8192, 5000, 300, 100},
	//26: lossless_config{8192, 5000, 500, 200},
	//27: lossless_config{8192, 5000, 1000, 500},
	//28: lossless_config{8192, 5000, 1500, 1000},
	//29: lossless_config{8192, 5000, 2000, 1500},
	//30: lossless_config{8192, 5000, 2500, 2000},
}

var port = "eth-17-1"
var pri_grp = 7
func main(){
	inv49 := "172.17.2.49"
	bp10 := "172.17.2.110"

	dir := "/home/"+os.Getenv("USER")+"/flow_control_test"

	os.Mkdir(dir,0766)
	filename := "test_results_"+time.Now().Format("02012006-150405")+".csv"
	filepath := filepath.Join(dir,filename)
	file,err := os.Create(filepath)
	if err != nil {
		panic("failed to open/create a file..ERROR:"+err.Error())
	}
	defer file.Close()
	title := fmt.Sprintf("%-13s %-15s %-7s %-6s %-20s %-19s %-12s %-29s %-23s %-22s %-11s \n","shared_limit", "headroom_limit", "offset", "floor", "eth-17-1_rx_packets", "eth-2-3_tx_packets", "rx_tx_drops","eth-17-1_mmu_rx_drop_packets", "eth-17-1_pause_packets",
		"bp10_rx_packets_delta","Throughput")
	file.WriteString(title)
	clientConfig := &ssh.ClientConfig{
		User: "nvmf",
		Auth: []ssh.AuthMethod{
			ssh.Password("nvmf"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if clientConfig == nil {
		log.Println("clientConfig is nil...")
		return
	}
	inv49_client, err := ssh.Dial("tcp", inv49+":22", clientConfig)
	if err != nil {
		log.Println("Failed to dial: " + err.Error())
		return
	}
	defer inv49_client.Close()
	bp10_client, err := ssh.Dial("tcp", bp10+":22", clientConfig)
	if err != nil {
		log.Println("Failed to dial: " + err.Error())
		return
	}
	defer bp10_client.Close()

	cmds_inv49 := []string{
		"sudo ip netns exec net1 bash -c \"/home/nvmf/mahesh/pause_control_with_packet_capture/goes-platina-mk1 install\"",
		"sudo ip netns exec net1 bash -c \"/home/nvmf/i49_net.sh\"",
		"sudo ip netns exec net1 bash -c \"/home/nvmf/start_quagga.sh\"",
	}
	cmds_inv49_2 := []string{
		"sudo goes hgetall platina | grep packets | egrep 'eth-17-1|eth-2-3' | egrep '(tx|rx)_packets'",
		"sudo goes hgetall platina|egrep 'drop|port_(rx|tx)_flow_control_packet|xon'|egrep -v ': 0|multicast'",
	}
	cmds_bp10 := []string{
		"sudo nvme connect -t rdma -a 172.28.17.10 -n dev-subsys51",
		"lsblk",
		"ethtool -S enp94s0|grep \"rx_vport_rdma_unicast_packets\"",
		"sudo fio /home/nvmf/mahesh/band_seqread_7partition.fio",
		"ethtool -S enp94s0|grep \"rx_vport_rdma_unicast_packets\"",
		"sudo nvme disconnect -n dev-subsys51",
	}

	for _,conf := range lossless_confs{
		var out []byte
		var rx_packets_before,rx_packets_after, rx_packets_delta, eth_17_1_rx_packets, eth_2_3_tx_packets, eth_17_1_mmu_rx_drop_packets, eth_17_1_pause_packets,rx_tx_drops  int
		var throughput string
		fio_run_finish := false

		log.Println("------------- conf(shared_limit headroom_limit offset floor):(",conf.shared_limit,conf.headroom_limit,conf.offset,conf.floor,")------------------")
		for _, c := range cmds_inv49{
			log.Println("------------- executing on inv49:",c,"------------------------")
			try := 0
			try_max := 5
			for {
				inv49_session, err := inv49_client.NewSession()
				if err != nil {
					panic("Failed to create session: " + err.Error())
				}

				_, err = inv49_session.Output(c)
				if err != nil{
					if try != try_max {
						try++
					}else {
						inv49_session.Close()
						panic("failed to execute the remote cmd:" + c + " after max tries...error:" + err.Error())
					}
				}else {
					inv49_session.Close()
					break
				}
				inv49_session.Close()
				time.Sleep(time.Millisecond * 300)
			}
			time.Sleep(time.Second * 1)
		}
		c := fmt.Sprintf("sudo goes vnet set fe1 lossless config port_name %v pri_grp %v shared_limit %v " +
			"headroom_enable 1 headroom_limit %v offset %v floor %v",port,pri_grp,conf.shared_limit,conf.headroom_limit,conf.offset,conf.floor)

		inv49_session, err := inv49_client.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())
		}
		err = inv49_session.Run(c)
		if err != nil{
			inv49_session.Close()
			panic("failed to run cmd: "+c)
		}
		inv49_session.Close()

		time.Sleep(time.Second*60)

		for _,c := range cmds_bp10 {

			if strings.Contains(c,"fio") {
				time.Sleep(time.Second * 5)
			}else {
				time.Sleep(time.Second * 1)
			}
			log.Println("------------- executing on bp10:",c,"--------------------")
			try := 0
			try_max := 5
			for {
				bp10_session, err := bp10_client.NewSession()
				if err != nil {
					panic("Failed to create session: " + err.Error())
				}

				out, err := bp10_session.Output(c)
				if err != nil {
					if try != try_max {
						try++
					} else {
						bp10_session.Close()
						panic("failed to execute the remote cmd:" + c + " after max tries...error:" + err.Error())
					}
				}else {
					out_str := string(out)
					if strings.Contains(c,"fio") {
						throughput = strings.Split(strings.Split(out_str, "/s (")[1], ")(")[0]
						log.Println("throughput:", throughput)
						fio_run_finish = true
					}
					if strings.Contains(c,"ethtool"){
						rx_packets,_ := strconv.Atoi(strings.TrimSuffix(strings.Split(out_str,"rx_vport_rdma_unicast_packets: ")[1],"\n"))
						if fio_run_finish{
							rx_packets_after = rx_packets
							rx_packets_delta = rx_packets_after - rx_packets_before
						}else {
							rx_packets_before = rx_packets
						}
						log.Println(out_str)
					}
					bp10_session.Close()
					break
				}
				bp10_session.Close()
				time.Sleep(time.Millisecond * 300)
			}

		}
		for i, c := range cmds_inv49_2 {
			log.Println("------------- executing on inv49:",c,"---------------------")
			try := 0
			try_max := 5
			for {
				log.Println("try:",try+1)
				inv49_session, err = inv49_client.NewSession()
				if err != nil {
					panic("Failed to create session: " + err.Error())
				}
				out, err = inv49_session.Output(c)
				if err != nil{
					if try != try_max {
						try++
					}else {
						inv49_session.Close()
						panic("failed to execute the remote cmd:" + c + " after max tries...error:" + err.Error())
					}
				}else {
					out_str := string(out)
					if i == 0{
						eth_17_1_rx_packets, _ = strconv.Atoi(strings.Split(strings.Split(out_str, "eth-17-1.port_rx_packets: ")[1], "\n")[0])
						eth_2_3_tx_packets, _ = strconv.Atoi(strings.Split(strings.Split(out_str, "eth-2-3.port_tx_packets: ")[1], "\n")[0])
						rx_tx_drops = eth_17_1_rx_packets - eth_2_3_tx_packets
					}else if i == 1 {
						eth_17_1_mmu_rx_drop_packets, _ = strconv.Atoi(strings.Split(strings.Split(out_str, "eth-17-1.mmu_rx_threshold_drop_packets: ")[1], "\n")[0])
						eth_17_1_pause_packets, _ = strconv.Atoi(strings.Split(strings.Split(out_str, "eth-17-1.port_tx_flow_control_packets: ")[1], "\n")[0])
					}
					inv49_session.Close()
					log.Println(out_str)
					break
				}
				time.Sleep(time.Millisecond * 300)
			}
		}
		result := fmt.Sprintf("%-13d %-15d %-7d %-6d %-20d %-19d %-12d %-29d %-23d %-22d %-11s \n",
			conf.shared_limit,conf.headroom_limit,conf.offset,conf.floor,eth_17_1_rx_packets,eth_2_3_tx_packets,rx_tx_drops,
			eth_17_1_mmu_rx_drop_packets,eth_17_1_pause_packets,rx_packets_delta,throughput)
		file.WriteString(result)
	}
}
