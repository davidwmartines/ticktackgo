package main

import (
	"testing"
)

func TestGetNeighborsCount(t *testing.T) {

	res := getNeighbors(0, 0)
	if len(res) != 3 {
		t.Errorf("0, 0 should return 3 neighbors")
	}

	res1 := getNeighbors(0, 1)
	if len(res1) != 5 {
		t.Errorf("0, 1 should return 5 neighbors")
	}

	res2 := getNeighbors(1, 1)
	if len(res2) != 8 {
		t.Errorf("1, 1 should return 8 neighbors")
	}

	res3 := getNeighbors(1, 2)
	if len(res3) != 5 {
		t.Errorf("1, 2 should return 5 neighbors")
	}

	res4 := getNeighbors(2, 2)
	if len(res4) != 3 {
		t.Errorf("2, 2 should return 3 neighbors")
	}

}
