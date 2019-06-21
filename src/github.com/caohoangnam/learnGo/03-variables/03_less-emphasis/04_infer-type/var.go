package main

import "fmt"

func main() {
	var message = "Hello world"
	var a, b, c = 1, 4.12123123, "caonam"
	var d, e, f int
	d = 4444

	fmt.Println(message, a, b, c, d, e, f)
	fmt.Printf("%T %T %T %T \n", message, a, b, c)

}
