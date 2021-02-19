package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	health *Health
	mu     sync.Mutex
)

type Health struct {
	healthOn bool
	mu       *sync.RWMutex
}

func GetHealth() *Health {
	mu.Lock()
	if health == nil {
		health = newHealth()
	}
	mu.Unlock()
	return health
}

func newHealth() *Health {
	return &Health{
		healthOn: false,
		mu:       new(sync.RWMutex),
	}
}

// SetReady will set the application ready to serve traffic
func (h *Health) SetReady() {
	h.mu.Lock()
	h.healthOn = true
	h.mu.Unlock()
}

// UnsetReady will set the application as not ready to serve traffic
func (h *Health) UnsetReady() {
	h.mu.Lock()
	h.healthOn = false
	h.mu.Unlock()
}

func (h *Health) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.mu.RLock()
	health := h.healthOn
	h.mu.RUnlock()
	// log.Message(log.Info, "Recived health request")
	if health {
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

type TestHandler struct{}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := mux.NewRouter()
	h2s := &http2.Server{}
	port := "9001"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(router, h2s),
	}
	GetHealth().SetReady()
	router.Handle("/{ueId}/routing-info", GetHealth())
	// router.Handle("/", &TestHandler{})
	router.NotFoundHandler = &TestHandler{}

	server.ListenAndServe()
}
