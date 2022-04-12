package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const MAX_DEPTH = 20

type node struct {
	isEnd   bool
	isStart bool
	key     interface{}
	value   interface{}
	next    []*node
	count   []int
	prev    []*node
}

func (n *node) Next() Node {
	if n.next[0] == nil || n.next[0].isEnd {
		return nil
	}
	return n.next[0]
}

func (n *node) Prev() Node {
	if n.prev[0] == nil || n.prev[0].isStart {
		return nil
	}
	return n.prev[0]
}

func (n *node) Value() interface{} {
	return n.value
}

func (n *node) Key() interface{} {
	return n.key
}

func (n *node) String() string {
	return fmt.Sprintf("%v", *n)
}

func newNode(depth int, key interface{}, value interface{}) *node {
	return &node{
		false,
		false,
		key,
		value,
		make([]*node, depth+1),
		make([]int, depth+1),
		make([]*node, depth+1),
	}
}

type skiplist struct {
	cc   Comparator
	r    *rand.Rand
	root *node
	end  *node
}

func (s *skiplist) nodeDepth() int {
	x := MAX_DEPTH - int(math.Log2(s.r.Float64()*(1<<MAX_DEPTH)))
	if x > MAX_DEPTH {
		return MAX_DEPTH
	}
	return x
}

func (s *skiplist) compare(node *node, key interface{}) int {
	if node.isStart {
		return -1
	}
	if node.isEnd {
		return 1
	}
	return s.cc.Compare(node.key, key)
}

func (s *skiplist) path(key interface{}, depth int) []*node {
	res := make([]*node, depth+1)
	cur := s.root
	for depth >= 0 {
		res[depth] = cur
		if s.compare(cur.next[depth], key) == 1 {
			depth--
		} else {
			cur = cur.next[depth]
		}
	}

	return res
}

func (s *skiplist) Add(key interface{}, value interface{}) {
	depth := s.nodeDepth()
	//fmt.Println(depth)
	node := newNode(depth, key, value)

	p := s.path(key, MAX_DEPTH)

	for i := 0; i < len(node.next); i++ {
		node.next[i] = p[i].next[i]
		node.prev[i] = p[i]
		p[i].next[i].prev[i] = node
		p[i].next[i] = node
	}

	count := 0
	cur := p[0]
	steps := 0
	for d := 0; d <= MAX_DEPTH; d++ {
		if d <= depth {
			node.count[d] = cur.count[d] - count - steps
			cur.count[d] = count + steps
		} else {
			cur.count[d] += 1
		}

		for d+1 < len(p) && p[d+1] != cur {
			steps++
			cur = cur.prev[d]
			count += cur.count[d]
		}
	}
}

func (s *skiplist) Remove(key interface{}) bool {
	path := s.path(key, MAX_DEPTH)
	if s.compare(path[0], key) != 0 {
		return false
	}

	node := path[0]

	for i := 0; i < len(path); i++ {
		if i < len(node.next) {
			node.prev[i].next[i] = node.next[i]
			node.prev[i].count[i] += node.count[i]
			node.next[i].prev[i] = node.prev[i]
		} else {
			path[i].count[i]--
		}
	}

	return true
}

func (s *skiplist) Get(k interface{}) Node {
	p := s.path(k, MAX_DEPTH)
	return p[0]
}

func (s *skiplist) GetAt(i int) Node {
	i++
	depth := MAX_DEPTH
	cur := s.root

	for depth >= 0 && i > 0 {
		if i >= cur.count[depth]+1 {
			i -= cur.count[depth] + 1
			cur = cur.next[depth]
		} else {
			depth--
		}
	}

	return cur
}

func (s *skiplist) String() string {
	n := 0
	for cur := Node(s.root); cur.Next() != nil; cur = cur.Next() {
		n++
	}

	res := make([][]byte, MAX_DEPTH+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, n+2)

		for j := 0; j < len(res[i]); j++ {
			res[i][j] = ' '
		}
	}

	i := 0
	for cur := s.root; cur != nil; cur = cur.next[0] {
		for j := 0; j < len(cur.count); j++ {
			res[j][i] = byte(cur.count[j]) + '0'
		}
		if cur.Value() != nil {
			res[0][i] = byte(cur.Value().(int)) + '0'
		}
		i++
	}

	x := ""
	for i := 0; i < len(res); i++ {
		x += string(res[i]) + "\n"
	}

	return x
}

func (s *skiplist) First() Node {
	return s.root.Next()
}

func (s *skiplist) Last() Node {
	return s.end.Prev()
}

func (i IntComparator) Compare(l, r interface{}) int {
	left := l.(int)
	right := r.(int)
	if left == right {
		return 0
	} else if left < right {
		return -1
	}
	return 1
}

func NewSkiplist(c Comparator) Skiplist {
	start := newNode(MAX_DEPTH, nil, nil)
	start.isStart = true
	end := newNode(MAX_DEPTH, nil, nil)
	end.isEnd = true

	for i := 0; i <= MAX_DEPTH; i++ {
		start.next[i] = end
		end.prev[i] = start
	}

	return &skiplist{
		c,
		rand.New(rand.NewSource(time.Now().UnixNano())),
		start,
		end,
	}
}

type Comparator interface {
	// Returns -1 if smaller, 0 if equal, 1 if bigger
	Compare(left, right interface{}) int
}

type IntComparator struct{}

type Node interface {
	// Return previous node or nil if it doesn't exist
	Prev() Node
	// Return next node or nil if it doesn't exist
	Next() Node
	Value() interface{}
	Key() interface{}
}

type Skiplist interface {
	Add(k interface{}, v interface{})
	Remove(k interface{}) bool
	// Get the node if key is present, or the node after which Key will be inserted
	Get(k interface{}) Node
	// Gets the node at zero based index i in the sorted order.
	GetAt(i int) Node
	First() Node
	Last() Node
	String() string
}
