package longest_increasing_subsequence_300

// 出处：https://labuladong.gitee.io/algo/3/23/72/
// 题目：https://leetcode-cn.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	var max int
	dp := make([]int, len(nums))
	// base case ，数组全部初始化为 1，假设都是以自己开头，所以是 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	// 遍历整个 nums
	for i := 0; i < len(nums); i++ {
		// 从 i 的位置，往前看，看看谁
		for j := 0; i > j; j++ {
			if nums[i] > nums[j] {
				if (dp[j] + 1) > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	// 找出 dp 里最大值
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}
