package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	var vals []int

	t.Run("basic", func(t *testing.T) {
		vals = []int{3, 5, 6, 3, 1, 4}
		QuickSort(vals)
		for i := 0; i < len(vals)-1; i++ {
			assert.True(t, vals[i] <= vals[i+1])
		}
	})

	t.Run("empty", func(t *testing.T) {
		vals = []int{}
		QuickSort(vals)
		for i := 0; i < len(vals)-1; i++ {
			assert.True(t, vals[i] <= vals[i+1])
		}
	})

	t.Run("reversed", func(t *testing.T) {
		vals = []int{6, 5, 4, 3, 2, 1}
		QuickSort(vals)
		for i := 0; i < len(vals)-1; i++ {
			assert.True(t, vals[i] <= vals[i+1])
		}
	})

	t.Run("sorted", func(t *testing.T) {
		vals = []int{1, 2, 3, 4, 5, 6}
		QuickSort(vals)
		for i := 0; i < len(vals)-1; i++ {
			assert.True(t, vals[i] <= vals[i+1])
		}
	})
}
