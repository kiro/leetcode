package leetcode

import (
	"container/heap"
)

type item struct {
	item  Item
	index int
}

type heapImpl struct {
	p               *pq
	items           map[interface{}]*item
	key             func(Item) interface{}
	DisableIndexing bool
}

func (h *heapImpl) Push(i Item) {
	it := &item{i, 0}
	h.addItem(it)
	heap.Push(h.p, it)
}

func (h *heapImpl) Pop() Item {
	res := heap.Pop(h.p).(*item)
	h.removeItem(res.item)
	return res.item
}

func (h *heapImpl) Top() Item {
	if len(h.p.items) > 0 {
		return h.p.items[0].item
	}
	return nil
}

func (h *heapImpl) Len() int {
	return len(h.p.items)
}

func (h *heapImpl) Update(o Item, n Item) {
	old := h.items[h.key(o)]
	new := &item{n, old.index}
	h.removeItem(old.item)
	h.p.items[old.index] = new
	heap.Fix(h.p, old.index)
	h.addItem(new)
}

func (h *heapImpl) Remove(i Item) {
	index := h.items[h.key(i)].index
	h.removeItem(i)
	heap.Remove(h.p, index)
}

func (h *heapImpl) Get(key interface{}) Item {
	return h.items[key].item
}

func (h *heapImpl) removeItem(item Item) {
	if !h.DisableIndexing {
		delete(h.items, h.key(item))
	}
}

func (h *heapImpl) addItem(item *item) {
	if !h.DisableIndexing {
		h.items[h.key(item.item)] = item
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type pq struct {
	items []*item
}

func (pq pq) Len() int { return len(pq.items) }

func (pq pq) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.items[i].item.Compare(pq.items[j].item)
}

func (pq pq) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *pq) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *pq) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}

func NewHeapWithKey(getKey func(Item) interface{}) Heap {
	return &heapImpl{
		&pq{
			[]*item{},
		},
		make(map[interface{}]*item),
		getKey,
		false,
	}
}

func NewHeap() Heap {
	return NewHeapWithKey(func(i Item) interface{} { return i })
}

type Item interface {
	Compare(left Item) bool
}

type Heap interface {
	Push(i Item)
	Pop() Item
	Top() Item
	Len() int
	Update(old Item, new Item)
	Remove(item Item)
	Get(key interface{}) Item
}

type MinInt int

func (i MinInt) Compare(j Item) bool {
	return i < j.(MinInt)
}

type MaxInt int

func (i MaxInt) Compare(j Item) bool {
	return i > j.(MaxInt)
}
