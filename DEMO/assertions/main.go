package main

import "fmt"

func main() {
	var i interface{}
	i = 1111

	var s string

	if v, ok := i.(string); ok {
		s = v
		fmt.Printf("%T %v", s, s)
	}

	var a, b Shap
	a = rectangle{w: 2, l: 2}
	b = triangle{w: 8, h: 8}

	if v, ok := b.(rectangle); ok {
		fmt.Println(v)
	}

	fmt.Println(a.Area())
	fmt.Println(b.Area())

	whichType(s)
}

func whichType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("this is a string %s \n", v)
	case int:
		fmt.Printf("this is a int %d \n", v)
	default:
		fmt.Printf("I don't know")
	}
}

type Shap interface {
	Area() float32
}

type rectangle struct {
	w, l float32
}

type triangle struct {
	w, h float32
}

func (r rectangle) Area() float32 {
	return (r.w * r.l)
}

func (t triangle) Area() float32 {
	return (0.5) * t.w * t.h
}
