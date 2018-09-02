package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *App) InitializeRoutes() {
	//api := a.Router.PathPrefix("/exp2").Subrouter()
	// exp.Use(authMiddleWareJWT)
	a.Router.HandleFunc("/todo", a.CreateTodo).Methods("POST")
	a.Router.HandleFunc("/todo", a.GetTodo).Methods("GET")
	a.Router.HandleFunc("/todo/{id}", a.UpdateTodo).Methods("PUT")
	a.Router.HandleFunc("/todo/{id}", a.DeleteTodo).Methods("DELETE")
	a.Router.HandleFunc("/todo/save", a.saveToJSON).Methods("GET")
	// health checks
	a.Router.HandleFunc("/", a.health).Methods("GET")

}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["msg"] = "OK"
	respondWithJSON(w, http.StatusOK, m)
}

func (a *App) GetTodo(w http.ResponseWriter, r *http.Request) {
	m, err := store.TodoGetter()
	if err != nil {
		respondWithError(w, err.Error())
	}
	respondWithJSON(w, http.StatusOK, m)
}

func (a *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	t := Todo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	if err := store.TodoCreater(&t); err != nil {
		respondWithError(w, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, fmt.Sprintf(`{"reference_id":%v}`, t.Id))
}

func (a *App) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idRaw := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		respondWithError(w, fmt.Sprintf("%v not a valid id", idRaw))
		return
	}
	err = store.TodoDeleter(id)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (a *App) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idRaw := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		respondWithError(w, fmt.Sprintf("%v not a valid id", idRaw))
		return
	}
	defer r.Body.Close()
	t := Todo{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, "Invalid request payload")
		return
	}
	err = store.TodoUpdater(id, &t)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}

func (a *App) saveToJSON(w http.ResponseWriter, r *http.Request) {
	d, err := store.TodoGetter()
	if err != nil {
		respondWithError(w, err.Error())
		return
	}
	var b []byte
	b, err = json.Marshal(d)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}
	ioutil.WriteFile("database.json", b, 0666)
	respondWithJSON(w, http.StatusOK, nil)
}
