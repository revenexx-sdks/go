package orders

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Orders service
type Orders struct {
	client client.Client
}

func New(clt client.Client) *Orders {
	return &Orders{
		client: clt,
	}
}


// OrdersList
func (srv *Orders) OrdersList()(*interface{}, error) {
	path := "/v1/orders"
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

// OrdersNumberRangesList
func (srv *Orders) OrdersNumberRangesList()(*interface{}, error) {
	path := "/v1/orders/number-ranges"
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
type OrdersNumberRangesCreateOptions struct {
	ChannelId string
	Counter int
	Metadata interface{}
	Padding int
	PositionStep int
	Prefix string
	Step int
	Suffix string
	enabledSetters map[string]bool
}
func (options OrdersNumberRangesCreateOptions) New() *OrdersNumberRangesCreateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"Counter": false,
		"Metadata": false,
		"Padding": false,
		"PositionStep": false,
		"Prefix": false,
		"Step": false,
		"Suffix": false,
	}
	return &options
}
type OrdersNumberRangesCreateOption func(*OrdersNumberRangesCreateOptions)
func (srv *Orders) WithOrdersNumberRangesCreateChannelId(v string) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreateCounter(v int) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Counter = v
		o.enabledSetters["Counter"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreateMetadata(v interface{}) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreatePadding(v int) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Padding = v
		o.enabledSetters["Padding"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreatePositionStep(v int) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.PositionStep = v
		o.enabledSetters["PositionStep"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreatePrefix(v string) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreateStep(v int) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Step = v
		o.enabledSetters["Step"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesCreateSuffix(v string) OrdersNumberRangesCreateOption {
	return func(o *OrdersNumberRangesCreateOptions) {
		o.Suffix = v
		o.enabledSetters["Suffix"] = true
	}
}
			
// OrdersNumberRangesCreate
func (srv *Orders) OrdersNumberRangesCreate(Code string, optionalSetters ...OrdersNumberRangesCreateOption)(*models.NumberRange, error) {
	path := "/v1/orders/number-ranges"
	options := OrdersNumberRangesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Counter"] {
		params["counter"] = options.Counter
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Padding"] {
		params["padding"] = options.Padding
	}
	if options.enabledSetters["PositionStep"] {
		params["position_step"] = options.PositionStep
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
	}
	if options.enabledSetters["Step"] {
		params["step"] = options.Step
	}
	if options.enabledSetters["Suffix"] {
		params["suffix"] = options.Suffix
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

		parsed := models.NumberRange{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.NumberRange
	parsed, ok := resp.Result.(models.NumberRange)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// OrdersNumberRangesDefaults
func (srv *Orders) OrdersNumberRangesDefaults()(*interface{}, error) {
	path := "/v1/orders/number-ranges/defaults"
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
	
// OrdersNumberRangesDelete
func (srv *Orders) OrdersNumberRangesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/number-ranges/{id}")
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
	
// OrdersNumberRangesGet
func (srv *Orders) OrdersNumberRangesGet(Id string)(*models.NumberRange, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/number-ranges/{id}")
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

		parsed := models.NumberRange{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.NumberRange
	parsed, ok := resp.Result.(models.NumberRange)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersNumberRangesUpdateOptions struct {
	ChannelId string
	Code string
	Counter int
	Metadata interface{}
	Padding int
	PositionStep int
	Prefix string
	Step int
	Suffix string
	enabledSetters map[string]bool
}
func (options OrdersNumberRangesUpdateOptions) New() *OrdersNumberRangesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"Code": false,
		"Counter": false,
		"Metadata": false,
		"Padding": false,
		"PositionStep": false,
		"Prefix": false,
		"Step": false,
		"Suffix": false,
	}
	return &options
}
type OrdersNumberRangesUpdateOption func(*OrdersNumberRangesUpdateOptions)
func (srv *Orders) WithOrdersNumberRangesUpdateChannelId(v string) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdateCode(v string) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdateCounter(v int) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Counter = v
		o.enabledSetters["Counter"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdateMetadata(v interface{}) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdatePadding(v int) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Padding = v
		o.enabledSetters["Padding"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdatePositionStep(v int) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.PositionStep = v
		o.enabledSetters["PositionStep"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdatePrefix(v string) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdateStep(v int) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Step = v
		o.enabledSetters["Step"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesUpdateSuffix(v string) OrdersNumberRangesUpdateOption {
	return func(o *OrdersNumberRangesUpdateOptions) {
		o.Suffix = v
		o.enabledSetters["Suffix"] = true
	}
}
			
// OrdersNumberRangesUpdate
func (srv *Orders) OrdersNumberRangesUpdate(Id string, optionalSetters ...OrdersNumberRangesUpdateOption)(*models.NumberRange, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/number-ranges/{id}")
	options := OrdersNumberRangesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Counter"] {
		params["counter"] = options.Counter
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Padding"] {
		params["padding"] = options.Padding
	}
	if options.enabledSetters["PositionStep"] {
		params["position_step"] = options.PositionStep
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
	}
	if options.enabledSetters["Step"] {
		params["step"] = options.Step
	}
	if options.enabledSetters["Suffix"] {
		params["suffix"] = options.Suffix
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

		parsed := models.NumberRange{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.NumberRange
	parsed, ok := resp.Result.(models.NumberRange)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersPlaceOptions struct {
	BillingAddress interface{}
	Buyer interface{}
	CartId string
	ChannelId string
	ContactId string
	Currency string
	CustomerOrderNumber string
	GrandTotal float64
	MarketId string
	Metadata interface{}
	OrganizationId string
	Payment interface{}
	Shipping interface{}
	ShippingAddress interface{}
	ShippingTotal float64
	UserData interface{}
	enabledSetters map[string]bool
}
func (options OrdersPlaceOptions) New() *OrdersPlaceOptions {
	options.enabledSetters = map[string]bool{
		"BillingAddress": false,
		"Buyer": false,
		"CartId": false,
		"ChannelId": false,
		"ContactId": false,
		"Currency": false,
		"CustomerOrderNumber": false,
		"GrandTotal": false,
		"MarketId": false,
		"Metadata": false,
		"OrganizationId": false,
		"Payment": false,
		"Shipping": false,
		"ShippingAddress": false,
		"ShippingTotal": false,
		"UserData": false,
	}
	return &options
}
type OrdersPlaceOption func(*OrdersPlaceOptions)
func (srv *Orders) WithOrdersPlaceBillingAddress(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.BillingAddress = v
		o.enabledSetters["BillingAddress"] = true
	}
}
func (srv *Orders) WithOrdersPlaceBuyer(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.Buyer = v
		o.enabledSetters["Buyer"] = true
	}
}
func (srv *Orders) WithOrdersPlaceCartId(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *Orders) WithOrdersPlaceChannelId(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Orders) WithOrdersPlaceContactId(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Orders) WithOrdersPlaceCurrency(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Orders) WithOrdersPlaceCustomerOrderNumber(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.CustomerOrderNumber = v
		o.enabledSetters["CustomerOrderNumber"] = true
	}
}
func (srv *Orders) WithOrdersPlaceGrandTotal(v float64) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.GrandTotal = v
		o.enabledSetters["GrandTotal"] = true
	}
}
func (srv *Orders) WithOrdersPlaceMarketId(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Orders) WithOrdersPlaceMetadata(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersPlaceOrganizationId(v string) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Orders) WithOrdersPlacePayment(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.Payment = v
		o.enabledSetters["Payment"] = true
	}
}
func (srv *Orders) WithOrdersPlaceShipping(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.Shipping = v
		o.enabledSetters["Shipping"] = true
	}
}
func (srv *Orders) WithOrdersPlaceShippingAddress(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.ShippingAddress = v
		o.enabledSetters["ShippingAddress"] = true
	}
}
func (srv *Orders) WithOrdersPlaceShippingTotal(v float64) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.ShippingTotal = v
		o.enabledSetters["ShippingTotal"] = true
	}
}
func (srv *Orders) WithOrdersPlaceUserData(v interface{}) OrdersPlaceOption {
	return func(o *OrdersPlaceOptions) {
		o.UserData = v
		o.enabledSetters["UserData"] = true
	}
}
			
// OrdersPlace
func (srv *Orders) OrdersPlace(Items []models.OrderItemCreateRequest, optionalSetters ...OrdersPlaceOption)(*models.OrderDetail, error) {
	path := "/v1/orders/place"
	options := OrdersPlaceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["BillingAddress"] {
		params["billing_address"] = options.BillingAddress
	}
	if options.enabledSetters["Buyer"] {
		params["buyer"] = options.Buyer
	}
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["CustomerOrderNumber"] {
		params["customer_order_number"] = options.CustomerOrderNumber
	}
	if options.enabledSetters["GrandTotal"] {
		params["grand_total"] = options.GrandTotal
	}
	if options.enabledSetters["MarketId"] {
		params["market_id"] = options.MarketId
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Payment"] {
		params["payment"] = options.Payment
	}
	if options.enabledSetters["Shipping"] {
		params["shipping"] = options.Shipping
	}
	if options.enabledSetters["ShippingAddress"] {
		params["shipping_address"] = options.ShippingAddress
	}
	if options.enabledSetters["ShippingTotal"] {
		params["shipping_total"] = options.ShippingTotal
	}
	if options.enabledSetters["UserData"] {
		params["user_data"] = options.UserData
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

		parsed := models.OrderDetail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderDetail
	parsed, ok := resp.Result.(models.OrderDetail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrdersGet
func (srv *Orders) OrdersGet(Id string)(*models.OrderDetail, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}")
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

		parsed := models.OrderDetail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderDetail
	parsed, ok := resp.Result.(models.OrderDetail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersUpdateOptions struct {
	BillingAddress interface{}
	Buyer interface{}
	CustomerOrderNumber string
	Metadata interface{}
	ShippingAddress interface{}
	UserData interface{}
	enabledSetters map[string]bool
}
func (options OrdersUpdateOptions) New() *OrdersUpdateOptions {
	options.enabledSetters = map[string]bool{
		"BillingAddress": false,
		"Buyer": false,
		"CustomerOrderNumber": false,
		"Metadata": false,
		"ShippingAddress": false,
		"UserData": false,
	}
	return &options
}
type OrdersUpdateOption func(*OrdersUpdateOptions)
func (srv *Orders) WithOrdersUpdateBillingAddress(v interface{}) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.BillingAddress = v
		o.enabledSetters["BillingAddress"] = true
	}
}
func (srv *Orders) WithOrdersUpdateBuyer(v interface{}) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.Buyer = v
		o.enabledSetters["Buyer"] = true
	}
}
func (srv *Orders) WithOrdersUpdateCustomerOrderNumber(v string) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.CustomerOrderNumber = v
		o.enabledSetters["CustomerOrderNumber"] = true
	}
}
func (srv *Orders) WithOrdersUpdateMetadata(v interface{}) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersUpdateShippingAddress(v interface{}) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.ShippingAddress = v
		o.enabledSetters["ShippingAddress"] = true
	}
}
func (srv *Orders) WithOrdersUpdateUserData(v interface{}) OrdersUpdateOption {
	return func(o *OrdersUpdateOptions) {
		o.UserData = v
		o.enabledSetters["UserData"] = true
	}
}
			
