package battleship

import (
	"crypto/rand"
	"fmt"
	"io"
)

// Fake KV database of all games in progress
var Games = make(map[string]*Game)

type Game struct {
	GameID string `json:"gameId"`
	Grid   string `json:"grid"`
}

func NewGame() *Game {
	uuid, _ := uuid()
	game := &Game{
		GameID: uuid,
		Grid:   generateGrid(),
	}

	Games[game.GameID] = game

	return game
}

func (g *Game) ReceiveShot(c *Coord) bool {
	position := c.Position()
	character := fmt.Sprintf("%c", g.Grid[position])

	if character != "0" {
		return true
	}

	return false
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
