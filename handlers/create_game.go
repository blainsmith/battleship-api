package handlers

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship-api/battleship"
	"net/http"
	"net/url"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	game := battleship.NewGame()

	battleship.Games[game.GameID] = game

	helm.RespondWithJSON(w, game, 201)
}
