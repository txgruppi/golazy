package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{})
		assert.Equal(t, []float64{}, golazy.ToArray(iterator))
	})

	t.Run("one", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{1.0})
		assert.Equal(t, []float64{1.0}, golazy.ToArray(iterator))
	})

	t.Run("many", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{1.0, 2.0, 3.0})
		assert.Equal(t, []float64{1.0, 2.0, 3.0}, golazy.ToArray(iterator))
	})
}

func TestSum(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		iterator := golazy.FromRange(0, 0, 0)
		assert.Equal(t, 0, golazy.Sum(iterator))
	})

	t.Run("one", func(t *testing.T) {
		iterator := golazy.FromRange(1, 2, 1)
		assert.Equal(t, 1, golazy.Sum(iterator))
	})

	t.Run("many", func(t *testing.T) {
		iterator := golazy.FromRange(1, 4, 1)
		assert.Equal(t, 6, golazy.Sum(iterator))
	})
}
