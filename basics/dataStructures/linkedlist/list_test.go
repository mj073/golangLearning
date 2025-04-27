package linkedlist

import (
	"fmt"
	"testing"
)

var (
	list = &List{}
	dll = &DoublyLinkedList{}
)

func TestList_InsertAtHead(t *testing.T) {
	for i:=0; i<10; i++ {
		list.InsertAtHead(i)
	}
	fmt.Println(list)
}

func TestList_InsertAtEnd(t *testing.T) {
	for i:=0; i<10; i++ {
		list.InsertAtEnd(i)
	}
	fmt.Println(list)
}

func TestList_InsertANode(t *testing.T) {
	for i:=0; i<10; i++ {
		n := &Node{
			Data: i,
		}
		list.InsertANode(n)
	}
	fmt.Println(list)
}

func TestList_DeleteFromHead(t *testing.T) {
	TestList_InsertAtEnd(t)
	list.DeleteFromHead(6)
	fmt.Println(list)
}

func TestList_ReverseList(t *testing.T) {
	TestList_InsertAtEnd(t)
	list.RevereList()
	fmt.Println(list)
}

func TestList_DetectLoop(t *testing.T) {
	TestList_InsertANode(t)
	list.InsertANode(list.Head)
	if list.DetectLoop() {
		fmt.Println("Loop Detected..List:", list)
	} else {
		t.Fatalf("Loop Not Detected")
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	for i:=0; i<10; i++ {
		dll.Insert(i)
	}
	fmt.Println(dll)
}