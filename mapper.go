package golazy

type MapperIterator[I any, O any] struct {
	iterator Iterator[I]
	mapper   func(I) O
	curr     O
}

func (t *MapperIterator[I, O]) Next() bool {
	if t.iterator.Next() {
		t.curr = t.mapper(t.iterator.Value())
		return true
	}
	return false
}

func (t *MapperIterator[I, O]) Value() O {
	return t.curr
}

func Map[I any, O any](iterator Iterator[I], mapper func(I) O) *MapperIterator[I, O] {
	return &MapperIterator[I, O]{iterator: iterator, mapper: mapper}
}
