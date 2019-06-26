package leetcode

func solveNQueens(n int) [][]string {
	data := initData(n)
	cur := make([]int, n)
	solveNQueensBacktrack(data, cur, 0)
	return data.res
}

type solveNQueensData struct {
	res         [][]string
	col, d1, d2 []bool
	n           int
}

func initData(n int) *solveNQueensData {
	data := &solveNQueensData{}
	data.res = make([][]string, 0)
	data.col = make([]bool, n)
	data.d1 = make([]bool, 2*n-1)
	data.d2 = make([]bool, 2*n-1)
	data.n = n
	return data
}

func boardFromSlice(s []int) []string {
	n := len(s)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		a := make([]byte, n)
		for j := 0; j < n; j++ {
			a[j] = '.'
		}
		a[s[i]] = 'Q'
		res[i] = string(a)
	}
	return res
}

func solveNQueensBacktrack(data *solveNQueensData, current []int, row int) {
	if row == data.n {
		data.res = append(data.res, boardFromSlice(current))
		return
	}
	for col := 0; col < data.n; col++ {
		if data.col[col] {
			continue
		}
		d1Index := row + col
		if data.d1[d1Index] {
			continue
		}
		d2Index := col - row + data.n - 1
		if data.d2[d2Index] {
			continue
		}
		current[row] = col
		data.col[col] = true
		data.d1[d1Index] = true
		data.d2[d2Index] = true
		solveNQueensBacktrack(data, current, row+1)
		data.col[col] = false
		data.d1[d1Index] = false
		data.d2[d2Index] = false
	}
}
