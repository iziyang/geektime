package ch2

import (
	"flag"
	"fmt"
)

var name string
var age *string

func init() {
	flag.StringVar(&name, "name", "everyone", "The hello world.")
	age = flag.String("age", "18", "The age of name.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello %s.\n", name)
	fmt.Printf("My name is %s\n", *age)
	fmt.Println(age)

}
