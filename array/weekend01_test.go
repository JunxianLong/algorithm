package array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	count := RemoveDuplicates(nums)
	if count != 2 {
		t.Log("测试不通过")
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	count = RemoveDuplicates(nums)
	if count != 5 {
		t.Log("测试不通过")
	}
}
