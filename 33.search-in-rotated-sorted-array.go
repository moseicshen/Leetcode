/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	higher := 0
	lower := len(nums) - 1
	if nums[higher] == target {
		return 0
	}
	if nums[lower] == target {
		return len(nums) - 1
	}
	if target < nums[higher] && target > nums[lower] {
		return -1
	}
	//lower higher is num; up_bound, down_bound is position
	medium, up_bound, down_bound := -1, lower, higher
	if nums[higher] > target {
		for up_bound >= down_bound {
			medium = down_bound + (up_bound-down_bound)/2
			if nums[medium] == target {
				return medium
			} else {
				if nums[medium] > nums[lower] || nums[medium] < target {
					down_bound = medium + 1
				} else {
					up_bound = medium - 1
				}
			}
		}
	}
	if nums[higher] < target {
		for up_bound >= down_bound {
			medium = down_bound + (up_bound-down_bound)/2
			if nums[medium] == target {
				return medium
			} else {
				if nums[medium] < nums[higher] || nums[medium] > target {
					up_bound = medium - 1
				} else {
					down_bound = medium + 1
				}
			}
		}
	}
	return -1
}

// @lc code=end

