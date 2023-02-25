package hashmap

import (
	"github.com/google/go-cmp/cmp"
	"github.com/neuralchecker/go-adt/iterable"
)

type Hashable interface {
	Hash() int
}

const lambda float64 = 0.75

// HashMap is a map for use when you don't have a natively hashable type
type HashMap[K Hashable, V any] interface {
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

type pair[K Hashable, V any] struct {
	key   K
	value V
}

type hMap[K Hashable, V any] struct {
	arr      [][]pair[K, V]
	elements int
}

func NewHashMap[K Hashable, V any]() HashMap[K, V] {
	return &hMap[K, V]{
		arr: make([][]pair[K, V], 0, 49),
	}
}

func NewHashMapSize[K Hashable, V any](size int) HashMap[K, V] {
	return &hMap[K, V]{
		arr: make([][]pair[K, V], 0, size*2-1),
	}
}

// Clear implements HashMap
func (m *hMap[K, V]) Clear() {
	m.arr = make([][]pair[K, V], 0, 49)
	m.elements = 0
}

// Get implements HashMap
func (m *hMap[K, V]) Get(key K) (value V, ok bool) {
	index := key.Hash() % len(m.arr)
	for _, p := range m.arr[index] {
		if cmp.Equal(p.key, key) {
			return p.value, true
		}
	}
	return value, false
}

// IsEmpty implements HashMap
func (m *hMap[K, V]) IsEmpty() bool {
	return m.elements == 0
}

// Keys implements HashMap
func (m *hMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.elements)
	for _, bucket := range m.arr {
		for _, p := range bucket {
			keys = append(keys, p.key)
		}
	}
	return keys
}

// Remove implements HashMap
func (m *hMap[K, V]) Remove(key K) {
	index := key.Hash() % len(m.arr)
	for i, p := range m.arr[index] {
		if cmp.Equal(p.key, key) {
			m.arr[index] = append(m.arr[index][:i], m.arr[index][i+1:]...)
			m.elements--
			return
		}
	}
}

// Set implements HashMap
func (m *hMap[K, V]) Set(key K, value V) {
	index := key.Hash() % len(m.arr)
	for i, p := range m.arr[index] {
		if cmp.Equal(p.key, key) {
			m.arr[index][i].value = value
			return
		}
	}
	m.arr[index] = append(m.arr[index], pair[K, V]{key: key, value: value})
	m.elements++

	if float64(m.elements)/float64(cap(m.arr)) > lambda {
		m.rehash()
	}
}

// Size implements HashMap
func (m *hMap[K, V]) Size() int {
	return m.elements
}

// Values implements HashMap
func (m *hMap[K, V]) Values() []V {
	values := make([]V, 0, m.elements)
	for _, bucket := range m.arr {
		for _, p := range bucket {
			values = append(values, p.value)
		}
	}
	return values
}

func (m *hMap[K, V]) rehash() {
	oldIt := m.Iterator()
	m.arr = make([][]pair[K, V], 0, cap(m.arr)*2-1)
	m.elements = 0
	for oldIt.HasNext() {
		p := oldIt.Next()
		m.Set(p.key, p.value)
	}
}

func (m *hMap[K, V]) Iterator() iterable.Iterator[pair[K, V]] {
	pairs := make([]pair[K, V], 0, m.elements)
	for _, bucket := range m.arr {
		pairs = append(pairs, bucket...)
	}
	return iterable.SliceIterator(pairs)
}
