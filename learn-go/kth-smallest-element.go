/**
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
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

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		traverse(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		traverse(root.Right, res)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	traverse(root, &res)
	return res[k-1]
}
