/**
Given a singly linked list, group all odd nodes together followed by the even nodes.
Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL

*/
package exercise

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	even_beg := head.Next
	evenhead := even_beg
	oddhead := head

	for evenhead != nil && evenhead.Next != nil {
		// get all odds and assign to odd head next
		oddhead.Next = evenhead.Next
		oddhead = oddhead.Next
		// get all evens and assign to even head next
		evenhead.Next = oddhead.Next
		evenhead = evenhead.Next
	}

	oddhead.Next = even_beg

	return head
}
