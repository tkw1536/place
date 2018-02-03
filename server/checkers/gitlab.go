package checkers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GitLabChecker runs the Hook if it receives a valid GitHub web hook
type GitLabChecker struct {
	token        string
	ref          string
	events       []string
	eventsString string // for printing only
}

// GitLab API event names
const (
	GitLabPushEvent = "Push Hook"
)

// Create a GitLabChecker instance from a parameter.
func (gl *GitLabChecker) Create(token string, ref string) {
	gl.token = token
	gl.ref = ref

	gl.events = []string{
		GitLabPushEvent,
	}
}

// Turn GitHubChecker into a string
func (gl *GitLabChecker) String() string {
	return fmt.Sprintf("GitLabChecker{%q,%s,%q}", gl.token, gl.ref, gl.eventsString)
}

// the header names for Github
const (
	xGitLabEvent string = "X-Gitlab-Event"
	xGitLabToken string = "X-Gitlab-Token"
)

// Check that a valid GitLab has been sent
func (gl *GitLabChecker) Check(req *http.Request) error {
	if err := checkPOSTMethod(req); err != nil {
		return err
	}

	if err := gl.containsEvent(req.Header.Get(xGitLabEvent)); err != nil {
		return err
	}

	if err := gl.validToken(req.Header.Get(xGitLabToken)); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return gl.validBranch(body)
}

// containsEvent checks if a given event is contained within the events
// that should be listened to
func (gl *GitLabChecker) containsEvent(event string) error {
	for _, e := range gl.events {
		if strings.EqualFold(event, e) {
			return nil
		}
	}

	return fmt.Errorf("%s header should be one of %s, not %q", xGitLabEvent, gl.eventsString, event)
}

// Checks if a given signature is valid
func (gl *GitLabChecker) validToken(token string) error {
	if gl.token != token {
		return fmt.Errorf("%s header should be %q, not %q", xGitLabToken, gl.token, token)
	}

	return nil
}

// checks if the right branch was pushed
func (gl *GitLabChecker) validBranch(payload []byte) error {
	var event struct {
		Ref string `json:"ref"`
	}
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	if !strings.EqualFold(event.Ref, gl.ref) {
		return fmt.Errorf("expected %q field to be %q, not %q", "refs", gl.ref, event.Ref)
	}

	return nil
}
