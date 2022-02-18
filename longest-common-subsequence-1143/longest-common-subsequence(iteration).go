package longest_common_subsequence_1143

// 自底向上
// 迭代
func longestCommonSubsequence2(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	// memo 存储的是到 [i][j] 的 lcs，从左到右依次增大
	memo = make([][]int, l1+1)
	for i := range memo {
		memo[i] = make([]int, l2+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				memo[i+1][j+1] = memo[i][j] + 1
			} else {
				memo[i+1][j+1] = max(memo[i+1][j], memo[i][j+1])
			}
		}
	}

	return memo[l1][l2]
}
