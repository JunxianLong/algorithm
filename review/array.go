package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// removeDuplicates 移除重复元素
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if (i == 0) || (nums[i-1] != nums[j]) {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return i
}

// isPalindrome 回文窜
func isPalindrome(s string) bool {
	var str string
	for i := 0; i < len(s); i++ {
		if isAlNum(s[i]) {
			str += string(s[i])
		}
	}

	str = strings.ToUpper(str)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func isAlNum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

// ,ab,
// isPalindrome2 回文窜2
func isPalindrome2(s string) bool {
	s = strings.ToUpper(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlNum(s[i]) {
			i++
		}
		for i < j && !isAlNum(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}
	return true
}

// isSubsequence 是否子序列
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

//输入：nums1 = [4,5,6,0,0,0], m = 3, nums2 = [2,3,6], n = 3
//输出：[1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	for i := 0; i < p2+1; i++ {
		nums1[i] = nums2[i]
	}
}

// merge2 合并两个有序数组
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, n+m-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for index := 0; index <= j; index++ {
		nums1[index] = nums2[index]
	}
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			//nums[i], nums[j] = nums[j], nums[i]
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func majorityElement(nums []int) int {
	var num, count int
	for i := range nums {
		if count == 0 {
			num = nums[i]
		}
		count += flag(num == nums[i])
	}
	return num
}

func flag(e bool) int {
	if e {
		return 1
	}
	return -1
}

// maxProfit 买股票的最佳时机 超时！！！
func maxProfit2(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			n := prices[j] - prices[i]
			if n > max {
				max = n
			}
		}
	}
	return max
}

// maxProfit 买股票的最佳时机
func maxProfit(prices []int) int {
	min := math.MaxInt
	var max int
	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > max {
			max = p - min
		}
	}
	return max
}
func romanToInt(s string) int {
	nums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var count int
	for i := range s {
		n := nums[s[i]]
		if i < len(s)-1 && n < nums[s[i+1]] {
			count -= n
		} else {
			count += n
		}
	}
	return count
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		fmt.Println(1)
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				fmt.Println(2)
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

//pattern = "abba", s = "dog cat cat dog"
//pattern = "abba", s = "dog cat cat fish"

func wordPattern(pattern string, s string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}

	for i := range pattern {
		p, ok1 := p2s[pattern[i]]
		s, ok2 := s2p[strs[i]]
		if ok1 && p != strs[i] {
			return false
		}
		if ok2 && s != pattern[i] {
			return false
		}
		p2s[pattern[i]] = strs[i]
		s2p[strs[i]] = pattern[i]
	}
	return true
}

func getNum(num int) int {
	var n int
	for num > 0 {
		n += num % 10 * num % 10
		num /= 10
	}
	return n
}
func isHappy(n int) bool {
	nums := make(map[int]bool)
	for ; n != 1 && !nums[n]; n, nums[n] = getNum(n), true {

	}
	return n == 1
}

// nums = [1,2,3,1], k = 3
// nums = [1,0,1,1], k = 1
// containsNearbyDuplicate 重复元素
func containsNearbyDuplicate(nums []int, k int) bool {
	indexs := make(map[int]int)
	for i, num := range nums {

		if _, ok := indexs[num]; ok {
			return true
		}
		if len(indexs) >= k {
			delete(indexs, nums[i-k])
		}
		indexs[num] = i
	}
	return false
}

/*
输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]

输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
*/
// summaryRanges 汇总区间
func summaryRanges(nums []int) []string {
	var strs []string
	for i := 0; i < len(nums); {
		left := i
		for i++; i < len(nums) && nums[i] == nums[i-1]+1; i++ {

		}
		str := fmt.Sprintf("%d", nums[left])
		if left < i-1 {
			// 证明跨元素了
			str += fmt.Sprintf("->%d", nums[i-1])
		}
		strs = append(strs, str)
	}
	return strs
}

