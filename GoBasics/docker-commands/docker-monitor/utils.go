package main



//func Pipeline(cmds...*exec.Cmd){
//	fmt.Print("number of commands:",len(cmds))
//	var reader = make([]*io.PipeReader,len(cmds)-1)
//	var writer = make([]*io.PipeWriter,len(cmds)-1)
//	for i := 0;i<len(cmds)-1;i++ {
//		reader[i],writer[i] = io.Pipe()
//	}
//	var out bytes.Buffer
//	for i := 0;i<len(cmds)-1;i++{
//		fmt.Println("i:",i)
//		cmds[i].Stdout = writer[i]
//		cmds[i+1].Stdin = reader[i]
//	}
//	cmds[len(cmds)-1].Stdout = &out
//	for i := 0;i<len(cmds);i++{
//		fmt.Println("starting cmd:",cmds[i].Path," arguments: ",cmds[i].Args)
//		cmds[i].Start()
//	}
//	for i := 0; i<len(cmds)-1;i++{
//		cmds[i].Wait()
//		writer[i].Close()
//	}
//	cmds[len(cmds)-1].Wait()
//
//	total := out.String()
//	fmt.Println("final output:",total)
//}
