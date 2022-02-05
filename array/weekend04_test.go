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
