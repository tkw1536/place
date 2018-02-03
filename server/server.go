package server

import (
	"net/http"
	"net/http/httputil"

	"github.com/tkw1536/place/config"
	"github.com/tkw1536/place/utils"
)

// StartServer starts a server with the given configuration
func StartServer(cfg *config.Config) error {

	// create a hook handler
	r := http.NewServeMux()
	r.Handle(cfg.WebhookPath, NewHookHandler(cfg))

	if cfg.StaticPath != "" {
		r.Handle("/", http.FileServer(http.Dir(cfg.StaticPath)))
	} else {
		r.Handle("/", httputil.NewSingleHostReverseProxy(cfg.ProxyURL))
	}

	http.ListenAndServe(cfg.BindAddress, NewLoggingHandler(utils.Logger, r))
	return nil
}
