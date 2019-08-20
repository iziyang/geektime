package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
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
	tsFn := timeSpent(slowFun)
	t.Log(tsFn(10))

}
