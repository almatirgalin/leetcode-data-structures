package linkedlist

func hasCycle(head *Node) bool {
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer, fastPointer = slowPointer.Next, fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
