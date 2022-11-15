package main

import (
	"os"
)

func main() {
	writeFile()
}

func writeFile() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close() // Not work yet while end function

	if _, err = f.WriteString("hello test write file"); err != nil {
		panic(err)
	}
}
