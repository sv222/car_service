package main

import (
	"car_informer/internal/app"
	"log"
)

func main() {
	serv := app.NewServer()
	log.Fatal(serv.Start())
}
