package linkedlist

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.AddAtTail(4)

	l.head.Next.Next.Next = l.head.Next

	isCycle := hasCycle(l.head)

	if !isCycle {
		t.Error("Expected true, got ", isCycle)
	}

	l.head.Next.Next.Next = nil

	isCycle = hasCycle(l.head)

	if isCycle {
		t.Error("Expected false, got ", isCycle)
	}
}
