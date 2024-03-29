package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi"

	"github.com/Api-go-postgres/models"
)

func Delete(w http.ResponseWriter, r *http.Request){

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error in parse of id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error on remove register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d registers not removeds ", rows)
	}

	resp := map[string]any{

		"Error": false,
		"Message": "Remove data sucefully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
