package main

import (
	"fmt"
	"sort"
)

func main() {
	r := rectangle{w: 10, l: 5}
	fmt.Printf("%s %0.2f\n", r.toString(), r.volume())

	t := triangle{w: 10, h: 5}
	fmt.Printf("%s %0.2f\n", t.toString(), t.volume())

	customers := []customer{
		{name: "John", age: 18},
		{name: "Jin", age: 20},
		{name: "Jin", age: 5},
	}

	sort.Slice(customers, func(i, j int) bool {
		return customers[i].age < customers[j].age
	})

	for _, c := range customers {
		fmt.Printf("%s %d\n", c.name, c.age)
	}
}

type customer struct {
	name string
	age  int
}

type shape interface {
	toString() string
	volume() float64
}

type rectangle struct {
	w, l float64
}

type triangle struct {
	w, h float64
}

func (r rectangle) toString() string {
	return fmt.Sprintf("It's a rectangle")
}

func (r rectangle) volume() float64 {
	return r.w * r.l
}

func (t triangle) toString() string {
	return fmt.Sprintf("It's a triangle")
}

func (r triangle) volume() float64 {
	return 0.5 * r.w * r.h
}
