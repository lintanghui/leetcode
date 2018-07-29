package binary

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res = make([]int, 0, 0)
	view(root, &res, 0)
	return res
}
func view(root *TreeNode, res *[]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*res) {
		*res = append(*res, root.Val)
	}
	view(root.Right, res, depth+1)
	view(root.Left, res, depth+1)

}
