package lru

type LRUCache struct {
	limit int
	size  int
	data  map[int]*Node
	head  *Node
	tail  *Node
}

type Node struct {
	val  int
	key  int
	pre  *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:  0,
		limit: capacity,
		data:  make(map[int]*Node),
		head:  new(Node),
		tail:  new(Node),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.data[key]
	if !ok {
		return -1
	}
	// remove
	v.pre.next = v.next
	v.next.pre = v.pre
	//  set haeader
	this.head.next.pre = v
	v.next = this.head.next
	this.head.next = v
	v.pre = this.head

	return v.val
}

func (this *LRUCache) romoveLast() {
	p := this.head
	for p.next != nil {
		p = p.next
	}
	p.next = nil
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.data[key]; ok {
		v.pre.next = v.next
		v.next.pre = v.pre
		delete(this.data, key)
		this.size--
	} else if this.size >= this.limit {
		v := this.tail.pre
		v.pre.next = v.next
		v.next.pre = v.pre
		// this.tail.next = this.tail.next.next
		delete(this.data, v.key)
		this.size--
		// fmt.Println("remove")
	}
	n := &Node{
		val: value,
		key: key,
	}
	// set header
	this.head.next.pre = n
	n.next = this.head.next
	this.head.next = n
	n.pre = this.head

	this.data[key] = n
	this.size++
	return
}
