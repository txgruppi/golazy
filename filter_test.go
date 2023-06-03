package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterIterator(t *testing.T) {
	t.Run("empty source", func(t *testing.T) {
		iterator := golazy.FromEmpty[int]()
		iterator = golazy.Filter(iterator, func(v int) bool { return v > 0 })
		assert.False(t, iterator.Next())
	})

	t.Run("no match", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Filter(iterator, func(v int) bool { return v > 3 })
		assert.False(t, iterator.Next())
	})

	t.Run("one match", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Filter(iterator, func(v int) bool { return v == 2 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("many matches", func(t *testing.T) {
		iterator := golazy.FromRange(0, 5, 1)
		iterator = golazy.Filter(iterator, func(v int) bool { return v > 1 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 3, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 4, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("multiple calls to .Value() return the same result", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Filter(iterator, func(v int) bool { return v == 2 })
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.Equal(t, 2, iterator.Value())
		assert.Equal(t, 2, iterator.Value())
	})
}
