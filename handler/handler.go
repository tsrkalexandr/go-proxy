package handler

import (
	"net/http"

	"github.com/google/uuid"
)

// Handler request handler interface
type Handler interface {
	Match(url string) bool
	Handler(reqID uuid.UUID) http.Handler
}
