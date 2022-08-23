package leetcode_69

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start

// func mySqrt(x int) int {
// 	left, right := 1, x // [left, right]

// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if mid*mid < x {
// 			left = mid + 1
// 		} else if mid*mid > x {
// 			right = mid - 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return right
// }

// func mySqrt(x int) int {
// 	left, right := 0, x // (left, right]

// 	for left < right {
// 		mid := left + (right-left)/2 + 1
// 		if mid*mid < x {
// 			left = mid
// 		} else if mid*mid > x {
// 			right = mid - 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return left
// }

func mySqrt(x int) int {
	left, right := 0, x+1 // (left, right)

	for right-left > 1 {
		mid := left + (right-left)/2
		if mid*mid < x {
			left = mid
		} else if mid*mid > x {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

// @lc code=end

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(1))
}
