package main

import (
	"log"
	"net/http"

	"github.com/jm-menon/Jahnavi-s-Portfolio/server"
)

func main() {

	srv := server.NewServer()
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
