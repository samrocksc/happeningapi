package main

import (
	"testing"
)

func TestSampleUserTakesNameString(t *testing.T) {
	u := SampleUser{Name: "Test"}
	u2 := SampleUser{Name: "Test"}
	if u != u2 {
		t.Fatalf("SampleUser() failed")
	}
}
