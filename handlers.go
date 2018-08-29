package main

import (
	_ "encoding/json"
	_ "fmt"
	"net/http"
)

func (a *App) InitializeRoutes() {
	//api := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	a.Router.HandleFunc("/todo", a.CreateTodo).Methods("POST")
	a.Router.HandleFunc("/todo", a.GetTodo).Methods("GET")
	a.Router.HandleFunc("/", a.health).Methods("GET")
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["msg"] = "OK"
	respondWithJSON(w, http.StatusOK, m)
}

func (a *App) GetTodo(w http.ResponseWriter, r *http.Request) {
	m := store.TodoGetter()
	respondWithJSON(w, http.StatusOK, m)
}

func (a *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	inc++
	t := Todo{
		Status:      false,
		Description: "Call",
		CreatedDate: "Today",
		Id:          inc,
	}
	store.TodoCreater(&t)
	// respond to http call
	respondWithJSON(w, http.StatusOK, "")
}
