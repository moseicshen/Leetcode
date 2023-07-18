// given a roman numeral, convert it to an integer

package leetcode

func romanToInt(a string) int {
	//var num [15]int
	result := 0

	var num = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	j := 0
	for j = 0; j < len(a)-1; j++ {
		if num[a[j]] < num[a[j+1]] {
			result -= num[a[j]]
		} else {
			result += num[a[j]]
		}
	}
	return result + num[a[j]]
}
