package handler

import (
	"go-dns-server/resolver"
	"go-dns-server/utils"
	"log"
	"net"
)

func HandleRequest(conn net.PacketConn, addr net.Addr, data []byte) {
	log.Printf("Request received from %v", addr)

	ipAddress := resolver.ResolveDomain("example.com") // Example URL

	response := utils.CreateDNSResponse(ipAddress)

	_, err := conn.WriteTo(response, addr)
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
