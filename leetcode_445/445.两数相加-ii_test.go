/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode.cn/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (59.98%)
 * Likes:    548
 * Dislikes: 0
 * Total Accepted:    108.8K
 * Total Submissions: 181.2K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 *
 *
 * 示例1：
 *
 *
 *
 *
 * 输入：l1 = [7,2,4,3], l2 = [5,6,4]
 * 输出：[7,8,0,7]
 *
 *
 * 示例2：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[8,0,7]
 *
 *
 * 示例3：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表的长度范围为 [1, 100]
 * 0 <= node.val <= 9
 * 输入数据保证链表代表的数字无前导 0
 *
 *
 *
 *
 * 进阶：如果输入链表不能翻转该如何解决？
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

// 1. 构建两个桟
// 2. 生成返回链表时，直接反序
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var prev *ListNode
	var carry int
	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		sum := carry
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		carry = sum / 10

		prev = &ListNode{
			Val:  sum % 10,
			Next: prev,
		}
	}
	return prev
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	l1 = reverse(l1)
// 	l2 = reverse(l2)

// 	dummy := &ListNode{}
// 	carry, curr := 0, dummy
// 	for l1 != nil || l2 != nil {
// 		sum := carry
// 		if l1 != nil {
// 			sum += l1.Val
// 			l1 = l1.Next
// 		}

// 		if l2 != nil {
// 			sum += l2.Val
// 			l2 = l2.Next
// 		}

// 		carry = sum / 10
// 		curr.Next = &ListNode{Val: sum % 10}
// 		curr = curr.Next
// 	}

// 	if carry > 0 {
// 		curr.Next = &ListNode{Val: carry}
// 	}

// 	return reverse(dummy.Next)
// }

// func reverse(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	for head != nil {
// 		prev, head, head.Next = head, head.Next, prev
// 	}
// 	return prev
// }

// @lc code=end

