/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	root, low, high := 0, 0, x/2
	for high >= low {
		if x == 1 {
			return 1
		}
		root = (high-low)/2 + low
		if root*root == x || (root*root < x && (root+1)*(root+1) > x) {
			return root
		} else if root*root < x {
			low = root + 1
		} else {
			high = root - 1
		}
	}
	return low

}

// @lc code=end

