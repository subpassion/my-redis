package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	con, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to connect to tcp server")
		os.Exit(1)
	}

	for {
		fmt.Println("SENDING PING")
		con.Write([]byte("PING"))
		time.Sleep(1 * time.Second)
	}
}
