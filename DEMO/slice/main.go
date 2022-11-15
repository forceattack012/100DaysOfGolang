package main

import "fmt"

func main() {
	/*var s []int

	if s == nil {
		fmt.Println("It's nil")
	}

	a := [...]int{1, 2, 3, 4, 5}
	s = a[:]

	fmt.Printf("%d %d %p %p\n", len(s), cap(s), s, &a)

	s = append(s, 12, 13)
	fmt.Printf("%d %d %p %p\n", len(s), cap(s), s, &a)*/

	abc()
	efg()
	cde()
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
	s = s[:3]

	fmt.Print(s)
	// [apple banana coconut]
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	s = s[4:]

	fmt.Print(s)
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed
	s = s[2:5]
	fmt.Print(s)
	// * [coconut durian elderberries]
}
