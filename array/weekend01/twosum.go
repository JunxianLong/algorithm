package weekend01

// url:https://leetcode.cn/problems/two-sum/
// twoSum 两数之和 输入无序数组
func twoSum(nums []int, target int) []int {
	numInfo := make(map[int]int)
	for i, num := range nums {
		if j, ok := numInfo[target-num]; ok {
			return []int{i, j}
		}
		numInfo[num] = i
	}
	return nil
}
