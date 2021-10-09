package leetcode

type substr struct {
	hash int
	i, j int
	str  string
}

func hash(s string, n int) []substr {
	res := make([]substr, 0)
	m := 1000000007
	p := 7

	mp := 1
	for i := 0; i < n; i++ {
		mp = (mp * p) % m
	}

	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (hash*p + int(s[i])) % m
		if i >= n {
			hash = (hash - int(s[i-n])*mp) % m
			if hash < 0 {
				hash += m
			}
		}

		if i >= n-1 {
			res = append(res, substr{
				hash,
				i - n + 1,
				i + 1,
				s[i-n+1 : i+1],
			})
		}
	}

	return res
}

func repeated(s string, n int) string {
	h := make(map[int][]string)

	for _, sub := range hash(s, n) {
		for _, v := range h[sub.hash] {
			if v == sub.str {
				return v
			}
		}

		h[sub.hash] = append(h[sub.hash], sub.str)
	}

	return ""
}
