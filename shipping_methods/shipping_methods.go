package shipping_methods

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ShippingMethods service
type ShippingMethods struct {
	client client.Client
}

func New(clt client.Client) *ShippingMethods {
	return &ShippingMethods{
		client: clt,
	}
}

type ShippingMethodsListOptions struct {
	Limit int
	Offset int
	Order string
	Code string
	Enabled bool
	PricingType string
	CarrierId string
	Carrier string
	TaxClass string
	enabledSetters map[string]bool
}
func (options ShippingMethodsListOptions) New() *ShippingMethodsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Code": false,
		"Enabled": false,
		"PricingType": false,
		"CarrierId": false,
		"Carrier": false,
		"TaxClass": false,
	}
	return &options
}
type ShippingMethodsListOption func(*ShippingMethodsListOptions)
func (srv *ShippingMethods) WithShippingMethodsListLimit(v int) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListOffset(v int) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListOrder(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListCode(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListEnabled(v bool) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListPricingType(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.PricingType = v
		o.enabledSetters["PricingType"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListCarrierId(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.CarrierId = v
		o.enabledSetters["CarrierId"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListCarrier(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsListTaxClass(v string) ShippingMethodsListOption {
	return func(o *ShippingMethodsListOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
	
// ShippingMethodsList filterable by exact column value — `?code=`,
// `?enabled=`, `?pricing_type=`, `?carrier_id=`, `?carrier=` and
// `?tax_class=` are applied as equalities and echoed back in `filter`.
// `?carrier_id=` and `?carrier=` are the two halves of one question: the
// first finds the methods holding a reference, the second the ones still
// resolving through the legacy code text. A query key that names no column of
// this entity is SILENTLY IGNORED — `?status=` on this route is the trap,
// since carriers have a status and methods do not: the page comes back
// unfiltered, 200, with an empty `filter`.
func (srv *ShippingMethods) ShippingMethodsList(optionalSetters ...ShippingMethodsListOption)(*models.Error, error) {
	path := "/v1/shipping/methods"
	options := ShippingMethodsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["PricingType"] {
		params["pricing_type"] = options.PricingType
	}
	if options.enabledSetters["CarrierId"] {
		params["carrier_id"] = options.CarrierId
	}
	if options.enabledSetters["Carrier"] {
		params["carrier"] = options.Carrier
	}
	if options.enabledSetters["TaxClass"] {
		params["tax_class"] = options.TaxClass
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingMethodsCreateOptions struct {
	Carrier string
	CarrierId string
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
	QuoteAbove float64
	TaxClass string
	enabledSetters map[string]bool
}
func (options ShippingMethodsCreateOptions) New() *ShippingMethodsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Carrier": false,
		"CarrierId": false,
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
		"QuoteAbove": false,
		"TaxClass": false,
	}
	return &options
}
type ShippingMethodsCreateOption func(*ShippingMethodsCreateOptions)
func (srv *ShippingMethods) WithShippingMethodsCreateCarrier(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateCarrierId(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.CarrierId = v
		o.enabledSetters["CarrierId"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateCountries(v []string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateCurrency(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateDescription(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateEnabled(v bool) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateEtaDaysMax(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateEtaDaysMin(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateFreeAbove(v float64) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.FreeAbove = v
		o.enabledSetters["FreeAbove"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateLabels(v interface{}) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateMatrixAttribute(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.MatrixAttribute = v
		o.enabledSetters["MatrixAttribute"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateMatrixBasis(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.MatrixBasis = v
		o.enabledSetters["MatrixBasis"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateMetadata(v interface{}) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreatePosition(v int) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreatePrice(v float64) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreatePricingType(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.PricingType = v
		o.enabledSetters["PricingType"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateQuoteAbove(v float64) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.QuoteAbove = v
		o.enabledSetters["QuoteAbove"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsCreateTaxClass(v string) ShippingMethodsCreateOption {
	return func(o *ShippingMethodsCreateOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
					
// ShippingMethodsCreate a shipping method is the line a buyer picks in the
// checkout: a pricing model ('fixed', 'free' or 'matrix'), the countries it
// may be offered into, a free-above threshold, and the carrier it ships with.
// The method owns the PRICE; the delivery promise — tracking template,
// cut-off, handling and transit days — is inherited from the carrier
// wherever the method states none of its own. A create cannot omit `code` and
// `name`; every other column is optional or defaulted by the database. Two
// rows of this tenant may not share `code` — that is the 409. The new
// method is quoted by nobody until two further things are true: `enabled`
// defaults to FALSE, and a 'matrix' method has no tiers yet — until POST or
// PUT …/tiers gives it some it appears in `excluded` with 'matrix has no
// rate tiers configured' rather than in the rates. `carrier_id` and the
// legacy `carrier` code are both accepted and neither is verified against the
// carrier table here: an unmatched code is a plain carrier name on the rate,
// not an error.
func (srv *ShippingMethods) ShippingMethodsCreate(Code string, Name string, optionalSetters ...ShippingMethodsCreateOption)(*models.Error, error) {
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
	if options.enabledSetters["CarrierId"] {
		params["carrier_id"] = options.CarrierId
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
	if options.enabledSetters["QuoteAbove"] {
		params["quote_above"] = options.QuoteAbove
	}
	if options.enabledSetters["TaxClass"] {
		params["tax_class"] = options.TaxClass
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ShippingMethodsDefaults runs the carrier seed first, then creates any
// missing method: the three lines a shop is expected to offer — standard,
// express and pickup. The app runs this itself on `app.installed`, so a fresh
// install already has them; calling it by hand afterwards is how a tenant
// that deleted one gets it back, and calling it twice costs nothing, because
// it reconciles rather than seeds. The seeded methods deliberately name no
// carrier: which carrier carries the standard method is a contract, not a
// default, and a method that says 'dhl' resolves to the seeded DHL row
// anyway.
func (srv *ShippingMethods) ShippingMethodsDefaults()(*interface{}, error) {
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
	
// ShippingMethodsDelete deleting one takes every `shipping_rate_tiers` row
// that points at it with it — the foreign keys decide that, not this route.
// So the whole rate matrix goes with the method, which is also why this never
// answers a conflict and why there is no way to recover the table afterwards
// — for a method a checkout may still be holding in a session, `enabled:
// false` is the safer edit.
func (srv *ShippingMethods) ShippingMethodsDelete(Id string)(*models.Error, error) {
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ShippingMethodsGet a shipping method is the line a buyer picks in the
// checkout: a pricing model ('fixed', 'free' or 'matrix'), the countries it
// may be offered into, a free-above threshold, and the carrier it ships with.
// The method owns the PRICE; the delivery promise — tracking template,
// cut-off, handling and transit days — is inherited from the carrier
// wherever the method states none of its own. This is the CONFIGURATION of
// one, by row id — not what a buyer would be charged. A matrix method's
// prices are not in here at all: they are its rate tiers, GET
// /shipping/methods/{method_id}/tiers, and the price for a given basket is
// POST /shipping/rates, which is the only place free-above thresholds,
// country restrictions, the carrier's reach and tax are applied. A checkout
// that reads `price` off this row prices a matrix method at 0.
func (srv *ShippingMethods) ShippingMethodsGet(Id string)(*models.Error, error) {
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingMethodsUpdateOptions struct {
	Carrier string
	CarrierId string
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
	QuoteAbove float64
	TaxClass string
	enabledSetters map[string]bool
}
func (options ShippingMethodsUpdateOptions) New() *ShippingMethodsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Carrier": false,
		"CarrierId": false,
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
		"QuoteAbove": false,
		"TaxClass": false,
	}
	return &options
}
type ShippingMethodsUpdateOption func(*ShippingMethodsUpdateOptions)
func (srv *ShippingMethods) WithShippingMethodsUpdateCarrier(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateCarrierId(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.CarrierId = v
		o.enabledSetters["CarrierId"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateCode(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateCountries(v []string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateCurrency(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateDescription(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateEnabled(v bool) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateEtaDaysMax(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateEtaDaysMin(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateFreeAbove(v float64) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.FreeAbove = v
		o.enabledSetters["FreeAbove"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateLabels(v interface{}) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateMatrixAttribute(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.MatrixAttribute = v
		o.enabledSetters["MatrixAttribute"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateMatrixBasis(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.MatrixBasis = v
		o.enabledSetters["MatrixBasis"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateMetadata(v interface{}) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateName(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdatePosition(v int) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdatePrice(v float64) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdatePricingType(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.PricingType = v
		o.enabledSetters["PricingType"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateQuoteAbove(v float64) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.QuoteAbove = v
		o.enabledSetters["QuoteAbove"] = true
	}
}
func (srv *ShippingMethods) WithShippingMethodsUpdateTaxClass(v string) ShippingMethodsUpdateOption {
	return func(o *ShippingMethodsUpdateOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
			
// ShippingMethodsUpdate a shipping method is the line a buyer picks in the
// checkout: a pricing model ('fixed', 'free' or 'matrix'), the countries it
// may be offered into, a free-above threshold, and the carrier it ships with.
// The method owns the PRICE; the delivery promise — tracking template,
// cut-off, handling and transit days — is inherited from the carrier
// wherever the method states none of its own. A partial update — send only
// what changes, whether that is taking the method in or out of the checkout,
// its pricing, the countries it is restricted to or the delivery estimate it
// states of its own; a payload carrying no column at all is refused rather
// than answering a row it did not touch. Flipping `enabled` is what puts the
// method in front of a buyer or takes it away, and a disabled method is
// reported in the rate answer's `excluded` rather than hidden. Changing
// `pricing_type` away from 'matrix' does NOT delete the tier table — it
// stops being read, and changing back reinstates the old prices, so a method
// switched to 'fixed' and back quotes what it quoted before. Two rows of this
// tenant may not share `code` — that is the 409.
func (srv *ShippingMethods) ShippingMethodsUpdate(Id string, optionalSetters ...ShippingMethodsUpdateOption)(*models.Error, error) {
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
	if options.enabledSetters["CarrierId"] {
		params["carrier_id"] = options.CarrierId
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
	if options.enabledSetters["QuoteAbove"] {
		params["quote_above"] = options.QuoteAbove
	}
	if options.enabledSetters["TaxClass"] {
		params["tax_class"] = options.TaxClass
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingTiersListOptions struct {
	Limit int
	Offset int
	Order string
	FromValue float64
	enabledSetters map[string]bool
}
func (options ShippingTiersListOptions) New() *ShippingTiersListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"FromValue": false,
	}
	return &options
}
type ShippingTiersListOption func(*ShippingTiersListOptions)
func (srv *ShippingMethods) WithShippingTiersListLimit(v int) ShippingTiersListOption {
	return func(o *ShippingTiersListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersListOffset(v int) ShippingTiersListOption {
	return func(o *ShippingTiersListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersListOrder(v string) ShippingTiersListOption {
	return func(o *ShippingTiersListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersListFromValue(v float64) ShippingTiersListOption {
	return func(o *ShippingTiersListOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
			
// ShippingTiersList the rate matrix of one method — every `from_value`
// threshold with the price charged at or above it — lowest threshold first.
// Filterable by `?from_value=` — the unique index is (tenant_id, method_id,
// from_value), so that addresses one row of the matrix by the threshold it
// prices rather than by an id a bulk replace has already discarded. The
// applied filters are echoed in `filter`, which always carries the
// `method_id` taken from the path.
func (srv *ShippingMethods) ShippingTiersList(MethodId string, optionalSetters ...ShippingTiersListOption)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers")
	options := ShippingTiersListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
	}
	if options.enabledSetters["FromValue"] {
		params["from_value"] = options.FromValue
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
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
func (srv *ShippingMethods) WithShippingTiersCreateFromValue(v float64) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersCreatePosition(v int) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersCreatePrice(v float64) ShippingTiersCreateOption {
	return func(o *ShippingTiersCreateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
			
// ShippingTiersCreate a rate tier is one row of a matrix method's price
// table: a `from_value` threshold and the price charged at or above it. The
// bound is INCLUSIVE and the winning tier is the one with the highest
// `from_value` at or below the measured value, so a measure of exactly 10 is
// priced by the tier at 10. What the number measures is the method's
// `matrix_basis` — kilograms in the market's own weight unit, items, money
// in the method's currency, or a named attribute — and the last tier has no
// upper bound. This adds ONE row to the table of the method in the path,
// leaving the rest alone — the edit for a merchant who has added a heavier
// bracket. To lay a whole table down at once use PUT …/tiers (set
// semantics) or POST …/tiers/ladder (evenly stepped), and note that both of
// those DISCARD the ids of the rows they replace. Two rows of this tenant may
// not share the combination of `method_id` + `from_value` — that is the
// 409. `method_id` is taken from the path on every write, so a body naming a
// different method is ignored rather than obeyed.
func (srv *ShippingMethods) ShippingTiersCreate(MethodId string, optionalSetters ...ShippingTiersCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId)
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// ShippingTiersReplace the write behind a table editor: a merchant edits the
// whole matrix on screen and saves it in one call, rather than diffing it
// into a row added here and a row deleted there. Set semantics, and it
// replaces EVERY tier the method had: the tiers this method has afterwards
// are exactly the ones handed in, positions derived from the array order. An
// empty `tiers` array clears the table — and a matrix method with no tiers
// quotes nothing, with a reason.
func (srv *ShippingMethods) ShippingTiersReplace(MethodId string, Tiers []models.ShippingRateTierReplaceItem)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId)
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingTiersLadderOptions struct {
	FromValue float64
	Replace bool
	StepPrice float64
	enabledSetters map[string]bool
}
func (options ShippingTiersLadderOptions) New() *ShippingTiersLadderOptions {
	options.enabledSetters = map[string]bool{
		"FromValue": false,
		"Replace": false,
		"StepPrice": false,
	}
	return &options
}
type ShippingTiersLadderOption func(*ShippingTiersLadderOptions)
func (srv *ShippingMethods) WithShippingTiersLadderFromValue(v float64) ShippingTiersLadderOption {
	return func(o *ShippingTiersLadderOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersLadderReplace(v bool) ShippingTiersLadderOption {
	return func(o *ShippingTiersLadderOptions) {
		o.Replace = v
		o.enabledSetters["Replace"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersLadderStepPrice(v float64) ShippingTiersLadderOption {
	return func(o *ShippingTiersLadderOptions) {
		o.StepPrice = v
		o.enabledSetters["StepPrice"] = true
	}
}
									
// ShippingTiersLadder the tier table a merchant describes in words — "0 to
// 30 kg, every 5 kg, €4.90 plus €2 a step" — without typing every row.
// Replaces the method's tiers by default (set replace=false to append).
func (srv *ShippingMethods) ShippingTiersLadder(MethodId string, BasePrice float64, Step float64, ToValue float64, optionalSetters ...ShippingTiersLadderOption)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId)
	path := r.Replace("/v1/shipping/methods/{method_id}/tiers/ladder")
	options := ShippingTiersLadderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["method_id"] = MethodId
	params["base_price"] = BasePrice
	params["step"] = Step
	params["to_value"] = ToValue
	if options.enabledSetters["FromValue"] {
		params["from_value"] = options.FromValue
	}
	if options.enabledSetters["Replace"] {
		params["replace"] = options.Replace
	}
	if options.enabledSetters["StepPrice"] {
		params["step_price"] = options.StepPrice
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// ShippingTiersDelete a rate tier is one row of a matrix method's price
// table: a `from_value` threshold and the price charged at or above it. The
// bound is INCLUSIVE and the winning tier is the one with the highest
// `from_value` at or below the measured value, so a measure of exactly 10 is
// priced by the tier at 10. What the number measures is the method's
// `matrix_basis` — kilograms in the market's own weight unit, items, money
// in the method's currency, or a named attribute — and the last tier has no
// upper bound. Removing a tier in the MIDDLE of a table is harmless — the
// measures it used to cover fall to the highest remaining threshold below
// them. Removing the LOWEST one is not: a measure under the new lowest
// threshold matches no tier at all, and the method is then left out of POST
// /shipping/rates with 'no tier covers measure …' instead of being quoted
// at 0, so an entire band of baskets silently stops being offered this
// method. Deleting the last tier takes the method out of the checkout
// altogether. Rebuilding the table wholesale is PUT …/tiers or POST
// …/tiers/ladder; deleting the method deletes its tiers on its own.
func (srv *ShippingMethods) ShippingTiersDelete(MethodId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId, "{id}", Id)
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// ShippingTiersGet a rate tier is one row of a matrix method's price table: a
// `from_value` threshold and the price charged at or above it. The bound is
// INCLUSIVE and the winning tier is the one with the highest `from_value` at
// or below the measured value, so a measure of exactly 10 is priced by the
// tier at 10. What the number measures is the method's `matrix_basis` —
// kilograms in the market's own weight unit, items, money in the method's
// currency, or a named attribute — and the last tier has no upper bound.
// This reads one row of that table by id, under the method that owns it; a
// tier id belonging to another method is a 404 rather than somebody else's
// price. A tier id is not durable: PUT …/tiers and POST …/tiers/ladder
// replace the table by deleting and recreating it, so an id read before
// either of them names nothing afterwards. Where a caller wants a stable
// handle, address the row by what it MEANS — GET …/tiers?from_value=…
// — since (method_id, from_value) is unique.
func (srv *ShippingMethods) ShippingTiersGet(MethodId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId, "{id}", Id)
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
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
func (srv *ShippingMethods) WithShippingTiersUpdateFromValue(v float64) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.FromValue = v
		o.enabledSetters["FromValue"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersUpdatePosition(v int) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingMethods) WithShippingTiersUpdatePrice(v float64) ShippingTiersUpdateOption {
	return func(o *ShippingTiersUpdateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
					
// ShippingTiersUpdate a tier id is not stable across a bulk edit: `PUT
// …/tiers` and `POST …/tiers/ladder` replace the table by deleting and
// recreating it, so an id read before either of them is gone afterwards.
func (srv *ShippingMethods) ShippingTiersUpdate(MethodId string, Id string, optionalSetters ...ShippingTiersUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{method_id}", MethodId, "{id}", Id)
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ShippingRatesOptions struct {
	At string
	Attributes interface{}
	Country string
	Currency string
	MarketId string
	OrderValue float64
	OrderValueGross float64
	OrderValueNet float64
	Quantity float64
	Weight float64
	WeightUnit string
	enabledSetters map[string]bool
}
func (options ShippingRatesOptions) New() *ShippingRatesOptions {
	options.enabledSetters = map[string]bool{
		"At": false,
		"Attributes": false,
		"Country": false,
		"Currency": false,
		"MarketId": false,
		"OrderValue": false,
		"OrderValueGross": false,
		"OrderValueNet": false,
		"Quantity": false,
		"Weight": false,
		"WeightUnit": false,
	}
	return &options
}
type ShippingRatesOption func(*ShippingRatesOptions)
func (srv *ShippingMethods) WithShippingRatesAt(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.At = v
		o.enabledSetters["At"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesAttributes(v interface{}) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Attributes = v
		o.enabledSetters["Attributes"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesCountry(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesCurrency(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesMarketId(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesOrderValue(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.OrderValue = v
		o.enabledSetters["OrderValue"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesOrderValueGross(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.OrderValueGross = v
		o.enabledSetters["OrderValueGross"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesOrderValueNet(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.OrderValueNet = v
		o.enabledSetters["OrderValueNet"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesQuantity(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesWeight(v float64) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.Weight = v
		o.enabledSetters["Weight"] = true
	}
}
func (srv *ShippingMethods) WithShippingRatesWeightUnit(v string) ShippingRatesOption {
	return func(o *ShippingRatesOptions) {
		o.WeightUnit = v
		o.enabledSetters["WeightUnit"] = true
	}
}
	
// ShippingRates the question a checkout asks, and the only route that answers
// a PRICE. Hand in the buyer context — the destination country, the order
// value, and whatever the matrix methods measure: a weight, a quantity or a
// named product attribute — and this comes back with the methods that may
// be offered and what each of them costs, free-above thresholds, country
// restrictions, the carrier's delivery promise and tax already applied. A
// method that does not apply is never an error: it moves to `excluded` with a
// reason. So is a tax rate that cannot be resolved — `tax.resolved: false`
// means the rates are UNKNOWN, not untaxed.
func (srv *ShippingMethods) ShippingRates(optionalSetters ...ShippingRatesOption)(*models.Error, error) {
	path := "/v1/shipping/rates"
	options := ShippingRatesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["At"] {
		params["at"] = options.At
	}
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
	if options.enabledSetters["OrderValueGross"] {
		params["order_value_gross"] = options.OrderValueGross
	}
	if options.enabledSetters["OrderValueNet"] {
		params["order_value_net"] = options.OrderValueNet
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Weight"] {
		params["weight"] = options.Weight
	}
	if options.enabledSetters["WeightUnit"] {
		params["weight_unit"] = options.WeightUnit
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ShippingTaxClassesUsage markets.tax_classes is the source of record for the
// rate and this app points at it by CODE from two places: a method's own
// tax_class and the tenant's shipping_tax_class fallback. Neither is a
// foreign key and neither could be — a cross-app FK is what ADR-0055
// forbids — so integrity is a question one app asks the other, and this is
// the answering half. It is asked before a destructive edit: markets calls it
// when an operator tries to delete a tax class, and a count above zero is
// what stops the delete rather than leaving these methods pointing at a code
// nobody serves. Matched as a CODE, not a row: a tax class is unique per
// market, so 'reduced' may exist in several and a method naming it does not
// say which one it meant. Reports at most 500 methods and names the first 20.
// Every code answers, used or not — a code nobody points at is `in_use:
// false`, never a 404.
func (srv *ShippingMethods) ShippingTaxClassesUsage(Code string)(*models.ShippingTaxClassUsage, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/v1/shipping/tax-classes/{code}/usage")
	params := map[string]interface{}{}
	params["code"] = Code
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ShippingTaxClassUsage{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingTaxClassUsage
	parsed, ok := resp.Result.(models.ShippingTaxClassUsage)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
