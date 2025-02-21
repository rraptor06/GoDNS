package main

import (
	"fmt"
	"go-dns-server/server"
	"log"
)

func main() {
	address := "0.0.0.0:53"
	server := server.NewServer(address)

	fmt.Println("Starting the DNS server on", address)
	err := server.Start()
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
