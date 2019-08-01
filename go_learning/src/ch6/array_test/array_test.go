package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr1 = [4]int{1, 2, 2, 3}
	var arr2 [3]int
	arr3 := [4]int{1, 2, 3, 4}
	arr4 := [...]int{1, 2, 3}
	t.Log(arr1, arr2, arr3, arr4)
	t.Log(arr4[1], arr4[2])
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for _, elem := range arr {
		t.Log(elem)
	}
	arr1 := arr[:]
	t.Log(arr1)

}
