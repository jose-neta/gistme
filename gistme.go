package gistme

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// FIXME -- 2014-11-30 12:33 UTC --JPN --
// move the token out of here, for example keep it inside an ENV var.
const (
	TOKEN        = "<PUT YOUR TOKEN HERE>"
	API_ENDPOINT = "https://api.github.com/gists"
)

type Gist struct {
	//ua *http.Client
	*http.Client
	Description string                       `json:"description"`
	Public      bool                         `json:"public"`
	Files       map[string]map[string]string `json:"files"`
}

//
func NewGist() Gist {
	g := new(Gist)

	//g.ua = &http.Client{}
	g.Client = &http.Client{}

	return *g
}

type X interface{}

func (g Gist) All() string {
	req, err := http.NewRequest("GET", API_ENDPOINT, nil)
	req.Header.Add("Basic", TOKEN+":x-oauth-basic")

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

// Create creates a new gist
// TODO -- 2014-11-30 19:34 UTC --JPN --
// add support for multiple files inside the gist
// XXX Create only creates anonymous gists. However when you do a simple
// curl -u <YOUR TOKEN HERE>:x-oauth-basic -XPOST https://api.github.com/gists -d@/tmp/data
// the gist is as expected created in your account. So I'm probably doing
// something wrong.
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
	req.Header.Add("Basic", TOKEN+":x-oauth-basic")

	resp, err := g.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	createdGist, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var x map[string]X
	json.Unmarshal(createdGist, &x)

	return x["html_url"].(string)
}
