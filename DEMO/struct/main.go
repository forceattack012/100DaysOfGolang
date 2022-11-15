package main

import "fmt"

type rectangle struct {
	w, l float64
}

func (r rectangle) getArea() float64 {
	return r.w * r.l
}

type Book struct {
	Name   string
	Author string
}

func (b *Book) SetName(name string) {
	b.Name = name
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s", b.Name, b.Author)
}

func main() {
	fmt.Println("hello")

	rec := rectangle{w: 10, l: 11}
	rec.w = 200

	fmt.Println(rec.w)
	fmt.Println(rec.l)
	fmt.Println(rec.getArea())

	b := Book{Name: "Harry Potter", Author: "J. K. Rowling"}
	b.SetName("HarryFORMATION")
	fmt.Println(b.String())
}
