package battleship

import (
	"crypto/rand"
	"fmt"
	"io"
)

// Fake KV database of all games in progress
var Games = make(map[string]*Game)

type Game struct {
	GameID     string `json:"gameId"`
	Grid       string `json:"grid"`
	carrier    int
	battleship int
	cruiser    int
	submarine  int
	destroyer  int
}

func NewGame() *Game {
	uuid, _ := uuid()
	game := &Game{
		GameID:     uuid,
		Grid:       generateGrid(),
		carrier:    5,
		battleship: 4,
		cruiser:    3,
		submarine:  3,
		destroyer:  2,
	}

	return game
}

func (g *Game) ReceiveShot(c *Coord) int {
	result := 1
	position := c.Position()
	character := fmt.Sprintf("%c", g.Grid[position])

	switch character {
	case "0":
		result = 0
	case "1":
		g.carrier -= 1
		if g.carrier <= 0 {
			result = 2
		}
	case "2":
		g.battleship -= 1
		if g.battleship <= 0 {
			result = 2
		}
	case "3":
		g.cruiser -= 1
		if g.cruiser <= 0 {
			result = 2
		}
	case "4":
		g.submarine -= 1
		if g.submarine <= 0 {
			result = 2
		}
	case "5":
		g.destroyer -= 1
		if g.destroyer <= 0 {
			result = 2
		}
	}

	return result
}

func uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func generateGrid() string {
	return "0300222200030000000003100000000010005000001000500000100444000010000000000000000000000000000000000000"
}
