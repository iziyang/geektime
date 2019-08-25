package groutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {  // 这种情况下则是值传递，会把i 的值全部复制一份
			fmt.Println(i)
		}(i)
		go func() {  // 协程的竞争条件，这样会导致每个协程只能获得最终的i的值的地址需要用锁的机制来完成
			sync.Mutex.Lock()
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Microsecond * 50)
}
