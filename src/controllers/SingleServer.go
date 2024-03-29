package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/xXNurioXx/simple-golang-restapi-experiment/data"
)

func SingleServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range data.Servers {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
