package leetcode125

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.89%)
 * Likes:    563
 * Dislikes: 0
 * Total Accepted:    395.2K
 * Total Submissions: 843.1K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 解释："amanaplanacanalpanama" 是回文串
 *
 *
 * 示例 2:
 *
 *
 * 输入: "race a car"
 * 输出: false
 * 解释："raceacar" 不是回文串
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 字符串 s 由 ASCII 字符组成
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	var ss []rune
	for _, ch := range strings.ToLower(s) {
		if validRune(ch) {
			ss = append(ss, ch)
		}
	}

	left, right := 0, len(ss)-1
	for left < right {
		if ss[left] != ss[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func validRune(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// @lc code=end

func TestXxx(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	assert.Equal(t, true, isPalindrome(s))
}
