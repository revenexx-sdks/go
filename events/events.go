package events

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Events service
type Events struct {
	client client.Client
}

func New(clt client.Client) *Events {
	return &Events{
		client: clt,
	}
}

type EventsGetCatalogOptions struct {
	Fields string
	enabledSetters map[string]bool
}
func (options EventsGetCatalogOptions) New() *EventsGetCatalogOptions {
	options.enabledSetters = map[string]bool{
		"Fields": false,
	}
	return &options
}
type EventsGetCatalogOption func(*EventsGetCatalogOptions)
func (srv *Events) WithEventsGetCatalogFields(v string) EventsGetCatalogOption {
	return func(o *EventsGetCatalogOptions) {
		o.Fields = v
		o.enabledSetters["Fields"] = true
	}
}
	
// EventsGetCatalog every event type this tenant's installed apps and platform
// services declare — what can be published and subscribed to, independent
// of whether one has fired yet. Each entry says what causes it (`trigger`)
// and what it carries (`sample`, `data_schema`).
func (srv *Events) EventsGetCatalog(optionalSetters ...EventsGetCatalogOption)(*interface{}, error) {
	path := "/v1/events/catalog"
	options := EventsGetCatalogOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Fields"] {
		params["fields"] = options.Fields
	}
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
