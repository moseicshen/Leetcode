/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	right := len(nums) - 1
	left := 0
	if nums[right] > nums[right-1] {
		return right
	}
	if nums[left+1] < nums[left] {
		return left
	}
	high := right
	low := left
	for high >= low {
		medium := low + (high-low)/2
		// medium on any of the end?
		//fmt.Println("medium", medium)
		//fmt.Println("high", high)
		//fmt.Println("low", low)
		if medium == left {
			if nums[medium] > nums[medium+1] {
				return medium
			} else {
				low = medium + 1
			}
		} else if medium == right {
			if nums[medium] > nums[medium-1] {
				return medium
			} else {
				high = medium - 1
			}
		} else {
			if nums[medium-1] < nums[medium] && nums[medium] > nums[medium+1] {
				return medium
			} else if nums[medium-1] < nums[medium] && nums[medium] < nums[medium+1] {
				low = medium + 1
			} else if nums[medium-1] > nums[medium] && nums[medium] > nums[medium+1] {
				high = medium - 1
			} else {
				low = medium + 1
			}
		}
	}
	return -1
}

// @lc code=end

