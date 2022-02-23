package coin_change_2_518

// 若只使用前 i 个物品（可以重复使用），当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包
var dp [][]int

// 	状态：可选择的物品（可重复使用）、背包的容量
// 	选择：物品放还是不放进背包
func change(amount int, coins []int) int {
	n := len(coins)
	// base case
	// dp[0][..] = 0，因为如果不使用任何硬币面值，就无法凑出任何金额
	// dp[..][0] = 1，如果凑出的目标金额为 0，那么不取任何一个硬币就是唯一的一种凑法（空数组）
	// 我们最终想得到的答案就是 dp[n][amount]
	dp = make([][]int, n+1)
	for i := range dp {
		item := make([]int, amount+1)
		item[0] = 1
		dp[i] = item
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			// 如果当前硬币面额大于所求金额，没有意义
			if (j - coins[i-1]) >= 0 {
				// 如果你不把这第 i 个物品装入背包，也就是说你不使用 coins[i] 这个面值的硬币，那么凑出面额 j 的方法数 dp[i][j] 应该等于 dp[i-1][j]，继承之前的结果。
				// 如果你把这第 i 个物品装入了背包，也就是说你使用 coins[i] 这个面值的硬币，那么 dp[i][j] 应该等于 dp[i][j-coins[i-1]]。
				// dp[i][j-coins[i-1]] 也不难理解，如果你决定使用这个面值的硬币，那么就应该关注如何凑出金额 j - coins[i-1]。
				// 比如说，你想用面值为 2 的硬币凑出金额 5，那么如果你知道了凑出金额 3 的方法，再加上一枚面额为 2 的硬币，不就可以凑出 5 了嘛。
				// 综上就是两种选择，而我们想求的 dp[i][j] 是「共有多少种凑法」，所以 dp[i][j] 的值应该是以上两种选择的结果之和：
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][amount]
}
