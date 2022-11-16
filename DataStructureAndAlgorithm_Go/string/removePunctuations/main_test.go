package main

import "testing"

func TestRemovePunctuations(t *testing.T) {
	give := "Welcome???@@##$ to#$% Geeks%$^for$%^&Geeks"
	want := "Welcome to GeeksforGeeks"
	get := RemovePunctuations(give)

	if want != get {
		t.Errorf("FAILED REMOVE THAT")
	}
}
