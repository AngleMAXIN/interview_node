package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _isBalanced(n *TreeNode) int {
	if n == nil {
		return 0
	}
	left := _isBalanced(n.Left)
	right := _isBalanced(n.Right)
	max := right
	if left > right {
		max = left
	}
	return max + 1
}

func isBalanced_1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := _isBalanced(root.Left)
	right := _isBalanced(root.Right)
	sub := right - left
	return (sub > -2 && sub < 2) && isBalanced(root.Right) && isBalanced(root.Left)
}

func depth(n *TreeNode) int {
	if n == nil {
		return 0
	}
	left := depth(n.Left)
	right := depth(n.Right)
	if left == -1 || right == -1 {
		return -1
	}
	sub := right - left
	if sub < -1 || sub > 1 {
		return -1
	} else {
		max := left
		if right > max {
			max = right
		}
		return max + 1
	}

}
func isBalanced(root *TreeNode) bool {
	res := depth(root)
	if res == -1 {
		return false
	}

	return true
}
