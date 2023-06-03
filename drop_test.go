package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropIterator(t *testing.T) {
	t.Run("zero from nothing", func(t *testing.T) {
		iterator := golazy.FromEmpty[int]()
		iterator = golazy.Drop(iterator, 0)
		assert.False(t, iterator.Next())
	})

	t.Run("zero from one", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Drop(iterator, 0)
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("zero from many", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Drop(iterator, 0)
		assert.True(t, iterator.Next())
		assert.Equal(t, 0, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 1, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("one from nothing", func(t *testing.T) {
		iterator := golazy.FromEmpty[int]()
		iterator = golazy.Drop(iterator, 1)
		assert.False(t, iterator.Next())
	})

	t.Run("one from one", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Drop(iterator, 1)
		assert.False(t, iterator.Next())
	})

	t.Run("one from many", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Drop(iterator, 1)
		assert.True(t, iterator.Next())
		assert.Equal(t, 1, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 2, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("many from nothing", func(t *testing.T) {
		iterator := golazy.FromEmpty[int]()
		iterator = golazy.Drop(iterator, 3)
		assert.False(t, iterator.Next())
	})

	t.Run("many from one", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Drop(iterator, 3)
		assert.False(t, iterator.Next())
	})

	t.Run("all from many", func(t *testing.T) {
		iterator := golazy.FromRange(0, 3, 1)
		iterator = golazy.Drop(iterator, 3)
		assert.False(t, iterator.Next())
	})

	t.Run("many from many", func(t *testing.T) {
		iterator := golazy.FromRange(0, 6, 1)
		iterator = golazy.Drop(iterator, 3)
		assert.True(t, iterator.Next())
		assert.Equal(t, 3, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 4, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Equal(t, 5, iterator.Value())
		assert.False(t, iterator.Next())
	})

	t.Run("multiple calls to .Value() return the same result", func(t *testing.T) {
		iterator := golazy.FromRange(0, 1, 1)
		iterator = golazy.Drop(iterator, 1)
		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
		assert.False(t, iterator.Next())
	})
}
