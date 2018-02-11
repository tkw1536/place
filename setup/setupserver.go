package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tkw1536/place/config"
	"github.com/tkw1536/place/setup/auth"
	"github.com/tkw1536/place/utils"
	"github.com/tkw1536/place/utils/verify"
)

// jsonp writes a JSONP response to writer
func jsonp(json string, w http.ResponseWriter, r *http.Request) {
	callback := r.FormValue("callback")
	// if we have a c
	if callback != "" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		fmt.Fprintf(w, "%s(%s)", callback, json)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(json))
	}
}

func toJSONString(success bool, fields []verify.FieldError, message string) string {
	var fieldNames []string
	for _, field := range fields {
		fieldNames = append(fieldNames, field.Field())
	}

	fieldsBytes, _ := json.Marshal(fieldNames)
	messageBytes, _ := json.Marshal(message)
	return fmt.Sprintf("{\"success\": %t, \"fields\": %s, \"message\": %s}", success, string(fieldsBytes), string(messageBytes))
}

func validPlainConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Only PUT and POST", http.StatusMethodNotAllowed)
		return
	}

	conf := config.Config{}
	err := conf.Read(r.Body)
	if err != nil {
		jsonp(toJSONString(false, []verify.FieldError{}, err.Error()), w, r)
		return
	}

	ferr := conf.Validate()
	if ferr != nil {
		jsonp(toJSONString(false, ferr.Messages(), ferr.Error()), w, r)
		return
	}

	jsonp(toJSONString(true, []verify.FieldError{}, ""), w, r)
}

func savePlainConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Only PUT and POST", http.StatusMethodNotAllowed)
		return
	}

	conf := config.Config{}
	err := conf.Read(r.Body)
	if err != nil {
		jsonp(toJSONString(false, []verify.FieldError{}, err.Error()), w, r)
		return
	}

	ferr := conf.Validate()
	if ferr != nil {
		jsonp(toJSONString(false, ferr.Messages(), ferr.Error()), w, r)
		return
	}

	serr := conf.Save(configPath)

	if serr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		jsonp(toJSONString(true, []verify.FieldError{}, "Configuration saved"), w, r)
	}
}

func setupServerMux(st chan<- bool) *http.ServeMux {
	mux := http.NewServeMux()

	// save a configuration file
	mux.HandleFunc("/setup/save", savePlainConfigHandler)

	// save a configuration file
	mux.HandleFunc("/setup/validate", validPlainConfigHandler)

	// shutdown the server
	mux.HandleFunc("/setup/finish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST", http.StatusMethodNotAllowed)
			return
		}
		w.(http.Flusher).Flush()

		st <- true
	})

	// static everything else
	staticFiles := http.FileServer(http.Dir("static"))
	mux.Handle("/setup/", http.StripPrefix("/setup/", staticFiles))

	return mux
}

// runs the setup server
func setupServer() {
	// a channel to shutdown the server in
	st := make(chan bool)

	// the server that will handle all the connections
	srv := &http.Server{
		Addr:         listenAddr,
		WriteTimeout: time.Second,
		Handler: &auth.StickyUser{
			Next: setupServerMux(st),
		},
	}

	// go listen and server stuff
	go func() {
		utils.Logger.Printf("Starting setup server at %s\n", listenAddr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			utils.Logger.Printf("http server error: %v", err)
			os.Exit(1)
		}
	}()

	// and wait for shutdown
	<-st
	utils.Logger.Printf("Received finalize event, restarting server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		utils.Logger.Printf("http server error: %v", err)
		os.Exit(1)
	}
}
