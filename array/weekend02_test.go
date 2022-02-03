package array

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	t.Log(nums)
}

func TestFindKthPositive2(t *testing.T) {
	nums := []int{2, 3, 4, 7, 11}
	result := FindKthPositive(nums, 5)
	t.Log(result)
}
