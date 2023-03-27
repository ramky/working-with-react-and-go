package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	var app application

	// read from command line

	// connect to database
	app.Domain = "example.com"
	log.Println("Starting application on port: ", port)

	// start a webserver
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
