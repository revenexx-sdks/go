package shipping

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Shipping service
type Shipping struct {
	client client.Client
}

func New(clt client.Client) *Shipping {
	return &Shipping{
		client: clt,
	}
}


// ShippingMethodsList
func (srv *Shipping) ShippingMethodsList()(*interface{}, error) {
	path := "/v1/shipping/methods"
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
type ShippingMethodsCreateOptions struct {
	Carrier string
	Countries []string
	Currency string
	Description string
	Enabled bool
	EtaDaysMax int
	EtaDaysMin int
	FreeAbove float64
	Labels interface{}
	MatrixAttribute string
	MatrixBasis string
	Metadata interface{}
	Position int
	Price float64
	PricingType string
	enabledSetters map[string]bool
}
func (options ShippingMethodsCreateOptions) New() *ShippingMethodsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Carrier": false,
		"Countries": false,
		"Currency": false,
		"Description": false,
		"Enabled": false,
		"EtaDaysMax": false,
		"EtaDaysMin": false,
		"FreeAbove": false,
		"Labels": false,
		"MatrixAttribute": false,
		"MatrixBasis": false,
		"Metadata": false,
		"Position": false,
		"Price": false,
		"PricingType": false,
	}
	return &options
}
type ShippingMethodsCreateOption func(*ShippingMethodsCreateOptions)
func (srv *Shipping) WithShippingMethodsCreateCarrier(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateCountries(v []string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateCurrency(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateDescription(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateEnabled(v bool) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateEtaDaysMax(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateEtaDaysMin(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateFreeAbove(v float64) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.FreeAbove = v
		o.enabledSetters["FreeAbove"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateLabels(v interface{}) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateMatrixAttribute(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.MatrixAttribute = v
		o.enabledSetters["MatrixAttribute"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateMatrixBasis(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.MatrixBasis = v
		o.enabledSetters["MatrixBasis"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreateMetadata(v interface{}) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreatePosition(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreatePrice(v float64) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *Shipping) WithShippingMethodsCreatePricingType(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.PricingType = v
		o.enabledSetters["PricingType"] = true
	}
}
					
// ShippingMethodsCreate
func (srv *Shipping) ShippingMethodsCreate(Code string, Name string, optionalSetters ...ShippingMethodsCreateOption)(*models.ShippingMethod, error) {
	path := "/v1/shipping/methods"
	options := ShippingMethodsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Carrier"] {
		params["carrier"] = options.Carrier
	}
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["EtaDaysMax"] {
		params["eta_days_max"] = options.EtaDaysMax
	}
	if options.enabledSetters["EtaDaysMin"] {
		params["eta_days_min"] = options.EtaDaysMin
	}
	if options.enabledSetters["FreeAbove"] {
		params["free_above"] = options.FreeAbove
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MatrixAttribute"] {
		params["matrix_attribute"] = options.MatrixAttribute
	}
	if options.enabledSetters["MatrixBasis"] {
		params["matrix_basis"] = options.MatrixBasis
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
	}
	if options.enabledSetters["PricingType"] {
		params["pricing_type"] = options.PricingType
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

		parsed := models.ShippingMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingMethod
	parsed, ok := resp.Result.(models.ShippingMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ShippingMethodsDefaults
func (srv *Shipping) ShippingMethodsDefaults()(*interface{}, error) {
	path := "/v1/shipping/methods/defaults"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
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
	
// ShippingMethodsDelete
func (srv *Shipping) ShippingMethodsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/methods/{id}")
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
	
// ShippingMethodsGet
func (srv *Shipping) ShippingMethodsGet(Id string)(*models.ShippingMethod, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/methods/{id}")
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

		parsed := models.ShippingMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingMethod
	parsed, ok := resp.Result.(models.ShippingMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingMethodsUpdateOptions struct {
	Carrier string
	Code string
	Countries []string
	Currency string
	Description string
	Enabled bool
	EtaDaysMax int
	EtaDaysMin int
	FreeAbove float64
	Labels interface{}
	MatrixAttribute string
	MatrixBasis string
	Metadata interface{}
	Name string
	Position int
	Price float64
	PricingType string
	enabledSetters map[string]bool
}
func (options ShippingMethodsUpdateOptions) New() *ShippingMethodsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Carrier": false,
		"Code": false,
		"Countries": false,
		"Currency": false,
		"Description": false,
		"Enabled": false,
		"EtaDaysMax": false,
		"EtaDaysMin": false,
		"FreeAbove": false,
		"Labels": false,
		"MatrixAttribute": false,
		"MatrixBasis": false,
		"Metadata": false,
		"Name": false,
		"Position": false,
		"Price": false,
		"PricingType": false,
	}
	return &options
}
type ShippingMethodsUpdateOption func(*ShippingMethodsUpdateOptions)
func (srv *Shipping) WithShippingMethodsUpdateCarrier(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateCode(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateCountries(v []string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateCurrency(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateDescription(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateEnabled(v bool) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateEtaDaysMax(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateEtaDaysMin(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateFreeAbove(v float64) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.FreeAbove = v
		o.enabledSetters["FreeAbove"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateLabels(v interface{}) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateMatrixAttribute(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.MatrixAttribute = v
		o.enabledSetters["MatrixAttribute"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateMatrixBasis(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.MatrixBasis = v
		o.enabledSetters["MatrixBasis"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateMetadata(v interface{}) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdateName(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdatePosition(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdatePrice(v float64) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *Shipping) WithShippingMethodsUpdatePricingType(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.PricingType = v
		o.enabledSetters["PricingType"] = true
	}
}
			
// ShippingMethodsUpdate
func (srv *Shipping) ShippingMethodsUpdate(Id string, optionalSetters ...ShippingMethodsUpdateOption)(*models.ShippingMethod, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/methods/{id}")
	options := ShippingMethodsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Carrier"] {
		params["carrier"] = options.Carrier
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["EtaDaysMax"] {
		params["eta_days_max"] = options.EtaDaysMax
	}
	if options.enabledSetters["EtaDaysMin"] {
		params["eta_days_min"] = options.EtaDaysMin
	}
	if options.enabledSetters["FreeAbove"] {
		params["free_above"] = options.FreeAbove
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MatrixAttribute"] {
		params["matrix_attribute"] = options.MatrixAttribute
	}
	if options.enabledSetters["MatrixBasis"] {
		params["matrix_basis"] = options.MatrixBasis
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
	}
	if options.enabledSetters["PricingType"] {
		params["pricing_type"] = options.PricingType
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

		parsed := models.ShippingMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingMethod
	parsed, ok := resp.Result.(models.ShippingMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ShippingTiersList
func (srv *Shipping) ShippingTiersList(MethodId string)(*interface{}, error) {
	r := strings.NewReplacer("{methodId}", MethodId)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers")
	params := map[string]interface{}{}
	params["method_id"] = MethodId
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
type ShippingTiersCreateOptions struct {
	FromValue float64
	Position int
	Price float64
	enabledSetters map[string]bool
}
func (options ShippingTiersCreateOptions) New() *ShippingTiersCreateOptions {
	options.enabledSetters = map[string]bool{
		"FromValue": false,
		"Position": false,
		"Price": false,
	}
	return &options
}
type ShippingTiersCreateOption func(*ShippingTiersCreateOptions)
func (srv *Shipping) WithShippingTiersCreateFromValue(v float64) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
func (srv *Shipping) WithShippingTiersCreatePosition(v int) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Shipping) WithShippingTiersCreatePrice(v float64) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
			
// ShippingTiersCreate
func (srv *Shipping) ShippingTiersCreate(MethodId string, optionalSetters ...ShippingTiersCreateOption)(*models.ShippingRateTier, error) {
	r := strings.NewReplacer("{methodId}", MethodId)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers")
	options := ShippingTiersCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	if options.enabledSetters["FromValue"] {
		params["from_value"] = options.FromValue
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
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

		parsed := models.ShippingRateTier{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingRateTier
	parsed, ok := resp.Result.(models.ShippingRateTier)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// ShippingTiersReplace
func (srv *Shipping) ShippingTiersReplace(MethodId string, Tiers []models.ShippingRateTierReplaceItem)(*interface{}, error) {
	r := strings.NewReplacer("{methodId}", MethodId)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers")
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	params["tiers"] = Tiers
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
			
// ShippingTiersDelete
func (srv *Shipping) ShippingTiersDelete(MethodId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{methodId}", MethodId, "{id}", Id)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers/{id}")
	params := map[string]interface{}{}
	params["method_id"] = MethodId
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
			
// ShippingTiersGet
func (srv *Shipping) ShippingTiersGet(MethodId string, Id string)(*models.ShippingRateTier, error) {
	r := strings.NewReplacer("{methodId}", MethodId, "{id}", Id)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers/{id}")
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ShippingRateTier{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingRateTier
	parsed, ok := resp.Result.(models.ShippingRateTier)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingTiersUpdateOptions struct {
	FromValue float64
	Position int
	Price float64
	enabledSetters map[string]bool
}
func (options ShippingTiersUpdateOptions) New() *ShippingTiersUpdateOptions {
	options.enabledSetters = map[string]bool{
		"FromValue": false,
		"Position": false,
		"Price": false,
	}
	return &options
}
type ShippingTiersUpdateOption func(*ShippingTiersUpdateOptions)
func (srv *Shipping) WithShippingTiersUpdateFromValue(v float64) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
func (srv *Shipping) WithShippingTiersUpdatePosition(v int) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Shipping) WithShippingTiersUpdatePrice(v float64) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
					
// ShippingTiersUpdate
func (srv *Shipping) ShippingTiersUpdate(MethodId string, Id string, optionalSetters ...ShippingTiersUpdateOption)(*models.ShippingRateTier, error) {
	r := strings.NewReplacer("{methodId}", MethodId, "{id}", Id)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers/{id}")
	options := ShippingTiersUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	params["id"] = Id
	if options.enabledSetters["FromValue"] {
		params["from_value"] = options.FromValue
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
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

		parsed := models.ShippingRateTier{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingRateTier
	parsed, ok := resp.Result.(models.ShippingRateTier)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingRatesOptions struct {
	Attributes interface{}
	Country string
	Currency string
	MarketId string
	OrderValue float64
	Quantity float64
	Weight float64
	enabledSetters map[string]bool
}
func (options ShippingRatesOptions) New() *ShippingRatesOptions {
	options.enabledSetters = map[string]bool{
		"Attributes": false,
		"Country": false,
		"Currency": false,
		"MarketId": false,
		"OrderValue": false,
		"Quantity": false,
		"Weight": false,
	}
	return &options
}
type ShippingRatesOption func(*ShippingRatesOptions)
func (srv *Shipping) WithShippingRatesAttributes(v interface{}) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Attributes = v
		o.enabledSetters["Attributes"] = true
	}
}
func (srv *Shipping) WithShippingRatesCountry(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Shipping) WithShippingRatesCurrency(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Shipping) WithShippingRatesMarketId(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Shipping) WithShippingRatesOrderValue(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.OrderValue = v
		o.enabledSetters["OrderValue"] = true
	}
}
func (srv *Shipping) WithShippingRatesQuantity(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Shipping) WithShippingRatesWeight(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Weight = v
		o.enabledSetters["Weight"] = true
	}
}
	
// ShippingRates
func (srv *Shipping) ShippingRates(optionalSetters ...ShippingRatesOption)(*interface{}, error) {
	path := "/v1/shipping/rates"
	options := ShippingRatesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Attributes"] {
		params["attributes"] = options.Attributes
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["MarketId"] {
		params["market_id"] = options.MarketId
	}
	if options.enabledSetters["OrderValue"] {
		params["order_value"] = options.OrderValue
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Weight"] {
		params["weight"] = options.Weight
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
