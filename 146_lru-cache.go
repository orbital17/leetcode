package leetcode

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

type LRUCache struct {
	cap   int
	size  int
	hash  map[int]*LRUNode
	start *LRUNode
	end   *LRUNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{cap: capacity}
	cache.hash = make(map[int]*LRUNode)
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hash[key]
	if !ok {
		return -1
	}
	if node == this.start {
		return node.val
	}
	if node == this.end {
		this.end = node.prev
		this.end.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = this.start
	this.start.prev = node
	this.start = node
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) > 0 {
		node, _ := this.hash[key]
		node.val = value
		return
	}
	node := &LRUNode{
		val:  value,
		key:  key,
		next: this.start,
		prev: nil,
	}
	if this.start == nil {
		this.end = node
		this.hash[key] = node
	} else {
		this.start.prev = node
	}
	this.start = node
	this.hash[key] = node
	this.size++
	if this.size > this.cap {
		this.size--
		delete(this.hash, this.end.key)
		this.end.prev.next = nil
		this.end = this.end.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
