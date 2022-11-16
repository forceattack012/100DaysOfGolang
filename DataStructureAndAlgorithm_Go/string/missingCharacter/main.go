package main

import (
	"fmt"
	"strings"
)

func missingCharacterPanagram(s string) string {
	panagram := make([]bool, 26)
	result := ""
	s = strings.ToLower(s)
	index := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			index = int(s[i] - 'a')
		}
		panagram[index] = true
		fmt.Printf("%d ", index)
	}

	for j := 0; j < 26; j++ {
		if panagram[j] == false {
			result += fmt.Sprintf("%c", j+'a')
		}
	}

	return result
}
