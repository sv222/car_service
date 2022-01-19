package main

import (
	"car_informer/internal/app"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	serv := app.NewServer()
	logFile, err := os.OpenFile("./logs/log.txt", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("could not open log file %v", err)
	}
	defer logFile.Close()
	serv.ConfigureLogger(logrus.InfoLevel, logFile)
	serv.Router.HandleFunc("/", app.MainHandler)
	log.Fatal(serv.Start())
}
