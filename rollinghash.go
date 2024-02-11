package leetcode

type substr struct {
	hash int64
	i, j int
	str  string
}

const modulo int64  =  288230376151711615 // 1152921504606846975 //9223372036854775837

func hash(s string, n int) []substr {
	res := make([]substr, 0)
    var m int64 = modulo
    p := int64(7)

    mp := int64(1)
	for i := 0; i < n; i++ {
		mp = (mp * p) % m
	}

    hash := int64(0)
	for i := 0; i < len(s); i++ {
		hash = (hash*p + int64(s[i])) % m
		if i >= n {
			hash = (hash - int64(s[i-n])*mp) % m
		}
		if hash < 0 {
			hash += m
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
