# GoLang experimental project to test '[gorilla/mux](https://github.com/gorilla/mux)' library

I used this project to test the '[gorilla/mux](https://github.com/gorilla/mux)' library. It's a very simple RESTful API.

#### Response struct definition: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-restapi-experiment/blob/master/src/structs/MinecraftServer.go)

```go
type MinecraftServer struct {
	Id            string `json:"id"`
	Domain        string `json:"domain"`
	Score         int    `json:"score"`
	Image         int    `json:"image"`
	OnlinePlayers int    `json:"onlinePlayers"`
	MaxPlayers    int    `json:"maxPlayers"`
}
```

#### Request router: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-restapi-experiment/blob/master/src/main.go)
```go
func main() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/", HomePage)
	apiRouter.HandleFunc("/servers", AllServers)
	apiRouter.HandleFunc("/server", NewServer).Methods("POST")
	apiRouter.HandleFunc("/server/{id}", DeleteServer).Methods("DELETE")
	apiRouter.HandleFunc("/server/{id}", SingleServer)
	log.Fatal(http.ListenAndServe(":8080", apiRouter))
}
```

#### Example rest controller: [Click here to see the source code](https://github.com/xXNurioXx/simple-golang-restapi-experiment/blob/master/src/controllers/AllServers.go)
```go
func AllServers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Servers)
}
```
