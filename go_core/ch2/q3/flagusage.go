package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age *string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The hello world.")

	cmdLine.StringVar(&name, "name", "everyone", "The hello world.")

	age = flag.String("age", "18", "The age of name.")
}

func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
	//	flag.PrintDefaults()
	//}
	flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello %s.\n", name)
	fmt.Printf("My name is %s\n", *age)
	fmt.Println(age)

}
