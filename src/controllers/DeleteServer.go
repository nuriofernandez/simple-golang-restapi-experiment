package controllers

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/xXNurioXx/simple-golang-restapi-experiment/data"
)

func DeleteServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range data.Servers {
		if article.Id == id {
			data.Servers = append(data.Servers[:index], data.Servers[index+1:]...)
		}
	}

}
