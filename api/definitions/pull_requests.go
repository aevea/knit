package apidefinitions

// PullRequestService provides all data surrounding pull request data
type PullRequestService interface {
	Oldest(RepositoryRequest) OldestResponse
	AverageByWeek(RepositoryRequest) AverageByWeekResponse
}

// RepositoryRequest is a barebones request that just passes data about which repository to query
type RepositoryRequest struct {
	Repository string
}

// OldestResponse returns data about the oldest PR
type OldestResponse struct {
	Title   string
	URL     string
	OpenFor string
}

// AverageByWeekResponse returns PR data calculated by week
type AverageByWeekResponse struct {
	GeneratedAt string
	Averages    []PRAverage
}

// PRAverage is a representation of averages for PRs. Durations need to be passed as integers for JS to calculate it correctly.
type PRAverage struct {
	Week              string
	MeanTimeToMerge   int64
	MedianTimeToMerge int64
}
