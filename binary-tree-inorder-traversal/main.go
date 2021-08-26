package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	array := []int{}
	if root == nil {
		return array
	}
	if root.Left != nil {
		array = inorderTraversal(root.Left)
	}
	array = append(array, root.Val)

	if root.Right != nil {
		array = append(array, inorderTraversal(root.Right)...)
	}
	return array
}
