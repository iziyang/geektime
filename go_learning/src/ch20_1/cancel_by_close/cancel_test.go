package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false

	}

}

func cancel_1(cancelChan chan struct{}) {
	var p = &cancelChan
	fmt.Println(p)
	cancelChan <- struct{}{}
	//cancelChan <- struct{}{}
	//cancelChan <- struct{}{}
	//cancelChan <- struct{}{}
	//cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)

}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{})
	var p = &cancelChan
	fmt.Println(p)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			var p = &cancelCh
			fmt.Println(i,p)
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	//cancel_1(cancelChan)
	time.Sleep(time.Second * 1)
}
