package headers

import (
	"bytes"
	"errors"
)

type Headers map[string]string

const SEPARATOR = "\r\n\r\n"

func NewHeadersA() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	//TODO: this has to have it's own buffer as well
	//
	idx := bytes.Index(data, []byte(SEPARATOR))

	if idx == -1 {
		return 0, false, nil
	}

	content := bytes.Split(data[:idx], []byte(":"))

	if (len(content)) > 2 {
		return 0, false, errors.New("Too many fields within the header")
	}

	h[string(content[0])] = string(content[1])

	return 0, false, nil

}
