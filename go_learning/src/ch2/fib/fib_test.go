package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a, b = 1, 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = tmp + a

	}

}

func TestExchange(t *testing.T) {
	a, b := 1, 1
	a, b = b, a
	t.Log(a, b)
}
