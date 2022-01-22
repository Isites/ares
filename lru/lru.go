package lru

import (
	"math/rand"
	"sync"
)

type lruCacheItem struct {
	prev, next *lruCacheItem
	data       interface{}
	key        string
}

type lruc struct {
	head, tail *lruCacheItem
	lruMap     map[string]*lruCacheItem
	rw         sync.RWMutex
	size       int64
}

func New(size int64) *lruc {
	if size < 0 {
		size = 100
	}
	lru := &lruc{
		head:   new(lruCacheItem),
		tail:   new(lruCacheItem),
		lruMap: make(map[string]*lruCacheItem),
		size:   size,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *lruc) Get(key string) (interface{}, bool) {
	lru.rw.RLock()
	v, ok := lru.lruMap[key]
	lru.rw.RUnlock()
	if ok {
		// 当缓存数量达到一定程度时才将访问的数据移至队头
		// 为了尽可能减少锁竞争，尝试引入随机数：访问越频繁被移动到队头的概率越大
		if len(lru.lruMap) > int(lru.size)-1 && rand.Int()%100 == 1 {
			lru.rw.Lock()
			v.prev.next = v.next
			v.next.prev = v.prev

			v.prev = lru.head
			v.next = lru.head.next
			lru.head.next = v
			lru.rw.Unlock()
		}
		return v.data, true
	}
	return nil, false
}

func (lru *lruc) Set(key string, v interface{}) {
	// fast path
	if _, exist := lru.lruMap[key]; exist {
		return
	}
	node := &lruCacheItem{
		data: v,
		prev: lru.head,
		next: lru.head.next,
		key:  key,
	}
	// add first
	lru.rw.Lock()
	// double check
	if _, exist := lru.lruMap[key]; !exist {
		lru.lruMap[key] = node
		lru.head.next = node
		node.next.prev = node
	}
	if len(lru.lruMap) > int(lru.size) {
		// delete tail
		prev := lru.tail.prev
		prev.prev.next = lru.tail
		lru.tail.prev = prev.prev
		delete(lru.lruMap, prev.key)
	}
	lru.rw.Unlock()
}
