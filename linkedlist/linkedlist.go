package linkedlist

import "fmt"

type Node struct {
	next *Node
	val  int
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
			node = node.next
		}

		return node.val
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
			newNode.next = l.head
		}

		l.head = &newNode
	} else {
		node := l.head

		for i := 1; i < index; i++ {
			node = node.next
		}

		if node.next != nil {
			newNode.next = node.next
		}

		node.next = &newNode
	}

	l.len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < l.len {
		prevNode := l.head

		if index == 0 && prevNode.next != nil {
			l.head = prevNode.next
		}

		for i := 1; i < index; i++ {
			prevNode = prevNode.next
		}

		if prevNode.next != nil && prevNode.next.next != nil {
			prevNode.next = prevNode.next.next
		} else {
			prevNode.next = nil
		}

		l.len--
	}
}

func (l *MyLinkedList) Print() {
	fmt.Println("len")
	fmt.Println(l.len)
	node := l.head
	fmt.Println(node)

	for node != nil && node.next != nil {
		node = node.next
		fmt.Println(node)
	}
}
