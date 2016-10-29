package main

import "testing"

func TestStringSliceReverse(t *testing.T) {
	mySlice := []string{"the", "sky", "is", "blue"}
	reverseSlice(mySlice)

	if mySlice[0] != "blue" &&
		mySlice[1] != "is" &&
		mySlice[2] != "sky" &&
		mySlice[3] != "the" {
		t.Error("Invalid Slice return")
	}
}

func TestStringReverse(t *testing.T) {
	result := reverse("the sky is blue")
	expected := "blue is sky the"
	if result != expected {
		t.Errorf("Expected: %v, got %v", expected, result)
	}
}
