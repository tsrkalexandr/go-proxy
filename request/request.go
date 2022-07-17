package request

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Request struct to hold internal request data.
// Contains information sends by client.
type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

// New create new instance of Request.
func New() Request {
	return Request{
		Headers: make(map[string]string),
	}
}

// Validate check is request is valid.
func Validate(req Request) error {
	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return fmt.Errorf("invalid url in request url=\"%s\": %w", req.URL, err)
	}

	for header := range req.Headers {
		if header == "" {
			return errors.New("empty header name")
		}
	}

	return nil
}

// Send send real http request based on parameters in `Request`.
func Send(req Request) (*http.Response, error) {
	r, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request object: %w", err)
	}

	for header, value := range req.Headers {
		r.Header.Add(header, value)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}
