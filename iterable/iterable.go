package iterable

type Iterable[T any] interface {
	Iterator() Iterator[T]
}
