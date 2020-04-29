package github

import (
	"time"

	"github.com/shurcooL/githubv4"
)

type mergedPRQuery struct {
	Repository struct {
		PullRequests struct {
			Nodes []struct {
				Title     string
				CreatedAt time.Time
				MergedAt  time.Time
				URL       string
			}
			PageInfo struct {
				StartCursor     githubv4.String
				HasPreviousPage bool
			}
		} `graphql:"pullRequests(last: 100, before: $prCursor, states: MERGED)"`
	} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
}
