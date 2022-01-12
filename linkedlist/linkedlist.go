package linkedlist

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type MyLinkedList struct {
	head *Node
	len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) GetLen() int {
	return l.len
}

func (l *MyLinkedList) Get(index int) int {
	if index < l.len {
		node := l.head

		for i := 0; i != index; i++ {
			node = node.Next
		}

		return node.Val
	}

	return -1
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.len, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.len {
		return
	}

	newNode := Node{nil, val}

	if index == 0 {
		if l.head != nil {
			newNode.Next = l.head
		}

		l.head = &newNode
	} else {
		node := l.head

		for i := 1; i < index; i++ {
			node = node.Next
		}

		if node.Next != nil {
			newNode.Next = node.Next
		}

		node.Next = &newNode
	}

	l.len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < l.len {
		prevNode := l.head

		if index == 0 && prevNode.Next != nil {
			l.head = prevNode.Next
		}

		for i := 1; i < index; i++ {
			prevNode = prevNode.Next
		}

		if prevNode.Next != nil && prevNode.Next.Next != nil {
			prevNode.Next = prevNode.Next.Next
		} else {
			prevNode.Next = nil
		}

		l.len--
	}
}

func (l *MyLinkedList) Print() {
	fmt.Println("len")
	fmt.Println(l.len)
	node := l.head
	fmt.Println(node)

	for node != nil && node.Next != nil {
		node = node.Next
		fmt.Println(node)
	}
}
