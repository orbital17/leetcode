package leetcode

import . "./lib"

import "container/list"

type LCAState int

const (
	BP LCAState = iota
	LD
	BD
)

type LCAItem struct {
	node  *TreeNode
	state LCAState
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	stack := list.New()
	stack.PushBack(&LCAItem{root, BP})

	oneFound := false
	var lcaIndex int
	var lca *LCAItem

	for stack.Len() != 0 {
		e := stack.Back().Value.(*LCAItem)
		switch {
		case e.node == nil || e.state == BD:
			stack.Remove(stack.Back())
			if oneFound && stack.Len() == lcaIndex {
				lcaIndex--
				lca = stack.Back().Value.(*LCAItem)
			}
		case e.state == BP:
			if e.node == p || e.node == q {
				if oneFound {
					return lca.node
				} else {
					oneFound = true
					lcaIndex = stack.Len() - 1
					lca = e
				}
			}
			stack.PushBack(&LCAItem{e.node.Left, BP})
			e.state = LD
		case e.state == LD:
			stack.PushBack(&LCAItem{e.node.Right, BP})
			e.state = BD
		}
	}
	panic("")
}

func lowestCommonAncestorRec(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	lowestCommonAncestorHelper(root, p, q, &ans)
	return ans
}

func lowestCommonAncestorHelper(node, p, q *TreeNode, ans **TreeNode) bool {
	if node == nil {
		return false
	}
	l := lowestCommonAncestorHelper(node.Left, p, q, ans)
	r := lowestCommonAncestorHelper(node.Right, p, q, ans)
	mid := node == p || node == q
	if l && r || l && mid || r && mid {
		*ans = node
	}
	return l || r || mid
}
