/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cur, finder := 0, 1
	for cur < len(nums) {
		for ; finder < len(nums); finder++ {
			if nums[finder] != nums[cur] {
				cur++
				nums[cur] = nums[finder]
			}
		}
		if finder >= len(nums) {
			break
		}
	}
	return cur + 1
}

// @lc code=end

