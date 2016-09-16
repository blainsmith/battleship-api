package battleship_test

import (
	"github.com/blainsmith/battleship-api/battleship"
	"testing"
)

func TestPosition(t *testing.T) {
	A1 := &battleship.Coord{Letter: "A", Number: "1"}
	D10 := &battleship.Coord{Letter: "D", Number: "10"}

	if A1.Position() != 0 {
		t.Error("wrong position")
	}
	if D10.Position() != 39 {
		t.Error("wrong position")
	}
}
