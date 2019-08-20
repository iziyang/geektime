package _interface

import "testing"

type Programmer interface { // 接口定义
	WriteHelloWorld() string
}

//接口实现，不依赖于接口定义
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"

}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}
