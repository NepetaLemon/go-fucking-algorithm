package climbing_stairs_70

// dp[n]：第 n 阶有多少种爬法
var dp []int

func climbStairs(n int) int {
	if n < 3 {
		dp = make([]int, 3)
	} else {
		dp = make([]int, n+1)
	}
	// 	状态：当前楼梯的阶数
	// 	选择：爬一步还是两步
	// base case
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
