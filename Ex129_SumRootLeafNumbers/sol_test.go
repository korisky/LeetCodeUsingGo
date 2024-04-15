package Ex129_SumRootLeafNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0, 0)
}

// dfs Go语言里面要慎用global variable, 基本都会有并发问题, 最好避免使用
func dfs(root *TreeNode, prev, sum int) int {

	if root == nil {
		return sum
	}

	curNum := prev*10 + root.Val
	sum = dfs(root.Left, curNum, sum)
	sum = dfs(root.Right, curNum, sum)

	if root.Left == nil && root.Right == nil {
		sum += curNum
	}

	return sum
}
