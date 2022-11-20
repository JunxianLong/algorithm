package weekend01

// url:https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// removeDuplicates 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left := 1
	for right := 1; right < n; right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
