package main

import "strings"

func main() {

}

func isPanagram(s string) bool {
	mark := make([]bool, 26)
	s = strings.ToLower(s)
	index := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			index = int(s[i] - 'a')
		}

		mark[index] = true
	}

	for i := 0; i < 26; i++ {
		if mark[i] != true {
			return false
		}
	}

	return true

}
