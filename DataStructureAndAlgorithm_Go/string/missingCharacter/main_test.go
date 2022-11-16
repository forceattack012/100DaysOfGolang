package main

import "testing"

func TestMissingCharacterPanagram(t *testing.T) {
	give := "The quick brown fox jumps"
	want := "adglvyz"

	get := missingCharacterPanagram(give)

	if get != want {
		t.Errorf("FAILED")
	}
}
