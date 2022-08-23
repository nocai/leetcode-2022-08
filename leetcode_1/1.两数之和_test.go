/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// func twoSum(nums []int, target int) []int {
// 	for i, num := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if num+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, num := range nums {
		if p, exist := m[target-num]; exist {
			return []int{p, index}
		}
		m[num] = index
	}

	return nil
}

// @lc code=end

