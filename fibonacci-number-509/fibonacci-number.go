package fibonacci_number_509

// 出处：https://labuladong.gitee.io/algo/3/23/71/
// 题目：https://leetcode-cn.com/problems/fibonacci-number/
// 这里使用「自底向下」+「DP 表」的方式
func fib(n int) int {
	// 这里的 n 必须加 1，因为 0 占了一个位置
	dp := make([]int, n+1)
	// base case
	dp[0] = 0
	if n > 0 {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
