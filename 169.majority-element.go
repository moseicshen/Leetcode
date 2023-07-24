/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	//Boyer-Moore Majority Vote Algorithm
	count := 0
	majority := nums[0]
	for num := range nums {
		if count == 0 {
			majority = nums[num]
			count = 1
		} else {
			if majority == nums[num] {
				count++
			} else {
				count--
			}
		}
	}
	return majority
}

// @lc code=end

