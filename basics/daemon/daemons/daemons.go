package daemons

import (
	"os"
	"fmt"
	"os/exec"
)

type Command struct {
	Daemons []string
}

func (c *Command) Main(s ...string) error{
	fmt.Println("in daemons Main")
	for _, d := range c.Daemons{
		start(d)
	}
	return nil
}
func start(args string) {
/*	_, wout, err := os.Pipe()
	defer func() {
		if err != nil {
			fmt.Print("daemon", "err:",  err)
		}
	}()
	if err != nil {
		return
	}
	_, werr, err := os.Pipe()
	if err != nil {
		return
	}*/
	p := Fork(args)
	p.Stdin = nil
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr
	p.Dir = "/"
	p.Env = []string{
		"PATH=" + "/home/mahesh",
		"TERM=linux",
	}
	fmt.Println("daemons cmd:",p)
	err := p.Start()
	if err != nil {
		fmt.Println("error while starting cmd:",p,"ERROR:",err)
		return
	}
	fmt.Println("cmd:",args,"pid:",p.Process.Pid)
	//id := fmt.Sprintf("%s.%s[%d]", prog.Base(), args[0], p.Process.Pid)
	/*daemons.mutex.Lock()
	daemons.cmdsByPid[p.Process.Pid] = p
	daemons.mutex.Unlock()
	go log.LinesFrom(rout, id, "info")
	go log.LinesFrom(rerr, id, "err")*/
	/*go func(p *exec.Cmd, wout, werr *os.File) {
		if err := p.Wait(); err != nil {
			fmt.Fprintln(werr, err)
		} else {
			fmt.Fprintln(wout, "done")
		}
		wout.Sync()
		werr.Sync()
		wout.Close()
		werr.Close()
		*//*daemons.mutex.Lock()
		delete(daemons.cmdsByPid, p.Process.Pid)
		daemons.mutex.Unlock()
		daemons.done <- struct{}{}*//*
	}(p, wout, werr)*/
	go p.Wait()
	return
}
func Fork(args string) *exec.Cmd {
	x := exec.Command("/home/mahesh/daemon")
	x.Args[0] = args
	return x
}