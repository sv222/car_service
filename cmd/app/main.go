package main

import (
	"car_informer/internal/app"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	serv := app.NewServer()
	// TODO: implement own logger with built-in logging
	l := serv.Logger
	r := serv.Router
	logFile, err := os.OpenFile("./logs/log.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		l.Fatalf("could not open log file %v", err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			l.Fatalf("problem with closing log file: %v", err)
		}
	}(logFile)

	serv.ConfigureLogger(logrus.InfoLevel, logFile)

	fs := http.FileServer(http.Dir("./web/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// TODO make logging for all handlers
	r.HandleFunc("/", app.LoggingRequest(l, app.MainHandler)).Methods("GET")
	r.HandleFunc("/sign-up", app.SignUpHandler).Methods("GET")
	r.HandleFunc("/sign-up", app.CreateUserHandler).Methods("POST")
	r.HandleFunc("/sign-in", app.SignInHandler).Methods("GET")
	r.HandleFunc("/{title}", app.PageHandler).Methods("GET")

	// TODO: implement graceful shutdown
	l.Fatal(serv.Start())
}
