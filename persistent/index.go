package persistent

import "math"

type node struct {
	v     int
	left  *node
	right *node
}

type indexTree struct {
	n            int
	root         *node
	f            func(i, j int) int
	defaultValue int
}

type Aggregation interface {
	DefaultValue() int
	Aggregate(left, right int) int
}

func (m MinIndex) DefaultValue() int { return math.MaxInt32 }
func (m MinIndex) Aggregate(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}

func (m MaxIndex) DefaultValue() int { return -1 }
func (m MaxIndex) Aggregate(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func (m SumIndex) DefaultValue() int             { return 0 }
func (m SumIndex) Aggregate(left, right int) int { return left + right }

func (i *indexTree) Add(k, v int) IndexTree {
	newRoot := i.traverse(i.root, 0, i.n, true, func(cur *node, b, e int) *node {
		if b <= k && e >= k {
			next := &node{0, nil, nil}
			next.v = i.f(cur.v, v)
			return next
		}
		return nil
	})

	return &indexTree{i.n, newRoot, i.f, i.defaultValue}
}

func (i *indexTree) Get(qb, qe int) int {
	res := i.defaultValue
	i.traverse(i.root, 0, i.n, false, func(cur *node, b, e int) *node {
		if b >= qb && e <= qe {
			res = i.f(res, cur.v)
			return nil
		}
		if !(qe < b || qb > e) {
			return &node{0, nil, nil}
		} else {
			return nil
		}
	})
	return res
}

func (i *indexTree) traverse(cur *node, b, e int, add bool, do func(node *node, b, e int) *node) *node {
	x := do(cur, b, e)

	if x == nil {
		return cur
	}

	if cur.left == nil {
		cur.left = &node{i.defaultValue, nil, nil}
	}
	if cur.right == nil {
		cur.right = &node{i.defaultValue, nil, nil}
	}

	x.left = cur.left
	x.right = cur.right

	m := (b + e) / 2

	if b != e {
		x.left = i.traverse(cur.left, b, m, add, do)
		x.right = i.traverse(cur.right, m+1, e, add, do)
	}

	return x
}

// Aggregation types
type MinIndex struct{}
type MaxIndex struct{}
type SumIndex struct{}

type IndexTree interface {
	// Adds v to index i
	Add(i, v int) IndexTree
	// Gets the aggregation in the interval [b, e]
	Get(b, e int) int
}

func NewIndexTree(n int, a Aggregation) IndexTree {
	root := &node{a.DefaultValue(), nil, nil}

	return &indexTree{n, root, a.Aggregate, a.DefaultValue()}
}
