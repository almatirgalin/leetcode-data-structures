package linkedlist

import (
	"testing"
)

func TestGet(t *testing.T) {
	l := Constructor()
	node := Node{nil, 2}
	l.head = &node
	l.len = 1

	val := l.Get(0)

	if val != 2 {
		t.Error("Expected 2, got ", val)
	}

	node = Node{nil, 3}

	l.head.Next = &node
	l.len = 2

	val = l.Get(1)

	if val != 3 {
		t.Error("Expected 3, got ", val)
	}

	val = l.Get(2)

	if val != -1 {
		t.Error("Expected -1, got ", val)
	}
}

func TestAddAtHead(t *testing.T) {
	l := Constructor()
	l.AddAtHead(20)

	firstVal := l.Get(0)

	if firstVal != 20 {
		t.Error("Expected 20, got ", firstVal)
	}

	l.AddAtHead(40)

	firstVal = l.Get(0)

	if firstVal != 40 {
		t.Error("Expected 40, got ", firstVal)
	}

	secondVal := l.Get(1)

	if secondVal != 20 {
		t.Error("Expected 20, got ", secondVal)
	}

	emptyVal := l.Get(2)

	if emptyVal != -1 {
		t.Error("Expected -1, got ", emptyVal)
	}

	if l.len != 2 {
		t.Error("Expected len 2, got ", l.len)
	}
}

func TestAddAtTail(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)

	firstVal := l.Get(0)

	if firstVal != 1 {
		t.Error("Expected 1, got ", firstVal)
	}

	l.AddAtTail(2)

	firstVal = l.Get(0)

	if firstVal != 1 {
		t.Error("Expected 1, got ", firstVal)
	}

	secondVal := l.Get(1)

	if secondVal != 2 {
		t.Error("Expected 2, got ", secondVal)
	}
}

func TestAddAtIndex(t *testing.T) {
	l := Constructor()

	l.AddAtIndex(1, 3)

	firstVal := l.Get(0)

	if firstVal != -1 {
		t.Error("Expected -1, got ", firstVal)
	}

	l.AddAtIndex(0, 2)

	firstVal = l.Get(0)

	if firstVal != 2 {
		t.Error("Expected 2, got ", firstVal)
	}

	if l.len != 1 {
		t.Error("Expected len 1, got ", l.len)
	}

	l.AddAtIndex(0, 3)

	firstVal = l.Get(0)

	if firstVal != 3 {
		t.Error("Expected 3, got ", firstVal)
	}

	l.AddAtIndex(2, 4)

	lastVal := l.Get(2)

	if lastVal != 4 {
		t.Error("Expected 4, got ", lastVal)
	}
}

func TestAll(t *testing.T) {
	l := Constructor()
	l.AddAtHead(2)
	l.DeleteAtIndex(1)
	l.AddAtHead(2)
	l.AddAtHead(7)
	l.AddAtHead(3)
	l.AddAtHead(2)
	l.AddAtHead(5)
	l.AddAtTail(5)
	fifthVal := l.Get(5)

	if fifthVal != 2 {
		t.Error("Expected 2, got ", fifthVal)
	}

	l.DeleteAtIndex(6)
	l.DeleteAtIndex(4)
}
