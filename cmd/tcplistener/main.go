package main

import (
	"alex_wang/internal/request"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":42069")

	if err != nil {
		log.Fatal(err) // prints error and calls os.Exit(1)
	}

	conn, err := listen.Accept()

	if err != nil {
		log.Fatal(err) // prints error and calls os.Exit(1)
	} else {
		fmt.Println("connection accepted")
	}

	req, err := request.RequestFromReader(conn)

	fmt.Println(
		req.RequestLine.Method,
		req.RequestLine.RequestTarget,
		req.RequestLine.HttpVersion,
		req.State,
	)
}
