package health

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Health service
type Health struct {
	client client.Client
}

func New(clt client.Client) *Health {
	return &Health{
		client: clt,
	}
}


// HealthLive answers as long as the process is running. Never touches a
// dependency, so it stays 200 while the gateway is degraded — use readiness
// to decide whether to send traffic.
func (srv *Health) HealthLive()(*interface{}, error) {
	path := "/health/live"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// HealthReady answers 200 once the gateway's registry source is reachable,
// 503 until then.
func (srv *Health) HealthReady()(*interface{}, error) {
	path := "/health/ready"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
