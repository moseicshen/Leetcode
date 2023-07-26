/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	flag := 0
	for i := range nums {
		if nums[i] == val {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return len(nums)
	}
	cur, finder := 0, 0
	for finder < len(nums) && cur < len(nums) {
		//fmt.Println("cur", cur)
		//fmt.Println("finder", finder)
		if nums[cur] == val {
			for finder = cur; finder < len(nums); finder++ {
				if nums[finder] != val {
					nums[cur], nums[finder] = nums[finder], nums[cur]
					break
				}
			}
		}
		cur++
	}
	return cur - 1
}

// @lc code=end

