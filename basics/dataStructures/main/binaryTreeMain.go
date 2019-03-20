package main

import (
	"basics/dataStructures/binaryTree"
	"fmt"
)

func main(){
	var tree = &binaryTree.Tree{}
	arr := []int{43,22,23,64,76,242,65,87,34,83,86}
	for _,x:=range arr{
		_ = tree.Insert(x)
	}
	tree.RootNode.DisplayTree()
	found, n := tree.Search(43)
	if found{
		fmt.Println("tree of found node..")
		n.DisplayTree()
	}else {
		fmt.Print("not found...")
	}
}
