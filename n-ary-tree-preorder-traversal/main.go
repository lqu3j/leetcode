package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	results := []int{}

	if root == nil {
		return results
	}

	results = append(results, root.Val)
	for _, v := range root.Children {
		results = append(results, preorder(v)...)
	}
	return results
}
