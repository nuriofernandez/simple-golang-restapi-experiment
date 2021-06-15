package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../data"
	. "../structs"
)

func NewServer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article MinecraftServer
	json.Unmarshal(reqBody, &article)

	data.Servers = append(data.Servers, article)

	json.NewEncoder(w).Encode(article)
}
