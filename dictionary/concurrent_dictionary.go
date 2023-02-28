package dictionary

import (
	"sync"

	"github.com/neuralchecker/go-adt/iterable"
)

type concurrenthDictionary[K Hashable, V any] struct {
	Dictionary[K, V]
	lock *sync.RWMutex
}

func NewConcurrentUnordered[K Hashable, V any]() Dictionary[K, V] {
	return &concurrenthDictionary[K, V]{
		Dictionary: NewUnordered[K, V](),
		lock:       &sync.RWMutex{},
	}
}

func NewConcurrentUnorderedSize[K Hashable, V any](size int) Dictionary[K, V] {
	return &concurrenthDictionary[K, V]{
		Dictionary: NewUnorderedSize[K, V](size),
		lock:       &sync.RWMutex{},
	}
}

func (d *concurrenthDictionary[K, V]) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.Dictionary.Clear()
}

func (d *concurrenthDictionary[K, V]) Get(key K) (value V, ok bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.Get(key)
}

func (d *concurrenthDictionary[K, V]) IsEmpty() bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.IsEmpty()
}

func (d *concurrenthDictionary[K, V]) Keys() []K {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.Keys()
}

func (d *concurrenthDictionary[K, V]) Remove(key K) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.Dictionary.Remove(key)
}

func (d *concurrenthDictionary[K, V]) Set(key K, value V) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.Dictionary.Set(key, value)
}

func (d *concurrenthDictionary[K, V]) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.Size()
}

func (d *concurrenthDictionary[K, V]) Values() []V {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.Values()
}

func (d *concurrenthDictionary[K, V]) Iterator() iterable.Iterator[pair[K, V]] {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.Dictionary.Iterator()
}
