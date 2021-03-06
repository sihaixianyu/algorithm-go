package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func FromSlice(vals []int) LinkedList {
	if len(vals) == 0 {
		return LinkedList{
			Head: nil,
			Tail: nil,
		}
	}

	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}

	curr := head
	for _, val := range vals[1:] {
		newNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		curr.Next = newNode
		curr = newNode
	}

	return LinkedList{
		Head: head,
		Tail: curr,
	}
}

func (l *LinkedList) Traverse() []int {
	var nums []int

	curr := l.Head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	return nums
}

func Traverse(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
