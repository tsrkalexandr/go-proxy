package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Response struct to hold internal response data.
// Contains information which sends to client.
type Response struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Length  int64             `json:"length"`
	Headers map[string]string `json:"headers"`
}

// prepareResponse prepare `Response` object from `http.response`
func prepareResponse(id uuid.UUID, r *http.Response) (resp Response, err error) {
	if r == nil {
		return resp, errors.New("empty response")
	}

	resp.ID = id.String()
	resp.Status = r.StatusCode
	resp.Length = r.ContentLength
	resp.Headers = make(map[string]string)

	for header, value := range r.Header {
		resp.Headers[header] = strings.Join(value, ",")
	}

	return resp, nil
}

// WriteResponse prepare response data, write to client.
func WriteResponse(rw http.ResponseWriter, reqID uuid.UUID, httpResp *http.Response) {
	// prepare `Response` object from `http.response`
	resp, err := prepareResponse(reqID, httpResp)
	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to prepare response: %s", err.Error()), 500)

		return
	}

	// convert `Response` to json string
	bytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to prepare json: %s", err.Error()), 500)

		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.Write(bytes)
}
