package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/Api-go-postgres/models"
	"github.com/go-chi/chi"
)


func Get(w http.ResponseWriter, r http.Request){

	id, err := strconv.Atoi(chi.URLParam(&r, "id"))
	if err != nil {
		log.Printf("Error in parse of id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id),)
	if err != nil {
		log.Printf("Error on update register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}


	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}