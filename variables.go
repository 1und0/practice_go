package main

import (
	"fmt"
	"math"
)

const pi float32 = 3.14159

func main() {

	var a int = 1
	var b, c int = 2, 3
	var d = true
	var e string = "is this"
	short := "thats neat"

	fmt.Println(a + b + c)
	fmt.Println(e, d)
	fmt.Println(math.Sin(pi))
	fmt.Println(short)
}