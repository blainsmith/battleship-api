package handlers

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship-api/battleship"
	"net/http"
	"net/url"
)

func DeleteGameHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	delete(battleship.Games, params["gameId"][0])

	helm.RespondWithJSON(w, struct{}{}, 201)
}
