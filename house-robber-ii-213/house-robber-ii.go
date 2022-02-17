package house_robber_ii_213

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 状态：房子索引 i
	// 选择：打劫或者不打劫
	// dp 定义：对于 dp[i] 指的是从 i 开始打劫，最高能打劫到多少钱

	// 	题目定义了房子是一圈，也就是说有 3 种情况
	// 0 和 1 都不打劫（但是这个没有意义，因为所有打劫金额都是正数）
	// 打劫 0 ，不打劫 1
	// 打劫 1 ，不打劫 0
	// 由第二和第三种情况可得，存在两种打劫区间：[0, n-2] 和 [1,n-1]
	foo := dp(nums, 0, n-2)
	bar := dp(nums, 1, n-1)
	if foo > bar {
		return foo
	} else {
		return bar
	}
}

func dp(nums []int, start, end int) int {
	if end < start {
		return 0
	}
	newNums := nums[start : end+1]
	n := len(newNums)
	memo := make([]int, n+2) // n+2 主要是方便处理边界，其实从 n 开始打劫，能打劫到的金额已经是 0
	for i := n - 1; i >= 0; i-- {
		next := memo[i+1]
		nextTwo := memo[i+2]
		if next > (nextTwo + newNums[i]) {
			memo[i] = next
		} else {
			memo[i] = nextTwo + newNums[i]
		}
	}

	return memo[0]
}
