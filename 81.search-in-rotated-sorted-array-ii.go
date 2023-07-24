/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 1 && nums[0] != target {
		return false
	}
	higher := 0
	lower := len(nums) - 1
	if nums[higher] == target || nums[lower] == target {
		return true
	}
	if target < nums[higher] && target > nums[lower] {
		return false
	}
	//lower higher is num; up_bound, down_bound is position
	medium, up_bound, down_bound := -1, lower, higher
	if nums[higher] > target {
		for up_bound >= down_bound {
			medium = down_bound + (up_bound-down_bound)/2
			if nums[medium] == target {
				return true
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
				return true
			} else {
				if nums[medium] < nums[higher] || nums[medium] > target {
					up_bound = medium - 1
				} else {
					down_bound = medium + 1
				}
			}
		}
	}
	return false
}

// @lc code=end

