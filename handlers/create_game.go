package handlers

import (
	"fmt"
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship/battleship"
	"net/http"
	"net/url"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	game := battleship.NewGame()

	helm.RespondWithJSON(w, game, 201)
}
