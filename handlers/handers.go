package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/josnunezg/twittor/middlew"
	"github.com/josnunezg/twittor/routers"
	"github.com/rs/cors"
)

// Handlers set own port, the cors handle and server listend and serve
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/sign_up", middlew.CheckingBD(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("GO_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	corsHandler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, corsHandler))
}
