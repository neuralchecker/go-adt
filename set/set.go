package set

import "github.com/neuralchecker/go-adt/iterable"

type Set[T any] interface {
	// Add adds an element to the set.
	Add(element T)
	// Clear removes all elements from the set.
	Clear()
	// Contains returns true if the set contains the given element.
	Contains(element T) bool
	// IsEmpty returns true if the set is empty.
	IsEmpty() bool
	// Remove removes the given element from the set.
	Remove(element T)
	// Size returns the number of elements in the set.
	Size() int
	ToSlice() []T
	Iterator() iterable.Iterator[T]
}
