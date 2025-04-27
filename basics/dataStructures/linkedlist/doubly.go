package linkedlist

type DoublyLinkedList struct {
	List
}

func (l *DoublyLinkedList) Insert(d interface{}) {
	n := &Node{
		Data: d,
	}
	if l.Head == nil {
		l.Head = n
	} else {
		var prev *Node
		curr := l.Head
		for ;curr.Next != nil; curr = curr.Next {}
		curr.Next = n
		curr.Prev = curr
	}
	l.Length++
}