package leetcode

import (
	"sort"
	"testing"
)

func TestSkiplist(t *testing.T) {
	s := NewSkiplist(IntComparator{})

	s.Add(1, 2)
	node := s.Get(1)

	eq(t, 2, node.Value())
	eq(t, 1, int(node.Key().(int)))
	eq(t, node.Prev(), nil)
	eq(t, node.Next(), nil)

	s.Add(4, 3)

	node = s.Get(4)
	eq(t, 3, node.Value())

	node = s.Get(3)
	eq(t, 2, node.Value())
	node = s.Get(5)
	eq(t, 3, node.Value())

	eq(t, false, s.Remove(2))
	eq(t, true, s.Remove(1))

	node = s.Get(1)
	eq(t, nil, node.Value())
	node = s.Get(3)
	eq(t, nil, node.Value())
	node = s.Get(4)
	eq(t, 3, node.Value())
}

func TestGetAt(t *testing.T) {
	nums := []int{5, 8, 1, 3, 6, 2, 4, 7, 9}

	s := NewSkiplist(IntComparator{})
	for _, v := range nums {
		s.Add(v, v)
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		eq(t, s.GetAt(i).Value(), nums[i])
	}

	s.Remove(1)
	s.Remove(6)
	s.Remove(9)

	nums = []int{2, 3, 4, 5, 7, 8}
	for i := 0; i < len(nums); i++ {
		eq(t, s.GetAt(i).Value(), nums[i])
	}
}

func eq(t *testing.T, left, right interface{}) {
	if left != right {
		t.Fatalf("Expected %v = %v", left, right)
	}
}

func BenchmarkSkiplist_Add(b *testing.B) {
	s := NewSkiplist(IntComparator{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}
	b.StopTimer()
}

func BenchmarkSkiplist_Remove(b *testing.B) {
	s := NewSkiplist(IntComparator{})

	for i := 0; i < b.N; i++ {
		s.Add(i, i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
	b.StopTimer()
}
