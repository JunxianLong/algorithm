package main

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func QuickSort(nums []int, left, right int) {
	if left < right {
		middle := partition(nums, left, right)
		QuickSort(nums, left, middle-1)
		QuickSort(nums, middle+1, right)
	}
}

// partition 分区
func partition(nums []int, left, right int) int {
	p := nums[left]
	for left < right {
		for ; left < right && nums[right] > p; right-- {
		}
		nums[left] = nums[right]
		for ; left < right && nums[left] < p; left++ {
		}
		nums[right] = nums[left]
	}
	nums[left] = p
	return left
}
