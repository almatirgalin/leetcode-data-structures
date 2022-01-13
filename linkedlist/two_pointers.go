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

func detectCycle(head *Node) *Node {
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer, fastPointer = slowPointer.Next, fastPointer.Next.Next

		if slowPointer == fastPointer {
			slowPointer = head

			for slowPointer != fastPointer {
				slowPointer = slowPointer.Next
				fastPointer = fastPointer.Next
			}

			return fastPointer
		}
	}

	return nil
}
