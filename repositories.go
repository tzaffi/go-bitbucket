package bitbucket

import (
	"fmt"
)

type Repositories struct {
	c                  *Client
	PullRequests       *PullRequests
	Repository         *Repository
	Commits            *Commits
	Diff               *Diff
	BranchRestrictions *BranchRestrictions
	Webhooks           *Webhooks
	repositories
}

func (r *Repositories) ListForAccount(ro *RepositoriesOptions) interface{} {
	url := r.c.requestUrl("/repositories/%s", ro.Owner)
	if ro.Role != "" {
		url += "?role=" + ro.Role
	}
	return r.c.execute("GET", url, "")
}

func (r *Repositories) ListForTeam(ro *RepositoriesOptions, pages ...uint) interface{} {
	fmt.Println("Inside ListForTeam")
	fmt.Printf("pages = %v\n", pages)
	if(pages == nil){
		pages = []uint{1}
	}
	if(len(pages) == 1) {
		fmt.Printf("pages = %v\n", pages)
		url := r.c.requestUrl("/repositories/%s?page=%d", ro.Owner, pages[0])
		if ro.Role != "" {
			url += "?role=" + ro.Role
		}
		return r.c.execute("GET", url, "")
	} else {
		first := pages[0]
		last := pages[1]-1
		result := make([]interface{}, last-first+1)
		for page := first; page <= last; page++ {
			result[page-first] = r.ListForTeam(ro, []uint{page}...)
		}
		return result
	}
}

func (r *Repositories) ListPublic() interface{} {
	url := r.c.requestUrl("/repositories/", "")
	return r.c.execute("GET", url, "")
}
