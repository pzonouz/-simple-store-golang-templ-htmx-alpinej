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
	menuItems := []myTypes.MenuItemType{
		{
			Title: "خانه",
			Slug:  "/",
		},
		{
			Title: "پارامترها",
			Children: {
				myTypes.MenuItemType{Title: "گروه بندی", Slug: "/parameters-groups"},
				myTypes.MenuItemType{},
			},
		},
	}
	mux.Handle("/admin", templ.Handler(pages.AdminPage(menuItems)))
}
