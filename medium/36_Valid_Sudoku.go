/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells
need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the
digits 1-9 without repetition.

Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.


Example 1:


Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true


Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner
being modified to 8. Since there are two 8's in the top left 3x3 sub-box,
it is invalid.


Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make(map[string]bool)
	cols := make(map[string]bool)
	squares := make(map[string]bool)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			value := board[c][r]

			if string(value) == "." {
				continue
			}

			rowV := fmt.Sprintf("row %v %v", r,string(value))
			colV := fmt.Sprintf("col %v %v", c,string(value))
			squareV := fmt.Sprintf("row %v %v %v", r/3, c/3,string(value))

			if rows[rowV] || cols[colV] || squares[squareV] {
				return false
			}

			rows[rowV] = true
			cols[colV] = true
			squares[squareV] = true
		}
	}

	return true
}
