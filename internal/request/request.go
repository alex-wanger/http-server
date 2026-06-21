package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const SEPERATOR = "\r\n"

var MALFORMED_REQUEST_ERROR = errors.New("Malformed Request Line")

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(Reader io.Reader) (*Request, error) {
	input, err := io.ReadAll(Reader)

	if err != nil {
		return nil, err
	}
	requestLine, _, err := parseRequestLine(string(input))

	if err != nil {
		return nil, MALFORMED_REQUEST_ERROR
	}

	req := Request{
		RequestLine: *requestLine,
	}

	return &req, nil
}

func parseRequestLine(input string) (*RequestLine, string, error) {
	idx := strings.Index(input, SEPERATOR)

	if idx == -1 {
		fmt.Println("Seperator not found in input!")
		return nil, input, MALFORMED_REQUEST_ERROR
	}

	// goes up to the first seperator and takes a slice of the string
	var requestLine string = input[0 : idx+2]

	MethodTargetProtocol := strings.Split(requestLine, " ")

	if len(MethodTargetProtocol) != 3 {
		fmt.Println("Not enough seperators found in request line!")
		return nil, input, MALFORMED_REQUEST_ERROR
	}

	Method := MethodTargetProtocol[0]
	Target := MethodTargetProtocol[1]
	Protocol := strings.TrimSpace(MethodTargetProtocol[2])
	fmt.Println("MTP "+Method, Target, Protocol)

	if strings.ToUpper(Method) != Method {
		fmt.Println("Malformed METHOD in request line!")
		return nil, input, MALFORMED_REQUEST_ERROR
	}

	if Protocol != "HTTP/1.1" {
		fmt.Println("Protocol must be 1.1 in request line!" + " Recieved " + Protocol)
		return nil, input, errors.New("Protocol error, " + Protocol + " is not correct!")
	}

	Protocol = "1.1"

	requestLineStruct := RequestLine{
		HttpVersion:   Protocol,
		RequestTarget: Target,
		Method:        Method,
	}

	// succesfully return the new struct
	return &requestLineStruct, input[idx+len(SEPERATOR):], nil
}
