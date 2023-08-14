package weekend01

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	target := 0
	result := searchInsert(nums, target)
	t.Log(result)

}
