package app

import (
	"car_informer/internal/app/db"
	"car_informer/internal/app/model"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
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

	params := mux.Vars(r)
	title := params["title"]

	// TODO: check correct title
	if err := tmpl.Execute(w, title); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		log.Fatalf("could not execute template: %v", err)
	}

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("could not decode user: %v", err)
	}

	if err := user.Validate(); err != nil {
		json.NewEncoder(w).Encode("no email address or password provided")
		return
	}

	insertedID := db.InsertUser(user)

	res := response{
		ID:      insertedID,
		Message: "user created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Printf("couldn't convert string id to int: %v", err)
	}

	user, err := db.GetUserByID(id)

	if err != nil {
		log.Printf("couldn't get user from database: %v", err)
	}

	res := struct {
		ID    int64
		Email string
	}{
		ID:    user.ID,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(res)
}
