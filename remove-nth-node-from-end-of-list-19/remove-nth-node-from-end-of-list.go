package remove_nth_node_from_end_of_list_19

import (
	. "go-fucking-algorithm/listnode"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	// 因为是要删除倒数第 n 个节点，且这个是单链表无法知道上一个节点，所以我们要先找到倒数 n + 1 的节点
	x := findFromEnd(dummy, n+1)
	// 删除倒数第 N 个节点
	x.Next = x.Next.Next

	return dummy.Next
}

// 返回倒数第 k 个节点
func findFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	// 先让 p1 走 k 步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	// 让 p1 和 p2 一起走，直到 p1 到达尽头
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
