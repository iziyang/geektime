package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "service Done"

}

func AsynService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("Asynservice start")
		retCh <- ret //service done
		fmt.Println("Asynservice exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsynService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")

	}
}
