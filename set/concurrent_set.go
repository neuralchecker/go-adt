package set

import (
	"sync"

	"github.com/neuralchecker/go-adt/iterable"
)

type concurrentSet[T Hashable] struct {
	Set[T]
	lock *sync.RWMutex
}

func NewConcurrentUnordered[T Hashable]() Set[T] {
	return &concurrentSet[T]{
		Set:  NewUnordered[T](),
		lock: &sync.RWMutex{},
	}
}

func NewConcurrentUnorderedSize[T Hashable](size int) Set[T] {
	return &concurrentSet[T]{
		Set:  NewUnorderedSize[T](size),
		lock: &sync.RWMutex{},
	}
}

func (s *concurrentSet[T]) Add(element T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Set.Add(element)
}

func (s *concurrentSet[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Set.Clear()
}

func (s *concurrentSet[T]) Contains(element T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Set.Contains(element)
}

func (s *concurrentSet[T]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Set.IsEmpty()
}

func (s *concurrentSet[T]) Remove(element T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Set.Remove(element)
}

func (s *concurrentSet[T]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Set.Size()
}

func (s *concurrentSet[T]) ToSlice() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Set.ToSlice()
}

func (s *concurrentSet[T]) Iterator() iterable.Iterator[T] {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Set.Iterator()
}
