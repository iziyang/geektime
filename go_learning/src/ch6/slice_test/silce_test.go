package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 3)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 3)
	t.Log(len(s2), cap(s2))

}

func TestSliceShareMemory(t *testing.T) {
	var s0 = []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	s1 := s0[3:6]
	s2 := s0[4:6]
	t.Log(s1, s2)
	s2[0] = "unknown"
	t.Log(s0, s1, s2)
	t.Log(len(s1), cap(s1))

}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	//if a == b{
	//	t.Log("equal")
	//}
	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(a, b, c)

}
