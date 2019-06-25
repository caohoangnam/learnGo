package main

import "fmt"

var x = 42

func main() {
	y := 10
	fmt.Println("x", x)
	fmt.Println("y", y)
	foo()
}

func foo() {
	fmt.Println(x)
}
