package main

import (
	"fmt"
	"sync"
)

var i int
var mux sync.Mutex

type thing struct {
	i   int
	mux sync.Mutex
}

func main() {
	go func() {
		for {
			fmt.Printf("%d \n", read())
		}
	}()
	for i := 0; i < 10; i++ {
		write(i)
	}
}

func write(n int) {
	mux.Lock()
	i = n
	mux.Unlock()
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}
