package routes

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/myTypes"
	"github.com/pzonouz/simple-store-golang-templ-htmx-alpinejs/web/templates/pages"
)

//	func RegisterHomeRoute(mux *chi.Mux) {
//		mux.Handle("/", templ.Handler(ui.Base("je")))
//	}

func RegisterAdminRoute(mux *chi.Mux) {
	var menuItems []myTypes.MenuItemType
	menuItems = []myTypes.MenuItemType{{
		Title: "خانه",
		Slug:  "/",
	}}
	mux.Handle("/admin", templ.Handler(pages.AdminPage(menuItems)))
}
