package leetcode

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	n := 31
	s := NewSqrt(n, SqrtSum)

	equal(t, s.get(0, n), 0)
	s.set(5, 3)
	equal(t, s.get(0, n), 3)
	equal(t, s.get(5, n), 3)
	equal(t, s.get(0, 5), 0)
	equal(t, s.get(0, 6), 3)

	s.set(6, 2)
	equal(t, s.get(0, n), 5)
	equal(t, s.get(5, n), 5)
	equal(t, s.get(6, n), 2)
	equal(t, s.get(7, n), 0)

	s.set(10, 1)
	s.set(20, 2)
	equal(t, s.get(2, 19), 6)
	equal(t, s.get(2, 20), 6)
	equal(t, s.get(2, 21), 8)
}

func equal(t *testing.T, l, r int) {
	if l != r {
		t.Fatalf("Expected %v = %v", l, r)
	}
}
