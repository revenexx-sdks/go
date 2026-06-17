package greetings

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Greetings service
type Greetings struct {
	client client.Client
}

func New(clt client.Client) *Greetings {
	return &Greetings{
		client: clt,
	}
}


// GreetingsDigest
func (srv *Greetings) GreetingsDigest()(*interface{}, error) {
	path := "/v1/digest"
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

// GreetingsList
func (srv *Greetings) GreetingsList()(*interface{}, error) {
	path := "/v1/greetings"
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
type GreetingsCreateOptions struct {
	Locale string
	enabledSetters map[string]bool
}
func (options GreetingsCreateOptions) New() *GreetingsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Locale": false,
	}
	return &options
}
type GreetingsCreateOption func(*GreetingsCreateOptions)
func (srv *Greetings) WithGreetingsCreateLocale(v string) GreetingsCreateOption {
	return func(o *GreetingsCreateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
			
// GreetingsCreate
func (srv *Greetings) GreetingsCreate(Name string, optionalSetters ...GreetingsCreateOption)(*interface{}, error) {
	path := "/v1/greetings"
	options := GreetingsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
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
	
// GreetingsDelete
func (srv *Greetings) GreetingsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/greetings/{id}")
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
	
// GreetingsGet
func (srv *Greetings) GreetingsGet(Id string)(*models.Greeting, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/greetings/{id}")
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

		parsed := models.Greeting{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Greeting
	parsed, ok := resp.Result.(models.Greeting)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GreetingsUpdateOptions struct {
	Locale string
	Message string
	Metadata interface{}
	Name string
	enabledSetters map[string]bool
}
func (options GreetingsUpdateOptions) New() *GreetingsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Locale": false,
		"Message": false,
		"Metadata": false,
		"Name": false,
	}
	return &options
}
type GreetingsUpdateOption func(*GreetingsUpdateOptions)
func (srv *Greetings) WithGreetingsUpdateLocale(v string) GreetingsUpdateOption {
	return func(o *GreetingsUpdateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Greetings) WithGreetingsUpdateMessage(v string) GreetingsUpdateOption {
	return func(o *GreetingsUpdateOptions) {
		o.Message = v
		o.enabledSetters["Message"] = true
	}
}
func (srv *Greetings) WithGreetingsUpdateMetadata(v interface{}) GreetingsUpdateOption {
	return func(o *GreetingsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Greetings) WithGreetingsUpdateName(v string) GreetingsUpdateOption {
	return func(o *GreetingsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// GreetingsUpdate
func (srv *Greetings) GreetingsUpdate(Id string, optionalSetters ...GreetingsUpdateOption)(*models.Greeting, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/greetings/{id}")
	options := GreetingsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Message"] {
		params["message"] = options.Message
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Greeting{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Greeting
	parsed, ok := resp.Result.(models.Greeting)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
