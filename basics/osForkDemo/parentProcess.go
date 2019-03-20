package main

import (
	"os/exec"
	"os"
	"fmt"
	"time"
)
var progname = "ls"
var progpath,_ = exec.LookPath(progname)
func main(){
	start("/usr/bin")
	time.Sleep(time.Second * 1000)
}

func Fork(args ...string) *exec.Cmd{
	x := exec.Command(progname,args...)
	return x
}
func start(args ...string){
	_, wout,err := os.Pipe()
	if err != nil{
		panic(err)
	}
	_,werr,err := os.Pipe()
	if err != nil{
		panic(err)
	}
	p := Fork(args...)
	p.Stdin = nil
	p.Stdout = wout
	p.Stderr = werr
	p.Dir = "/"
	p.Env = []string{
		"PATH=" + progpath,
		"TERM=linux",
	}
	err = p.Start()
	if err != nil{
		panic(err)
	}

	go func(p *exec.Cmd, wout, werr *os.File) {
		fmt.Println("waiting for process to complete")
		if err := p.Wait(); err != nil {
			fmt.Fprintln(werr, err)
		} else {
			fmt.Fprintln(wout, "done")
		}
		b:= make([]byte,1000)
		wout.Read(b)
		fmt.Println(string(b))
		wout.Sync()
		werr.Sync()
		wout.Close()
		werr.Close()
		<- time.After(time.Second * 1000)
	}(p, wout, werr)
}