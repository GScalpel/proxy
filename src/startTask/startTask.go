package startTask

import (
	"log"
	"net"
)

func ListenAndServer(host string,port string) net.Listener {
	sock, err := net.Listen("tcp", host + ":" +port)
	if err != nil {
		log.Fatalf("server start error whether host is mistake or port has been used")
	}
	return sock
}

