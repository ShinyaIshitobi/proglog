package main

import (
	"log"

	"github.com/ShinyaIshitobi/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
