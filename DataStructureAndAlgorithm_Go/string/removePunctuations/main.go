package main

import (
	"fmt"
	"regexp"
)

func RemovePunctuations(s string) string {
	re := regexp.MustCompile(`[^\w\d\s]`)
	s = re.ReplaceAllString(s, "")

	fmt.Printf("%s", s)

	return s
}
