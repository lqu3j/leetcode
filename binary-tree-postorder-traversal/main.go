package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	results := []int{}

	if root == nil {
		return results
	}

	if root.Left != nil {
		results = append(results, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		results = append(results, postorderTraversal(root.Right)...)
	}

	return append(results, root.Val)
}
