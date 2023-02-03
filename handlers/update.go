package handlers

import (
	"net/http"
	"github.com/Api-go-postgres/models"
)

func Update(w http.ResponseWriter, r *http.Resquest){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error in parse of id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var todo models.todo


	decoder, err = json.NewDecoder(res.Body)
	decoder.Decode(&todo)
	
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

}
