/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode.cn/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (65.58%)
 * Likes:    646
 * Dislikes: 0
 * Total Accepted:    476.5K
 * Total Submissions: 726.2K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * s 和 t 仅包含小写字母
 *
 *
 *
 *
 * 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

// @lc code=start
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	cnt1 := map[rune]int{}
// 	for _, ch := range s {
// 		cnt1[ch]++
// 	}
// 	cnt2 := map[rune]int{}
// 	for _, ch := range t {
// 		cnt2[ch]++
// 	}

//		for ch, cnt := range cnt1 {
//			if cnt != cnt2[ch] {
//				return false
//			}
//		}
//		return true
//	}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

