package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Gedis...")
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	numBytes, err := conn.Write([]byte("+OK\r\n"))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("(%d) bytes sent to client\n", numBytes)
}
