package markets

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Markets service
type Markets struct {
	client client.Client
}

func New(clt client.Client) *Markets {
	return &Markets{
		client: clt,
	}
}


// MarketsList
func (srv *Markets) MarketsList()(*interface{}, error) {
	path := "/v1/markets"
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
type MarketsCreateOptions struct {
	Currency string
	IsDefault bool
	Labels interface{}
	Position int
	Status string
	enabledSetters map[string]bool
}
func (options MarketsCreateOptions) New() *MarketsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Currency": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Status": false,
	}
	return &options
}
type MarketsCreateOption func(*MarketsCreateOptions)
func (srv *Markets) WithMarketsCreateCurrency(v string) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsCreateIsDefault(v bool) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCreateLabels(v interface{}) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsCreatePosition(v int) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsCreateStatus(v string) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
					
// MarketsCreate
func (srv *Markets) MarketsCreate(Code string, Name string, optionalSetters ...MarketsCreateOption)(*models.Market, error) {
	path := "/v1/markets"
	options := MarketsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
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
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Market{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Market
	parsed, ok := resp.Result.(models.Market)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsDelete
func (srv *Markets) MarketsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
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
	
// MarketsGet
func (srv *Markets) MarketsGet(Id string)(*models.Market, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
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

		parsed := models.Market{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Market
	parsed, ok := resp.Result.(models.Market)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MarketsUpdateOptions struct {
	Code string
	Currency string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Status string
	enabledSetters map[string]bool
}
func (options MarketsUpdateOptions) New() *MarketsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Currency": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Status": false,
	}
	return &options
}
type MarketsUpdateOption func(*MarketsUpdateOptions)
func (srv *Markets) WithMarketsUpdateCode(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsUpdateCurrency(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsUpdateIsDefault(v bool) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsUpdateLabels(v interface{}) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsUpdateName(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsUpdatePosition(v int) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsUpdateStatus(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// MarketsUpdate
func (srv *Markets) MarketsUpdate(Id string, optionalSetters ...MarketsUpdateOption)(*models.Market, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
	options := MarketsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Market{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Market
	parsed, ok := resp.Result.(models.Market)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsContext
func (srv *Markets) MarketsContext(Id string)(*models.MarketContext, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/context")
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

		parsed := models.MarketContext{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketContext
	parsed, ok := resp.Result.(models.MarketContext)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsCurrenciesList
func (srv *Markets) MarketsCurrenciesList(MarketId string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/currencies")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsCurrenciesCreateOptions struct {
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsCurrenciesCreateOptions) New() *MarketsCurrenciesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsCurrenciesCreateOption func(*MarketsCurrenciesCreateOptions)
func (srv *Markets) WithMarketsCurrenciesCreateIsDefault(v bool) MarketsCurrenciesCreateOption {
	return func(o *MarketsCurrenciesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesCreatePosition(v int) MarketsCurrenciesCreateOption {
	return func(o *MarketsCurrenciesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsCurrenciesCreate
func (srv *Markets) MarketsCurrenciesCreate(MarketId string, Code string, optionalSetters ...MarketsCurrenciesCreateOption)(*models.MarketCurrency, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/currencies")
	options := MarketsCurrenciesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["code"] = Code
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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

		parsed := models.MarketCurrency{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketCurrency
	parsed, ok := resp.Result.(models.MarketCurrency)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// MarketsCurrenciesDelete
func (srv *Markets) MarketsCurrenciesDelete(MarketId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsCurrenciesGet
func (srv *Markets) MarketsCurrenciesGet(MarketId string, Id string)(*models.MarketCurrency, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MarketCurrency{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketCurrency
	parsed, ok := resp.Result.(models.MarketCurrency)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MarketsCurrenciesUpdateOptions struct {
	Code string
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsCurrenciesUpdateOptions) New() *MarketsCurrenciesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsCurrenciesUpdateOption func(*MarketsCurrenciesUpdateOptions)
func (srv *Markets) WithMarketsCurrenciesUpdateCode(v string) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesUpdateIsDefault(v bool) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesUpdatePosition(v int) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsCurrenciesUpdate
func (srv *Markets) MarketsCurrenciesUpdate(MarketId string, Id string, optionalSetters ...MarketsCurrenciesUpdateOption)(*models.MarketCurrency, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	options := MarketsCurrenciesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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

		parsed := models.MarketCurrency{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketCurrency
	parsed, ok := resp.Result.(models.MarketCurrency)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsLocalesList
func (srv *Markets) MarketsLocalesList(MarketId string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/locales")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsLocalesCreateOptions struct {
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsLocalesCreateOptions) New() *MarketsLocalesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsLocalesCreateOption func(*MarketsLocalesCreateOptions)
func (srv *Markets) WithMarketsLocalesCreateIsDefault(v bool) MarketsLocalesCreateOption {
	return func(o *MarketsLocalesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsLocalesCreatePosition(v int) MarketsLocalesCreateOption {
	return func(o *MarketsLocalesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
									
// MarketsLocalesCreate
func (srv *Markets) MarketsLocalesCreate(MarketId string, Code string, Country string, Language string, optionalSetters ...MarketsLocalesCreateOption)(*models.MarketLocale, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/locales")
	options := MarketsLocalesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["code"] = Code
	params["country"] = Country
	params["language"] = Language
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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

		parsed := models.MarketLocale{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketLocale
	parsed, ok := resp.Result.(models.MarketLocale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// MarketsLocalesDelete
func (srv *Markets) MarketsLocalesDelete(MarketId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsLocalesGet
func (srv *Markets) MarketsLocalesGet(MarketId string, Id string)(*models.MarketLocale, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MarketLocale{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketLocale
	parsed, ok := resp.Result.(models.MarketLocale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MarketsLocalesUpdateOptions struct {
	Code string
	Country string
	IsDefault bool
	Language string
	Position int
	enabledSetters map[string]bool
}
func (options MarketsLocalesUpdateOptions) New() *MarketsLocalesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Country": false,
		"IsDefault": false,
		"Language": false,
		"Position": false,
	}
	return &options
}
type MarketsLocalesUpdateOption func(*MarketsLocalesUpdateOptions)
func (srv *Markets) WithMarketsLocalesUpdateCode(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateCountry(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateIsDefault(v bool) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateLanguage(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Language = v
		o.enabledSetters["Language"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdatePosition(v int) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsLocalesUpdate
func (srv *Markets) MarketsLocalesUpdate(MarketId string, Id string, optionalSetters ...MarketsLocalesUpdateOption)(*models.MarketLocale, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	options := MarketsLocalesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Language"] {
		params["language"] = options.Language
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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

		parsed := models.MarketLocale{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketLocale
	parsed, ok := resp.Result.(models.MarketLocale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsTaxClassesList
func (srv *Markets) MarketsTaxClassesList(MarketId string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/tax_classes")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsTaxClassesCreateOptions struct {
	IsDefault bool
	Labels interface{}
	Position int
	Rate float64
	enabledSetters map[string]bool
}
func (options MarketsTaxClassesCreateOptions) New() *MarketsTaxClassesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Rate": false,
	}
	return &options
}
type MarketsTaxClassesCreateOption func(*MarketsTaxClassesCreateOptions)
func (srv *Markets) WithMarketsTaxClassesCreateIsDefault(v bool) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreateLabels(v interface{}) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreatePosition(v int) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreateRate(v float64) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Rate = v
		o.enabledSetters["Rate"] = true
	}
}
							
// MarketsTaxClassesCreate
func (srv *Markets) MarketsTaxClassesCreate(MarketId string, Code string, Name string, optionalSetters ...MarketsTaxClassesCreateOption)(*models.MarketTaxClass, error) {
	r := strings.NewReplacer("{marketId}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/tax_classes")
	options := MarketsTaxClassesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
	if options.enabledSetters["Rate"] {
		params["rate"] = options.Rate
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

		parsed := models.MarketTaxClass{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketTaxClass
	parsed, ok := resp.Result.(models.MarketTaxClass)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// MarketsTaxClassesDelete
func (srv *Markets) MarketsTaxClassesDelete(MarketId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsTaxClassesGet
func (srv *Markets) MarketsTaxClassesGet(MarketId string, Id string)(*models.MarketTaxClass, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MarketTaxClass{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketTaxClass
	parsed, ok := resp.Result.(models.MarketTaxClass)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MarketsTaxClassesUpdateOptions struct {
	Code string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Rate float64
	enabledSetters map[string]bool
}
func (options MarketsTaxClassesUpdateOptions) New() *MarketsTaxClassesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Rate": false,
	}
	return &options
}
type MarketsTaxClassesUpdateOption func(*MarketsTaxClassesUpdateOptions)
func (srv *Markets) WithMarketsTaxClassesUpdateCode(v string) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateIsDefault(v bool) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateLabels(v interface{}) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateName(v string) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdatePosition(v int) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateRate(v float64) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Rate = v
		o.enabledSetters["Rate"] = true
	}
}
					
// MarketsTaxClassesUpdate
func (srv *Markets) MarketsTaxClassesUpdate(MarketId string, Id string, optionalSetters ...MarketsTaxClassesUpdateOption)(*models.MarketTaxClass, error) {
	r := strings.NewReplacer("{marketId}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	options := MarketsTaxClassesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
	if options.enabledSetters["Rate"] {
		params["rate"] = options.Rate
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

		parsed := models.MarketTaxClass{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketTaxClass
	parsed, ok := resp.Result.(models.MarketTaxClass)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
