package operator_test

import (
	"testing"
)

const (
	readable = 1 << iota
	writeable
	readwrite
)

func TestCompareArry(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 4, 5} //为什么要这样赋值？
	//d := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	t.Log(a == c)
	//t.Log(a == d)

}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ readable
	a = a &^ writeable
	t.Log(a&readable == readable, a&writeable == writeable, a&readwrite == readwrite)
	t.Log(readable, writeable, readwrite)
}
