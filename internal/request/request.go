package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type parserState string

const (
	StateInit parserState = "init"
	StateDone parserState = "done"
)

const SEPARATOR = "\r\n"

var MALFORMED_REQUEST_ERROR = errors.New("Malformed Request Line")

type Request struct {
	RequestLine RequestLine
	State       parserState
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func newRequest() *Request {
	return &Request{
		State: StateInit,
	}
}

func (r *Request) done() bool {
	return r.State == StateDone
}

func (r *Request) parse(data []byte) (int, error) {
	if r.done() {
		return 0, nil
	}

	idx := bytes.Index(data, []byte(SEPARATOR))
	if idx == -1 {
		return 0, nil
	}

	line := data[:idx]
	parseRequestLine(string(line))

	return idx + len(SEPARATOR), nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()

	buffer := make([]byte, 1024)
	bufferIndex := 0

	for !request.done() {
		// read in n bytes into the buffer, only from the position AFTER the last read
		n, err := reader.Read(buffer[bufferIndex:])
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err // not EOF — genuine failure, bail immediately
			}
			if request.State != StateDone {
				return nil, fmt.Errorf("incomplete request: connection closed before request finished parsing")
			}
			break
		}

		bufferIndex += n

		readN, err := request.parse(buffer[:bufferIndex])

		if err != nil {
			return nil, err
		}

		copy(buffer, buffer[readN:bufferIndex])

		bufferIndex -= readN
	}

	return request, nil
}

func parseRequestLine(input string) (*RequestLine, int, error) {
	idx := strings.Index(input, SEPARATOR)

	if idx == -1 {
		fmt.Println("SEPARATOR not found in input!")
		return nil, 0, errors.New("SEPARATOR not found in input!")
		// not an error, just not enough input
	}

	// goes up to the first SEPARATOR and takes a slice of the string
	var requestLine string = input[0 : idx+2]

	MethodTargetProtocol := strings.Split(requestLine, " ")

	if len(MethodTargetProtocol) != 3 {
		fmt.Println("Not enough SEPARATORs found in request line!")
		return nil, 0, MALFORMED_REQUEST_ERROR
	}

	Method := MethodTargetProtocol[0]
	Target := MethodTargetProtocol[1]
	Protocol := strings.TrimSpace(MethodTargetProtocol[2])

	fmt.Println("MTP "+Method, Target, Protocol)

	if strings.ToUpper(Method) != Method {
		fmt.Println("Malformed METHOD in request line!")
		return nil, 0, MALFORMED_REQUEST_ERROR
	}

	if Protocol != "HTTP/1.1" {
		fmt.Println("Protocol must be 1.1 in request line!" + " Recieved " + Protocol)
		return nil, 0, errors.New("Protocol error, " + Protocol + " is not correct!")
	}

	Protocol = "1.1"

	requestLineStruct := RequestLine{
		HttpVersion:   Protocol,
		RequestTarget: Target,
		Method:        Method,
	}

	// succesfully return the new struct
	return &requestLineStruct, idx + len(SEPARATOR), nil
}
