package gistme

import (
	"fmt"
	"strings"
)

// Invoking the gist client
func ExampleGist() {
	x := NewGist("<your token here>")
	fmt.Printf("%v", x)
}

// NOTE -- 2014-12-04 03:07 UTC --JPN --
// The test will not run for ExampleGist_All because I've altered the name
// 'output:' to 'The output:'. This was done on purpose.

// Listing all your gists
func ExampleGist_All() {
	x := NewGist("<your token here>")
	fmt.Println(x.All())

	// The output:
	//  [
	//      {
	//          "url": "https://api.github.com/gists/ef9436173feb965ee019",
	//          "forks_url": "https://api.github.com/gists/ef9436173feb965ee019/forks",
	//          "commits_url": "https://api.github.com/gists/ef9436173feb965ee019/commits",
	//          "id": "ef9436173feb965ee019",
	//          "git_pull_url": "https://gist.github.com/ef9436173feb965ee019.git",
	//          "git_push_url": "https://gist.github.com/ef9436173feb965ee019.git",
	//          "html_url": "https://gist.github.com/ef9436173feb965ee019",
	//          "files": {
	//              "untrusted-lvl6-solution.js": {
	//                  "filename": "untrusted-lvl6-solution.js",
	//                  "type": "application/javascript",
	//                  "language": "JavaScript",
	//                  "raw_url": "https://gist.githubusercontent.com/anonymous/ef9436173feb965ee019/raw/3589b17b5d8a898d626ece9b33ca87f58df99d2e/untrusted-lvl6-solution.js",
	//                  "size": 1974
	//              }
	//          },
	//          "public": true,
	//          "created_at": "2014-12-04T00:11:03Z",
	//          "updated_at": "2014-12-04T00:11:03Z",
	//          "description": "Solution to level 6 in Untrusted: http://alex.nisnevich.com/untrusted/",
	//          "comments": 0,
	//          "user": null,
	//          "comments_url": "https://api.github.com/gists/ef9436173feb965ee019/comments"
	//      },
	//      {
	//          "url": "https://api.github.com/gists/19048d27576140f3d400",
	//          "forks_url": "https://api.github.com/gists/19048d27576140f3d400/forks",
	//          "commits_url": "https://api.github.com/gists/19048d27576140f3d400/commits",
	//          "id": "19048d27576140f3d400",
	//          "git_pull_url": "https://gist.github.com/19048d27576140f3d400.git",
	//          "git_push_url": "https://gist.github.com/19048d27576140f3d400.git",
	//          "html_url": "https://gist.github.com/19048d27576140f3d400",
	//          "files": {
	//              "gistfile1.md": {
	//                  "filename": "gistfile1.md",
	//                  "type": "text/plain",
	//                  "language": "Markdown",
	//                  "raw_url": "https://gist.githubusercontent.com/asdayhf/19048d27576140f3d400/raw/46c6f1a35c263a77ac81c25629ab34e283ce6d1b/gistfile1.md",
	//                  "size": 762
	//              }
	//          },
	//          "public": true,
	//          "created_at": "2014-12-04T00:11:01Z",
	//          "updated_at": "2014-12-04T00:11:02Z",
	//          "description": "",
	//          "comments": 0,
	//          "user": null,
	//          "comments_url": "https://api.github.com/gists/19048d27576140f3d400/comments",
	//          "owner": {
	//              "login": "asdayhf",
	//              "id": 10068136,
	//              "avatar_url": "https://avatars.githubusercontent.com/u/10068136?v=3",
	//              "gravatar_id": "",
	//              "url": "https://api.github.com/users/asdayhf",
	//              "html_url": "https://github.com/asdayhf",
	//              "followers_url": "https://api.github.com/users/asdayhf/followers",
	//              "following_url": "https://api.github.com/users/asdayhf/following{/other_user}",
	//              "gists_url": "https://api.github.com/users/asdayhf/gists{/gist_id}",
	//              "starred_url": "https://api.github.com/users/asdayhf/starred{/owner}{/repo}",
	//              "subscriptions_url": "https://api.github.com/users/asdayhf/subscriptions",
	//              "organizations_url": "https://api.github.com/users/asdayhf/orgs",
	//              "repos_url": "https://api.github.com/users/asdayhf/repos",
	//              "events_url": "https://api.github.com/users/asdayhf/events{/privacy}",
	//              "received_events_url": "https://api.github.com/users/asdayhf/received_events",
	//              "type": "User",
	//              "site_admin": false
	//          }
	//      }
	//  ]

}

// NOTE -- 2014-12-04 03:10 UTC --JPN --
// once again if you run the tests go test -v ./... you will not run them
// for ExampleGist_Create as well, because we are not defining the output.

// Creating a new gist
func ExampleGist_Create() {
	x := NewGist("<your token here>")

	content := []string{"line one", "line ...", "line n"}
	res := x.Create("my gist name", "my gist description", true, strings.Join(content, "\n"))
	fmt.Println(res)

	// The output:
	// https://gist.github.com/28d51bc36acdac7610a3
}
