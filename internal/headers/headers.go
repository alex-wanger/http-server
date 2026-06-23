package headers

import (
	"bytes"
	"errors"
	"strings"
)

type Headers map[string]string

const CRLF = "\r\n"

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	//TODO: this has to have it's own buffer as well
	idx := bytes.Index(data, []byte(CRLF))

	if idx == -1 {
		return 0, false, nil
	}

	// END OF HEADERS, PARSING IS DONE
	if idx == 0 {
		return len(CRLF), true, nil
	}

	key, value, found := bytes.Cut(data[:idx], []byte(":"))

	if !found {
		return 0, false, errors.New("malformed header: no colon found")
	}

	if bytes.HasPrefix(key, []byte(" ")) || bytes.HasSuffix(key, []byte(" ")) {
		return 0, false, errors.New("malformed header: spaces around key")
	}

	h[(string(key))] = strings.TrimSpace(string(value))

	return len(data[:idx]) + len(CRLF), false, nil
}
