package house_robber_198

func rob(nums []int) int {
	n := len(nums)
	// 	状态：房子索引
	// 	选择：当前屋子抢还是不抢，但是题目规定了不能连续抢 2 家
	// 	重复子问题：从 house[i] 开始抢
	//  局部最优解：
	// 从 dp[start] 能抢到的最大值
	// 这题的思路是，贼最多只会跨过 1 家进行抢劫，原因很简单，跨过 2 家还不如把 1 和 3 一起抢了（每户金额都是正数）
	// 该题还可以优化空间复杂度，但是方便理解没处理
	var dp = make([]int, len(nums)+2)
	for i := n - 1; i >= 0; i-- {
		one := dp[i+1] // 打劫下一家
		two := dp[i+2] // 打劫下下一家
		if one > (nums[i] + two) {
			dp[i] = one
		} else {
			dp[i] = nums[i] + two
		}
	}

	return dp[0]
}
