package weekend01

// url:https://leetcode.cn/problems/remove-element/
// removeElement 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var left int
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}
	return left
}
