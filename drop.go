package golazy

type DropIterator[V any] struct {
	iterator Iterator[V]
	n        int
}

func (t *DropIterator[V]) Next() bool {
	for ; t.n > 0 && t.iterator.Next(); t.n-- {
	}
	return t.iterator.Next()
}

func (t *DropIterator[V]) Value() V {
	return t.iterator.Value()
}

func Drop[V any](iterator Iterator[V], n int) *DropIterator[V] {
	return &DropIterator[V]{iterator: iterator, n: n}
}
