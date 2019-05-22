package leetcode

import "fmt"

var board = [][]byte{
	[]byte{'5','3','.','.','7','.','.','.','.'},
	[]byte{'6','.','.','1','9','5','.','.','.'},
	[]byte{'.','9','8','.','.','.','.','6','.'},
	[]byte{'8','.','.','.','6','.','.','.','3'},
	[]byte{'4','.','.','8','.','3','.','.','1'},
	[]byte{'7','.','.','.','2','.','.','.','6'},
	[]byte{'.','6','.','.','.','.','2','8','.'},
	[]byte{'.','.','.','4','1','9','.','.','5'},
	[]byte{'.','.','.','.','8','.','.','7','9'},
}

func ExampleIsValidSudoku() {
	fmt.Print(isValidSudoku(board))
	//Output:
	//true
}

func ExampleSolveSudoku() {
	solveSudoku(board)
	fmt.Print(board)
	//Output:
	//true
}