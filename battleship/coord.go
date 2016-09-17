package battleship

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Coord struct {
	Letter string `json:"letter"`
	Number string `json:"number"`
}

func RandomCoord() *Coord {
	letter := fmt.Sprintf("%c", random()+65)
	number := random()

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

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(9) + 1
}
