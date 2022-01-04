package main

import (
	"github.com/fpiwowarczyk/Distributed_GO/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
}
