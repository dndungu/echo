package main

import (
	"io"
	"log"
	"net"
)

const HOSTPORT = ":1024"

var CommitSha string

func main() {
	clients := make(chan net.Conn)
	server, err := net.Listen("tcp", HOSTPORT)
	if err != nil {
		log.Fatalf("Server error, could not start listener : %s", err.Error())
	}
	log.Printf("Echo version %s, listening on %s", CommitSha, HOSTPORT)
	go func() {
		client, err := server.Accept()
		if err == nil {
			clients <- client
		} else {
			log.Printf("Accept error : %s", err.Error())
		}
	}()
	for {
		go func(c net.Conn) {
			io.Copy(c, c)
		}(<-clients)
	}
}
