package minimum_depth_of_binary_tree_111

import (
	. "go-fucking-algorithm/treenode"
)

// 题目：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// 思路：https://labuladong.gitee.io/algo/4/29/115/

// 写 BFS 算法都是用「队列」这种数据结构，每次将一个节点周围的所有节点加入队
// BFS 找到的路径一定是最短的，但代价就是空间复杂度可能比 DFS 大很多
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// root 加入队列
	queue := []*TreeNode{root}
	// 0 的情况在上面已经处理了
	depth := 1

	// 原文中的 poll() 是「删除并返回第一个元素」
	// 原文中的 offer() 是「向链表末尾添加元素，返回是否成功，成功为 true，失败为 false」
	for len(queue) > 0 {
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			// 这里写得比较复杂，主要是不想引进其他库，核心就是从切片中弹出第一个元素
			// 注意，这里一定是取 0 ，因为我们要做的就是取第一个元素
			// 核心思想是，从一开始就计算好当前队列还有几个元素，然后处理几个元素，也就是说，后面加入的元素不会受到影响，处理的是「上一轮」加入的元素
			// ---实现 poll()---
			cur := queue[0]
			if len(queue) == 1 {
				queue = []*TreeNode{}
			} else {
				queue = queue[1:]
			}
			// ---实现 poll()---
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 在这里加步数
		depth++
	}

	return depth
}
