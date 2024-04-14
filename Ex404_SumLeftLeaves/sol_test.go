package Ex404_SumLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeaves(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, bool)
	dfs = func(node *TreeNode, flag bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && flag {
			res += node.Val
		}
		dfs(node.Left, true)
		dfs(node.Right, false)
	}
	dfs(root, false)
	return res
}
