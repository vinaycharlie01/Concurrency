package main

import (
	"io"
	"log"
	"net"
	"os"
)

func MustCopy(in io.Writer, src io.Reader) {
	_, err := io.Copy(in, src)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	MustCopy(os.Stdout, conn)
}
