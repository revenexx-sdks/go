package payments

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Payments service
type Payments struct {
	client client.Client
}

func New(clt client.Client) *Payments {
	return &Payments{
		client: clt,
	}
}


// PaymentsList
func (srv *Payments) PaymentsList()(*interface{}, error) {
	path := "/v1/payments"
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
type PaymentsCreateOptions struct {
	CartId string
	ContactId string
	Country string
	Currency string
	IdempotencyKey string
	Metadata interface{}
	OrderRef string
	ReturnUrl string
	enabledSetters map[string]bool
}
func (options PaymentsCreateOptions) New() *PaymentsCreateOptions {
	options.enabledSetters = map[string]bool{
		"CartId": false,
		"ContactId": false,
		"Country": false,
		"Currency": false,
		"IdempotencyKey": false,
		"Metadata": false,
		"OrderRef": false,
		"ReturnUrl": false,
	}
	return &options
}
type PaymentsCreateOption func(*PaymentsCreateOptions)
func (srv *Payments) WithPaymentsCreateCartId(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *Payments) WithPaymentsCreateContactId(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Payments) WithPaymentsCreateCountry(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Payments) WithPaymentsCreateCurrency(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Payments) WithPaymentsCreateIdempotencyKey(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.IdempotencyKey = v
		o.enabledSetters["IdempotencyKey"] = true
	}
}
func (srv *Payments) WithPaymentsCreateMetadata(v interface{}) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Payments) WithPaymentsCreateOrderRef(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *Payments) WithPaymentsCreateReturnUrl(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.ReturnUrl = v
		o.enabledSetters["ReturnUrl"] = true
	}
}
					
