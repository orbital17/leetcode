package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := InitGraph(numCourses)
	for _, v := range prerequisites {
		g.Add(v[0], v[1])
	}
	return !g.IsCycle()
}

type Graph struct {
	v [][]int
	n int
}

func InitGraph(n int) *Graph {
	g := &Graph{
		v: make([][]int, n),
		n: n,
	}
	return g
}

func (g *Graph) Add(i, j int) {
	g.v[i] = append(g.v[i], j)
}

func (g *Graph) IsCycle() bool {
	n := g.n
	indegree := make([]int, n)
	for _, i := range g.v {
		for _, j := range i {
			indegree[j]++
		}
	}
	for i := 0; i < n; i++ {
		j := 0
		for j < n && indegree[j] != 0 {
			j++
		}
		if j == n {
			return true
		}
		indegree[j] = -1
		for _, v := range g.v[j] {
			indegree[v]--
		}
	}
	return false
}
