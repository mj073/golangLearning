package binaryTree

import "fmt"

type Tree struct{
	RootNode *Node
}
type Node struct {
	Data int
	Left *Node
	Right *Node
}

func (t *Tree) Insert(data int) (bool){
	tempNode := &Node{}
	tempNode.Data = data
	tempNode.Left = nil
	tempNode.Right = nil
	if t.RootNode == nil{
		fmt.Println("tree is empty...creating tree now..")
		t.RootNode = tempNode
	}else {
		current := t.RootNode
		var parent *Node
		for {
			parent = current
			if(data < parent.Data){
				current = parent.Left
				if parent.Left == nil {
					parent.Left = tempNode
					return true				}
			}
			if (data > parent.Data) {
				current = parent.Right
				if parent.Right == nil {
					parent.Right = tempNode
					return true
				}
			}
		}
	}
	return false
}

func (n *Node) DisplayTree(){
	current := n
	parent := current
	if parent != nil {
		fmt.Print(parent.Data,"->")
		if parent.Left != nil{
			fmt.Print("L")
			current = parent.Left
			current.DisplayTree()
		}
		if parent.Right != nil{
			fmt.Print("R")
			current = parent.Right
			current.DisplayTree()
		}
	}
}

func (t *Tree) Search(data int) (bool,*Node){
	current := t.RootNode
	for{
		parent := current
		if parent != nil {
			if parent.Data == data {
				return true, parent
			} else if data < parent.Data {
				current = parent.Left
			} else if data > parent.Data {
				current = parent.Right
			}
		}else {
			break
		}
	}
	return false, nil
}