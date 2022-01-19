package app

import (
	"html/template"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	tmpls := []string{
		"./web/home.page.gohtml",
		"./web/base.layout.gohtml",
		"./web/footer.layout.gohtml",
		"./web/header.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(tmpls...)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		log.Fatalf("could not parse template: %v", err)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		log.Fatalf("could not execute template: %v", err)
	}
}
