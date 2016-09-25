package lib

import (
	crand "crypto/rand"
	"fmt"
	"io"
	mrand "math/rand"
	"time"
)

func Random(lower, upper int) int {
	mrand.Seed(time.Now().UTC().UnixNano())
	return mrand.Intn(upper-lower) + lower
}

func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
