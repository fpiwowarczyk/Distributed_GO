package main

import (
	"log"

	"github.com/fpiwowarczyk/Distributed_GO/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
