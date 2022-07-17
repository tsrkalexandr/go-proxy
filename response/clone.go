package response

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/tsrkalexandr/go-proxy/request"
)

// Clone create `http.Response` clone object.
func Clone(r *http.Response) *http.Response {
	clone := new(http.Response)
	*clone = *r

	clone.Status = r.Status
	clone.StatusCode = r.StatusCode
	clone.Proto = r.Proto
	clone.ProtoMajor = r.ProtoMajor
	clone.ProtoMinor = r.ProtoMinor
	clone.ContentLength = r.ContentLength
	clone.Close = r.Close
	clone.Uncompressed = r.Uncompressed

	clone.Request = request.Clone(r.Request)

	if r.Header != nil {
		clone.Header = r.Header.Clone()
	}

	if r.Trailer != nil {
		clone.Trailer = r.Trailer.Clone()
	}

	if r.TransferEncoding != nil {
		s2 := make([]string, len(r.TransferEncoding))
		copy(s2, r.TransferEncoding)

		clone.TransferEncoding = s2
	}

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
