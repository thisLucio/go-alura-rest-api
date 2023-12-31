package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thislucio/go-alura-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalidadeById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
