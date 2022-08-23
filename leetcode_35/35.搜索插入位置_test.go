package leetcode35

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode.cn/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.15%)
 * Likes:    1660
 * Dislikes: 0
 * Total Accepted:    894.3K
 * Total Submissions: 2M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 请必须使用时间复杂度为 O(log n) 的算法。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,3,5,6], target = 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,3,5,6], target = 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums = [1,3,5,6], target = 7
 * 输出: 4
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 为 无重复元素 的 升序 排列数组
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start

// func searchInsert(nums []int, target int) int {
// 	left, right := 0, len(nums) // [left, right)

// 	for left < right {
// 		mid := left + (right-left)/2
// 		if nums[mid] < target {
// 			left = mid + 1
// 		} else if nums[mid] > target {
// 			right = mid
// 		} else {
// 			return mid
// 		}
// 	}
// 	// return left
// 	return right
// }

func searchInsert(nums []int, target int) int {
	left, right := -1, len(nums) // (left, right)
	for right-left > 1 {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return right
}

// @lc code=end

func TestAAA(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	assert.Equal(t, 2, searchInsert(nums, 5))
}
