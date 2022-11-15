package main

import "fmt"

func main() {
	variadicString("a", "b", "c")

	s := []string{"z", "x", "y"}
	variadicString(s...)
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
