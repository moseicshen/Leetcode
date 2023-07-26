/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

func maxArea(height []int) int {
	start, end, curArea := 0, len(height)-1, 0
	for start < end {
		tempArea := 0
		if height[start] < height[end] {
			tempArea = height[start] * (end - start)
			start++
		} else {
			tempArea = height[end] * (end - start)
			end--
		}
		if tempArea > curArea {
			curArea = tempArea
		}
	}
	return curArea
}

// @lc code=end

