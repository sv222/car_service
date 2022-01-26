package app

import (
	"car_informer/internal/app/db"
	"car_informer/internal/app/model"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Handler struct {
}

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

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"./web/signup.page.gohtml",
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

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"./web/signin.page.gohtml",
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

func PageHandler(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("could not decode user: %v", err)
	}

	insertedID := db.InsertUser(user)

	res := response{
		ID:      insertedID,
		Message: "user created successfully",
	}

	json.NewEncoder(w).Encode(res)
}
