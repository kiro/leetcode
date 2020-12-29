package leetcode

type dsu struct {
	p []int
	d []int
}

func newdsu(n int) *dsu {
	res := &dsu{
		make([]int, n+1),
		make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		res.p[i] = i
	}
	return res
}

func (d *dsu) add(i, j int) {
	pi := d.get(i)
	pj := d.get(j)

	if d.d[pi] > d.d[pj] {
		d.p[pj] = pi
	} else {
		d.p[pi] = pj
		if d.d[pj] < d.d[pi]+1 {
			d.d[pj] = d.d[pi] + 1
		}
	}
}

func (d *dsu) get(i int) int {
	j := i
	for i != d.p[i] {
		i = d.p[i]
	}

	for j != d.p[j] {
		j = d.p[j]
		d.p[j] = i
	}

	return i
}
