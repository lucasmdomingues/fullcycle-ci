package main

import "testing"

func TestSum(t *testing.T) {
	want := 30
	got := Sum(15, 15)
	if got != want {
		t.Errorf("Invalid sum, Got %d. Want %d", got, want)
	}
}
