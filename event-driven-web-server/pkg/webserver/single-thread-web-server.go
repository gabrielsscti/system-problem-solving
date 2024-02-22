package webserver

import (
	"fmt"
	"net"
)

type SingleThreadWebServer struct {
	Protocol Protocol
	Port     int
}

func CreateSingleThreadWebServer(protocol Protocol, port int) *SingleThreadWebServer {
	return &SingleThreadWebServer{
		Protocol: protocol,
		Port:     port,
	}
}

func (s *SingleThreadWebServer) ListenAndServe() {
	ln, err := net.Listen(s.Protocol.String(), fmt.Sprintf(":%d", s.Port))
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Waiting for connection...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.handleConnection(conn)
	}
}

func (s *SingleThreadWebServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming data
	fmt.Println("Reading data...")
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)
}
