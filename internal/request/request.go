package request

import (
	"io"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HTTPVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(Reader io.Reader) (*Request, error) {
	//return nil
}
