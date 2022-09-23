package routes

import (
	"log"
	"net/http"

	"github.com/grbalmeida/curso-go-desenvolvendo-api-rest-alura/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
