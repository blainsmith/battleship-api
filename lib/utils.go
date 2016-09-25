package lib

import (
	"math/rand"
	"time"
)

func Random(lower, upper int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(upper-lower) + lower
}
