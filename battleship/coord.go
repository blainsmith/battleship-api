package battleship

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/blainsmith/battleship-api/lib"
)

type Coord struct {
	Letter string `json:"letter"`
	Number string `json:"number"`
}

func RandomCoord() *Coord {
	letter := fmt.Sprintf("%c", lib.Random(1, 10)+65)
	number := lib.Random(1, 10)

	return &Coord{
		Letter: letter,
		Number: strconv.Itoa(number),
	}
}

func (c *Coord) Position() int {
	// Get the character code and offset it to a zero-based int
	letter, _ := utf8.DecodeRuneInString(strings.ToUpper(c.Letter))
	letter -= 65

	// Convert the number into an actual number to set it to zero-based int
	number, _ := strconv.Atoi(c.Number)

	// Concat the position so A4 -> 3, C10 -> 29
	position, _ := strconv.Atoi(strconv.Itoa(int(letter)) + strconv.Itoa((number - 1)))

	return position
}
