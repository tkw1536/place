package checkers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GitHubChecker runs the Hook if it receives a valid GitHub web hook
type GitHubChecker struct {
	token        string
	ref          string
	events       []string
	eventsString string // for printing only
}

// GitHub API event names
const (
	GitHubPushEvent = "push"
)

// Create a GitHubChecker instance from a parameter.
// It is of the form "token[,ref[,events...]]", with the later components being optional
func (gh GitHubChecker) Create(param string) error {
	var params = strings.Split(param, ",")

	// token
	if len(params) == 0 {
		return fmt.Errorf("GitHubChecker needs at least a token")
	}
	gh.token = params[0]

	// ref = refs/heads/master"
	if len(params) == 1 {
		gh.ref = "refs/heads/master"
	} else {
		gh.ref = params[1]
	}

	// events = [Push]
	if len(params) >= 2 {
		gh.events = params[2:]
	} else {
		gh.events = make([]string, 1)
		gh.events[0] = GitHubPushEvent
	}

	gh.eventsString = strings.Join(gh.events, ",")

	return nil
}

// Turn GitHubChecker into a string
func (gh GitHubChecker) String() string {
	return fmt.Sprintf("GitHubChecker{%q,%s,%q}", gh.token, gh.ref, gh.eventsString)
}

// the header names for Github
const (
	xGitHubEvent     string = "X-GitHub-Event"
	xGitHubSignature string = "X-Hub-Signature"
)

// Check that a valid GitHubRequest has been sent
func (gh GitHubChecker) Check(req *http.Request) error {
	if err := checkPOSTMethod(req); err != nil {
		return err
	}

	if err := gh.containsEvent(req.Header.Get(xGitHubEvent)); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return gh.validSignature(body, req.Header.Get(xGitHubSignature))
}

// containsEvent checks if a given event is contained within the events
// that should be listened to
func (gh GitHubChecker) containsEvent(event string) error {
	for _, e := range gh.events {
		if strings.EqualFold(event, e) {
			return nil
		}
	}

	return fmt.Errorf("%s header should be one of %s, not %q", xGitHubEvent, gh.eventsString, event)
}

// Checks if a given signature is valid
func (gh GitHubChecker) validSignature(payload []byte, signature string) error {
	expected := gh.hashPayload(payload)

	signatureParts := strings.SplitN(signature, "=", 2)
	if len(signatureParts) != 2 {
		return fmt.Errorf("%s header should be of the form \"<type>=<hash>\", not %q", xGitHubSignature, signature)
	}

	tp := signatureParts[0]
	hash := signatureParts[1]

	if tp != "sha1" {
		return fmt.Errorf("%s header signature type should be \"sha1\", not %q", xGitHubSignature, signature)
	}

	if !hmac.Equal([]byte(hash), []byte(expected)) {
		return fmt.Errorf("%s header signature hash should be %q, not %q", xGitHubSignature, expected, hash)
	}

	return nil
}

// Hashes a given Payload
func (gh GitHubChecker) hashPayload(payload []byte) string {
	hm := hmac.New(sha1.New, []byte(gh.token))
	hm.Write(payload)
	sum := hm.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

// checks if the right branch was pushed
func (gh GitHubChecker) validBranch(payload []byte) error {
	var event struct {
		Ref string `json:"ref"`
	}
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	if !strings.EqualFold(event.Ref, gh.ref) {
		return fmt.Errorf("expected %q field to be %q, not %q", "refs", gh.ref, event.Ref)
	}

	return nil
}
