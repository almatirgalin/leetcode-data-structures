package linkedlist

func getIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	pointerA, pointerB := headA, headB

	for pointerA != pointerB {
		if pointerA == nil {
			pointerA = headB
		} else {
			pointerA = pointerA.Next
		}

		if pointerB == nil {
			pointerB = headA
		} else {
			pointerB = pointerB.Next
		}
	}

	return pointerA
}
