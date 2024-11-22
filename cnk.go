package leetcode

type cnk struct {
	mod  int
	f    []int
	invc map[int]int
}

func NewCnk(mod int) *cnk {
	return &cnk{
		mod,
		[]int{1},
		make(map[int]int),
	}
}

func (c *cnk) mul(arr ...int) int {
	res := 1
	for _, v := range arr {
		res = int((int64(res) * int64(v)) % int64(c.mod))
	}
	return res
}

func (c *cnk) pow(a, p int) int {
	if p == 0 {
		return 0
	}
	if p == 1 {
		return a
	}
	x := c.pow(a, p/2)
	x = c.mul(x, x)
	if p%2 == 1 {
		x = c.mul(x, a)
	}
	return x
}

func (c *cnk) fact(n int) int {
	for i := len(c.f); i <= n; i++ {
		c.f = append(c.f, c.mul(c.f[i-1]*i))
	}
	return c.f[n]
}

func (c *cnk) inv(n int) int {
	if v, ok := c.invc[n]; !ok { 
		res := c.pow(c.fact(n), c.mod-2)
		c.invc[n] = res
		return res
	} else {
		return v
	}
}

func (c *cnk) Calc(n int, k int) int {
	return c.mul(c.fact(n), c.inv(k), c.inv(n-k))
}
