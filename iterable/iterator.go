package iterable

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type sliceIterator[T any] struct {
	slice []T
	index int
}

func (it *sliceIterator[T]) HasNext() bool {
	return it.index < len(it.slice)
}

func (it *sliceIterator[T]) Next() T {
	if !it.HasNext() {
		panic("no more elements")
	}
	defer func() { it.index++ }()
	return it.slice[it.index]
}

func SliceIterator[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{slice: slice}
}
