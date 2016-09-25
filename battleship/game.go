package battleship

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
)

// Fake KV database of all games in progress
var Games = make(map[string]*Game)

type Ship struct {
	ID   string
	Name string
	Size int
}

type Game struct {
	GameID string `json:"gameId"`
	Grid   string `json:"grid"`
	grid   [][]string
	fleet  []Ship
}

type ShotResult struct {
	Result int `json:"result"`
}

func NewGame() *Game {
	uuid, _ := uuid()
	game := &Game{
		GameID: uuid,
		fleet: []Ship{
			{ID: "1", Name: "Carrier", Size: 5},
			{ID: "2", Name: "Battleship", Size: 4},
			{ID: "3", Name: "Cruiser", Size: 3},
			{ID: "4", Name: "Submarine", Size: 3},
			{ID: "5", Name: "Destroyer", Size: 2},
		},
	}

	game.generateGrid()

	return game
}

func (g *Game) ReceiveShot(c *Coord) *ShotResult {
	result := 1
	position := c.Position()
	character := fmt.Sprintf("%c", g.Grid[position])
	index, _ := strconv.Atoi(string(character))
	index -= 1

	if index == -1 {
		result = 0
	} else {
		g.fleet[index].Size -= 1
		if g.fleet[index].Size <= 0 {
			result = 2
		}
	}

	return &ShotResult{
		Result: result,
	}
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

func (g *Game) generateGrid() {
	g.Grid = "0300222200030000000003100000000010005000001000500000100444000010000000000000000000000000000000000000"
}
