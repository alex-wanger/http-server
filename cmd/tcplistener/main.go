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

// func getLinesChannel(f io.ReadCloser) <-chan string {
// 	ch := make(chan string)
// 	go func() {
// 		defer close(ch) // always unblock receivers
// 		defer f.Close()
// 		defer fmt.Println("channel closed!")

// 		var line string

// 		data := make([]byte, 8)

// 		for {
// 			n, err := f.Read(data)
// 			if n > 0 {
// 				chunk := string(data[:n])
// 				for {
// 					idx := strings.IndexByte(chunk, '\n')
// 					if idx == -1 {
// 						line += chunk
// 						break
// 					}
// 					ch <- line + chunk[:idx]
// 					line = ""
// 					chunk = chunk[idx+1:]
// 				}
// 			}
// 			if err != nil {
// 				break // io.EOF or real error
// 			}
// 		}

// 		// flush last line if file didn't end with \n
// 		if line != "" {
// 			ch <- line
// 		}
// 	}()
// 	return ch
// }
