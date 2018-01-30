package checkers

import (
	"fmt"
	"net/http"
)

// Checker represents an object which validates if a request should trigger a webh
type Checker interface {
	Create(param string) error
	String() string
	Check(req *http.Request) error
}

// CreateChecker creates a checker of the given name
// and the given parameter
func CreateChecker(name string, param string) Checker {
	var checker Checker
	switch name {
	case "github":
		checker = GitHubChecker{}
	case "gitlab":
		checker = GitLabChecker{}
	case "debug":
		checker = DebugChecker{}
	default:
		return nil
	}

	checker.Create(param)
	return checker
}

// checks that the POST method is used or returns on error
func checkPOSTMethod(req *http.Request) error {
	if req.Method != "POST" {
		return fmt.Errorf("request method should be \"POST\", not %q", req.Method)
	}

	return nil
}
