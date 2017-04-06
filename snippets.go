package bitbucket

import (
	"fmt"
)

type Snippets struct {
	c                  *Client
	PullRequests       *PullRequests
	Snippet            *Snippet
	Commits            *Commits
	Diff               *Diff
	BranchRestrictions *BranchRestrictions
	Webhooks           *Webhooks
	snippets
}

func (r *Snippets) ListForAccount(ro *SnippetsOptions) interface{} {
	url := r.c.requestUrl("/snippets/%s", ro.Owner)
	if ro.Role != "" {
		url += "?role=" + ro.Role
	}
	return r.c.execute("GET", url, "")
}

func (r *Snippets) ListForTeam(ro *SnippetsOptions, pages ...uint) interface{} {
	fmt.Println("Inside ListForTeam")
	fmt.Printf("pages = %v\n", pages)
	if pages == nil {
		pages = []uint{1}
	}
	if len(pages) == 1 {
		fmt.Printf("pages = %v\n", pages)
		url := r.c.requestUrl("/snippets/%s?page=%d", ro.Owner, pages[0])
		if ro.Role != "" {
			url += "?role=" + ro.Role
		}
		return r.c.execute("GET", url, "")
	} else {
		first := pages[0]
		last := pages[1] - 1
		result := make([]interface{}, last-first+1)
		for page := first; page <= last; page++ {
			result[page-first] = r.ListForTeam(ro, []uint{page}...)
		}
		return result
	}
}

func (r *Snippets) ListPublic() interface{} {
	url := r.c.requestUrl("/snippets/", "")
	return r.c.execute("GET", url, "")
}
