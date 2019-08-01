package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service Done"

}

func otherTask() {
	fmt.Println("otherTask start.")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done")

}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()

}

func AsynService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("Asynservice start")
		retCh <- ret //service done
		fmt.Println("Asynservice exited")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsynService()
	otherTask()
	fmt.Println(<-retCh)
	//time.Sleep(time.Second * 1)

}
