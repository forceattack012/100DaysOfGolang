package main

import "testing"

func TestBinarySearch(t *testing.T) {
	give := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 3
	want := 2

	get := BinarySearch(give, target)

	if get != want {
		t.Errorf("got %v, want %v", get, want)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	give := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 11
	want := -1

	get := BinarySearch(give, target)

	if get != want {
		t.Errorf("got %v, want %v", get, want)
	}
}
