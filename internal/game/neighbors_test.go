package game

import (
	"testing"
)

func TestNeighborsCount(t *testing.T) {

	res := neighbors(0, 0)
	if len(res) != 3 {
		t.Errorf("0, 0 should return 3 neighbors")
	}

	res1 := neighbors(0, 1)
	if len(res1) != 5 {
		t.Errorf("0, 1 should return 5 neighbors")
	}

	res2 := neighbors(1, 1)
	if len(res2) != 8 {
		t.Errorf("1, 1 should return 8 neighbors")
	}

	res3 := neighbors(1, 2)
	if len(res3) != 5 {
		t.Errorf("1, 2 should return 5 neighbors")
	}

	res4 := neighbors(2, 2)
	if len(res4) != 3 {
		t.Errorf("2, 2 should return 3 neighbors")
	}

}
