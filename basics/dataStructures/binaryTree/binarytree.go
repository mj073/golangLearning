package binarytree

import (
	"container/list"
	"fmt"
)

type Tree struct{
	RootNode *Node
}
type Node struct {
	Data interface{}
	Left *Node
	Right *Node
}

func (t *Tree) Insert(data int) bool {
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
			if data < parent.Data.(int) {
				current = parent.Left
				if parent.Left == nil {
					parent.Left = tempNode
					return true
				}
			}
			if data > parent.Data.(int) {
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
			} else if data < parent.Data.(int) {
				current = parent.Left
			} else if data > parent.Data.(int) {
				current = parent.Right
			}
		}else {
			break
		}
	}
	return false, nil
}

func (t *Tree) Dft() (result []interface{}) {
	if t.RootNode != nil {
		stack := []*Node{t.RootNode}
		pop := func() (d *Node, ok bool) {
			if l := len(stack); l != 0 {
				d = stack[l-1]
				stack = stack[:l-1]
				ok = true
			}
			return
		}
		for s, ok := pop(); ok; s, ok = pop() {
			 result = append(result, s.Data)
			 if s.Right != nil {
				 stack = append(stack, s.Right)
			 }
			 if s.Left != nil {
			 	stack = append(stack, s.Left)
			 }
		}
	}
	return
}

// Recursive
func (t *Tree) DftRecursive(n *Node) (result []interface{})                                                                                                                                                               {
	if n == nil {
		return
	}
	result = append(result, n.Data)
	result = append(result, t.DftRecursive(n.Left)...)
	result = append(result, t.DftRecursive(n.Right)...)
	return
}

// Bft Using Queue
func (t *Tree) Bft() (result []interface{}) {
	if t.RootNode == nil {
		return
	}
	q := list.New()
	q.PushBack(t.RootNode)
	for e := q.Front(); e != nil; e = q.Front() {
		q.Remove(e)
		n := e.Value.(*Node)
		result = append(result, n.Data)
		if n.Left != nil {
			q.PushBack(n.Left)
		}
		if n.Right != nil {
			q.PushBack(n.Right)
		}
	}
	return
}

// Bft2 Using Queue Slice
func (t *Tree) Bft2() (result []interface{}) {
	if t.RootNode == nil {
		return
	}
	var q []*Node
	q = append(q, t.RootNode) // PushBack
	for len(q) > 0 {
		n := q[0] 	// Front()
		q = q[1:] 	// RemoveFront() // Doesn't panic even if q[1] not present
		result = append(result, n.Data)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return
}

