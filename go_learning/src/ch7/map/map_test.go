package _map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	t.Log(m1[2])
	t.Logf("m1 len=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2 len=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("m3 len=%d", len(m3))

	t.Log(m1, m2, m3)

}

func TestAccessNoExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	//m1[2] = 0
	t.Log(m1[2])
	//m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3's value not exiting.")
	}

}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		t.Log(k, v)
	}

}
