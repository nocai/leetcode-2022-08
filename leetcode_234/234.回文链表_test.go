package leetcode234

import "testing"

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (52.25%)
 * Likes:    1470
 * Dislikes: 0
 * Total Accepted:    483.6K
 * Total Submissions: 925.1K
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
// func isPalindrome(head *ListNode) bool {
// 	vals := []int{}

// 	curr := head
// 	for curr != nil {
// 		vals = append(vals, curr.Val)
// 		curr = curr.Next
// 	}

// 	n := len(vals)
// 	for index, val := range vals[:n/2] {
// 		if val != vals[n-1-index] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func isPalindrome(head *ListNode) bool {
// 	frontPoint := head

// 	var check func(*ListNode) bool
// 	check = func(curr *ListNode) bool {
// 		if curr != nil {
// 			if !check(curr.Next) {
// 				return false
// 			}

// 			if curr.Val != frontPoint.Val {
// 				return false
// 			}

// 			frontPoint = frontPoint.Next
// 		}
// 		return true
// 	}
// 	return check(head)
// }

// func isPalindrome(head *ListNode) bool {
// 	nodes := []int{}
// 	for head != nil {
// 		nodes = append(nodes, head.Val)
// 		head = head.Next
// 	}

// 	left, right := 0, len(nodes)-1
// 	for left < right {
// 		if nodes[left] != nodes[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

func isPalindrome(head *ListNode) bool {
	half := half(head)
	list := reverse(half)

	for list != nil {
		if list.Val != head.Val {
			return false
		}
		head = head.Next
		list = list.Next
	}
	return true

}

func half(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev = &ListNode{
			Val:  head.Val,
			Next: prev,
		}
		head = head.Next
	}
	return prev
}

// @lc code=end

func TestXxx(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	isPalindrome(head)
}
