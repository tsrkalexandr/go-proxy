package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// Clone create `http.Request` clone object.
func Clone(r *http.Request) *http.Request {
	clone := r.Clone(context.Background())

	if r.Body != nil {
		// read all from request body.
		contents, err := io.ReadAll(r.Body)
		if err != nil {
			// skip return error, just log it
			// assume that contents is empty in case of read error
			fmt.Printf("failed to read request body on clone: %s", err)
		}

		clone.Body = io.NopCloser(bytes.NewReader(contents))

		// replace original request body with buffered one
		r.Body = io.NopCloser(bytes.NewReader(contents))
	}

	return clone
}
