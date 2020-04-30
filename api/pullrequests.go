package api

import (
	"context"
	"net/http"

	"github.com/aevea/merge-master/api/generated"
	"github.com/aevea/merge-master/internal/github"
)

type PullRequestService struct {
	HTTPClient *http.Client
}

func (service PullRequestService) Oldest(ctx context.Context, request generated.OldestRequest) (*generated.OldestResponse, error) {
	client, err := github.NewGithubClient(service.HTTPClient, request.Repository)

	if err != nil {
		return nil, err
	}

	oldestPR, err := client.OldestPR()

	if err != nil {
		return nil, err
	}

	return &generated.OldestResponse{OpenForDays: oldestPR.OpenFor.Milliseconds(), Title: oldestPR.Title, URL: oldestPR.URL}, nil
}
