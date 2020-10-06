package main

import (
	"log"
	"net/http"
	"os"

	".."
)

func main() {
	http.HandleFunc("/Gopher", hellogo.Gopher)
	http.HandleFunc("/HelloGopher", hellogo.HelloGopher)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
