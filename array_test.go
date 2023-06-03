package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayIterator(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{})
		assert.False(t, iterator.Next())
	})

	t.Run("one", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{1.0})
		assert.True(t, iterator.Next())
		assert.Equal(t, 1.0, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("many", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{1.0, 2.0, 3.0})
		assert.True(t, iterator.Next())
		assert.Equal(t, 1.0, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 2.0, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 3.0, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("multiple calls to .Value() return the same result", func(t *testing.T) {
		iterator := golazy.FromArray([]float64{1.0})
		assert.True(t, iterator.Next())
		assert.Equal(t, 1.0, iterator.Value())
		assert.Equal(t, 1.0, iterator.Value())
		assert.Equal(t, 1.0, iterator.Value())
	})
}
