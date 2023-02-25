package hashset

import (
	"github.com/google/go-cmp/cmp"
	"github.com/neuralchecker/go-adt/iterable"
)

type Hashable interface {
	Hash() int
}

type hashSet[T Hashable] struct {
	arr      [][]T
	elements int
}

type HashSet[T Hashable] interface {
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
	Iterator() iterable.Iterator[T]
}

func NewHashSet[T Hashable]() HashSet[T] {
	return &hashSet[T]{
		arr: make([][]T, 0, 49),
	}
}

func NewHashSetSize[T Hashable](size int) HashSet[T] {
	return &hashSet[T]{
		arr: make([][]T, 0, size*2-1),
	}
}

// Add implements HashSet
func (s *hashSet[T]) Add(element T) {
	index := element.Hash() % len(s.arr)
	for _, e := range s.arr[index] {
		if cmp.Equal(e, element) {
			return
		}
	}
	s.arr[index] = append(s.arr[index], element)
	s.elements++
}

// Clear implements HashSet
func (s *hashSet[T]) Clear() {
	s.arr = make([][]T, 0, 49)
	s.elements = 0
}

// Contains implements HashSet
func (s *hashSet[T]) Contains(element T) bool {
	index := element.Hash() % len(s.arr)
	for _, e := range s.arr[index] {
		if cmp.Equal(e, element) {
			return true
		}
	}
	return false
}

// IsEmpty implements HashSet
func (s *hashSet[T]) IsEmpty() bool {
	return s.elements == 0
}

// Remove implements HashSet
func (s *hashSet[T]) Remove(element T) {
	index := element.Hash() % len(s.arr)
	for i, e := range s.arr[index] {
		if cmp.Equal(e, element) {
			s.arr[index] = append(s.arr[index][:i], s.arr[index][i+1:]...)
			s.elements--
			return
		}
	}
}

// Size implements HashSet
func (s *hashSet[T]) Size() int {
	return s.elements
}

// Iterator implements HashSet
func (s *hashSet[T]) Iterator() iterable.Iterator[T] {
	ts := make([]T, 0, s.elements)
	for _, arr := range s.arr {
		ts = append(ts, arr...)
	}
	return iterable.SliceIterator(ts)
}
