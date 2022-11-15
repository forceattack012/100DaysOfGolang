package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	start := time.Now()
	go dosomething(wg)
	go dosomething(wg)
	go dosomething(wg)

	wg.Wait()

	fmt.Println(time.Since(start))

	//exmaple()
}

func dosomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("dosomething")
	wg.Done()
}

func exmaple() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()

	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()

	time.Sleep(time.Second)
}
