package api

import (
	"context"

	"github.com/aevea/merge-master/api/generated"
)

type HealthcheckService struct{}

func (HealthcheckService) Check(ctx context.Context, r generated.HealthcheckRequest) (*generated.HealthcheckResponse, error) {
	resp := &generated.HealthcheckResponse{
		Ok: "Ok",
	}
	return resp, nil
}
