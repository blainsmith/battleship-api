package handlers

import (
	"github.com/acmacalister/helm"
	"net/http"
	"net/url"
)

func RootHandler(w http.ResponseWriter, r *http.Request, params url.Values) {
	helm.RespondWithJSON(w, struct{}{}, 201)
}
