package longest_common_subsequence_1143

var memo [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	memo = make([][]int, l1)
	// 填充默认值
	for i := 0; i < l1; i++ {
		tmp := make([]int, l2)
		for j := 0; j < l2; j++ {
			tmp[j] = -1
		}
		memo[i] = tmp
	}

	return dp(text1, 0, text2, 0)
}

// 定义：计算 s1[i...] 和 s2[j...] 的最长公共子序列
func dp(s1 string, i int, s2 string, j int) int {
	// base case
	// 显然，如果从任意一个字符串的最后面（相当于空串）开始计算，那么 LCS 就为 0
	// 且 LCS 的理论最大值就是较短字符串的长度
	if len(s1) == i || len(s2) == j {
		return 0
	}
	// 已经计算过，直接返回
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// 如果相等，s1[i] 和 s2[j] 必然在 lcs 中
	if s1[i] == s2[j] {
		memo[i][j] = 1 + dp(s1, i+1, s2, j+1)
	} else {
		// s1[i] 和 s2[j] 中至少有一个字符不在 lcs 中
		// 穷举三种情况的结果，取其中的最大结果
		memo[i][j] = max(
			dp(s1, i+1, s2, j), // s1[i] 不在 lcs 中
			dp(s1, i, s2, j+1), // s2[j] 不在 lcs 中
			// dp(s1, i+1, s2, j+1), // 两个字符都不在 lcs 中，去计算这个没有意义
		)
	}

	return memo[i][j]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
