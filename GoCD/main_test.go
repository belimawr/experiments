package main

import "testing"

func Test_sum(t *testing.T) {
	s := sum(40, 2)

	if s != 42 {
		t.Errorf("Expected 42, not %d", s)
	}
}

func Test_even_odd(t *testing.T) {
	ret := even(1)

	if ret {
		t.Error("Expected a false")
	}
}
