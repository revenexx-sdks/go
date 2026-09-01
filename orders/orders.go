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

type OrdersListOptions struct {
	Id string
	Number string
	CustomerOrderNumber string
	ExternalRef string
	AcknowledgedAt string
	CartId string
	ContactId string
	OrganizationId string
	ChannelId string
	Currency string
	Status string
	PaymentStatus string
	FulfillmentStatus string
	OnHold bool
	HoldReason string
	ItemCount int
	Subtotal float64
	ShippingTotal float64
	TaxTotal float64
	GrandTotal float64
	PlacedAt string
	CompletedAt string
	CancelledAt string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrdersListOptions) New() *OrdersListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Number": false,
		"CustomerOrderNumber": false,
		"ExternalRef": false,
		"AcknowledgedAt": false,
		"CartId": false,
		"ContactId": false,
		"OrganizationId": false,
		"ChannelId": false,
		"Currency": false,
		"Status": false,
		"PaymentStatus": false,
		"FulfillmentStatus": false,
		"OnHold": false,
		"HoldReason": false,
		"ItemCount": false,
		"Subtotal": false,
		"ShippingTotal": false,
		"TaxTotal": false,
		"GrandTotal": false,
		"PlacedAt": false,
		"CompletedAt": false,
		"CancelledAt": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrdersListOption func(*OrdersListOptions)
