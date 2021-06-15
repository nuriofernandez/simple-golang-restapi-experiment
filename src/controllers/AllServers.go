package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xXNurioXx/simple-golang-restapi-experiment/data"
)

func AllServers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Servers)
}
