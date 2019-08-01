package _map

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {
	var m = map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	var MySet = map[int]bool{}
	MySet[1] = true
	var n int = 1
	if MySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	MySet[3] = true
	t.Log(len(MySet))
	delete(MySet, 1)

	if MySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
