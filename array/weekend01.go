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

// ContainsNearbyDuplicate 存在重复元素2
func ContainsNearbyDuplicate(nums []int, k int) bool {
	numsInfo := make(map[int]int)
	for i, num := range nums {
		if _, ok := numsInfo[num]; ok {
			return true
		}
		numsInfo[num] = i
		if len(numsInfo) > k {
			delete(numsInfo, nums[i-k])
		}
	}
	return false
}

// RemoveElement 移除元素
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int
	for _, v := range nums {
		if v == val {
			continue
		}
		nums[i] = v
		i++
	}
	return i
}
