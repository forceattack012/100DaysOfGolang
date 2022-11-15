package main

import "testing"

func TestLeftRotationArray(t *testing.T) {
	give := []int{1, 2, 3, 4, 5, 6, 7}
	want := []int{5, 6, 7, 1, 2, 3, 4}

	get := LeftRotationArray(give, 4)

	if !isEqual(get, want) {
		t.Errorf("Faild not equal")
	}

}

func TestRightRotationArray(t *testing.T) {
	give := []int{1, 2, 3, 4, 5}
	want := []int{4, 5, 1, 2, 3}

	get := RightRotationArray(give, 2)

	if !isEqual(get, want) {
		t.Errorf("Faild not equal")
	}

}

func TestLeftRotationArrayByReverse(t *testing.T) {
	give := []int{1, 2, 3, 4, 5, 6, 7}
	want := []int{5, 6, 7, 1, 2, 3, 4}

	get := LeftRotationArrayByReverse(give, 4)

	if !isEqual(get, want) {
		t.Errorf("Faild not equal")
	}

}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
