package server

import (
	"log"
	"net/http"
)

// LoggingHandler is a handler that logs requests
type LoggingHandler struct {
	logger  *log.Logger
	handler http.Handler
}

// NewLoggingHandler creates a new LoggingHandler
func NewLoggingHandler(logger *log.Logger, h http.Handler) LoggingHandler {
	return LoggingHandler{logger, h}
}

func (l LoggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	l.logger.Printf("%s %s", req.Method, req.URL)
	l.handler.ServeHTTP(w, req)
}
