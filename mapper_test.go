package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapperIterator(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		iterator := golazy.FromEmpty[int]()
		iterator = golazy.Map(iterator, func(v int) int { return v * 2 })
		assert.False(t, iterator.Next())
	})

	t.Run("one", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Map(iterator, func(v int) int { return v * 2 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("many", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Map(iterator, func(v int) int { return v * 2 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 4, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("multiple calls to .Value() return the same result", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Map(iterator, func(v int) int { return v * 2 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 0, iterator.Value())
	})

	t.Run("multiple calls to .Value() only call the mapper once", func(t *testing.T) {
		calls := 0
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Map(iterator, func(v int) int {
			calls++
			return v * 2
		})
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 1, calls)
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 1, calls)
		assert.Equal(t, 0, iterator.Value())
		assert.Equal(t, 1, calls)
	})
}
