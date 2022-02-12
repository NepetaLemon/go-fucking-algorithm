package generate_parentheses_26

import "strings"

var res []string

// 括号问题存在两个特性，
// 1. 一个「合法」括号组合的左括号数量一定等于右括号数量
// 2. 从左往右数，必然是左括号的数量大于等于右括号，因为左括号在左才是合法的（废话）
func generateParenthesis(n int) []string {
	res = make([]string, 0)
	track := make([]string, 2*n)
	backtrack(track, n, n)

	return res
}

func backtrack(track []string, leftAmount, rightAmount int) {
	// 数量小于 0 必然是不合法的
	if leftAmount < 0 || rightAmount < 0 {
		return
	}
	// 违反特性 2，因为我们是从左往右放，所以必须满足当前情况下左的小于等于右
	if leftAmount > rightAmount {
		return
	}
	// 满足条件，合法括号（左等于右，且括号已经用完）
	if leftAmount == 0 && rightAmount == 0 {
		res = append(res, strings.Join(track, ""))
		return
	}
	track = append(track, "(")
	backtrack(track, leftAmount-1, rightAmount)
	track = track[:len(track)-1]

	track = append(track, ")")
	backtrack(track, leftAmount, rightAmount-1)
	track = track[:len(track)-1]
}
