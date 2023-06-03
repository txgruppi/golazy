package golazy

type TakeIterator[V any] struct {
	iterator Iterator[V]
	n        int
}

func (t *TakeIterator[V]) Next() bool {
	if t.n == 0 {
		return false
	}
	t.n--
	return t.iterator.Next()
}

func (t *TakeIterator[V]) Value() V {
	return t.iterator.Value()
}

func Take[V any](iterator Iterator[V], n int) *TakeIterator[V] {
	return &TakeIterator[V]{iterator: iterator, n: n}
}
