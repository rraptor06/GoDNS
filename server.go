package server

import (
	"fmt"
	"go-dns-server/handler"
	"log"
	"net"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	return &Server{address: address}
}

func (s *Server) Start() error {
	conn, err := net.ListenPacket("udp", s.address)
	if err != nil {
		return fmt.Errorf("listening failure on %s : %v", s.address, err)
	}
	defer conn.Close()

	buffer := make([]byte, 512)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error reading packet: %v", err)
			continue
		}
		go handler.HandleRequest(conn, addr, buffer[:n])
	}
}
