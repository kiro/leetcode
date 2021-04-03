package leetcode

import (
	"math"
)

type sqrtType int

const (
	SqrtMin sqrtType = iota
	SqrtMax
	SqrtSum
)

type sqrt struct {
	n        int
	ns       int
	sqrt     []int
	arr      []int
	agg      func(i, j int) int
	defValue int
}

func NewSqrt(n int, t sqrtType) *sqrt {
	ns := int(math.Sqrt(float64(n))) + 1

	agg, defValue := getAgg(t)

	return &sqrt{
		n,
		ns + 1,
		mkArr(ns, defValue),
		mkArr(n, defValue),
		agg,
		defValue,
	}
}

func getAgg(t sqrtType) (func(i, j int) int, int) {
	switch t {
	case SqrtMin:
		return func(i, j int) int {
			if i < j {
				return i
			} else {
				return j
			}
		}, math.MaxInt32
	case SqrtMax:
		return func(i, j int) int {
			if i > j {
				return i
			} else {
				return j
			}
		}, math.MinInt32
	case SqrtSum:
		return func(i, j int) int { return i + j }, 0
	default:
		panic(t)
	}
}

func mkArr(n int, value int) []int {
	res := make([]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = value
	}
	return res
}

func (s *sqrt) set(i int, val int) {
	s.arr[i] = s.agg(s.arr[i], val)
	s.sqrt[i/s.ns] = s.agg(s.sqrt[i/s.ns], val)
}

func (s *sqrt) get(b, e int) int {
	res := s.defValue
	for i := b; i < e; {
		if i%s.ns == 0 && i+s.ns < e {
			res = s.agg(res, s.sqrt[i/s.ns])
			i += s.ns
		} else {
			res = s.agg(res, s.arr[i])
			i++
		}
	}

	return res
}
