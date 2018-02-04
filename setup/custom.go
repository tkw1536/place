package main

import (
	"github.com/tkw1536/place/config"
	myconf "github.com/tkw1536/place/setup/config"
	"net/http"
)

func savePlainConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Only PUT and POST", http.StatusMethodNotAllowed)
		return
	}

	conf := config.Config{}
	err := conf.Read(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = conf.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myconf.Set(&conf)
	err = myconf.WriteToPath(configPath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write([]byte("Config saved."))
	}
}
