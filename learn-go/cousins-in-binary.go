/**
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.
*/

package exercise

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sameParent(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	return ((root.Left != nil && root.Left.Val == x && root.Right != nil && root.Right.Val == y) ||
		(root.Left != nil && root.Left.Val == y && root.Right != nil && root.Right.Val == x) ||
		sameParent(root.Left, x, y) ||
		sameParent(root.Right, x, y))
}

func getLevel(root *TreeNode, value int, level int) int {
	if root == nil {
		return 0
	}
	if root.Val == value {
		return level
	}
	leftLevel := getLevel(root.Left, value, level+1)
	if leftLevel != 0 {
		return leftLevel
	}
	return getLevel(root.Right, value, level+1)
}

func isCousins(root *TreeNode, x int, y int) bool {
	if sameParent(root, x, y) {
		return false
	}
	xLevel := getLevel(root, x, 1)
	yLevel := getLevel(root, y, 1)
	if xLevel == yLevel {
		return true
	}
	return false
}
