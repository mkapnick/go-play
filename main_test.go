package main

import "testing"

func TestSum(t *testing.T) {
	c := &Circle{
		Point:  Point{1, 2},
		Radius: 1,
	}

	sum := c.Sum()

	if sum != 3 {
		t.Errorf("Incorrect sum, got: %d, want: %d", sum, 3)
	}

}
