package delete_operation_for_two_strings_583

// memo[i][j] 定义为：s1 前 i 个 char 和 s2 前 j 个 char 的 lcs
var memo [][]int

// 这题等同「最长公共子序列」
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	memo = make([][]int, l1+1)
	for i := range memo {
		memo[i] = make([]int, l2+2)
	}
	// base case
	memo[0][0] = 0

	lcs := longestCommonSubsequence2(word1, word2)

	return l1 - lcs + l2 - lcs
}

// 当复习一次 lcs 好了
func longestCommonSubsequence2(text1 string, text2 string) int {
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			// 如果 c1 和 c2 相等，则 c 为 lcs 的一部分
			// memo[i+1][j+1] 里 i + 1 和 j + 1 的原因是 i 和 j 是 c 的索引，并不是我们对 memo 的定义
			if c1 == c2 {
				memo[i+1][j+1] = memo[i][j] + 1
			} else {
				// 实际上就是把之前的结果继承下来，举个例子：
				// memo[3][4] 的 lcs 为 2，那么 memo[4][4] 和 memo[3][5] 的 lcs 必然也是大于等于 2
				// 假设 i = 3, j = 4，则：
				// memo[i+1][j+1] 为 memo[4][5]
				// memo[i+1][j]   为 memo[4][4]
				// memo[i][j+1]   为 memo[3][5]
				memo[i+1][j+1] = max(
					memo[i+1][j], // 之前就求过，s1[i] 不是 lcs 的一部分
					memo[i][j+1], // 之前就求过，s2[j] 不是 lcs 的一部分
				)
			}
		}
	}

	return memo[len(text1)][len(text2)]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
