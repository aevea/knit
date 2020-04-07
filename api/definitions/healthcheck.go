package apidefinitions

type Healthcheck interface {
	Check(HealthcheckRequest) HealthcheckResponse
}

type HealthcheckRequest struct {
}

type HealthcheckResponse struct {
	Ok string
}
