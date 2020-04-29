package github

import "time"

// PR is an internal struct presenting the information we care about in a PR
type PR struct {
	Title   string
	URL     string
	OpenFor time.Duration
}
