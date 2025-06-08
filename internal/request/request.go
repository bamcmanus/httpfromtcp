package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var requestMethods = map[string]struct{} {
	"GET": {},
	"POST": {},
	"DELETE": {},
	"PATCH": {},
	"PUT": {},
}

const httpVersionPart = "HTTP/1.1"

func RequestFromReader(reader io.Reader) (*Request, error) {
	reqBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	reqString := string(reqBytes)
	parts := strings.Split(reqString, "\r\n")

	if parts[0] == "" {
		return nil, fmt.Errorf("invalid empty request")
	}

	reqLine, err := parseRequestLine(parts[0])
	if err != nil {
		return nil, err
	}
	if reqLine == nil {
		return nil, fmt.Errorf("failed to parse request line: '%s'", reqString)
	}
	
	return &Request{
		RequestLine: *reqLine,
	}, nil
}

func parseRequestLine(line string) (*RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line part count")
	}

	if _, ok := requestMethods[parts[0]]; !ok {
		return nil, fmt.Errorf("invalid request method; got: %s", parts[0])
	}

	if !strings.HasPrefix(parts[1], "/") {
		return nil, fmt.Errorf("invalid target; got: %s", parts[1])
	}

	if parts[2] != httpVersionPart {
		return nil, fmt.Errorf("invalid http version; got: %s", parts[2])
	}
	
	return &RequestLine{
		Method: parts[0],
		RequestTarget: parts[1],
		HttpVersion: "1.1",
	}, nil
}
 
