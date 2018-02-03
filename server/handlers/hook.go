package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"../../utils"
	"../../utils/command"
	"../config"
)

// HookHandler implements a handler for Git-like-webhooks
type HookHandler struct {
	lock *sync.Mutex
	cfg  *config.Config
}

// NewHookHandler creates a new HookHandler
func NewHookHandler(cfg *config.Config) HookHandler {
	var lock sync.Mutex

	var handler HookHandler
	handler.lock = &lock
	handler.cfg = cfg
	return handler
}

func (hh HookHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, c := range hh.cfg.Checkers {
		name := c.String()
		if err := c.Check(req); err != nil {
			utils.Logger.Printf("%s checker failed: %s\n", name, err.Error())
		} else {
			utils.Logger.Printf("%s checker succeeded\n", name)
			fmt.Fprintf(res, "Success\n")
			go hh.runHook()
			return
		}
	}
	http.Error(res, "Failure\n", 500)
}

func (hh HookHandler) runHook() {
	utils.Logger.Println("queuing hook")

	// we only run one hook at a time
	hh.lock.Lock()
	defer hh.lock.Unlock()

	utils.Logger.Println("running hook")

	_, err := command.WithTimeout(hh.cfg.ScriptTimeout, hh.cfg.ScriptCommand...)

	// error handling
	if err != nil {
		utils.Logger.Printf("hook failed to run: %s\n", err.Error())
	} else {
		utils.Logger.Print("hook finished")
	}
}
