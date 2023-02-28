package dictionary

import (
	"github.com/google/go-cmp/cmp"
	"github.com/neuralchecker/go-adt/iterable"
)

type Hashable interface {
	Hash() int
}

const lambda float64 = 0.75

type pair[K any, V any] struct {
	Key   K
	Value V
}

func (p pair[K, V]) Unwrap() (K, V) {
	return p.Key, p.Value
}

type hashDictionary[K Hashable, V any] struct {
	arr      [][]pair[K, V]
	elements int
}

func NewUnordered[K Hashable, V any]() Dictionary[K, V] {
	return &hashDictionary[K, V]{
		arr: make([][]pair[K, V], 49),
	}
}

func NewUnorderedSize[K Hashable, V any](size int) Dictionary[K, V] {
	size = size*2 - 1
	if size < 49 {
		size = 49
	}
	return &hashDictionary[K, V]{
		arr: make([][]pair[K, V], size),
	}
}

// Clear implements HashMap
func (m *hashDictionary[K, V]) Clear() {
	m.arr = make([][]pair[K, V], 49)
	m.elements = 0
}

// Get implements HashMap
func (m *hashDictionary[K, V]) Get(key K) (value V, ok bool) {
	index := key.Hash() % len(m.arr)
	for _, p := range m.arr[index] {
		if cmp.Equal(p.Key, key) {
			return p.Value, true
		}
	}
	return value, false
}

// IsEmpty implements HashMap
func (m *hashDictionary[K, V]) IsEmpty() bool {
	return m.elements == 0
}

// Keys implements HashMap
func (m *hashDictionary[K, V]) Keys() []K {
	keys := make([]K, 0, m.elements)
	for _, bucket := range m.arr {
		for _, p := range bucket {
			keys = append(keys, p.Key)
		}
	}
	return keys
}

// Remove implements HashMap
func (m *hashDictionary[K, V]) Remove(key K) {
	index := key.Hash() % len(m.arr)
	for i, p := range m.arr[index] {
		if cmp.Equal(p.Key, key) {
			m.arr[index] = append(m.arr[index][:i], m.arr[index][i+1:]...)
			m.elements--
			return
		}
	}
}

// Set implements HashMap
func (m *hashDictionary[K, V]) Set(key K, value V) {
	index := key.Hash() % len(m.arr)
	for i, p := range m.arr[index] {
		if cmp.Equal(p.Key, key) {
			m.arr[index][i].Value = value
			return
		}
	}
	m.arr[index] = append(m.arr[index], pair[K, V]{Key: key, Value: value})
	m.elements++

	if float64(m.elements)/float64(len(m.arr)) > lambda {
		m.rehash()
	}
}

// Size implements HashMap
func (m *hashDictionary[K, V]) Size() int {
	return m.elements
}

// Values implements HashMap
func (m *hashDictionary[K, V]) Values() []V {
	values := make([]V, 0, m.elements)
	for _, bucket := range m.arr {
		for _, p := range bucket {
			values = append(values, p.Value)
		}
	}
	return values
}

func (m *hashDictionary[K, V]) rehash() {
	oldIt := m.Iterator()
	m.arr = make([][]pair[K, V], len(m.arr)*2-1)
	m.elements = 0
	for oldIt.HasNext() {
		p := oldIt.Next()
		m.Set(p.Key, p.Value)
	}
}

func (m *hashDictionary[K, V]) Iterator() iterable.Iterator[pair[K, V]] {
	pairs := make([]pair[K, V], 0, m.elements)
	for _, bucket := range m.arr {
		pairs = append(pairs, bucket...)
	}
	return iterable.SliceIterator(pairs)
}
