package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	var vals []int

	t.Run("case1", func(t *testing.T) {
		vals = []int{3, 1, 2, 4}
		InsertSort(vals)
		assert.Equal(t, []int{1, 2, 3, 4}, vals)
	})

	t.Run("case2", func(t *testing.T) {
		vals = []int{2, 11, 7, 15}
		InsertSort(vals)
		assert.Equal(t, []int{2, 7, 11, 15}, vals)
	})
}