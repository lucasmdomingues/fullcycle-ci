package main

import (
	"math/rand/v2"
	"testing"
)

func TestSum(t *testing.T) {
	want := rand.IntN(100)
	got := Sum(15, 15)
	if got != want {
		t.Errorf("Invalid sum, Got %d. Want %d", got, want)
	}
}
