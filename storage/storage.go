package storage

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

// Storage struct which holds user request and response.
type Storage struct {
	requests  map[uuid.UUID]*http.Request
	responses map[uuid.UUID]*http.Response

	mReq, mResp sync.RWMutex
}

// NewStorage create new Storage
func NewStorage() *Storage {
	return &Storage{
		requests:  make(map[uuid.UUID]*http.Request),
		responses: make(map[uuid.UUID]*http.Response),
		mReq:      sync.RWMutex{},
		mResp:     sync.RWMutex{},
	}
}

// SaveRequest add request to storage, overwrite if exist.
func (s *Storage) SaveRequest(id uuid.UUID, request *http.Request) {
	s.mReq.Lock()
	defer s.mReq.Unlock()

	s.requests[id] = request
}

// SaveResponse add response to storage, overwrite if exist.
func (s *Storage) SaveResponse(id uuid.UUID, response *http.Response) {
	s.mResp.Lock()
	defer s.mResp.Unlock()

	s.responses[id] = response
}

// GetRequest return request by ID.
func (s *Storage) GetRequest(id uuid.UUID) *http.Request {
	s.mReq.RLock()
	defer s.mReq.RUnlock()

	return s.requests[id]
}

// GetResponse return response by ID.
func (s *Storage) GetResponse(id uuid.UUID) *http.Response {
	s.mResp.RLock()
	defer s.mResp.RUnlock()

	return s.responses[id]
}

// Save add request and response to storage.
func (s *Storage) Save(id uuid.UUID, request *http.Request, response *http.Response) {
	s.mReq.Lock()
	s.mResp.Lock()
	defer s.mReq.Unlock()
	defer s.mResp.Unlock()

	s.requests[id] = request
	s.responses[id] = response
}

// Get return request and response by ID.
func (s *Storage) Get(id uuid.UUID) (*http.Request, *http.Response) {
	s.mResp.RLock()
	s.mReq.RLock()
	defer s.mReq.RUnlock()
	defer s.mResp.RUnlock()

	return s.requests[id], s.responses[id]
}
