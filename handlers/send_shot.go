package handlers

import (
	"github.com/acmacalister/helm"
	"github.com/blainsmith/battleship-api/battleship"
	"net/http"
	"net/url"
)

func SendShotHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	randomCoord := battleship.RandomCoord()

	helm.RespondWithJSON(w, randomCoord, 200)
}
