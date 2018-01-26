package utils

import (
	"net/http"
	"fmt"
)

const headerXGitLabEvent string = "X-Gitlab-Event"
const headerXGitLabEventValue string = "Push Hook"

const headerXGitLabToken string = "X-Gitlab-Token"

// check if a request is a valid GitLab request
func IsValidGitLabRequest(req *http.Request, secret string) error {

	// method has to be post
	if req.Method != "POST" {
		return fmt.Errorf("request method should be \"POST\", not %q", req.Method)
	}

	// X-GitLab-Event Header has to be for push hook
	if req.Header.Get(headerXGitLabEvent) != headerXGitLabEventValue {
		return fmt.Errorf("%s header should be %q", headerXGitLabEvent, headerXGitLabEventValue)
	}

	// X-GitLab-Token should be as intended
	token := req.Header.Get(headerXGitLabToken)
	if token != secret {
		return fmt.Errorf("%s header should be %q, not %q", headerXGitLabToken, secret, token)
	}

	// alright alright alright
	return nil
}