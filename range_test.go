package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeIterator(t *testing.T) {
	t.Run("ascending", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 1, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("descending", func(t *testing.T) {
		iterator := golazy.FromRange(5, 0, -2)
		assert.True(t, iterator.Next())
		assert.Equal(t, 5, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 3, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 1, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("empty if step is zero", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 0)
		assert.False(t, iterator.Next())
	})

	t.Run("empty if ascending and start >= end", func(t *testing.T) {
		iterator := golazy.FromRange(3, 0, 1)
		assert.False(t, iterator.Next())
	})

	t.Run("empty if descending and start <= end", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, -1)
		assert.False(t, iterator.Next())
	})

	t.Run("multiple calls to .Value() return the same result", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 0, iterator.Value())
	})
}
