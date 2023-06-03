package golazy

type EmptyIterator[V any] struct{}

func (t *EmptyIterator[V]) Next() bool {
	return false
}

func (t *EmptyIterator[V]) Value() (value V) {
	return
}

func FromEmpty[V any]() Iterator[V] {
	return &EmptyIterator[V]{}
}
