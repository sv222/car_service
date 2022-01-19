package main

import (
	"car_informer/internal/app"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	waitShutdown = time.Second * 15
)

func main() {
	serv := app.NewServer()
	l := serv.Logger
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
	serv.Router.HandleFunc("/", app.LoggingRequest(l, app.MainHandler)).Methods("GET")
	//serv.Router.HandleFunc("/register", app.RegisterHandler).Methods("GET")
	//serv.Router.HandleFunc("/sign-up", app.SignUpHandler).Methods("GET")

	go func() {
		err := serv.Start()
		if err != nil {
			l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), waitShutdown)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		l.Fatalf("problem with server shutdown occurred: %v", err)
	}

	l.Println("server shutdown completed successfully")
	os.Exit(0)
}
