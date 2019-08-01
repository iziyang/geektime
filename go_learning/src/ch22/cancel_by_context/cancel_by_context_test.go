package cancel_by_close

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool  {
	select {
	case <-ctx.Done():
		return true
	default:
		return false

	}

}

//func cancel_1(cancelChan chan struct{})  {
//	cancelChan <- struct{}{}
//}
//
//func cancel_2(cancelChan chan struct{})  {
//	close(cancelChan)
//
//}

func TestCancel(t *testing.T)  {
	//cancelChan := make(chan struct{}, 0)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	//cancel_2(cancelChan)
	//cancel_1(cancelChan)
	cancel()
	time.Sleep(time.Second * 1)
}