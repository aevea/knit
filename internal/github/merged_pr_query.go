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
				Reviews   struct {
					Nodes []struct {
						CreatedAt time.Time
					}
				} `graphql:"reviews(first: 20, states: APPROVED)"`
			}
			PageInfo struct {
				StartCursor     githubv4.String
				HasPreviousPage bool
			}
		} `graphql:"pullRequests(last: 100, before: $prCursor, states: MERGED, baseRefName: $baseRefName)"`
	} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
}
