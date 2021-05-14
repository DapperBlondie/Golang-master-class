package Chapter7_1

import "testing"

func TestAddition(t *testing.T) {

	got := addition(5, 10)
	if got != 15 {

		t.Error(`addition(5, 10) != 15`)
	}
}

func TestMakeLine(t *testing.T) {

	got := makeLine(12.4, 34.67, 98.8)
	if got != "y=12.89x+98" {

		t.Error(`MissMatch !`)
	}
}
