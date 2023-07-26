/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	result := 0

	var num = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	j := 0
	for j = 0; j < len(s)-1; j++ {
		if num[s[j]] < num[s[j+1]] {
			result -= num[s[j]]
		} else {
			result += num[s[j]]
		}
	}
	return result + num[s[j]]
}

// @lc code=end

