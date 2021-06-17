package leetcode

import "math"

type IndexTree struct {
	n            int
	arr          []int
	f            func(i, j int) int
	defaultValue int
}

type Aggregation interface {
	DefaultValue() int
	Aggregate(left, right int) int
}

type MinIndex struct{}

func (m MinIndex) DefaultValue() int { return math.MaxInt32 }
func (m MinIndex) Aggregate(left, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}

type MaxIndex struct{}

func (m MaxIndex) DefaultValue() int { return math.MinInt32 }
func (m MaxIndex) Aggregate(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

type SumIndex struct{}

func (m SumIndex) DefaultValue() int             { return 0 }
func (m SumIndex) Aggregate(left, right int) int { return left + right }

func NewIndexTree(n int, a Aggregation) *IndexTree {
	arr := make([]int, n*4)
	for i := 0; i < len(arr); i++ {
		arr[i] = a.DefaultValue()
	}

	return &IndexTree{n, arr, a.Aggregate, a.DefaultValue()}
}

func (i *IndexTree) Add(k, v int) {
	i.traverse(1, 0, i.n, func(node, b, e int) bool {
		if b <= k && e >= k {
			i.arr[node] = i.f(i.arr[node], v)
			return true
		}
		return false
	})
}

func (i *IndexTree) Get(qb, qe int) int {
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

func (i *IndexTree) traverse(node int, b, e int, do func(node, b, e int) bool) {
	if !do(node, b, e) {
		return
	}
	m := (b + e) / 2
	if b != e {
		i.traverse(node*2, b, m, do)
		i.traverse(node*2+1, m+1, e, do)
	}
}
