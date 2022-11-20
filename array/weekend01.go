package array

// RemoveDuplicates 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[j-1] {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	return i
}
