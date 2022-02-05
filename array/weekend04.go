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
