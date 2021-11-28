package main

import "testing"

func TestSum(t *testing.T) {
	if sum(2, 3) != 5 {
		t.Errorf("sum(2, 3) = %d, want 5", sum(2, 3))
	}
}
