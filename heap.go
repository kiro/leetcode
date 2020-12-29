package leetcode

import "container/heap"

type value int

type Item struct {
	v value
	p int
	i int
}

func item(v value, p int) *Item {
	return &Item{v, p, 0}
}

func Min(l, r *Item) bool {
	return l.p < r.p
}

func Max(l, r *Item) bool {
	return l.p > r.p
}

type Heap struct {
	p     *pq
	items map[value][]*Item
}

func NewHeap(compare func(l, r *Item) bool) *Heap {
	return &Heap{
		&pq{
			[]*Item{},
			compare,
		},
		make(map[value][]*Item),
	}
}

func (h *Heap) Push(i *Item) {
	h.addItem(i)
	heap.Push(h.p, i)
}

func (h *Heap) Pop() *Item {
	res := heap.Pop(h.p).(*Item)
	h.removeItem(res)
	return res
}

func (h *Heap) Top() *Item {
	if len(h.p.items) > 0 {
		return h.p.items[0]
	}
	return nil
}

func (h *Heap) Len() int {
	return len(h.p.items)
}

func (h *Heap) Update(item *Item, v value, p int) {
	h.removeItem(item)
	item.v = v
	item.p = p
	heap.Fix(h.p, item.i)
	h.addItem(item)
}

func (h *Heap) Remove(item *Item) {
	h.removeItem(item)
	heap.Remove(h.p, item.i)
}

func (h *Heap) Get(v value) []*Item {
	return h.items[v]
}

func (h *Heap) removeItem(res *Item) {
	x := h.items[res.v]
	for i, it := range x {
		if it == res {
			h.items[res.v] = append(x[:i], x[i+1:]...)
			break
		}
	}
}

func (h *Heap) addItem(item *Item) {
	h.items[item.v] = append(h.items[item.v], item)
}

// A PriorityQueue implements heap.Interface and holds Items.
type pq struct {
	items   []*Item
	compare func(l, r *Item) bool
}

func (pq pq) Len() int { return len(pq.items) }

func (pq pq) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.compare(pq.items[i], pq.items[j])
}

func (pq pq) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].i = i
	pq.items[j].i = j
}

func (pq *pq) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item)
	item.i = n
	pq.items = append(pq.items, item)
}

func (pq *pq) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.i = -1    // for safety
	pq.items = old[0 : n-1]
	return item
}
