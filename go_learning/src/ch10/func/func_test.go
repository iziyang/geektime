package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		var (
			start = time.Now()
			ret   = inner(n)
		)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret

	}

}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	var a, b = returnMultiValues()
	t.Log(a, b)

	tsFn := timeSpent(slowFun)
	t.Log(tsFn(10))

}

func Sum(ops ...int) int  {
	var ret = 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func Clear()  {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("Start")
	panic("err")
	//fmt.Println("Continue.")

}