package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	cache := make(map[*TreeNode][]*TreeNode)
	traverse(root, target, cache)

	return calc(cache, target, nil, k)
}

// 遍历树
func traverse(root *TreeNode, target *TreeNode, cache map[*TreeNode][]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		cache[root] = append(cache[root], root.Left)
		cache[root.Left] = append(cache[root.Left], root)
		traverse(root.Left, target, cache)
	}

	if root.Right != nil {
		cache[root] = append(cache[root], root.Right)
		cache[root.Right] = append(cache[root.Right], root)
		traverse(root.Right, target, cache)
	}
}

func calc(cache map[*TreeNode][]*TreeNode, target *TreeNode, lastTarget *TreeNode, k int) []int {
	results := []int{}
	if target == nil {
		return results
	}

	if k == 0 {
		return append(results, target.Val)
	}

	for _, v := range cache[target] {
		if v != lastTarget {
			results = append(results, calc(cache, v, target, k-1)...)
		}
	}
	return results
}
