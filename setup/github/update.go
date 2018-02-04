package github

import (
	"context"
	"encoding/json"
	github "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
)

var repoData []*github.Repository
var currentRepo *github.Repository
var webhookSecret string

func getAuthClient(token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	return tc
}

func writeState(w http.ResponseWriter, state *hookingState) {
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err := enc.Encode(state)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func readState(r *http.Request) (state *hookingState, err error) {
	state = new(hookingState)

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(state)

	if err != nil {
		return nil, err
	}

	return state, nil
}

func UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	new, err := readState(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeState(w, &currently)
		return
	}

	if new.AccessToken != currently.AccessToken {
		client := github.NewClient(getAuthClient(new.AccessToken))

		// Validate the access token by grabbing a list of repos for the current user
		repos, _, err := client.Repositories.List(context.Background(), "", nil)

		if err == nil {
			// The token is fine, so we pull all the repos and update our state accordingly
			currently.AccessToken = new.AccessToken
			currently.Repos = make([]string, len(repos))

			repoData = repos

			for idx, repo := range repos {
				if repo.HTMLURL != nil {
					currently.Repos[idx] = *repo.HTMLURL
				}
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if new.SelectedRepo != currently.SelectedRepo && new.SelectedRepo != "" {

		for _, repo := range repoData {
			if repo.HTMLURL == nil || repo.CloneURL == nil {
				continue
			}

			if *repo.HTMLURL == new.SelectedRepo {
				currently.SelectedRepo = *repo.CloneURL
				currentRepo = repo
			}
		}

		if currently.SelectedRepo != "" {
			currently.Branches = []string{currentRepo.GetDefaultBranch()}
			currently.SelectedBranch = currentRepo.GetDefaultBranch()
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	if currently.SelectedRepo != "" && currently.SelectedBranch != "" {
		webhookURL := &url.URL{
			Host:   r.Host,
			Scheme: "https",
			Path:   "/webhook",
		}
		log.Println(webhookURL)

		// Try and hook up a webhook
		/*
			secret, err := uuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}


			client := github.NewClient(getAuthClient(currently.AccessToken))

			hookURL := currentRepo.GetHooksURL()
		*/
	}

	writeState(w, &currently)
}
