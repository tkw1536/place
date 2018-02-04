package github

import (
	"net/http"
)

// DefaultServeMux is pre-loaded with all the GitHub HTTP handlers
var DefaultServeMux = http.NewServeMux()

type hookingState struct {
	AccessToken    string   `json:"access_token"` // AccessToken as received from the user
	Repos          []string `json:"repos"`        // slice of "org/repo" formatted strings
	SelectedRepo   string   `json:"repo"`
	Branches       []string `json:"branches"`
	SelectedBranch string   `json:"branch"`
}

var currently hookingState
