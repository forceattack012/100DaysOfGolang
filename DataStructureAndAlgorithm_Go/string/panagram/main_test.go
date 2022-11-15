package main

import "testing"

func TestIsPanagramIsFalse(t *testing.T) {
	give := "The quick brown fox jumps over the dog"
	want := false
	get := isPanagram(give)

	if want != get {
		t.Errorf("Not panahram")
	}
}

func TestIsPanagramIsTrue(t *testing.T) {
	give := "The quick brown fox jumps over the lazy dog"
	want := true
	get := isPanagram(give)

	if want != get {
		t.Errorf("Not panahram")
	}
}
