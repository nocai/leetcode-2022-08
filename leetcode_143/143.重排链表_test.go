/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode.cn/problems/reorder-list/description/
 *
 * algorithms
 * Medium (64.20%)
 * Likes:    994
 * Dislikes: 0
 * Total Accepted:    204.2K
 * Total Submissions: 317.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * 请将其重新排列后变为：
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[1,4,2,3]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[1,5,2,4,3]
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 5 * 10^4]
 * 1 <= node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reorderList(head *ListNode) {
// 	nodes := []*ListNode{}
// 	for curr := head; curr != nil; curr = curr.Next {
// 		nodes = append(nodes, curr)
// 	}

// 	left, right := 0, len(nodes)-1
// 	for left < right {
// 		nodes[left].Next = nodes[right]
// 		left++
// 		nodes[right].Next = nodes[left]
// 		right--
// 	}
// 	nodes[left].Next = nil
// }

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil

	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		// next := head.Next

		// head.Next = prev
		// prev = head
		// head = next
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

// @lc code=end

