package leetcode

import "math"

type indexTree struct {
	n            int
	arr          []int
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

func (m MaxIndex) DefaultValue() int { return math.MinInt32 }
func (m MaxIndex) Aggregate(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

func (m SumIndex) DefaultValue() int             { return 0 }
func (m SumIndex) Aggregate(left, right int) int { return left + right }

func (i *indexTree) Add(k, v int) {
	i.traverse(1, 0, i.n, func(node, b, e int) bool {
		if b <= k && e >= k {
			i.arr[node] = i.f(i.arr[node], v)
			return true
		}
		return false
	})
}

func (i *indexTree) Get(qb, qe int) int {
	res := i.defaultValue
	i.traverse(1, 0, i.n, func(node, b, e int) bool {
		if b >= qb && e <= qe {
			res = i.f(res, i.arr[node])
			return false
		}
		return !(qe < b || qb > e)
	})
	return res
}

func (i *indexTree) traverse(node int, b, e int, do func(node, b, e int) bool) {
	if !do(node, b, e) {
		return
	}
	m := (b + e) / 2
	if b != e {
		i.traverse(node*2, b, m, do)
		i.traverse(node*2+1, m+1, e, do)
	}
}

// Aggregation types
type MinIndex struct{}
type MaxIndex struct{}
type SumIndex struct{}

type IndexTree interface {
	// Adds v to index i
	Add(i, v int)
	// Gets the aggregation in the interval [b, e]
	Get(b, e int) int
}

func NewIndexTree(n int, a Aggregation) IndexTree {
	arr := make([]int, n*4)
	for i := 0; i < len(arr); i++ {
		arr[i] = a.DefaultValue()
	}

	return &indexTree{n, arr, a.Aggregate, a.DefaultValue()}
}
