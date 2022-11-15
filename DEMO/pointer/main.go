package main

import "fmt"

func main() {
	d := 2
	double(&d)

	fmt.Println(d)
}

func double(d *int) {
	*d = *d * 2
}
