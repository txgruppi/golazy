package golazy_test

import (
	"golazy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyIterator(t *testing.T) {
	iterator := golazy.FromEmpty[float64]()
	assert.False(t, iterator.Next())
}
