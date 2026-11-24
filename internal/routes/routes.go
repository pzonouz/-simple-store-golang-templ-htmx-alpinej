package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/internal/services"
)

func MountApp(mux *chi.Mux, srv *services.Service) {
	// RegisterHomeRoute(mux)
	RegisterCategoriesRoute(mux)
	RegisterAdminRoute(mux)
}
