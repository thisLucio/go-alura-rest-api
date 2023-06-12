package routes

import (
	"log"
	"net/http"

	"github.com/thislucio/go-alura-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
