package config

import (
	"encoding/json"
	"net/url"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
)

// MarshalableEndpoint represents a transport.Endpoint which can be marshaled and unmarshaled as a string
type MarshalableEndpoint transport.Endpoint

// MarshalJSON marshals this endpoint into a json string
func (me *MarshalableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.Endpoint().String())
}

// Endpoint returns this MarshalableEndpoint as a transport.Endpoint
func (me *MarshalableEndpoint) Endpoint() *transport.Endpoint {
	te := transport.Endpoint(*me)
	return &te
}

// UnmarshalJSON unmarshals this endpoint from a json string
func (me *MarshalableEndpoint) UnmarshalJSON(b []byte) error {

	// load a string
	var tes string
	if err := json.Unmarshal(b, &tes); err != nil {
		return err
	}

	// parse it
	tep, err := transport.NewEndpoint(tes)
	if err != nil {
		return err
	}

	// and turn it into an endpoint
	*me = MarshalableEndpoint(*tep)
	return nil
}

// IsEmpty checks if this endpoint has empty content
func (me *MarshalableEndpoint) IsEmpty() bool {
	if me == nil {
		return true
	}
	ep := me.Endpoint()
	return ep.Host == "" && ep.Path == "" && ep.Port == 0
}

// MarshalableURL represents a url.URL which can be marshaled and unmarshaled as a string
type MarshalableURL url.URL

// MarshalJSON marshals this url into a json string
func (mu *MarshalableURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(mu.URL().String())
}

// URL returns this MarshalableURL as a url.URL
func (mu *MarshalableURL) URL() *url.URL {
	ul := url.URL(*mu)
	return &ul
}

// UnmarshalJSON unmarshals this url from a json string
func (mu *MarshalableURL) UnmarshalJSON(b []byte) error {

	// load a string
	var urls string
	if err := json.Unmarshal(b, &urls); err != nil {
		return err
	}

	// parse it
	urlp, err := url.Parse(urls)
	if err != nil {
		return err
	}

	// and turn it into an url
	*mu = MarshalableURL(*urlp)
	return nil
}
