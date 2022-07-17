package url

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/tsrkalexandr/go-proxy/handler"
	"github.com/tsrkalexandr/go-proxy/request"
	"github.com/tsrkalexandr/go-proxy/response"
)

const (
	paramNameURL      = "url"
	paramNameMethod   = "method"
	paramPrefixHeader = "header-"
)

// URLHandler url handler struct
type URLHandler struct {
	url string
}

var _ handler.Handler = (*URLHandler)(nil)

// NewHandler create new JSON handler.
func NewHandler(url string) *URLHandler {
	return &URLHandler{url}
}

// Match checks if request to url argument should be handled by the handler.
func (h *URLHandler) Match(url string) bool {
	return h.url == url
}

// Handler return http.Handler
func (h *URLHandler) Handler(reqID uuid.UUID) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// get `Request` struct from url params
		req, err := prepareRequestURL(r)
		if err != nil {
			http.Error(rw, fmt.Sprintf("failed to prepare request: %s", err.Error()), 500)

			return
		}

		// send request
		httpResp, err := request.Send(req)
		if err != nil {
			http.Error(rw, fmt.Sprintf("failed to process request: %s", err.Error()), 500)

			return
		}

		// prepare response, write it to the client
		response.WriteResponse(rw, reqID, httpResp)
	})
}

func prepareRequestURL(r *http.Request) (req request.Request, err error) {
	if r == nil {
		return req, http.ErrBodyNotAllowed
	}

	req = request.New()

	// iterate over url params, fill `req` struct
	for paramName, values := range r.URL.Query() {
		switch {
		case paramName == paramNameURL:
			req.URL = strings.Join(values, ",")

		case paramName == paramNameMethod:
			req.Method = strings.Join(values, ",")

		case strings.HasPrefix(paramName, paramPrefixHeader):
			header := paramName[len(paramPrefixHeader):]
			req.Headers[header] = strings.Join(values, ",")
		}
	}

	if err := request.Validate(req); err != nil {
		return req, fmt.Errorf("request not valid: %w", err)
	}

	return req, nil
}
