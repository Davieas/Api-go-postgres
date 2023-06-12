package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/Api-go-postgres/models"
	"github.com/go-chi/chi"
	
)

func Update(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error in parse of id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var todo models.Todo



	 err = json.NewDecoder(r.Body).Decode(&todo)

	
	if err != nil {
		log.Printf("Error in decoder json: %v", err)
	  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	  return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error on update register: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d registers is updates", rows)
	}

	resp := map[string]any{

		"Error": false,
		"Message": "Update data sucefully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
