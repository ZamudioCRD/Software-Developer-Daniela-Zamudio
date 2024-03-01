package server

import (
	"encoding/json"
	"final-project/zincsearch"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func ConfigureRoutes() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{

		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},

		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/debug", middleware.Profiler())

	r.Get("/search/{term}", Rutas)

	http.ListenAndServe(":8080", r)
}

func Rutas(w http.ResponseWriter, r *http.Request) {
	termParam := chi.URLParam(r, "term")

	if termParam == "" {
		http.Error(w, "Se requiere un parámetro 'term'", http.StatusBadRequest)
		return
	}

	response, err := zincsearch.SearchZs(termParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al realizar la busqueda: %s", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al escribir la respuesta JSON: %s", err), http.StatusInternalServerError)
		return
	}
}
