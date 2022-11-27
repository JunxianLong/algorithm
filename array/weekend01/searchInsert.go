package weekend01

// searchInsert 搜索插入位置
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	var mid int
	for i <= j {
		mid = (i + j) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

// searchInsert2 搜索插入位置
func searchInsert2(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}

	}
	return i
}
