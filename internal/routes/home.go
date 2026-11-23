package routes

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/web/templates"
)

func RegisterHomeRoute(mux *chi.Mux) {
	mux.Handle("/", templ.Handler(templates.Base("je")))

}
