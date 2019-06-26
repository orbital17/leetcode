package leetcode

func totalNQueens(n int) int {
	data := totalNQueensInitData(n)
	totalNQueensBacktrack(data, 0)
	return data.res
}

type totalNQueensData struct {
	res         int
	col, d1, d2 []bool
	n           int
}

func totalNQueensInitData(n int) *totalNQueensData {
	data := &totalNQueensData{}
	data.col = make([]bool, n)
	data.d1 = make([]bool, 2*n-1)
	data.d2 = make([]bool, 2*n-1)
	data.n = n
	return data
}

func totalNQueensBacktrack(data *totalNQueensData, row int) {
	if row == data.n {
		data.res++
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
		data.col[col] = true
		data.d1[d1Index] = true
		data.d2[d2Index] = true
		totalNQueensBacktrack(data, row+1)
		data.col[col] = false
		data.d1[d1Index] = false
		data.d2[d2Index] = false
	}
}
