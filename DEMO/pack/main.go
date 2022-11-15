package main

import (
	"fmt"

	books "github/forceattack012/pack/book"
)

func main() {
	b := books.New()
	b.Name = "hello"

	fmt.Println(b)
}
