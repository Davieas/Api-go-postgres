package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Api-go-postgres/configs"
	"github.com/Api-go-postgres/handlers"
)

func main() {

	configs.Load()

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Get(w, *r)
	})


	log.Printf("Server running on port: %s", configs.GetServerPort())
	err := http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Fatal(err)
	}
}
