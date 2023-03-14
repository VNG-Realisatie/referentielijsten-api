package main

import (
	"github.com/VNG-Realisatie/referentielijsten-api/api"
	"github.com/VNG-Realisatie/referentielijsten-api/data"
	"log"
	"net/http"
	"os"
)

const (
	standardPort string = ":8000"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = standardPort
	}

	log.Println("starting up.....")
	log.Println("loading data....")

	loader, err := data.LoadData()
	if err != nil {
		log.Fatalf("could not load any data %v", err)
	}
	log.Println("data loaded")

	srv := api.InitRoutes(loader)

	portStart := port[0:1]
	if portStart != ":" {
		port = ":" + port
	}

	log.Printf("%s : %s", "running on port", port)

	err = http.ListenAndServe(port, srv)
	if err != nil {
		log.Fatal(err)
	}
}
