package main

func main() {
	lru := Constructor(2)
	lru.Get(2)
	lru.Put(2, 6)
	lru.Get(1)
	lru.Put(1, 5)
	lru.Put(1, 2)
	lru.Get(1)
	lru.Get(2)

}

type LRUCache struct {
	capacity int
	time     int
	cache    map[int]timedValue
}

type timedValue struct {
	Val  int
	Time int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		0,
		make(map[int]timedValue),
	}
}

func (this *LRUCache) Get(key int) int {
	this.time++
	val, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.cache[key] = timedValue{val.Val, this.time}
	return val.Val
}

func (this *LRUCache) Put(key int, value int) {
	this.time++
	_, ok := this.cache[key]

	if !ok && len(this.cache) >= this.capacity {
		minTime := this.time
		minKey := -1
		for i, e := range this.cache {
			if e.Time < minTime {
				minTime = e.Time
				minKey = i
			}
		}
		delete(this.cache, minKey)
	}

	this.cache[key] = timedValue{value, this.time}
}
