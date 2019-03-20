package genericTree

import (
	"github.com/gotk3/gotk3/gdk"
	"fmt"
)

type Tree struct {
	RootNode *Node
}
type Node struct {
	Data interface{}
	ParentId int
	Children []*Node
}

type Gender int
const (
	Male  Gender = 0
	Female  = 1
)
type Person struct {
	Id int
	Name string
	Age int
	Gender Gender
	isAParent bool
}
func (t *Tree) Insert(parentId int, data interface{}) bool{
	tempNode := &Node{}
	tempNode.Data = data
	tempNode.ParentId = parentId
	tempNode.Children = nil
	if parentId == 0 && t.RootNode == nil{
		t.RootNode = tempNode
	}else {
		current := t.RootNode
		for i:=0;i<=len(current.Children);i++{
			parent := current
			if parent.ParentId == parentId{
				parent.Children = append(parent.Children,tempNode)
			}else {
				current = parent.Children[i]
			}
		}
	}
	return false
}

func (t *Tree) DisplayTree(){
	current := t.RootNode
	for i:=0;i<=len(current.Children);i++{
		fmt.Print(current.Data)
	}
}

func (t *Tree) GetNodeByPerson(p Person) *Node{
	return nil
}
func (t *Tree) GetPersonById(id int) Person{
	return nil
}