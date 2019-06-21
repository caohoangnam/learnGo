package main

import "fmt"

func main() {
	var message = "Hello world"

	var a, b, c int = 1, 2, 3

	fmt.Printf("%T %T %T %T \n", message, a, b, c)
	fmt.Println(message, a, b, c)

}
