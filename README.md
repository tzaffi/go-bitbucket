# go-bitbucket

> Bitbucket-API library for golang.

Support Bitbucket API v2.0. 

And the response type is json format defined Bitbucket API.

- Bitbucket API v2.0 <https://developer.atlassian.com/bitbucket/api/2/reference/>
- Swagger for API v2.0 <https://api.bitbucket.org/swagger.json>

## Install

```sh
go get github.com/ktrysmt/go-bitbucket
```

## How to use

```go
import "github.com/ktrysmt/go-bitbucket"
```

## Godoc

- <http://godoc.org/github.com/ktrysmt/go-bitbucket>


## Example

```go
package main

import (
        "github.com/ktrysmt/go-bitbucket" 
        "fmt"
)

func main() {

        c := bitbucket.NewBasicAuth("username", "password")

        opt := &bitbucket.PullRequestsOptions{
                Owner:      "your-team",
                Repo_slug:  "awesome-project",
                Source_branch: "develop",
                Destination_branch: "master",
                Title: "fix bug. #9999",
                Close_source_branch: true,
        }
        res := c.Repositories.PullRequests.Create(opt)

        fmt.Println(res) // receive the data as json format
}
```

## FAQ

### Support Bitbucket API v1.0 ?

It does not correspond yet. Because there are many differences between v2.0 and v1.0.

- Bitbucket API v1.0 <https://confluence.atlassian.com/bitbucket/version-1-423626337.html>

It is officially recommended to use v2.0.  
But unfortunately Bitbucket Server (formerly: Stash) API is still v1.0.   
And The API v1.0 covers resources that the v2.0 API and API v2.0 is yet to cover.

## Development

### Install Dependencies

```sh
go get github.com/golang/dep/...
cd ./go-bitbucket
dep ensure 
```

### Testing

Set your available user account to Global Env.

```sh
export BITBUCKET_TEST_USERNAME=<your_username> 
export BITBUCKET_TEST_PASSWORD=<your_password> 
export BITBUCKET_TEST_OWNER=<your_repo_owner>  
export BITBUCKET_TEST_REPOSLUG=<your_repo_name>
```

Refs; URL Syntax is `https://<your_username>:<your_password>@bitbucket.org/<your_repo_owner>/<your_repo_name>.git`. 

And just run,

```sh
make test
```

## License

[Apache License 2.0](./LICENSE)

## Author

[ktrysmt](https://github.com/ktrysmt)
