package start

import (
	"os/exec"
	"syscall"
	"fmt"
	"os"
)

type Command struct {

}

func (cmd *Command) Main(s ...string) error {
	fmt.Println("start Main")
	start(s...)
	return nil
}
func start(s ...string) {
	fmt.Println("start start()")
	cmd := exec.Command("/home/mahesh/daemon","daemons")
	cmd.Stdin = nil
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	cmd.Env = []string{
		"PATH=" + "/home/mahesh/",
		"TERM=linux",
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
		Pgid:   0,
	}
	fmt.Println("start cmd:",cmd)
	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd.Start() ERROR:",err)
	}
	go cmd.Wait()
	os.Exit(0)
}
