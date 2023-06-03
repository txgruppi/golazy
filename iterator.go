package golazy

type Iterator[V any] interface {
	Next() bool
	Value() V
}
