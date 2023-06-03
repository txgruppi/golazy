package golazy

type ArrayIterator[V any] struct {
	array []V
	index int
}

func (t *ArrayIterator[V]) Next() bool {
	t.index++
	return t.index < len(t.array)
}

func (t *ArrayIterator[V]) Value() (value V) {
	if t.index < len(t.array) {
		value = t.array[t.index]
	}
	return
}

func FromArray[V any](array []V) Iterator[V] {
	return &ArrayIterator[V]{array: array, index: -1}
}
