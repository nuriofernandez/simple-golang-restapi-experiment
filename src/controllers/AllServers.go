package controllers

import (
	"encoding/json"
	"net/http"

	"../data"
)

func AllServers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Servers)
}
