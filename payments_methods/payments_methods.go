package payments_methods

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// PaymentsMethods service
type PaymentsMethods struct {
	client client.Client
}

func New(clt client.Client) *PaymentsMethods {
	return &PaymentsMethods{
		client: clt,
	}
}

type PaymentsMethodsListOptions struct {
	Limit int
	Offset int
	Order string
	Code string
	Kind string
	Enabled bool
	Provider string
	enabledSetters map[string]bool
}
func (options PaymentsMethodsListOptions) New() *PaymentsMethodsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Code": false,
		"Kind": false,
		"Enabled": false,
		"Provider": false,
	}
	return &options
}
type PaymentsMethodsListOption func(*PaymentsMethodsListOptions)
func (srv *PaymentsMethods) WithPaymentsMethodsListLimit(v int) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListOffset(v int) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListOrder(v string) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListCode(v string) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListKind(v string) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListEnabled(v bool) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsListProvider(v string) PaymentsMethodsListOption {
	return func(o *PaymentsMethodsListOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
	
// PaymentsMethodsList every method this tenant has configured, enabled or not
// — what the Cockpit's Payment methods screen shows and how an integration
// finds out which codes exist. It answers CONFIGURATION, never an offer:
// nothing here is evaluated against a buyer, so a method restricted to
// Germany, one whose order-value bounds exclude this basket and one whose PSP
// was never set up all come back the same way. The call a checkout makes is
// POST /payments/methods/eligible. Rows come back in whatever order the
// database returns them, so a storefront-shaped list needs
// `?order=position.asc` — `position` is the merchant's intended sequence
// and nothing sorts by it here on its own.
func (srv *PaymentsMethods) PaymentsMethodsList(optionalSetters ...PaymentsMethodsListOption)(*interface{}, error) {
	path := "/v1/payments/methods"
	options := PaymentsMethodsListOptions{}.New()
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
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Provider"] {
		params["provider"] = options.Provider
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
type PaymentsMethodsCreateOptions struct {
	Countries []string
	Description string
	Enabled bool
	FeeAmount float64
	FeeCurrency string
	FeeType string
	Kind string
	Labels interface{}
	MaxOrderValue float64
	Metadata interface{}
	MinOrderValue float64
	Position int
	Provider string
	ProviderMethod string
	enabledSetters map[string]bool
}
func (options PaymentsMethodsCreateOptions) New() *PaymentsMethodsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Countries": false,
		"Description": false,
		"Enabled": false,
		"FeeAmount": false,
		"FeeCurrency": false,
		"FeeType": false,
		"Kind": false,
		"Labels": false,
		"MaxOrderValue": false,
		"Metadata": false,
		"MinOrderValue": false,
		"Position": false,
		"Provider": false,
		"ProviderMethod": false,
	}
	return &options
}
type PaymentsMethodsCreateOption func(*PaymentsMethodsCreateOptions)
func (srv *PaymentsMethods) WithPaymentsMethodsCreateCountries(v []string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateDescription(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateEnabled(v bool) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateFeeAmount(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeAmount = v
		o.enabledSetters["FeeAmount"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateFeeCurrency(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeCurrency = v
		o.enabledSetters["FeeCurrency"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateFeeType(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeType = v
		o.enabledSetters["FeeType"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateKind(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateLabels(v interface{}) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateMaxOrderValue(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.MaxOrderValue = v
		o.enabledSetters["MaxOrderValue"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateMetadata(v interface{}) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateMinOrderValue(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.MinOrderValue = v
		o.enabledSetters["MinOrderValue"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreatePosition(v int) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateProvider(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsCreateProviderMethod(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.ProviderMethod = v
		o.enabledSetters["ProviderMethod"] = true
	}
}
					
// PaymentsMethodsCreate adds a line a checkout can offer. A create cannot
// omit `code` and `name`; every other column is optional or defaulted by the
// database. Two rows of this tenant may not share `code` — that is the 409.
// Two defaults are worth knowing before the first call: `enabled` is false,
// so a new method reaches no checkout until it is switched on, and `kind` is
// 'self_managed' — a card or wallet method needs `kind: "psp"` plus a
// `provider` the catalog carries, or it falls back to the tenant's
// `default_provider` at payment time and fails there if none is set. The
// `code` is the value every payment, every checkout and every ERP will name
// this method by from now on, and once a single payment has been made under
// it a rename is refused with 409: choose it once.
func (srv *PaymentsMethods) PaymentsMethodsCreate(Code string, Name string, optionalSetters ...PaymentsMethodsCreateOption)(*models.Error, error) {
	path := "/v1/payments/methods"
	options := PaymentsMethodsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FeeAmount"] {
		params["fee_amount"] = options.FeeAmount
	}
	if options.enabledSetters["FeeCurrency"] {
		params["fee_currency"] = options.FeeCurrency
	}
	if options.enabledSetters["FeeType"] {
		params["fee_type"] = options.FeeType
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MaxOrderValue"] {
		params["max_order_value"] = options.MaxOrderValue
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["MinOrderValue"] {
		params["min_order_value"] = options.MinOrderValue
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Provider"] {
		params["provider"] = options.Provider
	}
	if options.enabledSetters["ProviderMethod"] {
		params["provider_method"] = options.ProviderMethod
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

// PaymentsMethodsDefaults writes the four methods a shop starts with —
// invoice and prepayment as self-managed, card and PayPal routed at the mock
// PSP so a fresh install can complete a checkout end to end — together with
// the four provider rows behind them: the built-in mock plus Stripe, PayPal
// and Novalnet, the three connectors this app opens outbound. The app already
// runs this for itself when it is installed (it listens on app.installed), so
// calling the route is for the second time and after: a method someone
// deleted, or a row a later release added that an existing install never got.
// Stripe, PayPal and Novalnet arrive disabled, in test mode and without
// credentials — the operator fills those in — while the mock arrives
// enabled, because it moves no money. Re-running is safe by design: it never
// duplicates a row and never overwrites an existing one, so nothing an
// operator has set can be undone by calling it again. Only genuinely missing
// option keys (a logo added after the first install) are filled, and those
// rows are reported as "updated" rather than created.
func (srv *PaymentsMethods) PaymentsMethodsDefaults()(*interface{}, error) {
	path := "/v1/payments/methods/defaults"
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
type PaymentsMethodsEligibleOptions struct {
	Amount float64
	Country string
	Currency string
	enabledSetters map[string]bool
}
func (options PaymentsMethodsEligibleOptions) New() *PaymentsMethodsEligibleOptions {
	options.enabledSetters = map[string]bool{
		"Amount": false,
		"Country": false,
		"Currency": false,
	}
	return &options
}
type PaymentsMethodsEligibleOption func(*PaymentsMethodsEligibleOptions)
func (srv *PaymentsMethods) WithPaymentsMethodsEligibleAmount(v float64) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Amount = v
		o.enabledSetters["Amount"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsEligibleCountry(v string) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsEligibleCurrency(v string) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
	
// PaymentsMethodsEligible the checkout's question — "what can THIS buyer
// pay with?" — answered server-side before any PSP is involved, so the
// storefront never renders a method the create would then refuse with 422. It
// evaluates the buyer context against every configured method: disabled, a
// country outside `countries`, an amount outside
// `min_order_value`/`max_order_value`. Restriction dimensions are ANDed and
// entries within one are ORed, and an empty dimension means unrestricted.
// Eligible methods come back sorted by `position` with their fee already
// computed for this amount; everything else lands in `excluded` with the
// reason in words, which is what makes a support question answerable. It
// reads only — nothing is written and no provider is called. Two things it
// does NOT check: whether the method's PSP is configured and enabled (a
// method whose provider is switched off is still offered here and fails at
// POST /payments — a provider a method names can no longer be deleted,
// which closes the other half of the same gap), and anything about the buyer
// beyond country and amount. A context that matches nothing is 200 with an
// empty `methods` list, never 404.
func (srv *PaymentsMethods) PaymentsMethodsEligible(optionalSetters ...PaymentsMethodsEligibleOption)(*interface{}, error) {
	path := "/v1/payments/methods/eligible"
	options := PaymentsMethodsEligibleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Amount"] {
		params["amount"] = options.Amount
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	
// PaymentsMethodsDelete payments.method_code is a CODE, not a foreign key: a
// payment records what happened and has to survive the configuration it was
// made with. The cost of that looseness is that deleting a method turns every
// payment made with it into a row naming something that no longer exists. So
// the count is taken HERE and answered as 409 with the number, rather than
// left to whoever is about to click delete — a client that pre-counts asks
// a second question whose answer disagrees the moment a payment lands between
// the two calls. Disabling the method (enabled: false) is what an operator
// usually meant and stays available.
func (srv *PaymentsMethods) PaymentsMethodsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/methods/{id}")
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
	
// PaymentsMethodsGet one configuration, every column, addressed by its row id
// — the edit form's read. It is addressed by ID and there is no route that
// takes a `code`, which matters because the CODE is what a checkout, a
// payment and an ERP name a method by: to resolve one, filter the list (`GET
// /payments/methods?code=invoice`), which answers a page of at most one row
// because (tenant_id, code) is unique. Reading a method says nothing about
// whether a buyer may use it — that is POST /payments/methods/eligible —
// and nothing about whether its PSP can transact, which is under the provider
// configuration.
func (srv *PaymentsMethods) PaymentsMethodsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/methods/{id}")
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
type PaymentsMethodsUpdateOptions struct {
	Code string
	Countries []string
	Description string
	Enabled bool
	FeeAmount float64
	FeeCurrency string
	FeeType string
	Kind string
	Labels interface{}
	MaxOrderValue float64
	Metadata interface{}
	MinOrderValue float64
	Name string
	Position int
	Provider string
	ProviderMethod string
	enabledSetters map[string]bool
}
func (options PaymentsMethodsUpdateOptions) New() *PaymentsMethodsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Countries": false,
		"Description": false,
		"Enabled": false,
		"FeeAmount": false,
		"FeeCurrency": false,
		"FeeType": false,
		"Kind": false,
		"Labels": false,
		"MaxOrderValue": false,
		"Metadata": false,
		"MinOrderValue": false,
		"Name": false,
		"Position": false,
		"Provider": false,
		"ProviderMethod": false,
	}
	return &options
}
type PaymentsMethodsUpdateOption func(*PaymentsMethodsUpdateOptions)
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateCode(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateCountries(v []string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateDescription(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateEnabled(v bool) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateFeeAmount(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeAmount = v
		o.enabledSetters["FeeAmount"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateFeeCurrency(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeCurrency = v
		o.enabledSetters["FeeCurrency"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateFeeType(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeType = v
		o.enabledSetters["FeeType"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateKind(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateLabels(v interface{}) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateMaxOrderValue(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.MaxOrderValue = v
		o.enabledSetters["MaxOrderValue"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateMetadata(v interface{}) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateMinOrderValue(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.MinOrderValue = v
		o.enabledSetters["MinOrderValue"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateName(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdatePosition(v int) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateProvider(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *PaymentsMethods) WithPaymentsMethodsUpdateProviderMethod(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.ProviderMethod = v
		o.enabledSetters["ProviderMethod"] = true
	}
}
			
// PaymentsMethodsUpdate a PUT that PATCHES: only the keys in the body are
// written and every omitted column keeps its value, so `{"enabled": false}`
// is the whole request for taking a method out of checkout. A body with no
// writable key is refused with 400 rather than treated as a no-op. This is
// the route for all three things an operator changes about a method after it
// exists — the `enabled` switch that puts it in or out of checkout, the fee
// it charges (`fee_type`, `fee_amount`, `fee_currency`) and the restrictions
// that decide who is offered it (`countries`, `min_order_value`,
// `max_order_value`) — alongside its labels, description and `position`.
// `enabled: false` is the safe way to retire one — it disappears from POST
// /payments/methods/eligible immediately and stays on every payment ever made
// with it. The one write this route refuses is a rename of `code` while the
// ledger still names the old one. The three tables of this app carry no
// foreign keys at all: a payment names its method by `method_code` and its
// acquirer by `provider`, both plain text, because a payment records what
// happened and has to survive the configuration it was made with. So the
// database will not stop this — whatever the ledger still names, it goes on
// naming. A rename would therefore leave every recorded payment pointing at a
// code no configuration carries, which is the same harm DELETE on this row
// answers 409 for — so it answers the same 409, with the same
// `method_in_use` code and the same count. Renaming a method nothing has been
// paid with is still free, and so is every other column at any time.
func (srv *PaymentsMethods) PaymentsMethodsUpdate(Id string, optionalSetters ...PaymentsMethodsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/methods/{id}")
	options := PaymentsMethodsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FeeAmount"] {
		params["fee_amount"] = options.FeeAmount
	}
	if options.enabledSetters["FeeCurrency"] {
		params["fee_currency"] = options.FeeCurrency
	}
	if options.enabledSetters["FeeType"] {
		params["fee_type"] = options.FeeType
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MaxOrderValue"] {
		params["max_order_value"] = options.MaxOrderValue
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["MinOrderValue"] {
		params["min_order_value"] = options.MinOrderValue
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Provider"] {
		params["provider"] = options.Provider
	}
	if options.enabledSetters["ProviderMethod"] {
		params["provider_method"] = options.ProviderMethod
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
