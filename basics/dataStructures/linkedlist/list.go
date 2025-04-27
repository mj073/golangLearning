package linkedlist

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type List struct {
	Head   *Node
	Length uint
}

func (l *List) String() (str string) {
	headCount := 0
	for current := l.Head; current != nil; current = current.Next {
		if current == l.Head {
			headCount++
		}
		str = fmt.Sprintf("%s->%v", str, current.Data)
		if headCount == 2 {
			break
		}
	}
	return
}

func (l *List) InsertAtHead(data interface{}) {
	n := &Node{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	l.Length++
}

func (l *List) InsertAtEnd(data interface{}) {
	n := &Node{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = n
	} else {
		cur:=l.Head
		for ; cur.Next != nil; cur = cur.Next {}
		cur.Next = n
	}
	l.Length++
}

func (l *List) InsertANode(n *Node) {
	if l.Head == nil {
		l.Head = n
	} else {
		cur:=l.Head
		for ; cur.Next != nil; cur = cur.Next {}
		cur.Next = n
	}
	l.Length++
}

func (l *List) DeleteFromHead(data interface{}) {
	if l.Head != nil {
		if l.Head.Data == data {
			n := l.Head.Next
			l.Head = n
			l.Length--
		} else {
			curr, prev := l.Head.Next, l.Head
			for ; curr != nil; curr, prev = curr.Next, curr {
				if curr.Data == data {
					n := curr.Next
					prev.Next = n
					l.Length--
					break
				}
			}
		}
	}
}

func (l *List) RevereList() {
	var prev, next *Node
	curr := l.Head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func (l *List) DetectLoop() bool {
	if l.Head != nil {
		curr := l.Head
		for ; curr != nil; curr = curr.Next {
			if curr.Next == l.Head {
				return true
			}
		}
	}
	return false
}

