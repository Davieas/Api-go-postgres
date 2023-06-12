package handlers

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/Api-go-postgres/models"
	"net/http"
	
)

func Create(w http.ResponseWriter, r *http.Request){
	var todo models.Todo

	//If any error on json, verify here
	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		log.Printf("Error in decoder json: %v", err)
	  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	  return
	}

	_, err = models.Insert(todo)



	var resp map[string]any
	if err != nil {
	resp = map[string]any{
		"Error":   true,
		"Message": fmt.Sprintf("Any error occurred on insert: %v", err),
	}
		} else {
	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Task successfully inserted"),
	}
}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}