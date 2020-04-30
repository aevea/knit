package apidefinitions

type PullRequestService interface {
	Oldest(OldestRequest) OldestResponse
}

type OldestRequest struct {
	Repository string
}

type OldestResponse struct {
	Title       string
	URL         string
	OpenForDays int64
}
