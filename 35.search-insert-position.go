/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	medium := -1
	for low <= high {
		if high <= low && nums[low] != target {
			if nums[low] < target {
				medium = low + 1
			} else {
				medium = low
			}
			break
		}
		medium = low + (high-low)/2
		if nums[medium] == target {
			break
		} else if nums[medium] < target {
			low = medium + 1
		} else {
			high = medium - 1
		}
	}
	return medium
}

// @lc code=end

