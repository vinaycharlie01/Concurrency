package main

import (
	"io"
	"log"
	"net"
	"time"
)

func HandlerConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "Response from the server\n")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)

	}
	for {
		conn, err := lis.Accept()

		if err != nil {
			// log.Fatal(err)
			continue

		}
		go HandlerConn(conn)
	}

}
