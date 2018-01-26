package utils

// adapted from https://github.com/GitbookIO/go-github-webhook/blob/master/handler.go

import (
	"strings"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/http"
	"io/ioutil"
)

const headerXGitHubEvent string = "X-GitHub-Event"
const headerXGitHubEventValue string = "Push"

const headerXHubSignature string = "X-Hub-Signature"

// check if the request is a valid request for GitHub
func IsValidGithubRequest(req *http.Request, secret string) error {

	// method has to be post
	if req.Method != "POST" {
		return fmt.Errorf("request method should be \"POST\", not %q", req.Method)
	}

	// X-GitHub-Event should be for push
	if req.Header.Get(headerXGitHubEvent) != headerXGitHubEventValue {
		return fmt.Errorf("%s header should be %q", headerXGitHubEvent, headerXGitHubEventValue)
	}

	// body should be readable
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	// signature header should be type=hash
	signature := req.Header.Get(headerXHubSignature)
	signatureParts := strings.SplitN(req.Header.Get(headerXHubSignature), "=", 2)
	if len(signatureParts) != 2 {
		return fmt.Errorf("%s header should be of the form \"<type>=<hash>\", not %q", headerXHubSignature, signature)
	}

	// signature type should be sha1
	signatureType := signatureParts[0]
	if signatureType != "sha1" {
		return fmt.Errorf("%s header signature type should be \"sha1\", not %q", headerXHubSignature, signatureType)
	}

	// signature should match
	signatureHash := signatureParts[1]
	expected := hashPayload(body, secret)
	if !hmac.Equal( []byte(signatureHash), []byte(expected)) {
		return 	fmt.Errorf("%s header signature hash should be %q, not %q", headerXHubSignature, expected, signatureHash)
	}

	// alright alright alright
	return nil
}

// performs the hashing of a given body
func hashPayload(body []byte, secret string) string {
	hm := hmac.New(sha1.New, []byte(secret))
	hm.Write(body)
	sum := hm.Sum(nil)
	return fmt.Sprintf("%x", sum)
}