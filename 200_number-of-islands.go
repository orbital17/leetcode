package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	uf := InitUnionFind(n * m)

	check := func(i, j int) bool {
		if i < 0 || j < 0 || i >= n || j >= m {
			return false
		}
		return grid[i][j] == '1'
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				uf.Add(m*i + j)
				if check(i-1, j) {
					uf.Union(m*i+j, m*(i-1)+j)
				}
				if check(i, j-1) {
					uf.Union(m*i+j, m*i+j-1)
				}
			}
		}
	}
	return uf.Count()
}

type UnionFind struct {
	data  []int
	count int
}

func InitUnionFind(n int) *UnionFind {
	uf := &UnionFind{}
	uf.data = make([]int, n)
	for i := 0; i < n; i++ {
		uf.data[i] = -1
	}
	return uf
}

func (uf *UnionFind) Union(i, j int) {
	f1, f2 := uf.find(i), uf.find(j)
	if f1 != f2 {
		uf.data[f1] = f2
		uf.count--
	}
}

func (uf *UnionFind) find(i int) int {
	if i == uf.data[i] {
		return i
	}
	uf.data[i] = uf.find(uf.data[i])
	return uf.data[i]
}

func (uf *UnionFind) Add(i int) {
	if uf.data[i] < 0 {
		uf.data[i] = i
	}
	uf.count++
}

func (uf *UnionFind) Count() int {
	return uf.count
}
