package main

import (
	"log"

	"github.com/tomoki-yamamura/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}