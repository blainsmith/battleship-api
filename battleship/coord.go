package battleship

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

type Coord struct {
	Letter string `json:"letter"`
	Number string `json:"number"`
}

func (c *Coord) Position() int {
	// Get the character code and offset it to a zero-based int
	letter, _ := utf8.DecodeRuneInString(c.Letter)
	letter -= 65

	// Convert the number into an actual number to set it to zero-based int
	number, _ := strconv.Atoi(c.Number)

	// Concat the position so A4 -> 3, C10 -> 29
	position, _ := strconv.Atoi(strings.ToUpper(strconv.Itoa(int(letter))) + strconv.Itoa((number - 1)))

	return position
}
