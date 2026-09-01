package settings

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Settings service
type Settings struct {
	client client.Client
}

func New(clt client.Client) *Settings {
	return &Settings{
		client: clt,
	}
}

type SettingsGetAppSettingsOptions struct {
	Market string
	enabledSetters map[string]bool
}
func (options SettingsGetAppSettingsOptions) New() *SettingsGetAppSettingsOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
	}
	return &options
}
type SettingsGetAppSettingsOption func(*SettingsGetAppSettingsOptions)
func (srv *Settings) WithSettingsGetAppSettingsMarket(v string) SettingsGetAppSettingsOption {
	return func(o *SettingsGetAppSettingsOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
			
// SettingsGetAppSettings the tenant's effective settings for the app — the
// declared schema's defaults merged with stored tenant/market values.
// Sensitive settings are masked (listed in `masked`, omitted from
// `settings`).
func (srv *Settings) SettingsGetAppSettings(App string, optionalSetters ...SettingsGetAppSettingsOption)(*interface{}, error) {
	r := strings.NewReplacer("{app}", App)
	path := r.Replace("/v1/settings/apps/{app}")
	options := SettingsGetAppSettingsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["app"] = App
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
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
