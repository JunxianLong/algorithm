package main

// merge 合并两个有序数组
func merge(nums1 []int, nums2 []int) []int {
	/*
		使用归并排序来做，需要开启新的内存空间nums
	*/
	nums := make([]int, 0)
	var i int
	var j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			nums = append(nums, nums2[j])
			j++
		} else {
			nums = append(nums, []int{nums1[i], nums2[j]}...)
			i++
			j++
		}
	}

	if i < len(nums1) {
		nums = append(nums, nums1[i:]...)
	}
	if j < len(nums2) {
		nums = append(nums, nums2[j:]...)
	}
	return nums
}

// fib  斐波那契数
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// URL:https://leetcode-cn.com/problems/kth-missing-positive-number/
// Time:20220130
// FindKthPositive 第k个缺失的正整数
func FindKthPositive(arr []int, k int) int {

	if arr[0] > k {
		return k
	}
	if arr[len(arr)-1] < k {
		return k - (arr[len(arr)-1] - len(arr)) + arr[len(arr)-1]
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		value := arr[mid] - mid - 1
		if value >= k {
			right = mid - 1
		} else if value < k {
			left = mid + 1
		}
	}

	return k - (arr[left-1] - left) + arr[left-1]
}

// twoSum 两数之和 II - 输入有序数组 DONE
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[j] + numbers[i]
		if sum == target {
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}

// twoSum2 两数之和 输入无序数组
func twoSum2(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, num := range nums {
		if j, ok := tmp[target-num]; ok {
			return []int{i, j}
		}
		tmp[num] = i
	}
	return []int{}
}

// URL:https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
// Time:20220131
// ReplaceElements 将每个元素替换为右侧最大元素
func ReplaceElements(arr []int) []int {

	nums := make([]int, len(arr))
	nums[len(arr)-1] = -1

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > nums[i] {
			nums[i-1] = arr[i]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return nums
}

// URL:https://leetcode-cn.com/problems/move-zeroes/
// Time:20220131
// MoveZeroes 移动零
func MoveZeroes(nums []int) {
	zero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[zero], nums[i] = nums[i], nums[zero]
		zero++
	}
}

// URL:https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
// Time:20220201
// FindLengthOfLCIS 最长连续递增序列
func FindLengthOfLCIS(nums []int) int {
	var start, max int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] > nums[i] {
			start = i
		}
		if max < i-start+1 {
			max = i - start + 1
		}
	}
	return max
}

// URL:https://leetcode-cn.com/problems/degree-of-an-array/
// Time:20220201
// FindShortestSubArray 数组的度
func FindShortestSubArray(nums []int) int {

	var (
		numMap = make(map[int][]int)
		min    = len(nums) // 最短连续数组
		max    int         // 度
	)

	for index, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = make([]int, 0)
		}
		numMap[num] = append(numMap[num], index)
		if len(numMap[num]) > max {
			max = len(numMap[num])
		}
	}

	if max == 1 {
		return 1
	}
	for _, indexes := range numMap {
		if len(indexes) != max {
			continue
		}
		if min > indexes[len(indexes)-1]-indexes[0]+1 {
			min = indexes[len(indexes)-1] - indexes[0] + 1
		}
	}
	return min
}

// URL:https://leetcode-cn.com/problems/two-sum/
// Time:20220202
// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}

// URL:https://leetcode-cn.com/problems/sort-array-by-parity-ii/
// Time:20220202
// SortArrayByParityII 按奇偶排序数组II
func SortArrayByParityII(nums []int) []int {

	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 0 {
			continue
		}
		for nums[j]%2 == 1 {
			j += 2
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
