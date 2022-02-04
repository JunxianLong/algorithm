package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	result := RelativeSortArray(arr1, arr2)
	t.Log(result)
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	nums := []int{8, 1, 2, 2, 3}
	result := SmallerNumbersThanCurrent(nums)
	t.Log(result)
}

func TestGame(t *testing.T) {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	assert.Equal(t, 3, Game(guess, answer))
}

func TestPairSumArray(t *testing.T) {
	nums := []int{1, 4, 3, 2}
	assert.Equal(t, 4, PairSumArray(nums))
}
