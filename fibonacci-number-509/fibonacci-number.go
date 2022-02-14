package fibonacci_number_509

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
