package dictionary

import "github.com/neuralchecker/go-adt/iterable"

// Dictionary is a map for use when you don't have a natively hashable type
type Dictionary[K any, V any] interface {
	// Get returns the value associated with the given key.
	Get(key K) (V, bool)
	// Set sets the value associated with the given key.
	Set(key K, value V)
	// Remove removes the value associated with the given key.
	Remove(key K)
	// Keys returns the keys of the map.
	Keys() []K
	// Values returns the values of the map.
	Values() []V
	// Size returns the number of elements in the map.
	Size() int
	// IsEmpty returns true if the map is empty.
	IsEmpty() bool
	// Clear removes all elements from the map.
	Clear()

	Iterator() iterable.Iterator[pair[K, V]]
}
