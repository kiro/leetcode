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

type pos struct {
	node int
	i int
}

type lca struct {
	fst map[int]int
	idx IndexTree
	at map[int][]pos
}

type Lca interface {
	Get(u, v int) int
} 

func NewLca(g [][]int) Lca {
	x := make([]int, 0)
	d := make([]int, 0)

    n := len(g)
    vis := make([]bool, n)

	var euler func(v int, k int)
	euler = func(u int, k int) {
		x = append(x, u)
		d = append(d, k)

		vis[u] = true
		for _, v := range g[u] {
			if !vis[v] {
				euler(v, k + 1)
				x = append(x, u)
				d = append(d, k)
			}
		} 
	}

	euler(0, 0)

	fst := make(map[int]int)
	idx := NewIndexTree(len(d) + 1, MinIndex{})
	at := make(map[int][]pos)

	for i := 0; i < len(d); i++ {
		if _, ok := fst[x[i]]; !ok {
			fst[x[i]] = i
		}
		idx.Add(i, d[i])
		at[d[i]] = append(at[d[i]], pos{x[i], i})
	}
	
	return &lca{
		fst, idx, at,
	}
}

func (lca *lca) Get(u, v int) int {
	b := lca.fst[u]
	e := lca.fst[v]

	if b > e {
		b, e = e, b
	}

	k := lca.idx.Get(b, e)

	j := sort.Search(len(lca.at[k]), func(i int) bool {
		p := lca.at[k][i]  
		return !(p.i < b)  
	})

	return lca.at[k][j].node
}