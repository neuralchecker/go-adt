package set

import (
	"github.com/google/go-cmp/cmp"
	"github.com/neuralchecker/go-adt/iterable"
)

const lambda = 0.75

type Hashable interface {
	Hash() int
}

type hashSet[T Hashable] struct {
	arr      [][]T
	elements int
}

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

func NewUnordered[T Hashable]() Set[T] {
	return &hashSet[T]{
		arr: make([][]T, 49),
	}
}

func NewUnorderedSize[T Hashable](size int) Set[T] {
	size = size*2 - 1
	if size < 49 {
		size = 49
	}
	return &hashSet[T]{
		arr: make([][]T, size),
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

	if float64(s.elements)/float64(len(s.arr)) > lambda {
		s.rehash()
	}
}

// Clear implements HashSet
func (s *hashSet[T]) Clear() {
	s.arr = make([][]T, 49)
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

// ToSlice implements HashSet
func (s *hashSet[T]) ToSlice() []T {
	ts := make([]T, 0, s.elements)
	for _, arr := range s.arr {
		ts = append(ts, arr...)
	}
	return ts
}

// Iterator implements HashSet
func (s *hashSet[T]) Iterator() iterable.Iterator[T] {
	return iterable.SliceIterator(s.ToSlice())
}

func (s *hashSet[T]) rehash() {
	it := s.Iterator()
	s.arr = make([][]T, len(s.arr)*2-1)
	s.elements = 0
	for it.HasNext() {
		s.Add(it.Next())
	}
}
