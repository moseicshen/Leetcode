/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	right := len(nums) - 1
	left := 0
	if nums[left] < nums[right] {
		return nums[left]
	}
	high := right
	low := left
	for high >= low {
		if high == low {
			return nums[high]
		}
		if len(nums[low:high]) == 1 {
			if nums[low] > nums[high] {
				return nums[high]
			} else {
				return nums[low]
			}
		}
		medium := low + (high-low)/2
		if nums[medium] > nums[left] {
			low = medium
		} else {
			high = medium
		}
		//fmt.Println("medium", medium)
		//fmt.Println("high", high)
		//fmt.Println("low", low)
	}
	return -1
}

// @lc code=end

