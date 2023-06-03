package golazy

type FilterIterator[V any] struct {
	iterator  Iterator[V]
	predicate func(V) bool
}

func (t *FilterIterator[V]) Next() bool {
	for t.iterator.Next() {
		if t.predicate(t.iterator.Value()) {
			return true
		}
	}
	return false
}

func (t *FilterIterator[V]) Value() V {
	return t.iterator.Value()
}

func Filter[V any](iterator Iterator[V], predicate func(V) bool) *FilterIterator[V] {
	return &FilterIterator[V]{iterator: iterator, predicate: predicate}
}
