package github

import "time"

type oldestPRQuery struct {
	Repository struct {
		PullRequests struct {
			Nodes []struct {
				Title     string
				CreatedAt time.Time
				URL       string
			}
		} `graphql:"pullRequests(first: 1, states: OPEN)"`
	} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
}
