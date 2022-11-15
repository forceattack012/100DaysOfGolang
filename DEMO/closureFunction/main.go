package main

import "fmt"

func main() {
	fn := counter()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func counter() func() int {
	var i int

	return func() int {
		i++
		return i
	}
}
