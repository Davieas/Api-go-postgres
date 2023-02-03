package main

import (
	"fmt"

	"github.com/Api-go-postgres/configs"
	"github.com/Api-go-postgres/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

func main() {

	err := configs.load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListAndServer(fmt.Sprintf(":%s", configs.GetServerPort(), r))

}
