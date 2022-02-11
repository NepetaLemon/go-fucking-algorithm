package n_queens_51

import (
	"strings"
)

// 出处：https://labuladong.gitee.io/algo/4/29/110/
// 题目：https://leetcode-cn.com/problems/permutations/
var res [][]string

func solveNQueens(n int) [][]string {
	// 初始化棋盘
	res = make([][]string, 0)

	board := make([][]string, n)
	for i, _ := range board {
		s := make([]string, n)
		for j := 0; j < n; j++ {
			s[j] = "."
		}
		board[i] = s
	}

	backtrack(board, 0)

	return res
}

func backtrack(board [][]string, row int) {
	// 当走到最后一行，则默认完成
	n := len(board)
	if row == n {
		// 注意，Go 的二维切片 copy 得自己来，copy(dst,src) 就是坑
		done := make([]string, 0)
		for i, _ := range board {
			// 这里棋盘要压扁一级，就是 让 row 拼接成字符串
			s := strings.Join(board[i], "")
			done = append(done, s)
		}

		res = append(res, done)

		return
	}

	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrack(board, row+1)
		board[row][col] = "."
	}
}

// 皇后可以攻击同一行、同一列、左上左下右上右下四个「方向」的任意单位。注意「方向」两个字，也就是说一个点位有 4 条「直线」禁区
// 检查 board[row][col] 是否可以放置皇后
func isValid(board [][]string, row, col int) bool {
	boardLen := len(board)
	// 「行」不用检查，我们只需要保证一行下一个皇后就行
	// 检查「列」上是否有冲突
	for i := 0; i < boardLen; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 检查「左上」是否有冲突，注意是「线」，注意 「r >= 0」，是可以 = 0 的
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	// 检查「右上」是否有冲突，注意是「线」
	for r, c := row-1, col+1; r >= 0 && c < boardLen; r, c = r-1, c+1 {
		if board[r][c] == "Q" {
			return false
		}
	}
	//「右下」和「左下」不需要检查是因为我们从上往下下皇后，下面还没有下皇后，自然也不会有冲突

	return true
}
