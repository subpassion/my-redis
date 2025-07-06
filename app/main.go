package main

import (
	"log"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	conn_addr := "0.0.0.0:6379"
	l, err := net.Listen("tcp", conn_addr)
	if err != nil {
		log.Println("Failed to bind to port 6379: ", err.Error())
		os.Exit(1)
	}

	log.Println("Listening on:", conn_addr)
	con, err := l.Accept()
	if err != nil {
		log.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handle(con)
}

func handle(con net.Conn) {
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Println("Failed to close connection: ", err.Error())
			os.Exit(1)
		}
	}(con)

	for {
		data := make([]byte, 1024)
		n, err := con.Read(data)
		if err != nil {
			log.Println("Error reading bytes from connection: ", err.Error())
			os.Exit(1)
		}

		message := string(data[:n])
		if message == "PING" {
			_, err = con.Write([](byte)("+PONG\r\n"))
			if err != nil {
				log.Println("Filed to write PONG response: ", err.Error())
				os.Exit(1)
			}
		}
	}
}
