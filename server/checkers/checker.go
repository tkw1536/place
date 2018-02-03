package checkers

import (
	"fmt"
	"net/http"
)

// Checker represents an object which validates if a request should trigger a webh
type Checker interface {
	String() string
	Check(req *http.Request) error
}

// checks that the POST method is used or returns on error
func checkPOSTMethod(req *http.Request) error {
	if req.Method != "POST" {
		return fmt.Errorf("request method should be \"POST\", not %q", req.Method)
	}

	return nil
}
