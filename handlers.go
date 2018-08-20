package main

import (
	_ "encoding/json"
	_ "fmt"
	"net/http"
)

func (a *App) InitializeRoutes() {
	//api := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	a.Router.HandleFunc("/", a.health).Methods("GET")
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["msg"] = "OK"
	respondWithJSON(w, http.StatusOK, m)
}
