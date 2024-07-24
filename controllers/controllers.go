package controllers

import (
	"api-rest-golang/database"
	"api-rest-golang/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var newP models.Personality

	json.NewDecoder(r.Body).Decode(&newP)
	database.DB.Create(&newP)
	json.NewEncoder(w).Encode(newP)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	var p models.Personality

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var p models.Personality

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
