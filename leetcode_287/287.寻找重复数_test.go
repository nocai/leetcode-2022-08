package leetcode287

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// TODO: 没有思路
// @lc code=start
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				return num
			}
		}
	}
	return 0
}

// @lc code=end

func TestFindDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	assert.Equal(t, 2, findDuplicate(nums))
	nums = []int{3, 1, 3, 4, 2}
	assert.Equal(t, 3, findDuplicate(nums))
}
