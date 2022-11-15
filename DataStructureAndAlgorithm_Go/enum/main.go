package main

import "fmt"

func main() {
	printWeek(Wednesday)
}

func printWeek(w Week) {
	fmt.Printf("%s", w)
}

type Week int

const (
	Monday Week = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w Week) String() string {
	switch w {
	case Monday:
		return " Monday"
	case Tuesday:
		return " Tuesday"
	case Wednesday:
		return " Wednesday"
	case Thursday:
		return " Thursday"
	case Friday:
		return " Friday"
	case Saturday:
		return " Saturday"
	case Sunday:
		return " Sunday"
	}

	return "I don't know"
}
