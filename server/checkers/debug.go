package checkers

import (
	"net/http"
)

// DebugChecker is a Checker that runs the hook on any POST request
type DebugChecker struct{}

// Create a new DebugChecker
func (dbg DebugChecker) Create() {
}

// Turn the Checker into a string
func (dbg DebugChecker) String() string {
	return "DebugChecker"
}

// Check if a request should be run
func (dbg DebugChecker) Check(req *http.Request) error {
	return checkPOSTMethod(req)
}
