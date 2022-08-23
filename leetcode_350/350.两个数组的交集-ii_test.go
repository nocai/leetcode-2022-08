/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
// func intersect(nums1 []int, nums2 []int) []int {
// 	var result []int
// 	for _, num1 := range nums1 {
// 		for index, num2 := range nums2 {
// 			if num1 == num2 {
// 				result = append(result, num1)
// 				nums2 = append(nums2[:index], nums2[index+1:]...)
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	var ret []int
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num]--
		}
	}
	return ret
}

// @lc code=end

