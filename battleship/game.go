package battleship

import (
	"fmt"
	"strconv"

	"github.com/blainsmith/battleship-api/lib"
)

// Fake KV database of all games in progress
var Games = make(map[string]*Game)

type Ship struct {
	ID   int
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
	uuid, _ := lib.UUID()
	game := &Game{
		GameID: uuid,
		fleet: []Ship{
			{ID: 1, Name: "Carrier", Size: 5},
			{ID: 2, Name: "Battleship", Size: 4},
			{ID: 3, Name: "Cruiser", Size: 3},
			{ID: 4, Name: "Submarine", Size: 3},
			{ID: 5, Name: "Destroyer", Size: 2},
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

func (g *Game) generateGrid() {
	var grid [10][10]int

	x := 0
	y := 0
	d := 0
	shipPlaced := false
	overlap := false

	// Iterate over the fleet of ships to place them
	for index, ship := range g.fleet {
		shipPlaced = false // Ship has not been placed yet

		// Loop until the ship has been sucessfully placed
		for !shipPlaced {
			overlap = false // Ship is not overlapping existing ships yet

			// Get a random coord and direction
			x = lib.Random(0, 9)
			y = lib.Random(0, 9)
			d = lib.Random(0, 1) // This is not random enough, always 0

			// Does the ship start at an empty slot and fit within the bounds of the grid?
			if grid[x][y] == 0 && ((d == 0 && (x+ship.Size) <= 9) || (d == 1 && (y+ship.Size) <= 9)) {

				// Based on the direct and size of the ship, check the rest of the slots for a possible overlap
				for j := 0; j < index; j++ {
					if d == 0 && grid[x+j][y] != 0 {
						overlap = true
					} else if d == 1 && grid[x][y+j] != 0 {
						overlap = true
					}
				}

				// If there is no overlap fill the slots with the ID of the ship and
				if !overlap {
					for k := 0; k < ship.Size; k++ {
						if d == 0 {
							grid[x+k][y] = ship.ID
						} else {
							grid[x][y+k] = ship.ID
						}
					}
					shipPlaced = true // Stops the for loop to move onto the next ship
				}
			}
		}
	}

	var stringGrid []byte
	for r := 0; r <= 9; r++ {
		for c := 0; c <= 9; c++ {
			stringGrid = append(stringGrid, strconv.Itoa(grid[r][c])...)
		}
	}

	g.Grid = string(stringGrid)
}
