package godns

import (
	"log"
	"github.com/miekg/dns"
)

type DNSConfig struct {
	Address     string
	UpstreamDNS string
}

type DNSServer struct {
	config DNSConfig
}

func NewDNSServer(address string, upstreamDNS string) *DNSServer {
	return &DNSServer{
		config: DNSConfig{
			Address:     address,
			UpstreamDNS: upstreamDNS,
		},
	}
}

func (s *DNSServer) handleDNSRequest(w dns.ResponseWriter, req *dns.Msg) {
	client := new(dns.Client)
	response, _, err := client.Exchange(req, s.config.UpstreamDNS)
	if err != nil {
		log.Printf("Error to request to the DNS : %v", err)
		return
	}
	w.WriteMsg(response)
}

func (s *DNSServer) Start() error {
	dns.HandleFunc(".", s.handleDNSRequest)

	srvUDP := &dns.Server{Addr: s.config.Address, Net: "udp"}
	srvTCP := &dns.Server{Addr: s.config.Address, Net: "tcp"}

	go func() {
		log.Printf("DNS UDP Server listen to %s", s.config.Address)
		if err := srvUDP.ListenAndServe(); err != nil {
			log.Fatalf("Error DNS : %v", err)
		}
	}()

	log.Printf("DNS TCP Server listen to %s", s.config.Address)
	return srvTCP.ListenAndServe()
}
