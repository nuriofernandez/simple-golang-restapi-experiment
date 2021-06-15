package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/xXNurioXx/simple-golang-restapi-experiment/data"
	. "github.com/xXNurioXx/simple-golang-restapi-experiment/structs"
)

func NewServer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article MinecraftServer
	json.Unmarshal(reqBody, &article)

	data.Servers = append(data.Servers, article)

	json.NewEncoder(w).Encode(article)
}
