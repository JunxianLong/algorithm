package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthDistinct(t *testing.T) {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	k := 2
	assert.Equal(t, KthDistinct(arr, k), "a")
}

func TestFindEvenNumbers(t *testing.T) {
	digits := []int{2, 1, 3, 0}
	result := []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}
	assert.Equal(t, FindEvenNumbers(digits), result)
}

func TestCountWords(t *testing.T) {
	words1 := []string{"a", "ab"}
	words2 := []string{"a", "a", "a", "ab"}
	assert.Equal(t, CountWords(words1, words2), 1)
}
