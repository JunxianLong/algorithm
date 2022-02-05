package array

// majorityElement2 多数元素 一定存在结果的，没必要判断众数了
func majorityElement2(nums []int) int {
	var count int
	var tmp int
	for _, num := range nums {
		if count == 0 {
			tmp = num
		}
		count += flag(tmp == num, 1, -1)
	}
	return tmp
}

func flag(f bool, ret1, ret2 int) int {
	if f {
		return ret1
	}
	return ret2
}

// URL:https://leetcode-cn.com/problems/kth-distinct-string-in-an-array/
// Time:20220205
// KthDistinct 数组中第 K 个独一无二的字符串
func KthDistinct(arr []string, k int) string {

	numFreq := make(map[string]int)
	for _, str := range arr {
		numFreq[str]++
	}

	for _, str := range arr {
		if numFreq[str] != 1 {
			continue
		}
		k--
		if k == 0 {
			return str
		}
	}
	return ""
}

// URL:https://leetcode-cn.com/problems/finding-3-digit-even-numbers/
// Time:20220205
// FindEvenNumbers 找出3位偶数
func FindEvenNumbers(digits []int) []int {
	nums := [10]int{}
	for _, digit := range digits {
		nums[digit]++
	}
	result := make([]int, 0)
	for i := 100; i <= 998; i += 2 {
		tmp := i
		tmpNums := [10]int{}
		flag := true
		for tmp > 0 {
			c := tmp % 10
			tmpNums[c]++
			if tmpNums[c] > nums[c] {
				flag = false
				break
			}
			tmp /= 10
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}

// URL:https://leetcode-cn.com/problems/count-common-words-with-one-occurrence/
// Time:20220205
// CountWords 统计出现过一次的公共字符串（龟速。。。）
func CountWords(words1 []string, words2 []string) int {
	word1Freq := make(map[string]int)
	var count int

	for _, word := range words1 {
		word1Freq[word]++
	}

	for _, word := range words2 {
		freq, ok := word1Freq[word]
		if !ok {
			continue
		}
		if freq > 1 {
			delete(word1Freq, word)
		} else {
			word1Freq[word]--
		}
	}

	for _, freq := range word1Freq {
		if freq == 0 {
			count++
		}
	}
	return count
}
