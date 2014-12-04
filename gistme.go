package gistme

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// API endpoint
const (
	API_ENDPOINT = "https://api.github.com/gists"
)

// A Gist is an HTTP client which also represents a Gist.
type Gist struct {
	//ua *http.Client
	*http.Client
	Description string                       `json:"description"`
	Public      bool                         `json:"public"`
	Files       map[string]map[string]string `json:"files"`

	token string
}

// NewGist function creates a new gist client
func NewGist(tok string) Gist {
	g := new(Gist)

	//g.ua = &http.Client{}
	g.Client = &http.Client{}
	g.token = tok

	return *g
}

// All method lists all gists from your account
func (g Gist) All() string {
	req, err := http.NewRequest("GET", API_ENDPOINT, nil)
	req.Header.Add("Basic", g.token+":x-oauth-basic")

	resp, err := g.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	gists, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(gists)
}

// TODO -- 2014-11-30 19:34 UTC --JPN --
// add support for multiple files inside the same gist

// FIXME -- 2014-12-04 00:23 UTC --JPN --
// Create method should be returning a Gist type.

// Create method creates a new gist
func (g Gist) Create(name string, desc string, is_private bool, content string) string {

	gist_alpha := Gist{
		Description: desc,
		Public:      is_private,
		Files:       map[string]map[string]string{name: map[string]string{"content": content}},
	}
	b, err := json.Marshal(gist_alpha)
	if err != nil {
		log.Fatal("error:", err)
	}

	req, err := http.NewRequest("POST", API_ENDPOINT, strings.NewReader(string(b)))
	req.SetBasicAuth(g.token, "x-oauth-basic")

	resp, err := g.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	createdGist, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var x map[string]interface{}
	json.Unmarshal(createdGist, &x)

	return x["html_url"].(string)
}
