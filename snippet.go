package bitbucket

import (
	"encoding/json"
	"os"

	"github.com/k0kubun/pp"
	"github.com/mitchellh/mapstructure"
)

type Snippet struct {
	c *Client

	Project     Project
	Slug        string
	Full_name   string
	Description string
	Fork_policy string
	Type        string
	Owner       map[string]interface{}
	Links       map[string]interface{}
}

func (r *Snippet) Create(ro *SnippetOptions) (Snippet, error) {
	data := r.buildSnippetBody(ro)
	url := r.c.requestUrl("/repositories/%s/%s", ro.Owner, ro.Repo_slug)
	response := r.c.execute("POST", url, data)

	return decodeSnippet(response)
}

func (r *Snippet) Get(ro *SnippetOptions) (Snippet, error) {
	url := r.c.requestUrl("/repositories/%s/%s", ro.Owner, ro.Repo_slug)
	response := r.c.execute("GET", url, "")

	return decodeSnippet(response)
}

func (r *Snippet) Delete(ro *SnippetOptions) interface{} {
	url := r.c.requestUrl("/repositories/%s/%s", ro.Owner, ro.Repo_slug)
	return r.c.execute("DELETE", url, "")
}

func (r *Snippet) ListWatchers(ro *SnippetOptions) interface{} {
	url := r.c.requestUrl("/repositories/%s/%s/watchers", ro.Owner, ro.Repo_slug)
	return r.c.execute("GET", url, "")
}

func (r *Snippet) ListForks(ro *SnippetOptions) interface{} {
	url := r.c.requestUrl("/repositories/%s/%s/forks", ro.Owner, ro.Repo_slug)
	return r.c.execute("GET", url, "")
}

func (r *Snippet) buildSnippetBody(ro *SnippetOptions) string {
	body := map[string]interface{}{}

	if ro.Scm != "" {
		body["scm"] = ro.Scm
	}
	if ro.Is_private != "" {
		body["is_private"] = ro.Is_private
	}
	if ro.Description != "" {
		body["description"] = ro.Description
	}
	if ro.Language != "" {
		body["language"] = ro.Language
	}
	
	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
		os.Exit(9)
	}

	return string(data)
}

func decodeSnippet(json interface{}) (Snippet, error) {
	jsonMap := json.(map[string]interface{})

	if jsonMap["type"] == "error" {
		return Snippet{}, DecodeError(jsonMap)
	}

	var snippet Snippet
	err := mapstructure.Decode(jsonMap, &snippet)
	if err != nil {
		return Snippet{}, err
	}
	return snippet, nil
}
