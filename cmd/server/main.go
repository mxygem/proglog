package main

import (
	"log"

	"github.com/mxygem/proglog/internal/server"
)

const (
	port string = ":8080"
)

func main() {
	srv := server.NewHTTPServer(port)
	log.Fatal(srv.ListenAndServe())
}
