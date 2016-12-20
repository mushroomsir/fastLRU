package fastLRU

import "sync"

//Options ...
type Options struct {
	MaxSize uint32
}

//LRUCache ...
type LRUCache struct {
	ageToDiscard uint64
	currentAge   uint64
	maxSize      uint32
	lock         sync.Mutex
	store        map[string]*trackValue
}

type trackValue struct {
	value interface{}
	age   uint64
}

//New create a LRUCache with options.
func New(opts Options) (lru *LRUCache) {
	if opts.MaxSize == 0 {
		opts.MaxSize = 1000000
	}
	lru = &LRUCache{maxSize: opts.MaxSize}
	lru.store = make(map[string]*trackValue)
	return
}

//Get one cache item by key
func (l *LRUCache) Get(key string) (val interface{}, ok bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	value, ok := l.store[key]
	if ok {
		value.age = l.currentAge
	} else {
		return
	}
	return value.value, true
}

//Add one new cache item to storage
func (l *LRUCache) Add(key string, value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.adjust()
	l.currentAge++
	val := &trackValue{
		value: value,
		age:   l.currentAge,
	}
	l.store[key] = val
}

//Remove cache item by key
func (l *LRUCache) Remove(key string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	delete(l.store, key)
}

//Count get total count of cache items
func (l *LRUCache) Count() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return len(l.store)
}

func (l *LRUCache) adjust() {
	for uint32(len(l.store)) >= l.maxSize {
		for key, value := range l.store {
			if value.age <= l.ageToDiscard {
				delete(l.store, key)
				return
			}
		}
		l.ageToDiscard = l.ageToDiscard + 1
	}
}
