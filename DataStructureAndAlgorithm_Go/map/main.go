package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "apple"
	m["b"] = "banana"

	for k, v := range m {
		fmt.Printf("key : %s, value %s \n", k, v)
	}

	for _, v := range m {
		if v == "banana" {
			fmt.Printf("%s\n", v)
		}
	}

	cusM := map[int]Customer{
		1: {name: "John", age: 18},
		2: {name: "Jame", age: 20},
	}

	for k, v := range cusM {
		fmt.Printf("key : %d, value %s, %d \n", k, v.name, v.age)
	}

	mStr := map[string][]string{
		"a": {"1", "2"},
		"c": {"2", "3", "4"},
	}

	for k, v := range mStr {
		fmt.Printf("key : %s, value %s \n", k, v)
	}

	var sli []Customer

	sli = append(sli, Customer{name: "John", age: 10})
	sli = append(sli, Customer{name: "John 2", age: 30})
	sli = append(sli, Customer{name: "Jame", age: 20})

	for _, v := range sli[:2] {
		fmt.Printf("%s %d \n", v.name, v.age)
	}

	for _, v := range sli[0:1] {
		fmt.Printf("%s %d \n", v.name, v.age)
	}
}

type Customer struct {
	name string
	age  int
}
