package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start serverL %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		s.newClient(conn)
	}
}
