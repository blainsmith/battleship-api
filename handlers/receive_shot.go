package handlers

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship-api/battleship"
	"net/http"
	"net/url"
)

func ReceiveShotHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	// Get the values from the URL
	gameId := params["gameId"][0]
	letter := params["letter"][0]
	number := params["number"][0]

	// Load the game from memory
	game := battleship.Games[gameId]
	if game == nil {
		helm.RespondWithJSON(w, struct{}{}, 404)
	} else {
		// Create a Coord of the values
		coord := &battleship.Coord{
			Letter: letter,
			Number: number,
		}

		// Recieve the shot and respond with the result
		helm.RespondWithJSON(w, game.ReceiveShot(coord), 200)
	}
}