func (srv *Orders) WithOrdersListId(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Orders) WithOrdersListNumber(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Number = v
		o.enabledSetters["Number"] = true
	}
}
func (srv *Orders) WithOrdersListCustomerOrderNumber(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.CustomerOrderNumber = v
		o.enabledSetters["CustomerOrderNumber"] = true
	}
}
func (srv *Orders) WithOrdersListExternalRef(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.ExternalRef = v
		o.enabledSetters["ExternalRef"] = true
	}
}
func (srv *Orders) WithOrdersListAcknowledgedAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.AcknowledgedAt = v
		o.enabledSetters["AcknowledgedAt"] = true
	}
}
func (srv *Orders) WithOrdersListCartId(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *Orders) WithOrdersListContactId(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Orders) WithOrdersListOrganizationId(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Orders) WithOrdersListChannelId(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Orders) WithOrdersListCurrency(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Orders) WithOrdersListStatus(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Orders) WithOrdersListPaymentStatus(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.PaymentStatus = v
		o.enabledSetters["PaymentStatus"] = true
	}
}
func (srv *Orders) WithOrdersListFulfillmentStatus(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.FulfillmentStatus = v
		o.enabledSetters["FulfillmentStatus"] = true
	}
}
func (srv *Orders) WithOrdersListOnHold(v bool) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.OnHold = v
		o.enabledSetters["OnHold"] = true
	}
}
func (srv *Orders) WithOrdersListHoldReason(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.HoldReason = v
		o.enabledSetters["HoldReason"] = true
	}
}
func (srv *Orders) WithOrdersListItemCount(v int) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.ItemCount = v
		o.enabledSetters["ItemCount"] = true
	}
}
func (srv *Orders) WithOrdersListSubtotal(v float64) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Subtotal = v
		o.enabledSetters["Subtotal"] = true
	}
}
func (srv *Orders) WithOrdersListShippingTotal(v float64) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.ShippingTotal = v
		o.enabledSetters["ShippingTotal"] = true
	}
}
func (srv *Orders) WithOrdersListTaxTotal(v float64) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.TaxTotal = v
		o.enabledSetters["TaxTotal"] = true
	}
}
func (srv *Orders) WithOrdersListGrandTotal(v float64) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.GrandTotal = v
		o.enabledSetters["GrandTotal"] = true
	}
}
func (srv *Orders) WithOrdersListPlacedAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.PlacedAt = v
		o.enabledSetters["PlacedAt"] = true
	}
}
func (srv *Orders) WithOrdersListCompletedAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.CompletedAt = v
		o.enabledSetters["CompletedAt"] = true
	}
}
func (srv *Orders) WithOrdersListCancelledAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.CancelledAt = v
		o.enabledSetters["CancelledAt"] = true
	}
}
func (srv *Orders) WithOrdersListCreatedAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Orders) WithOrdersListUpdatedAt(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Orders) WithOrdersListLimit(v int) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orders) WithOrdersListOffset(v int) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orders) WithOrdersListOrder(v string) OrdersListOption {
	return func(o *OrdersListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// OrdersList the route behind every order overview: the open orders of one
// customer, everything on hold, everything a market placed last week, or the
// one order somebody is quoting a number for (?number=ORD-000123 — the
// number is not the id, and this is how one becomes the other). The order
// LIST: the order rows without their positions, shipments, returns or
// cancellations — read GET /orders/{id} for the aggregate of one. Every
// parameter below is an exact match on the column it names, and combining
// them is an AND. Two kinds of key are not offered: one that names NO column
// is dropped silently, so a mistyped ?stauts=placed answers 200 with the
// whole list (compare the 'filter' echo against what you sent — no status
// code reports it), and the jsonb columns buyer, billing_address,
// shipping_address, payment, shipping, user_data and metadata reach the
// database as a text comparison and answer 400 invalid_value for anything
// that is not a whole JSON document.
func (srv *Orders) OrdersList(optionalSetters ...OrdersListOption)(*interface{}, error) {
	path := "/v1/orders"
	options := OrdersListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Number"] {
		params["number"] = options.Number
	}
	if options.enabledSetters["CustomerOrderNumber"] {
		params["customer_order_number"] = options.CustomerOrderNumber
	}
	if options.enabledSetters["ExternalRef"] {
		params["external_ref"] = options.ExternalRef
	}
	if options.enabledSetters["AcknowledgedAt"] {
		params["acknowledged_at"] = options.AcknowledgedAt
	}
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["PaymentStatus"] {
		params["payment_status"] = options.PaymentStatus
	}
	if options.enabledSetters["FulfillmentStatus"] {
		params["fulfillment_status"] = options.FulfillmentStatus
	}
	if options.enabledSetters["OnHold"] {
		params["on_hold"] = options.OnHold
	}
	if options.enabledSetters["HoldReason"] {
		params["hold_reason"] = options.HoldReason
	}
	if options.enabledSetters["ItemCount"] {
		params["item_count"] = options.ItemCount
	}
	if options.enabledSetters["Subtotal"] {
		params["subtotal"] = options.Subtotal
	}
	if options.enabledSetters["ShippingTotal"] {
		params["shipping_total"] = options.ShippingTotal
	}
	if options.enabledSetters["TaxTotal"] {
		params["tax_total"] = options.TaxTotal
	}
	if options.enabledSetters["GrandTotal"] {
		params["grand_total"] = options.GrandTotal
	}
	if options.enabledSetters["PlacedAt"] {
		params["placed_at"] = options.PlacedAt
	}
	if options.enabledSetters["CompletedAt"] {
		params["completed_at"] = options.CompletedAt
	}
	if options.enabledSetters["CancelledAt"] {
		params["cancelled_at"] = options.CancelledAt
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
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
type OrdersNumberRangesListOptions struct {
	Id string
	Code string
	Prefix string
	Suffix string
	Padding int
	Counter int
	Step int
	PositionStep int
	ChannelId string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrdersNumberRangesListOptions) New() *OrdersNumberRangesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Prefix": false,
		"Suffix": false,
		"Padding": false,
		"Counter": false,
		"Step": false,
		"PositionStep": false,
		"ChannelId": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrdersNumberRangesListOption func(*OrdersNumberRangesListOptions)
func (srv *Orders) WithOrdersNumberRangesListId(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListCode(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListPrefix(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListSuffix(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Suffix = v
		o.enabledSetters["Suffix"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListPadding(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Padding = v
		o.enabledSetters["Padding"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListCounter(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Counter = v
		o.enabledSetters["Counter"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListStep(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Step = v
		o.enabledSetters["Step"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListPositionStep(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.PositionStep = v
		o.enabledSetters["PositionStep"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListChannelId(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListCreatedAt(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListUpdatedAt(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListLimit(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListOffset(v int) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orders) WithOrdersNumberRangesListOrder(v string) OrdersNumberRangesListOption {
	return func(o *OrdersNumberRangesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// OrdersNumberRangesList the counters this tenant numbers its orders,
// delivery notes and returns from — what an operator sees on the Number
// ranges settings page, and what a migration reads to check the prefixes and
// the padding before it imports anything. Every parameter below is an
// exact-match filter on the column it names (?code=order finds the order
// counter). Two things are not: a key that names NO column is dropped
// silently — the call answers 200 with the unfiltered page, so compare the
// 'filter' echo against what you sent — and the jsonb column 'metadata' is
// honoured by the router but refused by the database (400 invalid_value)
// unless the value is a whole JSON document, which is why it is not offered
// here. It does not draw a number: `counter` is the last number DRAWN, and
// only placing an order, a shipment or a return moves it.
func (srv *Orders) OrdersNumberRangesList(optionalSetters ...OrdersNumberRangesListOption)(*interface{}, error) {
	path := "/v1/orders/number-ranges"
	options := OrdersNumberRangesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
	}
	if options.enabledSetters["Suffix"] {
		params["suffix"] = options.Suffix
	}
	if options.enabledSetters["Padding"] {
		params["padding"] = options.Padding
	}
	if options.enabledSetters["Counter"] {
		params["counter"] = options.Counter
	}
	if options.enabledSetters["Step"] {
		params["step"] = options.Step
	}
	if options.enabledSetters["PositionStep"] {
		params["position_step"] = options.PositionStep
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
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
			
// OrdersNumberRangesCreate add a counter beyond the three a tenant is seeded
// with, and give it the shape a merchant's numbers actually have:
// {prefix}{counter padded to `padding`}{suffix}, moving by `step` per draw. A
// new range is what the order_number_range_code / delivery_number_range_code
// / return_number_range_code settings can then be pointed at — the code is
// the name those settings use, and a setting naming a code no range carries
// makes placing an order answer 422. `code` is unique per tenant, so this is
// a 409 for one that is taken rather than a second counter under the same
// name. It does not renumber anything that already exists, and setting
// `counter` to a value already issued re-issues those numbers, which the
// unique index on the order number then refuses.
func (srv *Orders) OrdersNumberRangesCreate(Code string, optionalSetters ...OrdersNumberRangesCreateOption)(*models.Error, error) {
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

// OrdersNumberRangesDefaults make sure the three codes this app draws from
// exist: 'order' (ORD-), 'delivery' (DEL-) and 'return' (RET-), each padded
// to six digits and stepping by one. The app runs it for you on install, so a
// fresh tenant needs nothing; call it by hand after a range was deleted, or
// to check what a tenant has. Idempotent: a code that already exists comes
// back under 'existing' and is left EXACTLY as it is, counter included, so a
// merchant who changed the prefix keeps their change. Answers 200, never 201
// — it is a reconcile, not a create — and it never repairs or renames a
// range that is already there.
func (srv *Orders) OrdersNumberRangesDefaults()(*models.OrderNumberRangesSeeded, error) {
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

		parsed := models.OrderNumberRangesSeeded{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderNumberRangesSeeded
	parsed, ok := resp.Result.(models.OrderNumberRangesSeeded)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrdersNumberRangesDelete remove a counter a tenant no longer numbers
// anything from. It touches nothing that was numbered out of it: existing
// orders, delivery notes and returns keep the numbers they were given,
// because a number is copied onto the row at place-time and is not a
// reference to this table. Deleting one of the three standard codes is
// allowed and is usually a mistake — the next draw against it answers 422
// 'number_range_missing', unless POST /orders/number-ranges/defaults or a
// reinstall seeds it again, which starts its counter back at 0.
func (srv *Orders) OrdersNumberRangesDelete(Id string)(*models.Error, error) {
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
	
// OrdersNumberRangesGet one counter with its whole configuration: the prefix
// and suffix around the number, how wide it is padded, how far each draw
// moves it, where it currently stands, and the position_step new order lines
// are numbered in. Reach for it when you hold the id — from the list, or
// from what a create answered — and want the row as it stands now. Reading
// does not draw a number and does not move `counter`; the id is the range's
// uuid, not its `code`, and a code is turned into a range through GET
// /orders/number-ranges?code=order.
func (srv *Orders) OrdersNumberRangesGet(Id string)(*models.Error, error) {
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
			
// OrdersNumberRangesUpdate change the format or the state of an existing
// counter: a new prefix or suffix, a wider padding, a different step, a
// different position_step for new order lines — or `counter` itself, which
// is state rather than configuration. Everything takes effect on the NEXT
// draw only: nothing that was already numbered is renumbered, so widening the
// padding leaves ORD-000123 and starts writing ORD-0000124. Moving `counter`
// forward skips numbers, and moving it back re-issues numbers that exist,
// which the unique index on the order number answers 409 for at place-time
// rather than here. Renaming `code` to one another range of this tenant
// already holds is a 409.
func (srv *Orders) OrdersNumberRangesUpdate(Id string, optionalSetters ...OrdersNumberRangesUpdateOption)(*models.Error, error) {
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
type OrdersPlaceOptions struct {
	BillingAddress interface{}
	Buyer interface{}
	CartId string
	ChannelId string
	ContactId string
	Currency string
	CustomerOrderNumber string
	GrandTotal float64
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
			
// OrdersPlace the way an order comes into existence — the call a checkout,
// a punch-out or an ERP import makes once the basket is final. The body is a
// SNAPSHOT: items with their product copies, plus the buyer, the addresses
// and the payment and shipping choices frozen as they were at this moment, so
// the order stays readable when the catalogue or the customer changes
// underneath it. The app draws the order number from the tenant's order
// range, numbers the positions, computes subtotal, tax and grand_total from
// the lines, and writes the order.placed event that carries the order onto
// the bus. It does not reserve stock, take payment or talk to an ERP: those
// are separate capabilities, and this route's job ends when the event is on
// the bus. Two things can turn a placement into a REQUEST awaiting approval,
// and both still answer 201 — with status='pending' and no placed_at: a
// principal holding only orders.request, and an order worth more than the
// tenant's require_approval_above_value (a principal holding orders.approve
// is exempt from the threshold). The order.requested event says which, in
// 'approval_reason'. The currency defaults to the market's default_currency
// setting and the position cap is the tenant's max_items_per_order.
func (srv *Orders) OrdersPlace(Items []models.OrderItemCreateRequest, optionalSetters ...OrdersPlaceOption)(*models.Error, error) {
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
type OrdersReportsCustomerRollupOptions struct {
	AsOf string
	Cursor string
	OrganizationIds []string
	Statuses []string
	enabledSetters map[string]bool
}
func (options OrdersReportsCustomerRollupOptions) New() *OrdersReportsCustomerRollupOptions {
	options.enabledSetters = map[string]bool{
		"AsOf": false,
		"Cursor": false,
		"OrganizationIds": false,
		"Statuses": false,
	}
	return &options
}
type OrdersReportsCustomerRollupOption func(*OrdersReportsCustomerRollupOptions)
func (srv *Orders) WithOrdersReportsCustomerRollupAsOf(v string) OrdersReportsCustomerRollupOption {
	return func(o *OrdersReportsCustomerRollupOptions) {
		o.AsOf = v
		o.enabledSetters["AsOf"] = true
	}
}
func (srv *Orders) WithOrdersReportsCustomerRollupCursor(v string) OrdersReportsCustomerRollupOption {
	return func(o *OrdersReportsCustomerRollupOptions) {
		o.Cursor = v
		o.enabledSetters["Cursor"] = true
	}
}
func (srv *Orders) WithOrdersReportsCustomerRollupOrganizationIds(v []string) OrdersReportsCustomerRollupOption {
	return func(o *OrdersReportsCustomerRollupOptions) {
		o.OrganizationIds = v
		o.enabledSetters["OrganizationIds"] = true
	}
}
func (srv *Orders) WithOrdersReportsCustomerRollupStatuses(v []string) OrdersReportsCustomerRollupOption {
	return func(o *OrdersReportsCustomerRollupOptions) {
		o.Statuses = v
		o.enabledSetters["Statuses"] = true
	}
}
	
// OrdersReportsCustomerRollup what each company has bought, as numbers
// another app can keep: order count, lifetime revenue, first and last order
// date, and the same count and revenue over the last 30, 90 and 365 days.
// This is what a customer segment like "bought for more than 100k last year"
// is built on, and the customers app materialises it into a local projection
// its segment rules query. It answers about ORGANIZATIONS only — a private
// or guest order carries none and is counted in orders_without_organization
// rather than attributed to anybody — and it converts nothing, so an
// organization that ordered in two currencies gets both listed and one summed
// number to read with care. Revenue lives in orders, customer segments live
// in the customers app, and the two may not join (ADR-0055: no cross-app FK,
// grant or view). This capability is the hand-over. Every number is additive
// (count/sum/min/max) so partial answers merge; the average order value is
// deliberately not returned — it is revenue_total / order_count over the
// merged parts. Windows are anchored at as_of, which is echoed back so a loop
// measures one consistent picture.
func (srv *Orders) OrdersReportsCustomerRollup(optionalSetters ...OrdersReportsCustomerRollupOption)(*models.OrderCustomerRollupResponse, error) {
	path := "/v1/orders/reports/customer-rollup"
	options := OrdersReportsCustomerRollupOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["AsOf"] {
		params["as_of"] = options.AsOf
	}
	if options.enabledSetters["Cursor"] {
		params["cursor"] = options.Cursor
	}
	if options.enabledSetters["OrganizationIds"] {
		params["organization_ids"] = options.OrganizationIds
	}
	if options.enabledSetters["Statuses"] {
		params["statuses"] = options.Statuses
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

		parsed := models.OrderCustomerRollupResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderCustomerRollupResponse
	parsed, ok := resp.Result.(models.OrderCustomerRollupResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// OrdersVocabulariesList which value sets this app will describe for you, by
// name — order statuses, payment statuses, fulfillment statuses, item
// types, return statuses and return resolutions — so a client can discover
// them instead of shipping its own copy of five statuses that goes stale one
// release later. The values themselves are deliberately NOT here: this is the
// index, and each set is fetched on its own. Discovery for the vocabulary
// routes. Names: cancellation-scopes, comment-visibilities,
// fulfillment-statuses, item-types, payment-statuses, return-resolutions,
// return-statuses, statuses. Fetch one with GET /orders/vocabularies/{name};
// a client holding the qualified pair 'orders.<name>' builds that URL from
// the pair alone. 'title' and 'description' are locale maps wherever somebody
// wrote the copy and plain strings where the fallback did — read both
// forms.
func (srv *Orders) OrdersVocabulariesList()(*models.OrderVocabularyIndex, error) {
	path := "/v1/orders/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.OrderVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderVocabularyIndex
	parsed, ok := resp.Result.(models.OrderVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrdersVocabulariesGet everything a UI needs to render one of this app's
// value sets without knowing it: every permitted value, in order, each with a
// title and description in the locales somebody wrote and a badge tone to
// colour it. Fetch it once and a status filter, a status badge and a
// resolution picker all stay correct through a lifecycle change, because the
// set served IS the set enforced. It answers about values, not about rows —
// nothing here says how many orders are in a status. The values are read out
// of the column's CHECK constraint, so the served set IS the enforced set and
// the two cannot drift — a value added to the constraint appears here even
// before anyone labels it, titled from its own key. Values come back in
// constraint order, which is lifecycle order for a status, and 'final' marks
// the values that END the lifecycle (completed, cancelled) so a client can
// ask "is this order still open?" instead of matching names it guessed. Every
// set is exhaustive ('closed' is always true); 'source' says who enforces it
// — 'schema' for a CHECK constraint, 'app' for 'return-resolutions', whose
// column carries none and whose words the return routes enforce instead.
// Those values additionally carry 'stage' (complete | reject): the transition
// that accepts them. 'title' and 'description' are locale maps where the copy
// was written and plain strings where the key-derived fallback answered, on
// the vocabulary and on every value alike. Names: cancellation-scopes,
// comment-visibilities, fulfillment-statuses, item-types, payment-statuses,
// return-resolutions, return-statuses, statuses.
func (srv *Orders) OrdersVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/orders/vocabularies/{name}")
	params := map[string]interface{}{}
	params["name"] = Name
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
	
// OrdersGet the single source of order information, and what an order detail
// screen is built from: the order row plus its positions, its shipments with
// the shipment_items each one booked, its returns and its cancellations —
// one call, no assembling five lists. A cancellation's and a return's
// 'positions' are ARRAYS of {order_item_id, quantity}; a return's entries
// additionally carry 'restock'. Two things it does not carry: the comments
// and the event trail, which are their own paginated routes because both grow
// without bound. Addressed by uuid — an order number goes through GET
// /orders?number=… first.
func (srv *Orders) OrdersGet(Id string)(*models.Error, error) {
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
			
// OrdersUpdate the narrow correction window a service desk needs: the
// customer gave the wrong delivery address, the buyer's name is misspelled,
// their purchase-order number was missing. Six columns and no others —
// customer_order_number, buyer, billing_address, shipping_address, user_data
// and metadata — and each is REPLACED whole, not merged, so send the entire
// address rather than the one line that changed. It moves nothing: status,
// payment_status, fulfillment_status and the quantities belong to the action
// routes, and a body carrying them is accepted with those keys quietly
// dropped. The window closes when the fulfilling system acknowledges the
// order, because from then on the ERP holds the copy that ships — unless
// the tenant set allow_modification_after_acknowledge. Every accepted change
// writes an order.updated event naming the columns it touched.
func (srv *Orders) OrdersUpdate(Id string, optionalSetters ...OrdersUpdateOption)(*models.Error, error) {
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
			
// OrdersAcknowledge the return channel for whatever fulfils the order. An
// Integration Studio workflow picks up order.placed, books the order into the
// ERP, and calls this with the id the ERP gave it — which lands in
// external_ref and makes the two systems mutually findable. It stamps
// acknowledged_at from the server's clock, and that timestamp is what closes
// the correction window: PUT /orders/{id} refuses afterwards, because the
// copy that ships now lives elsewhere. It is a handshake and nothing more —
// it does not change status, payment_status or fulfillment_status, and it
// does not ship anything. Once only: a second call is a 422 rather than a
// silent overwrite of the first system's reference.
func (srv *Orders) OrdersAcknowledge(Id string, optionalSetters ...OrdersAcknowledgeOption)(*models.Error, error) {
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
			
// OrdersCancel call the whole order off: every position's full quantity is
// booked as cancelled, the order moves to 'cancelled', a cancellation record
// is written with the reason and who gave it, and an order.cancelled event
// goes onto the bus. Only while NOTHING has shipped — once a single
// position has gone out the order is partly real and this answers 422; take
// the remaining quantities off with POST /orders/{id}/items/cancel instead,
// and handle what already shipped as a return. It refunds nothing and returns
// nothing to stock: payment travels through /payment-status and restocking is
// an explicit inventories call by the orchestrator. A tenant may require a
// reason (cancel_requires_reason), and a hold may block it (on_hold_blocks =
// 'shipping_and_cancel').
func (srv *Orders) OrdersCancel(Id string, optionalSetters ...OrdersCancelOption)(*models.Error, error) {
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
type OrdersCommentsListOptions struct {
	IdQuery string
	Body string
	Visibility string
	Author string
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrdersCommentsListOptions) New() *OrdersCommentsListOptions {
	options.enabledSetters = map[string]bool{
		"IdQuery": false,
		"Body": false,
		"Visibility": false,
		"Author": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrdersCommentsListOption func(*OrdersCommentsListOptions)
func (srv *Orders) WithOrdersCommentsListIdQuery(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.IdQuery = v
		o.enabledSetters["IdQuery"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListBody(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListVisibility(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Visibility = v
		o.enabledSetters["Visibility"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListAuthor(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Author = v
		o.enabledSetters["Author"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListCreatedAt(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListLimit(v int) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListOffset(v int) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orders) WithOrdersCommentsListOrder(v string) OrdersCommentsListOption {
	return func(o *OrdersCommentsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// OrdersCommentsList what people have written about this order, oldest first:
// the service desk's own notes and the messages meant for the customer, in
// one list. Filter by ?visibility=customer to build the version a customer
// may see, and by ?visibility=internal for the desk's own — the route does
// NOT decide that for you, so a customer-facing surface has to ask for the
// customer ones. Comments are prose about the order and never move it; the
// lifecycle lives in the event trail. Every parameter below is an exact match
// on the column it names. `order_id` is deliberately absent: the route fixes
// it from the path AFTER the query filter is read, so sending one is accepted
// and then overwritten — it filters nothing. DEPRECATED KEY: the response
// also repeats 'items' under 'comments' for compatibility with the
// pre-envelope shape. It is the same array; read 'items'. The alias is
// removed in the next minor version.
func (srv *Orders) OrdersCommentsList(Id string, optionalSetters ...OrdersCommentsListOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/comments")
	options := OrdersCommentsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["IdQuery"] {
		params["id"] = options.IdQuery
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Visibility"] {
		params["visibility"] = options.Visibility
	}
	if options.enabledSetters["Author"] {
		params["author"] = options.Author
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
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
					
// OrdersCommentsCreate write down what happened that the state machine cannot
// record: what the customer said on the phone, why an exception was made,
// what the warehouse found in the box. `visibility` decides who the note is
// for — 'internal' for the service desk, 'customer' for text meant to be
// shown to the buyer — and it defaults to the tenant's
// default_comment_visibility, which is 'internal' out of the box, so a note
// is never accidentally customer-facing. Adding one writes an
// order.comment.added event, so the trail shows that a note was made and its
// visibility, without copying the text onto the bus. It changes nothing about
// the order, and it sends nothing to anybody: this stores a comment, it does
// not email the customer.
func (srv *Orders) OrdersCommentsCreate(Id string, Body string, optionalSetters ...OrdersCommentsCreateOption)(*models.Error, error) {
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
type OrdersCompleteOptions struct {
	CompletedBy string
	enabledSetters map[string]bool
}
func (options OrdersCompleteOptions) New() *OrdersCompleteOptions {
	options.enabledSetters = map[string]bool{
		"CompletedBy": false,
	}
	return &options
}
type OrdersCompleteOption func(*OrdersCompleteOptions)
func (srv *Orders) WithOrdersCompleteCompletedBy(v string) OrdersCompleteOption {
	return func(o *OrdersCompleteOptions) {
		o.CompletedBy = v
		o.enabledSetters["CompletedBy"] = true
	}
}
			
// OrdersComplete declare the order finished, whatever the quantities say —
// the service was delivered, the download was fetched, or an operator has
// decided the rest is not coming. status moves to 'completed' and
// completed_at is stamped from the server's clock. It does NOT ship anything
// or change the quantities, so fulfillment_status stays whatever the
// positions make it, and an order completed with lines still open shows
// exactly that. A completed order is final: modification, shipping and
// cancellation all refuse afterwards, and only a return may still be
// registered against it. The counterpart of auto_complete_on = 'payment' |
// 'manual': something has to close an order that shipping no longer closes by
// itself, and it is also the honest end for a service or digital order that
// never ships. Writes an order_events row 'order.completed' with
// via='manual'.
func (srv *Orders) OrdersComplete(Id string, optionalSetters ...OrdersCompleteOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/complete")
	options := OrdersCompleteOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["CompletedBy"] {
		params["completed_by"] = options.CompletedBy
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
type OrdersEventsListOptions struct {
	IdQuery string
	Name string
	Actor string
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrdersEventsListOptions) New() *OrdersEventsListOptions {
	options.enabledSetters = map[string]bool{
		"IdQuery": false,
		"Name": false,
		"Actor": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrdersEventsListOption func(*OrdersEventsListOptions)
func (srv *Orders) WithOrdersEventsListIdQuery(v string) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.IdQuery = v
		o.enabledSetters["IdQuery"] = true
	}
}
func (srv *Orders) WithOrdersEventsListName(v string) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Orders) WithOrdersEventsListActor(v string) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.Actor = v
		o.enabledSetters["Actor"] = true
	}
}
func (srv *Orders) WithOrdersEventsListCreatedAt(v string) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Orders) WithOrdersEventsListLimit(v int) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orders) WithOrdersEventsListOffset(v int) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orders) WithOrdersEventsListOrder(v string) OrdersEventsListOption {
	return func(o *OrdersEventsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// OrdersEventsList everything that has ever happened to this order, oldest
// first: placed or requested, updated, acknowledged, shipped, held, paid,
// returned, completed, cancelled — each with the payload the action
// carried. This is the audit trail an operator reads to answer "why is this
// order in this state", and it is the same row the platform publishes as a
// domain event, so what a workflow reacted to and what a person sees here
// cannot diverge. It is append-only and this route is read-only: rows are
// written by the action routes and there is no way to add, edit or remove
// one. An order's trail grows for as long as the order lives, so it is
// paginated like every other list — 'page.hasMore' says whether more of it
// exists. Every parameter below is an exact match on the column it names;
// `order_id` is deliberately absent, because the route fixes it from the path
// after the query filter is read and a value sent for it is overwritten
// rather than honoured. The jsonb column 'payload' is not offered for the
// same reason it is not offered on the order list: the data plane answers 400
// for anything that is not a whole JSON document. DEPRECATED KEY: the
// response also repeats 'items' under 'events' for compatibility with the
// pre-envelope shape. It is the same array; read 'items'. The alias is
// removed in the next minor version.
func (srv *Orders) OrdersEventsList(Id string, optionalSetters ...OrdersEventsListOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/events")
	options := OrdersEventsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["IdQuery"] {
		params["id"] = options.IdQuery
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Actor"] {
		params["actor"] = options.Actor
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
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
			
// OrdersHold stop an order from moving while a human sorts something out —
// a credit check, a suspected fraud, an address nobody can deliver to. It
// sets a flag with the reason attached, and the flag is deliberately
// ORTHOGONAL to the lifecycle: the order keeps its status, its payment status
// and its quantities, and appears on a worklist as 'held' rather than being
// pushed into a state it will have to come back out of. How far the hold
// reaches is the tenant's setting on_hold_blocks: shipping only, shipping and
// cancellation (the credit-check case, where the order must move in neither
// direction), or nothing at all, which leaves the flag advisory. Holding an
// order twice is allowed and simply replaces the reason; releasing it is POST
// /orders/{id}/unhold.
func (srv *Orders) OrdersHold(Id string, optionalSetters ...OrdersHoldOption)(*models.Error, error) {
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
					
// OrdersItemsCancel take quantities off an order that is otherwise going
// ahead — three of the ten are discontinued, one line is out of stock and
// the customer would rather not wait. Each named quantity is booked onto its
// position as cancelled and guarded against the OPEN quantity (ordered −
// shipped − cancelled), so nothing already shipped can be cancelled away
// underneath a shipment. The order's fulfillment_status is re-derived
// afterwards, and when every position ends up fully cancelled the order
// itself moves to 'cancelled' — which is how this becomes a full cancel by
// arithmetic rather than by a second call. Positions are REQUIRED here,
// unlike on /ship and /return: cancelling an entire order by omitting a field
// is not something anybody should be able to do by accident; that is what
// POST /orders/{id}/cancel is for. Read GET /orders/{id}/shippable for the
// open quantity per position before calling.
func (srv *Orders) OrdersItemsCancel(Id string, Positions []models.OrderCancelPosition, optionalSetters ...OrdersItemsCancelOption)(*models.Error, error) {
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
					
// OrdersPaymentStatusUpdate payment is the one status dimension this app does
// not decide for itself: it is FED IN from whatever took the money — the
// payments app, a PSP webhook relayed by a workflow, or a finance clerk
// marking an invoice settled. This route writes that word onto the order and
// records the change as an order.payment_status.changed event carrying the
// previous value, so the trail shows the sequence and not just the current
// state. Optionally attach the payment_id of the transaction it came from. It
// takes no money, refunds none and validates nothing about the amount — it
// records a fact somebody else established, and any of the seven words may
// follow any other. The other half of auto_complete_on = 'payment': an order
// that has shipped in full is completed by this call when the status becomes
// 'paid'.
func (srv *Orders) OrdersPaymentStatusUpdate(Id string, Status string, optionalSetters ...OrdersPaymentStatusUpdateOption)(*models.Error, error) {
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
type OrdersReturnOptions struct {
	Metadata interface{}
	Positions []models.OrderReturnPosition
	Reason string
	Restock bool
	enabledSetters map[string]bool
}
func (options OrdersReturnOptions) New() *OrdersReturnOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"Positions": false,
		"Reason": false,
		"Restock": false,
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
func (srv *Orders) WithOrdersReturnPositions(v []models.OrderReturnPosition) OrdersReturnOption {
	return func(o *OrdersReturnOptions) {
		o.Positions = v
		o.enabledSetters["Positions"] = true
	}
}
func (srv *Orders) WithOrdersReturnReason(v string) OrdersReturnOption {
	return func(o *OrdersReturnOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *Orders) WithOrdersReturnRestock(v bool) OrdersReturnOption {
	return func(o *OrdersReturnOptions) {
		o.Restock = v
		o.enabledSetters["Restock"] = true
	}
}
			
// OrdersReturn open a return case: the customer has announced goods are
// coming back, and this is where that becomes a tracked thing with a return
// number of its own, drawn from the tenant's return range. Positions are
// guarded against what actually SHIPPED and has not already come back, so a
// return cannot exceed the goods that left. Each position carries a `restock`
// flag saying whether the item is expected to be sellable again — recorded
// now, acted on only when the return completes. Omitting `positions`
// registers everything still returnable, the 'the customer sent the whole
// delivery back' case. Nothing is booked yet: quantity_returned stays where
// it is and the order does not move — the return starts as 'registered' and
// travels through receive and complete or reject. Allowed on a completed
// order, refused on a cancelled one.
func (srv *Orders) OrdersReturn(Id string, optionalSetters ...OrdersReturnOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/return")
	options := OrdersReturnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Positions"] {
		params["positions"] = options.Positions
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Restock"] {
		params["restock"] = options.Restock
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
					
// OrdersReturnsComplete accept the return and close the case: the goods are
// taken back on the order's books and the settlement is recorded as one of
// the published words — refunded, credited, replaced and so on. This is the
// step a refund or a credit note hangs off, and the only step that moves
// quantity_returned. It does not refund money and does not put stock back
// itself: the answer's 'restock' array names what the orchestrator should
// hand to inventories.restock, and payment travels through /payment-status.
// Once completed the return is final — receive, complete and reject all
// refuse afterwards. The goods accounting moves here and nowhere else:
// quantity_returned is booked onto each position, completed_at is stamped by
// the SERVER, and positions flagged restock are reported back in the answer's
// 'restock' array for the orchestrator's inventories.restock call.
// 'resolution' is validated against the settlement words this app publishes
// (refund, partial_refund, replacement, repair, store_credit — see GET
// /orders/vocabularies/return-resolutions); anything else is refused rather
// than stored as a word no reader knows. It is checked before the positions
// are booked, so a rejected value leaves nothing behind.
func (srv *Orders) OrdersReturnsComplete(Id string, Rid string, optionalSetters ...OrdersReturnsCompleteOption)(*models.Error, error) {
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
					
// OrdersReturnsReceive the goods-in scan: the parcel is physically back,
// warehouse staff have it in their hands, and nobody has decided yet whether
// the customer gets their money. It moves the return from 'registered' to
// 'received' and stamps received_at, which is what separates 'announced' from
// 'here' on a returns worklist. It books nothing — quantity_returned is
// written by the complete step and by nothing else — so a return that
// arrives damaged can still be rejected afterwards. Only a registered return
// can be received; a second call, or one against a settled return, is a 422.
// This step is skippable: a return may be completed straight from
// 'registered' where a merchant does not scan goods in.
func (srv *Orders) OrdersReturnsReceive(Id string, Rid string, Data interface{})(*models.Error, error) {
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
					
// OrdersReturnsReject close the case against the customer: the goods came
// back used, outside the window, or were never covered in the first place.
// The return moves to 'rejected', rejected_at is stamped, and the refusal is
// recorded either as one of the published refusal words or as a sentence
// somebody wrote about this one return. The order is untouched — the
// quantities still count as shipped and not returned, which is the point: a
// rejected return must leave the books exactly as they were. Rejection is
// final, and it says nothing about where the physical goods go. Nothing is
// booked onto the positions. 'resolution' is validated against the refusal
// words (wear_and_tear, not_returnable); 'reason' stays free text — a
// sentence about this one return rather than a value out of a set — and is
// what is stored when no resolution is named.
func (srv *Orders) OrdersReturnsReject(Id string, Rid string, optionalSetters ...OrdersReturnsRejectOption)(*models.Error, error) {
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
			
// OrdersShip book goods out: which positions and how much of each, with the
// carrier and the tracking code that go to the customer. It draws a
// delivery-note number from the tenant's delivery range, books
// quantity_shipped onto every named position, re-derives the order's
// fulfillment_status from the arithmetic (unfulfilled → partial →
// fulfilled) and emits order.shipment.created. Omitting `positions` means
// everything still open, in full, which is the ordinary 'send the rest' case
// and the only one a UI without a line editor can express; the answer always
// names the quantities that actually went out. It does not print a label, buy
// postage or notify anybody — a shipping workflow reacts to the event.
// Whether a full shipment CLOSES the order is the tenant's call (setting
// auto_complete_on): 'shipment' completes it here, 'payment' leaves it
// in_fulfillment until payment_status becomes paid, 'manual' waits for
// orders.complete. The order.completed event follows the order, so it is only
// emitted when the order actually completed.
func (srv *Orders) OrdersShip(Id string, optionalSetters ...OrdersShipOption)(*models.Error, error) {
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
	
// OrdersShippable what a shipment dialog needs before it can offer anything:
// the open quantity per position, and one boolean saying whether a shipment
// would be accepted at all. Reach for it to fill a picking screen or to
// decide whether a 'create shipment' button is enabled, instead of
// subtracting the quantities client-side. It changes nothing and books
// nothing — it is the question POST /orders/{id}/ship answers with an
// action. The read half of orders.ship. The open quantity per position and
// the two guards (cancelled/completed order, hold) are the SAME code the ship
// route runs, so what this answers and what that accepts cannot drift — a
// client subtracting the quantities itself eventually offers a shipment the
// server refuses, or one it should have refused. 'shippable' is false with a
// 'blocked_reason' when the order is held, cancelled, completed or has
// nothing open.
func (srv *Orders) OrdersShippable(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orders/{id}/shippable")
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
			
// OrdersUnhold the whole of the release: the flag comes off, the reason is
// cleared, and an order.unheld event says the order may move again. Whatever
// the hold was blocking — shipping, and cancellation on tenants configured
// that way — is accepted from this call on. It restores nothing else and
// skips nothing: the order continues from exactly the status and quantities
// it had when it was held, and any shipping that was due meanwhile still has
// to be done by hand. An order that is not on hold answers 422 rather than
// pretending to release one, so this is safe to give to a worklist and not to
// a loop that calls it blindly.
func (srv *Orders) OrdersUnhold(Id string, Data interface{})(*models.Error, error) {
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
