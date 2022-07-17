package proxy

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tsrkalexandr/go-proxy/config"
	"github.com/tsrkalexandr/go-proxy/handler"
	"github.com/tsrkalexandr/go-proxy/handler/json"
	"github.com/tsrkalexandr/go-proxy/handler/url"
	"github.com/tsrkalexandr/go-proxy/storage"
)

type Server struct {
	addr     string
	logging  bool
	handlers []handler.Handler
}

// NewServer create new server object, init proper handlers.
func NewServer(cfg config.Config, store *storage.Storage) *Server {
	return &Server{
		addr:    fmt.Sprintf(":%d", cfg.Port),
		logging: true,
		handlers: []handler.Handler{
			json.NewHandler(cfg.JSONParamPath, store),
			url.NewHandler(cfg.URLParamPath),
		},
	}
}

// Start start request listener.
func (s *Server) Start() error {
	// create handler
	h := http.HandlerFunc(s.handlerFunc)

	// add logging wrapper if needed
	if s.logging {
		h = handler.NewLoggerWrapper(h)
	}

	fmt.Printf("start listening on http://[%s]\n", s.addr)

	// start listening for requests
	if err := http.ListenAndServe(s.addr, h); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// handlerFunc main handler function, process incoming requests.
func (s *Server) handlerFunc(rw http.ResponseWriter, r *http.Request) {
	// generate new request ID
	reqID := uuid.New()

	// find proper handler for current request URL
	var h http.Handler
	for _, hand := range s.handlers {
		if hand.Match(r.URL.Path) {
			h = hand.Handler(reqID)
		}
	}

	// return error if handler not found
	if h == nil {
		http.Error(rw, "no required handlers", 500)

		return
	}

	// handler request by founded handler
	h.ServeHTTP(rw, r)
}
