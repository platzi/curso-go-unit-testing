package main

import "testing"

func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)

	expect := 22

	if result != expect {
		t.Errorf("got %q, expected %d", result, expect)
	}
}
