package channels

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Channels service
type Channels struct {
	client client.Client
}

func New(clt client.Client) *Channels {
	return &Channels{
		client: clt,
	}
}


// ChannelsList
func (srv *Channels) ChannelsList()(*interface{}, error) {
	path := "/v1/channels"
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
type ChannelsCreateOptions struct {
	IsDefault bool
	Labels interface{}
	Position int
	Status string
	Type string
	enabledSetters map[string]bool
}
func (options ChannelsCreateOptions) New() *ChannelsCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Status": false,
		"Type": false,
	}
	return &options
}
type ChannelsCreateOption func(*ChannelsCreateOptions)
func (srv *Channels) WithChannelsCreateIsDefault(v bool) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsCreateLabels(v interface{}) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsCreatePosition(v int) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsCreateStatus(v string) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Channels) WithChannelsCreateType(v string) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// ChannelsCreate
func (srv *Channels) ChannelsCreate(Code string, Name string, optionalSetters ...ChannelsCreateOption)(*models.Channel, error) {
	path := "/v1/channels"
	options := ChannelsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Channel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Channel
	parsed, ok := resp.Result.(models.Channel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ChannelsDefaults
func (srv *Channels) ChannelsDefaults()(*models.ChannelDefaults, error) {
	path := "/v1/channels/defaults"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ChannelDefaults{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ChannelDefaults
	parsed, ok := resp.Result.(models.ChannelDefaults)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ChannelsDelete
func (srv *Channels) ChannelsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
	
// ChannelsGet
func (srv *Channels) ChannelsGet(Id string)(*models.Channel, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Channel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Channel
	parsed, ok := resp.Result.(models.Channel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ChannelsUpdateOptions struct {
	Code string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Status string
	Type string
	enabledSetters map[string]bool
}
func (options ChannelsUpdateOptions) New() *ChannelsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Status": false,
		"Type": false,
	}
	return &options
}
type ChannelsUpdateOption func(*ChannelsUpdateOptions)
func (srv *Channels) WithChannelsUpdateCode(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Channels) WithChannelsUpdateIsDefault(v bool) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsUpdateLabels(v interface{}) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsUpdateName(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Channels) WithChannelsUpdatePosition(v int) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsUpdateStatus(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Channels) WithChannelsUpdateType(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
			
// ChannelsUpdate
func (srv *Channels) ChannelsUpdate(Id string, optionalSetters ...ChannelsUpdateOption)(*models.Channel, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
	options := ChannelsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Channel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Channel
	parsed, ok := resp.Result.(models.Channel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
