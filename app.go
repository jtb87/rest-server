package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func (a *App) initApp() {
	// initialize router
	a.NewRouter()
	// initialize log
	initLog()
	// initialize database and store
	d := &dbStore{
		db: make(map[int]Todo),
	}
	err := d.loadFromJSON()
	if err != nil {
		fmt.Println(err)
	}
	InitStore(d)

}

type App struct {
	// DB     dbStore
	Router *mux.Router
	// Log string // possibly?
}

// Run starts the server
func (a *App) Run(port string) {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	s := &http.Server{
		Addr:    ":" + port,
		Handler: http.TimeoutHandler(handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router), time.Second*10, "timeout"),
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}
	log.Printf("server started on: http://localhost:%s", port)
	log.Fatal(s.ListenAndServe())
}

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
