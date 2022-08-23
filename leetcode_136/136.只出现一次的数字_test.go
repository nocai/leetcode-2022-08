package leetcode136

import (

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode.cn/problems/single-number/description/
 *
 * algorithms
 * Easy (72.25%)
 * Likes:    2538
 * Dislikes: 0
 * Total Accepted:    759.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,1]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: [4,1,2,1,2]
 * 输出: 4
 *
 */

// @lc code=start
// func singleNumber(nums []int) int {
// 	set := map[int]struct{}{}
// 	for _, num := range nums {
// 		if _, exist := set[num]; exist {
// 			delete(set, num)
// 		} else {
// 			set[num] = struct{}{}
// 		}
// 	}

//		for num := range set {
//			return num
//		}
//		return 0
//	}

// func singleNumber(nums []int) int {
// 	m := map[int]int{}
// 	for _, num := range nums {
// 		if _, exist := m[num]; exist {
// 			m[num]++
// 		} else {
// 			m[num] = 1
// 		}
// 	}

// 	for num, count := range m {
// 		if count == 1 {
// 			return num
// 		}
// 	}
// 	return 0
// }

func singleNumber(nums []int) int {
	set := map[int]struct{}{}

	var sum int
	for _, num := range nums {
		sum += num
		set[num] = struct{}{}
	}

	var sum2 int
	for num := range set {
		sum2 += num
	}

	return sum2*2 - sum
}

// @lc code=end

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1} 
	nu
	assert.Equal(t, 1, singleNumbe(nums))
}
