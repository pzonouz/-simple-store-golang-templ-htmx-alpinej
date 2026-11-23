package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterCategoriesRoute(mux *chi.Mux) {
	mux.Route("/categories", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	})

}
