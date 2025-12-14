package opnsense

type HealthCheckResponse struct {
	Metadata struct {
		System struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"System"`
		// TODO: I do not see these on my install
		CrashReporter struct {
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"CrashReporter"`
		Firewall struct {
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"Firewall"`
	} `json:"metadata"`
}

const (
	HealthCheckStatusOK = "OK"
)

// HealthCheck checks if the OPNsense is up and running.
func (c *Client) HealthCheck() (HealthCheckResponse, error) {
	var resp HealthCheckResponse

	path, ok := c.endpoints["healthCheck"]

	if !ok {
		return HealthCheckResponse{}, &APICallError{
			Endpoint:   "healthCheck",
			Message:    "endpoint not found",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", path, nil, &resp); err != nil {
		return HealthCheckResponse{}, err
	}

	return resp, nil
}
