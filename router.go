package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// NewRouter created a new mux.router and initializes the routes
func (a *App) NewRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	// Initialize different Routes
	a.InitializeRoutes()
	// MiddleWare to use
	a.Router.Use(logRequest)
}

// logging request middleware
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL.Path,
			"method": r.Method,
		}).Info("http-request")
		next.ServeHTTP(w, r)
	})
}
