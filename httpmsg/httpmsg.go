// Package httpmsg holds HTTP requests and responses as plain values. It does
// no I/O, so a project's core can build requests and parse responses, and
// its shell can declare a port in these types without importing net/http.
package httpmsg

// Request is an HTTP request to send.
type Request struct {
	Method string
	URL    string
	Header map[string][]string
	Body   []byte
}

// Response is an HTTP response, read in full.
type Response struct {
	Status int
	Header map[string][]string
	Body   []byte
}
