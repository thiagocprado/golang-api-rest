package routes

import (
	"api-rest-golang/controllers"
	"api-rest-golang/middleware"
	"github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/personalities/{id}", controllers.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalities", controllers.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/personalities/{id}", controllers.GetById).Methods(http.MethodGet)
	r.HandleFunc("/api/personalities/{id}", controllers.Update).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
