/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.cn/problems/powx-n/description/
 *
 * algorithms
 * Medium (37.94%)
 * Likes:    1021
 * Dislikes: 0
 * Total Accepted:    312.8K
 * Total Submissions: 824.5K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n^ ）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 *
 * 提示：
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * -10^4 <= x^n <= 10^4
 *
 *
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1 / pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	result := pow(x, n/2)
	result *= result

	if n%2 != 0 {
		result *= x
	}
	return result
}

// @lc code=end