// PaymentsCreate
func (srv *Payments) PaymentsCreate(Amount float64, MethodCode string, optionalSetters ...PaymentsCreateOption)(*models.Payment, error) {
	path := "/v1/payments"
	options := PaymentsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["amount"] = Amount
	params["method_code"] = MethodCode
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["IdempotencyKey"] {
		params["idempotency_key"] = options.IdempotencyKey
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["ReturnUrl"] {
		params["return_url"] = options.ReturnUrl
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

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PaymentsMethodsList
func (srv *Payments) PaymentsMethodsList()(*interface{}, error) {
	path := "/v1/payments/methods"
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
func (srv *Payments) WithPaymentsMethodsCreateCountries(v []string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateDescription(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateEnabled(v bool) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateFeeAmount(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeAmount = v
		o.enabledSetters["FeeAmount"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateFeeCurrency(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeCurrency = v
		o.enabledSetters["FeeCurrency"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateFeeType(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.FeeType = v
		o.enabledSetters["FeeType"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateKind(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateLabels(v interface{}) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateMaxOrderValue(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.MaxOrderValue = v
		o.enabledSetters["MaxOrderValue"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateMetadata(v interface{}) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateMinOrderValue(v float64) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.MinOrderValue = v
		o.enabledSetters["MinOrderValue"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreatePosition(v int) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateProvider(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsCreateProviderMethod(v string) PaymentsMethodsCreateOption {
	return func(o *PaymentsMethodsCreateOptions) {
		o.ProviderMethod = v
		o.enabledSetters["ProviderMethod"] = true
	}
}
					
// PaymentsMethodsCreate
func (srv *Payments) PaymentsMethodsCreate(Code string, Name string, optionalSetters ...PaymentsMethodsCreateOption)(*models.PaymentMethod, error) {
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

		parsed := models.PaymentMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentMethod
	parsed, ok := resp.Result.(models.PaymentMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PaymentsMethodsDefaults
func (srv *Payments) PaymentsMethodsDefaults()(*interface{}, error) {
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
func (srv *Payments) WithPaymentsMethodsEligibleAmount(v float64) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Amount = v
		o.enabledSetters["Amount"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsEligibleCountry(v string) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsEligibleCurrency(v string) PaymentsMethodsEligibleOption {
	return func(o *PaymentsMethodsEligibleOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
	
// PaymentsMethodsEligible
func (srv *Payments) PaymentsMethodsEligible(optionalSetters ...PaymentsMethodsEligibleOption)(*interface{}, error) {
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
	
// PaymentsMethodsDelete
func (srv *Payments) PaymentsMethodsDelete(Id string)(*interface{}, error) {
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
	
// PaymentsMethodsGet
func (srv *Payments) PaymentsMethodsGet(Id string)(*models.PaymentMethod, error) {
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

		parsed := models.PaymentMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentMethod
	parsed, ok := resp.Result.(models.PaymentMethod)
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
func (srv *Payments) WithPaymentsMethodsUpdateCode(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateCountries(v []string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateDescription(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateEnabled(v bool) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateFeeAmount(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeAmount = v
		o.enabledSetters["FeeAmount"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateFeeCurrency(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeCurrency = v
		o.enabledSetters["FeeCurrency"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateFeeType(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.FeeType = v
		o.enabledSetters["FeeType"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateKind(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateLabels(v interface{}) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateMaxOrderValue(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.MaxOrderValue = v
		o.enabledSetters["MaxOrderValue"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateMetadata(v interface{}) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateMinOrderValue(v float64) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.MinOrderValue = v
		o.enabledSetters["MinOrderValue"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateName(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdatePosition(v int) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateProvider(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *Payments) WithPaymentsMethodsUpdateProviderMethod(v string) PaymentsMethodsUpdateOption {
	return func(o *PaymentsMethodsUpdateOptions) {
		o.ProviderMethod = v
		o.enabledSetters["ProviderMethod"] = true
	}
}
			
// PaymentsMethodsUpdate
func (srv *Payments) PaymentsMethodsUpdate(Id string, optionalSetters ...PaymentsMethodsUpdateOption)(*models.PaymentMethod, error) {
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

		parsed := models.PaymentMethod{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentMethod
	parsed, ok := resp.Result.(models.PaymentMethod)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PaymentsProvidersList
func (srv *Payments) PaymentsProvidersList()(*interface{}, error) {
	path := "/v1/payments/providers"
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
type PaymentsProvidersCreateOptions struct {
	Credentials interface{}
	Enabled bool
	Name string
	Options interface{}
	TestMode bool
	WebhookSecret string
	enabledSetters map[string]bool
}
func (options PaymentsProvidersCreateOptions) New() *PaymentsProvidersCreateOptions {
	options.enabledSetters = map[string]bool{
		"Credentials": false,
		"Enabled": false,
		"Name": false,
		"Options": false,
		"TestMode": false,
		"WebhookSecret": false,
	}
	return &options
}
type PaymentsProvidersCreateOption func(*PaymentsProvidersCreateOptions)
func (srv *Payments) WithPaymentsProvidersCreateCredentials(v interface{}) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.Credentials = v
		o.enabledSetters["Credentials"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersCreateEnabled(v bool) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersCreateName(v string) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersCreateOptions(v interface{}) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersCreateTestMode(v bool) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.TestMode = v
		o.enabledSetters["TestMode"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersCreateWebhookSecret(v string) PaymentsProvidersCreateOption {
	return func(o *PaymentsProvidersCreateOptions) {
		o.WebhookSecret = v
		o.enabledSetters["WebhookSecret"] = true
	}
}
			
// PaymentsProvidersCreate
func (srv *Payments) PaymentsProvidersCreate(Provider string, optionalSetters ...PaymentsProvidersCreateOption)(*models.PaymentProvider, error) {
	path := "/v1/payments/providers"
	options := PaymentsProvidersCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["provider"] = Provider
	if options.enabledSetters["Credentials"] {
		params["credentials"] = options.Credentials
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
	}
	if options.enabledSetters["TestMode"] {
		params["test_mode"] = options.TestMode
	}
	if options.enabledSetters["WebhookSecret"] {
		params["webhook_secret"] = options.WebhookSecret
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

		parsed := models.PaymentProvider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentProvider
	parsed, ok := resp.Result.(models.PaymentProvider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PaymentsProvidersCatalog
func (srv *Payments) PaymentsProvidersCatalog()(*interface{}, error) {
	path := "/v1/payments/providers/catalog"
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
	
// PaymentsProvidersDelete
func (srv *Payments) PaymentsProvidersDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/providers/{id}")
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
	
// PaymentsProvidersGet
func (srv *Payments) PaymentsProvidersGet(Id string)(*models.PaymentProvider, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/providers/{id}")
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

		parsed := models.PaymentProvider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentProvider
	parsed, ok := resp.Result.(models.PaymentProvider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PaymentsProvidersUpdateOptions struct {
	Credentials interface{}
	Enabled bool
	Name string
	Options interface{}
	Provider string
	TestMode bool
	WebhookSecret string
	enabledSetters map[string]bool
}
func (options PaymentsProvidersUpdateOptions) New() *PaymentsProvidersUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Credentials": false,
		"Enabled": false,
		"Name": false,
		"Options": false,
		"Provider": false,
		"TestMode": false,
		"WebhookSecret": false,
	}
	return &options
}
type PaymentsProvidersUpdateOption func(*PaymentsProvidersUpdateOptions)
func (srv *Payments) WithPaymentsProvidersUpdateCredentials(v interface{}) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.Credentials = v
		o.enabledSetters["Credentials"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateEnabled(v bool) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateName(v string) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateOptions(v interface{}) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateProvider(v string) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateTestMode(v bool) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.TestMode = v
		o.enabledSetters["TestMode"] = true
	}
}
func (srv *Payments) WithPaymentsProvidersUpdateWebhookSecret(v string) PaymentsProvidersUpdateOption {
	return func(o *PaymentsProvidersUpdateOptions) {
		o.WebhookSecret = v
		o.enabledSetters["WebhookSecret"] = true
	}
}
			
// PaymentsProvidersUpdate
func (srv *Payments) PaymentsProvidersUpdate(Id string, optionalSetters ...PaymentsProvidersUpdateOption)(*models.PaymentProvider, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/providers/{id}")
	options := PaymentsProvidersUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Credentials"] {
		params["credentials"] = options.Credentials
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
	}
	if options.enabledSetters["Provider"] {
		params["provider"] = options.Provider
	}
	if options.enabledSetters["TestMode"] {
		params["test_mode"] = options.TestMode
	}
	if options.enabledSetters["WebhookSecret"] {
		params["webhook_secret"] = options.WebhookSecret
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

		parsed := models.PaymentProvider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PaymentProvider
	parsed, ok := resp.Result.(models.PaymentProvider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// PaymentsWebhooksIngest consumes the dispatch envelope from
// webhooks.revenexx.com: normalizes the provider callback (stripe payment
// intents + a generic shape), resolves the payment by psp_payment_id or
// order_ref and moves the ledger. Facts only move forward — provider
// retries and redeliveries are idempotent no-ops; unverified envelopes are
// refused.
func (srv *Payments) PaymentsWebhooksIngest(Provider string, Data interface{})(*interface{}, error) {
	r := strings.NewReplacer("{provider}", Provider)
	path := r.Replace("/v1/payments/webhooks/{provider}")
	params := map[string]interface{}{}
	params["provider"] = Provider
	params["data"] = Data
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
	
// PaymentsGet
func (srv *Payments) PaymentsGet(Id string)(*models.Payment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}")
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

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PaymentsCancel
func (srv *Payments) PaymentsCancel(Id string)(*models.Payment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/cancel")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PaymentsCapture
func (srv *Payments) PaymentsCapture(Id string)(*models.Payment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/capture")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PaymentsConfirm
func (srv *Payments) PaymentsConfirm(Id string)(*models.Payment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/confirm")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PaymentsRefund
func (srv *Payments) PaymentsRefund(Id string)(*models.Payment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/refund")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Payment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Payment
	parsed, ok := resp.Result.(models.Payment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
