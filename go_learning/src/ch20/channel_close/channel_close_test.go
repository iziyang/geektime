package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //关闭 channel
		wg.Done()
	}()

}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok { // 可以返回一个 bool 值，代表 channel 是否关闭
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	//多个 receiver 也只能收到 10 个数
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