func plusOne2(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]%10 == 0 {
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

func thirdMax(nums []int) int {
	first, second, third := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if num > first {
			first, second, third = num, first, second
		} else if num > second && num < first {
			third = second
			second = num
		} else if num > third && num < second {
			third = num
		}
	}

	if third == math.MinInt {
		return first
	}
	return third
}

// plusOne 加一
func plusOne(digits []int) []int {

	count := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += count
		if digits[i] >= 10 {
			digits[i] = 0
			continue
		}
		count = 0
	}
	if count == 1 {
		digits = append([]int{count}, digits...)
	}
	return digits
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			result = append(result, x)
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return result
}

// distributeCandies 分糖果
func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	var count int
	n := len(candyType) / 2
	for i := 0; i < len(candyType); i++ {
		if (i == 0 || candyType[i] != candyType[i-1]) && count < n {
			count++
		}
	}
	return count
}

// canPlaceFlowers 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	for i, flower := range flowerbed {
		if flower == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func findLengthOfLCIS(nums []int) int {
	var (
		max   int
		start int
	)
	for i := range nums {
		if i > 0 && nums[i-1] >= nums[i] {
			start = i
		}
		if max < i-start+1 {
			max = i - start + 1
		}
	}
	return max
}

func sortedSquares(nums []int) []int {
	negative := -1
	for i := range nums {
		if nums[i] < 0 {
			negative++
		}
	}

	positive := negative + 1
	var result []int
	for i, j := negative, positive; i >= 0 || j < len(nums); {
		if i < 0 {
			// 全是正数
			result = append(result, nums[j]*nums[j])
			j++
		} else if j == len(nums) {
			// 全是负数
			result = append(result, nums[i]*nums[i])
			i--
		} else if nums[j]*nums[j] > nums[i]*nums[i] {
			result = append(result, nums[i]*nums[i])
			i--
		} else {
			result = append(result, nums[j]*nums[j])
			j++
		}
	}
	return result
}

func addToArrayForm(num []int, k int) []int {

	var result []int
	for i := len(num) - 1; i >= 0; i-- {
		count := k % 10
		sum := num[i] + count
		k /= 10
		if sum > 10 {
			k++
			sum -= 10
		}
		result = append(result, sum)
	}
	for ; k > 0; k /= 10 {
		result = append(result, k%10)
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

func commonChars(words []string) []string {
	chars := [26]int{}
	for i := range chars {
		chars[i] = math.MaxInt
	}

	for _, word := range words {
		tmp := [26]int{}
		for _, w := range word {
			tmp[w-'a']++
		}
		for i, t := range tmp {
			if t < chars[i] {
				chars[i] = t
			}
		}
	}

	var result []string
	for char, num := range chars {
		for i := 0; i < num; i++ {
			result = append(result, string(char+'a'))
		}
	}
	return result
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arrs := [1001]int{}
	for _, arr := range arr1 {
		arrs[arr]++
	}

	fmt.Println(arrs)

	var result []int

	for _, arr := range arr2 {
		for arrs[arr] > 0 {
			fmt.Println(arrs[arr])
			result = append(result, arr)
			arrs[arr]--
		}
		//for i := 0; i < arrs[arr]; i++ {
		//}
		//arrs[arr] = 0
	}

	for i := range arrs {
		for j := 0; j < arrs[i]; j++ {
			result = append(result, i)
		}
		arrs[i] = 0
	}
	return result
}

func replaceElements(arr []int) []int {
	for i := range arr {

		max := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	return arr
}

func sumZero(n int) []int {
	var result []int
	var count int
	for i := 1; i < n; i++ {
		result = append(result, i-1)
		count += i - 1
	}
	return append(result, -count)
}

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	i, j := 0, len(arr)-1
	for i < j {
		cal := arr[i] * 2
		if cal == arr[j] {
			return false
		} else if cal < arr[j] {
			j--
		} else {
			i++
		}
	}
	return false
}
