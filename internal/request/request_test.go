package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestFromReader(t *testing.T) {
	// Test: Good GET Request line
	r, err := RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good GET Request line with path
	r, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Invalid number of parts in request line
	_, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.Error(t, err)

	// Good POST Request line
	_, err = RequestFromReader(strings.NewReader("POST /coffee HTTP/1.1\r\n"))
	assert.NoError(t, err)

	// Good DELETE Request line
	_, err = RequestFromReader(strings.NewReader("DELETE /coffee HTTP/1.1\r\n"))
	assert.NoError(t, err)

	// Good PUT Request line
	_, err = RequestFromReader(strings.NewReader("PUT /coffee HTTP/1.1\r\n"))
	assert.NoError(t, err)

	// Good PATCH Request line
	_, err = RequestFromReader(strings.NewReader("PATCH /coffee HTTP/1.1\r\n"))
	assert.NoError(t, err)

	// Invalid POST target missing '/'
	_, err = RequestFromReader(strings.NewReader("POST coffee HTTP/1.1\r\n"))
	assert.EqualError(t, err, "invalid target; got: coffee")

	// Invalid number of parts in request line
	_, err = RequestFromReader(strings.NewReader("GET HTTP/1.2\r\n"))
	assert.EqualError(t, err, "invalid request line part count")

	// Invalid method (out of order) Request line
	r, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1 POST\r\n"))
	assert.EqualError(t, err, "invalid request method; got: /coffee")

	// Invalid version in Request line
	_, err = RequestFromReader(strings.NewReader("GET / HTTP/1.2\r\n"))
	assert.EqualError(t, err, "invalid http version; got: HTTP/1.2")

	// Invalid empty request
	_, err = RequestFromReader(strings.NewReader(""))
	assert.EqualError(t, err, "invalid empty request")
}
