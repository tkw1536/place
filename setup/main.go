package main

import (
	"bytes"
	"encoding/json"
	"github.com/tkw1536/place/setup/auth"
	"github.com/tkw1536/place/setup/config"
	"github.com/tkw1536/place/setup/github"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	listenAddr = "0.0.0.0:9001"
	setUpDone  = false
	configPath = config.DefaultPath
	configured bool
)

func init() {
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	err := config.ReadFromPath(configPath)
	if err != nil {
		log.Printf("Failed to load config file from %v: %v", configPath, err)
	}

	// If the config could be loaded, we are already configured.
	configured = err == nil
}

func rootHander(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func dumpConfig(w http.ResponseWriter, r *http.Request) {
	buf := bytes.Buffer{}

	enc := json.NewEncoder(&buf)
	err := enc.Encode(config.Get())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = buf.WriteTo(w)
}

func writeConfig() {
	err := config.WriteToPath(configPath)
	if err != nil {
		log.Printf("failed to write config: %v", err)
	}
}

func runChild(cmd string, args []string) int {
	log.Printf("Running child process %v with args %v", cmd, args)

	child := exec.Command(cmd, args...)

	// Loop through the output of the child process
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	child.Stdin = os.Stdin

	err := child.Run()

	if err == nil {
		return 0
	}

	log.Printf("Child process failed to exit cleanly: %v", err)
	return 1
}

func main() {
	if configured {
		ret := runChild(os.Args[1], os.Args[2:])
		os.Exit(ret)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/gh", github.UpdateStateHandler)
	mux.HandleFunc("/raw", savePlainConfigHandler)

	mux.HandleFunc("/dump", dumpConfig)

	staticFiles := http.FileServer(http.Dir("static"))
	mux.Handle("/", http.StripPrefix("", staticFiles))

	var stickyUser = auth.StickyUser{
		Next: mux,
	}

	err := http.ListenAndServe(listenAddr, &stickyUser)
	if err != nil {
		log.Printf("http server error: %v", err)
	}
}
