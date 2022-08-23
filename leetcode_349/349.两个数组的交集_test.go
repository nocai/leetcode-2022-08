package leetcode349

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	set2 := map[int]struct{}{}
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	var result []int
	for num := range set1 {
		if _, exist := set2[num]; exist {
			result = append(result, num)
		}
	}
	return result
}

// @lc code=end
