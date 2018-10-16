package main

import (
	"log"
	"net/http"
	"os"

	"github.com/phaicom/golang-docker-example/handlers"
	"github.com/phaicom/golang-docker-example/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// handle route
	r.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile", handlers.ProfileHandler).Methods(http.MethodGet)
	r.HandleFunc("/profile/{username}", handlers.MyProfileHandler).Methods(http.MethodGet)

	// handle not found route
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

	// serve static file
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Middleware
	http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))

	log.Fatal(http.ListenAndServe(":8080", r))
}
