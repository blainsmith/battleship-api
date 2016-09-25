package battleship_test

import (
	"github.com/blainsmith/battleship-api/battleship"
	"testing"
)

func TestReceiveShot(t *testing.T) {
	game := battleship.NewGame()
	game.Grid = "0300222200030000000003100000000010005000001000500000100444000010000000000000000000000000000000000000"

	if game.ReceiveShot(&battleship.Coord{Letter: "A", Number: "2"}).Result != 1 {
		t.Error("expected a hit")
	}
	if game.ReceiveShot(&battleship.Coord{Letter: "C", Number: "2"}).Result != 1 {
		t.Error("expected a hit")
	}
	if game.ReceiveShot(&battleship.Coord{Letter: "B", Number: "2"}).Result != 2 {
		t.Error("expected a sink")
	}
	if game.ReceiveShot(&battleship.Coord{Letter: "A", Number: "2"}).Result != 2 {
		t.Error("expected a sink")
	}
	if game.ReceiveShot(&battleship.Coord{Letter: "D", Number: "2"}).Result != 0 {
		t.Error("expected a miss")
	}
}