// OrdersUpdate
func (srv *Orders) OrdersUpdate(Id string, optionalSetters ...OrdersUpdateOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}")
	options := OrdersUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["BillingAddress"] {
		params["billing_address"] = options.BillingAddress
	}
	if options.enabledSetters["Buyer"] {
		params["buyer"] = options.Buyer
	}
	if options.enabledSetters["CustomerOrderNumber"] {
		params["customer_order_number"] = options.CustomerOrderNumber
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["ShippingAddress"] {
		params["shipping_address"] = options.ShippingAddress
	}
	if options.enabledSetters["UserData"] {
		params["user_data"] = options.UserData
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersAcknowledgeOptions struct {
	ExternalRef string
	enabledSetters map[string]bool
}
func (options OrdersAcknowledgeOptions) New() *OrdersAcknowledgeOptions {
	options.enabledSetters = map[string]bool{
		"ExternalRef": false,
	}
	return &options
}
type OrdersAcknowledgeOption func(*OrdersAcknowledgeOptions)
func (srv *Orders) WithOrdersAcknowledgeExternalRef(v string) OrdersAcknowledgeOption {
	return func(o *OrdersAcknowledgeOptions) {
		o.ExternalRef = v
		o.enabledSetters["ExternalRef"] = true
	}
}
			
// OrdersAcknowledge
func (srv *Orders) OrdersAcknowledge(Id string, optionalSetters ...OrdersAcknowledgeOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/acknowledge")
	options := OrdersAcknowledgeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ExternalRef"] {
		params["external_ref"] = options.ExternalRef
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersCancelOptions struct {
	CancelledBy string
	Reason string
	enabledSetters map[string]bool
}
func (options OrdersCancelOptions) New() *OrdersCancelOptions {
	options.enabledSetters = map[string]bool{
		"CancelledBy": false,
		"Reason": false,
	}
	return &options
}
type OrdersCancelOption func(*OrdersCancelOptions)
func (srv *Orders) WithOrdersCancelCancelledBy(v string) OrdersCancelOption {
	return func(o *OrdersCancelOptions) {
		o.CancelledBy = v
		o.enabledSetters["CancelledBy"] = true
	}
}
func (srv *Orders) WithOrdersCancelReason(v string) OrdersCancelOption {
	return func(o *OrdersCancelOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// OrdersCancel
func (srv *Orders) OrdersCancel(Id string, optionalSetters ...OrdersCancelOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/cancel")
	options := OrdersCancelOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["CancelledBy"] {
		params["cancelled_by"] = options.CancelledBy
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrdersCommentsList
func (srv *Orders) OrdersCommentsList(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/comments")
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
type OrdersCommentsCreateOptions struct {
	Author string
	Visibility string
	enabledSetters map[string]bool
}
func (options OrdersCommentsCreateOptions) New() *OrdersCommentsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Author": false,
		"Visibility": false,
	}
	return &options
}
type OrdersCommentsCreateOption func(*OrdersCommentsCreateOptions)
func (srv *Orders) WithOrdersCommentsCreateAuthor(v string) OrdersCommentsCreateOption {
	return func(o *OrdersCommentsCreateOptions) {
		o.Author = v
		o.enabledSetters["Author"] = true
	}
}
func (srv *Orders) WithOrdersCommentsCreateVisibility(v string) OrdersCommentsCreateOption {
	return func(o *OrdersCommentsCreateOptions) {
		o.Visibility = v
		o.enabledSetters["Visibility"] = true
	}
}
					
// OrdersCommentsCreate
func (srv *Orders) OrdersCommentsCreate(Id string, Body string, optionalSetters ...OrdersCommentsCreateOption)(*models.OrderComment, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/comments")
	options := OrdersCommentsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["body"] = Body
	if options.enabledSetters["Author"] {
		params["author"] = options.Author
	}
	if options.enabledSetters["Visibility"] {
		params["visibility"] = options.Visibility
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

		parsed := models.OrderComment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderComment
	parsed, ok := resp.Result.(models.OrderComment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrdersEventsList
func (srv *Orders) OrdersEventsList(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/events")
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
type OrdersHoldOptions struct {
	Reason string
	enabledSetters map[string]bool
}
func (options OrdersHoldOptions) New() *OrdersHoldOptions {
	options.enabledSetters = map[string]bool{
		"Reason": false,
	}
	return &options
}
type OrdersHoldOption func(*OrdersHoldOptions)
func (srv *Orders) WithOrdersHoldReason(v string) OrdersHoldOption {
	return func(o *OrdersHoldOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// OrdersHold
func (srv *Orders) OrdersHold(Id string, optionalSetters ...OrdersHoldOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/hold")
	options := OrdersHoldOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersItemsCancelOptions struct {
	CancelledBy string
	Reason string
	enabledSetters map[string]bool
}
func (options OrdersItemsCancelOptions) New() *OrdersItemsCancelOptions {
	options.enabledSetters = map[string]bool{
		"CancelledBy": false,
		"Reason": false,
	}
	return &options
}
type OrdersItemsCancelOption func(*OrdersItemsCancelOptions)
func (srv *Orders) WithOrdersItemsCancelCancelledBy(v string) OrdersItemsCancelOption {
	return func(o *OrdersItemsCancelOptions) {
		o.CancelledBy = v
		o.enabledSetters["CancelledBy"] = true
	}
}
func (srv *Orders) WithOrdersItemsCancelReason(v string) OrdersItemsCancelOption {
	return func(o *OrdersItemsCancelOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
					
// OrdersItemsCancel
func (srv *Orders) OrdersItemsCancel(Id string, Positions []models.OrderCancelPosition, optionalSetters ...OrdersItemsCancelOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/items/cancel")
	options := OrdersItemsCancelOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["positions"] = Positions
	if options.enabledSetters["CancelledBy"] {
		params["cancelled_by"] = options.CancelledBy
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersPaymentStatusUpdateOptions struct {
	PaymentId string
	enabledSetters map[string]bool
}
func (options OrdersPaymentStatusUpdateOptions) New() *OrdersPaymentStatusUpdateOptions {
	options.enabledSetters = map[string]bool{
		"PaymentId": false,
	}
	return &options
}
type OrdersPaymentStatusUpdateOption func(*OrdersPaymentStatusUpdateOptions)
func (srv *Orders) WithOrdersPaymentStatusUpdatePaymentId(v string) OrdersPaymentStatusUpdateOption {
	return func(o *OrdersPaymentStatusUpdateOptions) {
		o.PaymentId = v
		o.enabledSetters["PaymentId"] = true
	}
}
					
// OrdersPaymentStatusUpdate
func (srv *Orders) OrdersPaymentStatusUpdate(Id string, Status string, optionalSetters ...OrdersPaymentStatusUpdateOption)(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/payment-status")
	options := OrdersPaymentStatusUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["status"] = Status
	if options.enabledSetters["PaymentId"] {
		params["payment_id"] = options.PaymentId
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersReturnOptions struct {
	Metadata interface{}
	Reason string
	enabledSetters map[string]bool
}
func (options OrdersReturnOptions) New() *OrdersReturnOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"Reason": false,
	}
	return &options
}
type OrdersReturnOption func(*OrdersReturnOptions)
func (srv *Orders) WithOrdersReturnMetadata(v interface{}) OrdersReturnOption {
	return func(o *OrdersReturnOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersReturnReason(v string) OrdersReturnOption {
	return func(o *OrdersReturnOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
					
// OrdersReturn
func (srv *Orders) OrdersReturn(Id string, Positions []models.OrderReturnPosition, optionalSetters ...OrdersReturnOption)(*models.OrderReturn, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/return")
	options := OrdersReturnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["positions"] = Positions
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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

		parsed := models.OrderReturn{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderReturn
	parsed, ok := resp.Result.(models.OrderReturn)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersReturnsCompleteOptions struct {
	Resolution string
	enabledSetters map[string]bool
}
func (options OrdersReturnsCompleteOptions) New() *OrdersReturnsCompleteOptions {
	options.enabledSetters = map[string]bool{
		"Resolution": false,
	}
	return &options
}
type OrdersReturnsCompleteOption func(*OrdersReturnsCompleteOptions)
func (srv *Orders) WithOrdersReturnsCompleteResolution(v string) OrdersReturnsCompleteOption {
	return func(o *OrdersReturnsCompleteOptions) {
		o.Resolution = v
		o.enabledSetters["Resolution"] = true
	}
}
					
// OrdersReturnsComplete
func (srv *Orders) OrdersReturnsComplete(Id string, Rid string, optionalSetters ...OrdersReturnsCompleteOption)(*models.OrderReturn, error) {
	r := strings.NewReplacer("{id}", Id, "{rid}", Rid)
	path := r.Replace("/v1/orders/{id}/returns/{rid}/complete")
	options := OrdersReturnsCompleteOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["rid"] = Rid
	if options.enabledSetters["Resolution"] {
		params["resolution"] = options.Resolution
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

		parsed := models.OrderReturn{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderReturn
	parsed, ok := resp.Result.(models.OrderReturn)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// OrdersReturnsReceive
func (srv *Orders) OrdersReturnsReceive(Id string, Rid string, Data interface{})(*models.OrderReturn, error) {
	r := strings.NewReplacer("{id}", Id, "{rid}", Rid)
	path := r.Replace("/v1/orders/{id}/returns/{rid}/receive")
	params := map[string]interface{}{}
	params["id"] = Id
	params["rid"] = Rid
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

		parsed := models.OrderReturn{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderReturn
	parsed, ok := resp.Result.(models.OrderReturn)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersReturnsRejectOptions struct {
	Reason string
	Resolution string
	enabledSetters map[string]bool
}
func (options OrdersReturnsRejectOptions) New() *OrdersReturnsRejectOptions {
	options.enabledSetters = map[string]bool{
		"Reason": false,
		"Resolution": false,
	}
	return &options
}
type OrdersReturnsRejectOption func(*OrdersReturnsRejectOptions)
func (srv *Orders) WithOrdersReturnsRejectReason(v string) OrdersReturnsRejectOption {
	return func(o *OrdersReturnsRejectOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *Orders) WithOrdersReturnsRejectResolution(v string) OrdersReturnsRejectOption {
	return func(o *OrdersReturnsRejectOptions) {
		o.Resolution = v
		o.enabledSetters["Resolution"] = true
	}
}
					
// OrdersReturnsReject
func (srv *Orders) OrdersReturnsReject(Id string, Rid string, optionalSetters ...OrdersReturnsRejectOption)(*models.OrderReturn, error) {
	r := strings.NewReplacer("{id}", Id, "{rid}", Rid)
	path := r.Replace("/v1/orders/{id}/returns/{rid}/reject")
	options := OrdersReturnsRejectOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["rid"] = Rid
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Resolution"] {
		params["resolution"] = options.Resolution
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

		parsed := models.OrderReturn{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderReturn
	parsed, ok := resp.Result.(models.OrderReturn)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type OrdersShipOptions struct {
	Carrier string
	Metadata interface{}
	Number string
	Positions []models.OrderShipmentPosition
	ShippedAt string
	TrackingCode string
	TrackingUrl string
	enabledSetters map[string]bool
}
func (options OrdersShipOptions) New() *OrdersShipOptions {
	options.enabledSetters = map[string]bool{
		"Carrier": false,
		"Metadata": false,
		"Number": false,
		"Positions": false,
		"ShippedAt": false,
		"TrackingCode": false,
		"TrackingUrl": false,
	}
	return &options
}
type OrdersShipOption func(*OrdersShipOptions)
func (srv *Orders) WithOrdersShipCarrier(v string) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.Carrier = v
		o.enabledSetters["Carrier"] = true
	}
}
func (srv *Orders) WithOrdersShipMetadata(v interface{}) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orders) WithOrdersShipNumber(v string) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.Number = v
		o.enabledSetters["Number"] = true
	}
}
func (srv *Orders) WithOrdersShipPositions(v []models.OrderShipmentPosition) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.Positions = v
		o.enabledSetters["Positions"] = true
	}
}
func (srv *Orders) WithOrdersShipShippedAt(v string) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.ShippedAt = v
		o.enabledSetters["ShippedAt"] = true
	}
}
func (srv *Orders) WithOrdersShipTrackingCode(v string) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.TrackingCode = v
		o.enabledSetters["TrackingCode"] = true
	}
}
func (srv *Orders) WithOrdersShipTrackingUrl(v string) OrdersShipOption {
	return func(o *OrdersShipOptions) {
		o.TrackingUrl = v
		o.enabledSetters["TrackingUrl"] = true
	}
}
			
// OrdersShip
func (srv *Orders) OrdersShip(Id string, optionalSetters ...OrdersShipOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/ship")
	options := OrdersShipOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Carrier"] {
		params["carrier"] = options.Carrier
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Number"] {
		params["number"] = options.Number
	}
	if options.enabledSetters["Positions"] {
		params["positions"] = options.Positions
	}
	if options.enabledSetters["ShippedAt"] {
		params["shipped_at"] = options.ShippedAt
	}
	if options.enabledSetters["TrackingCode"] {
		params["tracking_code"] = options.TrackingCode
	}
	if options.enabledSetters["TrackingUrl"] {
		params["tracking_url"] = options.TrackingUrl
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
			
// OrdersUnhold
func (srv *Orders) OrdersUnhold(Id string, Data interface{})(*models.Order, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/unhold")
	params := map[string]interface{}{}
	params["id"] = Id
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

		parsed := models.Order{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Order
	parsed, ok := resp.Result.(models.Order)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
