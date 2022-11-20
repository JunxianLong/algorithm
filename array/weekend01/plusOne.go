package weekend01

// url:https://leetcode.cn/problems/plus-one/
// plusOne 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]%10 != 0 {
			return digits
		}
		digits[i] = 0
	}
	tmp := []int{1}
	digits = append(tmp, digits...)
	return digits
}
