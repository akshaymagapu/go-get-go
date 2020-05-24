/**
Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.
Also recall that a preorder traversal displays the value of the node first,
then traverses node.left, then traverses node.right.)

It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.

Example 1:

Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
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

func add(root *TreeNode, val int) {
	if val <= root.Val {
		if root.Left != nil {
			add(root.Left, val)
		} else {
			root.Left = &TreeNode{val, nil, nil}
		}
	} else {
		if root.Right != nil {
			add(root.Right, val)
		} else {
			root.Right = &TreeNode{val, nil, nil}
		}
	}
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{preorder[0], nil, nil}
	for i := 1; i < len(preorder); i++ {
		add(root, preorder[i])
	}
	return root
}
