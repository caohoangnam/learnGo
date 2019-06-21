package main

import "fmt"

func main() {
	var message string
	var a, b, c int
	a = 1

	message = "Hello world"

	fmt.Printf("%T %T %T %T \n", a, b, c, message)
	fmt.Println("%T \n", b)
	fmt.Println("%T \n", c)
	fmt.Println("%T \n", message)

}
