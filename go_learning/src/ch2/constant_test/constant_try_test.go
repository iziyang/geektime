package constant_test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	readable = 1 << iota
	writeable
	readwrite
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&readable == readable, a&writeable == writeable, a&readwrite == readwrite)
}
