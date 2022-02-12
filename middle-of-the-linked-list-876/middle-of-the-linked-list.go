package middle_of_the_linked_list_876

import (
	. "go-fucking-algorithm/listnode"
)

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		// 慢指针走一步，快指针走两步，尽头的时候，慢指针就是中点
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
