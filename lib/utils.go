package lib

import (
	"math/rand"
	"time"
)

func Random(lower, upper int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(upper-lower) + lower
}
