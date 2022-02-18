package interleaving_string_97

// 定义 memo[i][j] 为 s1 「前」 i 个元素和 s2 「前」 j 个元素是否能组成 s3 「前」 i + j 个元素
// 一定要注意 memo 的定义，特别是「前」字
var memo [][]bool

func isInterleave(s1 string, s2 string, s3 string) bool {
	// 长度不等，必然不可能
	l1, l2, t := len(s1), len(s2), len(s3)
	if t != l1+l2 {
		return false
	}
	memo = make([][]bool, l1+1)
	for i := range memo {
		memo[i] = make([]bool, l2+1)
	}
	// 	base case
	memo[0][0] = true
	// 如果 s1[i] 和 s2[j] 能组成 s3[i+j]，分成两种情况
	// s1[i] == s3[i+j]，此时要求 s[i-1][j] 也为真，直到推导到 base case，例如 ab + cd = cdab，也会要求 a + cd = cda 成立
	// s2[j] == s3[i+j]，此时要求 s[i][j-1] 也为真，直到推导到 base case，例如 ab + cd = abcd，也会要求 ab + c = abc 成立
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			// 注意 memo 的定义
			p := i + j - 1
			if i > 0 {
				// 或 memo[i][j] 是因为 memo[i][j] 可能会被计算两次：s1[i] == s3[i+j] 和 s2[j] == s3[i+j]
				memo[i][j] = memo[i][j] || (memo[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				memo[i][j] = memo[i][j] || memo[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}

	return memo[l1][l2]
}
