package server

import (
	"net/http"
	"net/http/httputil"

	"./config"
	"./handlers"
)

// StartServer starts a server with the given configuration
func StartServer(cfg *config.Config) error {

	r := http.NewServeMux()
	r.Handle(cfg.HookPath, handlers.NewHookHandler(cfg))

	if cfg.StaticPath != "" {
		r.Handle("/", http.FileServer(http.Dir(cfg.StaticPath)))
	} else {
		r.Handle("/", httputil.NewSingleHostReverseProxy(cfg.ProxyURL))
	}

	http.ListenAndServe(cfg.BindAddress, handlers.NewLoggingHandler(cfg.Logger, r))
	return nil
}
