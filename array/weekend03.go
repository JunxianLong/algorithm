package array

import "sort"

// URL:https://leetcode-cn.com/problems/relative-sort-array/
// Time:20220112
// relativeSortArray 数组相对排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if v > upper {
			upper = v
		}
	}
	frequency := make([]int, upper+1)
	for _, v := range arr1 {
		frequency[v]++
	}

	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; frequency[v] > 0; frequency[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range frequency {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}

// runningSum 一维数组的动态和
func runningSum(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i+1] += nums[i]
	}
	return nums
}

// checkIfExist 检查整数及其两倍数是否存在
func checkIfExist(arr []int) bool {

	numsMap := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := numsMap[v]; ok {
			return true
		} else {
			if v%2 == 0 {
				numsMap[v/2] = struct{}{}
			}
			numsMap[v*2] = struct{}{}
		}
	}
	return false
}

// shuffle 重新排列数组
func shuffle(nums []int, n int) []int {
	return nums
}

// fairCandySwap 公平的糖果交换
func fairCandySwap(A []int, B []int) []int {

	var suma int
	var sumb int

	numMapA := make(map[int]int)
	numMapB := make(map[int]int)

	for _, v := range A {
		suma += v
		numMapA[v]++
	}

	for _, v := range B {
		sumb += v
		numMapB[v]++
	}

	resp := make([]int, 0)
	for _, x := range A {
		if _, ok := numMapB[(sumb-suma)/2+x]; ok {
			resp = append(resp, x)
			resp = append(resp, (sumb-suma)/2+x)
			break
		}
	}
	return resp
}

// kLengthApart 是否所有 1 都至少相隔 k 个元素
func kLengthApart(nums []int, k int) bool {
	/*
		维护一个k+1大小的窗口，如果这个窗口出现两个1，返回false
	*/

	numMap := make(map[int]struct{})
	count := 0
	for index, num := range nums {
		if _, ok := numMap[num]; ok && num == 1 {
			return false
		}
		numMap[num] = struct{}{}
		count++
		if count == k+1 {
			delete(numMap, nums[index-k])
			count--
		}
	}

	return true
}

// missingNumber 消失的数字
func missingNumber(nums []int) int {

	/*
		求和，然后相减
	*/

	a := len(nums)
	b := 0
	for i, v := range nums {
		a += i
		b += v

	}
	return a - b
}

// average 去掉最低工资和最高工资后的工资平均值
func average(salary []int) float64 {
	return 0
}

// addToArrayForm 数组形式的整数加法
func addToArrayForm(A []int, K int) []int {

	/*
		10的0次方，10的1次方，10的2次方
	*/

	ans := make([]int, 0)

	for i := len(A) - 1; i >= 0; i-- {
		sum := A[i] + K%10
		K /= 10
		if sum >= 10 {
			K++
			sum -= 10
		}
		ans = append(ans, sum)
	}

	for ; K > 0; K /= 10 {
		ans = append(ans, K%10)
	}

	reverse(ans)
	return ans
}

// 数组翻转
func reverse(A []int) {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}

// finalPrices 商品折扣后的最终价格
func finalPrices(prices []int) []int {
	/*
		单调栈
	*/

	return nil

}

// sumZero 和为零的N个唯一整数
func sumZero(n int) []int {

	ans := make([]int, 0)
	sum := 0
	for i := 1; i < n; i++ {
		ans = append(ans, i)
		sum -= i
	}
	ans = append(ans, sum)
	return ans
}

// canBeEqual 通过翻转子数组使两个数组相等
func canBeEqual(target []int, arr []int) bool {
	targetMap := make(map[int]int)
	arrMap := make(map[int]int)

	for i := 0; i < len(target); i++ {
		targetMap[target[i]] += 1
		arrMap[arr[i]] += 1
	}

	if len(targetMap) != len(arrMap) {
		return false
	}

	for k, v1 := range targetMap {
		if v2, ok := arrMap[k]; ok {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// minCostClimbingStairs 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {

	//min1 := cost[0]
	//min2 := cost[1]
	//for i := 2; i < len(cost); i += 2 {
	//	if cost[i]
	//}

	return 0
}

// URL:https://leetcode-cn.com/problems/relative-sort-array/submissions/
// Time:20220203
// RelativeSortArray 数组的相对排序
func RelativeSortArray(arr1 []int, arr2 []int) []int {

	order := make(map[int]int)
	for i, v := range arr2 {
		order[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		// i > j，return true就是需要调换位置
		n1, n2 := arr1[i], arr1[j]
		order1, ok1 := order[n1]
		order2, ok2 := order[n2]
		if ok1 && ok2 {
			// 如果两个都存在就按照arr2规定排序
			return order1 < order2
		}
		if ok1 || ok2 {
			// 如果只有一个存在arr2中，按照题意，不存在的应该在右边
			return ok1
		}

		return n1 < n2
	})
	return arr1
}

// URL:https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// Time:20220204
// SmallerNumbersThanCurrent 有多少小于当前数字的数字
func SmallerNumbersThanCurrent(nums []int) []int {
	arr := [101]int{}
	for _, num := range nums {
		arr[num]++
	}

	for i := 0; i < 100; i++ {
		arr[i+1] += arr[i]
	}

	result := make([]int, len(nums))

	for i, num := range nums {
		if num > 0 {
			result[i] = arr[num-1]
		}
	}
	return result
}

// URL:https://leetcode-cn.com/problems/guess-numbers/
// Time:20220204
// Game 猜数字
func Game(guess []int, answer []int) int {
	var count int
	for i, v := range guess {
		if answer[i] != v {
			continue
		}
		count++
	}
	return count
}

// URL:https://leetcode-cn.com/problems/array-partition-i/
// Time:20220204
// PairSumArray 数组拆分I
func PairSumArray(nums []int) int {
	sort.Ints(nums)
	var count int
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] > nums[i+1] {
			count += nums[i+1]
		} else {
			count += nums[i]
		}
	}
	return count
}
