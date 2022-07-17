package json

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tsrkalexandr/go-proxy/handler"
	"github.com/tsrkalexandr/go-proxy/request"
	"github.com/tsrkalexandr/go-proxy/response"
	"github.com/tsrkalexandr/go-proxy/storage"
)

// JSONHandler json handler struct
type JSONHandler struct {
	url   string
	store *storage.Storage
}

var _ handler.Handler = (*JSONHandler)(nil)

// NewHandler create new JSON handler.
func NewHandler(url string, store *storage.Storage) *JSONHandler {
	return &JSONHandler{
		url,
		store,
	}
}

// Match checks if user request to specific `url` argument should be handled by the JSONhandler.
func (h *JSONHandler) Match(url string) bool {
	return h.url == url
}

// Handler return http.Handler
func (h *JSONHandler) Handler(reqID uuid.UUID) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// save request copy to storage
		h.store.SaveRequest(reqID, request.Clone(r))

		// get `Request` struct from request json body
		req, err := prepareRequestJSON(r)
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

		// save response copy to storage
		h.store.SaveResponse(reqID, response.Clone(httpResp))

		// prepare response, write it to the client
		response.WriteResponse(rw, reqID, httpResp)
	})
}

// prepareRequestJSON parse request body.
func prepareRequestJSON(r *http.Request) (req request.Request, err error) {
	if r == nil {
		return req, http.ErrBodyNotAllowed
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		return req, fmt.Errorf("failed to parse request body: %w", err)
	}

	if err := request.Validate(req); err != nil {
		return req, fmt.Errorf("request not valid: %w", err)
	}

	return req, nil
}
