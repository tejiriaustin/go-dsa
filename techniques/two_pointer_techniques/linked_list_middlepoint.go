package two_pointer_techniques

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func Midpoint(root *LinkedList) *LinkedList {
	slow, fast := root, root

	for slow.Next != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
