package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum(5, 5) != 10 {
		t.Error("Expected 5 + 5 to equal 10");
	}
}