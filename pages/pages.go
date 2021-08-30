package pages

import (
	"log"
	"myapp/handlers"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	handlers.RenderPageWithTemplate(w, "home.page.tmpl")
	log.Println("HomePage was called")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	handlers.RenderPageWithTemplate(w, "about.page.tmpl")
	log.Println("AboutPage was called")
}
