package leetcode

func majorityElement(nums []int) int {
	var times = make(map[int]int)
	for num := range nums {
		times[nums[num]] += 1
	}
	result := nums[0]
	for num := range nums {
		if times[nums[num]] > times[result] {
			result = nums[num]
		}
	}
	return result
}
