package main

import "fmt"

func main() {
	fmt.Println(3 >> 1)
	fmt.Println("Hello,World!!!")
}

// containsNearbyDuplicate
func containsNearbyDuplicate(nums []int, k int) bool {
	numsInfo := make(map[int]struct{})
	for i := range nums {
		if _, ok := numsInfo[nums[i]]; ok {
			return true
		}
		numsInfo[nums[i]] = struct{}{}
		if len(numsInfo) > k {
			delete(numsInfo, nums[i-k])
		}
	}
	return false
}
