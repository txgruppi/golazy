package golazy

import "golang.org/x/exp/constraints"

func ToArray[V any](i Iterator[V]) []V {
	result := []V{}
	for i.Next() {
		result = append(result, i.Value())
	}
	return result
}

func Sum[V constraints.Integer | constraints.Float](i Iterator[V]) V {
	var result V
	for i.Next() {
		result += i.Value()
	}
	return result
}
