package golazy

import "golang.org/x/exp/constraints"

type RangeIterator[V constraints.Integer | constraints.Float] struct {
	start V
	end   V
	step  V
	curr  *V
}

func (t *RangeIterator[V]) Next() bool {
	if t.curr == nil {
		t.curr = &t.start
	} else {
		*t.curr += t.step
	}
	return (t.step > 0 && *t.curr < t.end) || (t.step < 0 && *t.curr > t.end)
}

func (t *RangeIterator[V]) Value() (value V) {
	value = t.start
	return
}

func FromRange[V constraints.Integer | constraints.Float](start V, end V, step V) Iterator[V] {
	return &RangeIterator[V]{start: start, end: end, step: step}
}
