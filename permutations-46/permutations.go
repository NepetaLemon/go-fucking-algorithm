package permutations_46

// 出处：https://labuladong.gitee.io/algo/4/29/110/
// 题目：https://leetcode-cn.com/problems/permutations/
var res [][]int

func permute(nums []int) [][]int {
	// 力扣中，使用全局变量务必要重置，不然在单测中保留上次运行的结果
	res = [][]int{}
	visited := map[int]bool{}
	track := make([]int, 0)
	backtrack(nums, track, visited)

	return res
}

// 因为数组中可能包含元素 0，所以这里用
func backtrack(nums []int, track []int, visited map[int]bool) {
	// 触发结束条件
	// 这里的条件是 track 的长度等于 nums 的长度
	// 这里直接将 track 的 copy 当做可行结果保存起来
	numsLen := len(nums)
	if len(track) == numsLen {
		// 复制一个 track，切片是引用传递，Go 上的坑
		tmp := make([]int, numsLen)
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := 0; i < numsLen; i++ {
		if visited[i] {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		// 进入下一层决策树
		visited[i] = true
		backtrack(nums, track, visited)
		// 取消选择
		track = track[:len(track)-1]
		visited[i] = false
	}
}
