package handlers

import (
	"net/http"
	"udemyCourse1/pkg/render"
)

// Home is a function that writes a response, main page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
