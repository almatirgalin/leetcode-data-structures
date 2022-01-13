package linkedlist

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	l1 := Constructor()

	l1.AddAtHead(1)
	l1.AddAtTail(9)
	l1.AddAtTail(1)
	l1.AddAtTail(2)
	l1.AddAtTail(4)

	l2 := Constructor()
	l2.AddAtHead(3)
	l2.head.Next = l1.head.Next.Next.Next

	intersected := getIntersectionNode(l1.head, l2.head)

	if intersected != l2.head.Next || intersected != l1.head.Next.Next.Next {
		t.Error("Expected intersection with val 2, got node with val", intersected.Val)
	}

	l3 := Constructor()
	l3.AddAtHead(2)
	l3.AddAtTail(6)
	l3.AddAtTail(4)

	l4 := Constructor()
	l4.AddAtHead(1)
	l4.AddAtTail(5)

	intersected = getIntersectionNode(l3.head, l4.head)

	if intersected != nil {
		t.Error("Expected nil, got ", intersected)
	}
}
