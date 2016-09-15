package handlers

import (
	"crypto/rand"
	"fmt"
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship/types"
	"io"
	"net/http"
	"net/url"
)

var grid string = "0300222200030000000003100000000010005000001000500000100444000010000000000000000000000000000000000000"

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

func CreateGameHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	uuid, _ := uuid()

	game := &types.Game{
		GameID: uuid,
		Grid:   grid,
	}

	helm.RespondWithJSON(w, game, 201)
}
