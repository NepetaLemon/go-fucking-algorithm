package coin_change_322

// 题目：https://leetcode-cn.com/problems/coin-change/
// 出处：https://labuladong.gitee.io/algo/3/23/71/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化，因为硬币最小的可能值为 1，所以我们认为数字最大的值的凑成数量为 1 块钱的数量，使用 amount+1表示无解
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	// 最优子结构其实就是，amount - coin，可以理解为每个数值只需要在 (amount - coin) + 1 就可以得到达到，逐步求解，直到到达 base case
	for k := 1; k < len(dp); k++ {
		for _, coin := range coins {
			// 无解，放弃
			if k-coin < 0 {
				continue
			}
			if (dp[k-coin] + 1) < dp[k] {
				dp[k] = dp[k-coin] + 1
			}
		}
	}

	// 前面填充了整个切片的值为 amount + 1，且如果无解则跳过，即是认为 amount + 1 为 -1
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}
