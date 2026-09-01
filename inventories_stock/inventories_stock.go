package inventories_stock

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// InventoriesStock service
type InventoriesStock struct {
	client client.Client
}

func New(clt client.Client) *InventoriesStock {
	return &InventoriesStock{
		client: clt,
	}
}

type InventoriesAdjustOptions struct {
	Items []models.InventoryAdjustItem
	LocationCode string
	ProductId string
	Quantity float64
	Reason string
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesAdjustOptions) New() *InventoriesAdjustOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"LocationCode": false,
		"ProductId": false,
		"Quantity": false,
		"Reason": false,
		"Sku": false,
	}
	return &options
}
type InventoriesAdjustOption func(*InventoriesAdjustOptions)
func (srv *InventoriesStock) WithInventoriesAdjustItems(v []models.InventoryAdjustItem) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAdjustLocationCode(v string) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAdjustProductId(v string) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAdjustQuantity(v float64) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAdjustReason(v string) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAdjustSku(v string) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
	
// InventoriesAdjust the batch correction route — a stocktake, breakage,
// shrinkage — and the manual way `on_hand` is ever put right. Quantities
// are SIGNED: a positive one adds to the balance, a negative one takes it
// away, and neither is written onto the row directly. Each item is booked
// into the movements ledger as an `adjustment` and the balance follows, so a
// correction leaves a record of who changed what and why instead of a number
// that silently differs from yesterday's. A reason is mandatory unless
// movement_reason_required is 'none'.
func (srv *InventoriesStock) InventoriesAdjust(optionalSetters ...InventoriesAdjustOption)(*models.Error, error) {
	path := "/v1/inventories/adjust"
	options := InventoriesAdjustOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
type InventoriesAvailabilityOptions struct {
	Items []models.InventoryAvailabilityItem
	LocationCode string
	ProductId string
	Quantity float64
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesAvailabilityOptions) New() *InventoriesAvailabilityOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"LocationCode": false,
		"ProductId": false,
		"Quantity": false,
		"Sku": false,
	}
	return &options
}
type InventoriesAvailabilityOption func(*InventoriesAvailabilityOptions)
func (srv *InventoriesStock) WithInventoriesAvailabilityItems(v []models.InventoryAvailabilityItem) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAvailabilityLocationCode(v string) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAvailabilityProductId(v string) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAvailabilityQuantity(v float64) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesAvailabilitySku(v string) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
	
// InventoriesAvailability tHE stock call of this app, and a batch one: name
// any number of items and each comes back with `on_hand`, `reserved` and the
// derived `available` (their difference, computed on read and stored
// nowhere), summed across the locations in scope and broken down per
// location, plus `orderable` — whether this much of it can be promised at
// this moment. An item this app has never seen is NOT an error: it comes back
// tracked:false, and the storefront decides whether an untracked item sells
// freely. It is also the most customised surface this product has in the
// field. A tenant whose stock really lives in an ERP — SAP live stock is
// the ordinary case, not the exotic one — replaces exactly this one
// capability, 1:1, with a custom app through the gateway's capability
// override, while every other route here keeps doing the stock-keeping CRUD
// unchanged. That is why the request and response shapes below read as a
// contract to be implemented rather than as an implementation detail:
// whatever ends up answering this path has to answer in these terms.
func (srv *InventoriesStock) InventoriesAvailability(optionalSetters ...InventoriesAvailabilityOption)(*models.Error, error) {
	path := "/v1/inventories/availability"
	options := InventoriesAvailabilityOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
type InventoriesMovementsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	LocationId string
	ProductId string
	Sku string
	Type string
	Quantity float64
	OrderRef string
	Reason string
	Metadata string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options InventoriesMovementsListOptions) New() *InventoriesMovementsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"LocationId": false,
		"ProductId": false,
		"Sku": false,
		"Type": false,
		"Quantity": false,
		"OrderRef": false,
		"Reason": false,
		"Metadata": false,
		"CreatedAt": false,
	}
	return &options
}
type InventoriesMovementsListOption func(*InventoriesMovementsListOptions)
func (srv *InventoriesStock) WithInventoriesMovementsListLimit(v int) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListOffset(v int) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListOrder(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListId(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListLocationId(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.LocationId = v
		o.enabledSetters["LocationId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListProductId(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListSku(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListType(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListQuantity(v float64) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListOrderRef(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListReason(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListMetadata(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesMovementsListCreatedAt(v string) InventoriesMovementsListOption {
	return func(o *InventoriesMovementsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// InventoriesMovementsList the movements ledger, read end to end. Every stock
// change this app has ever made is a booking row in it — a receipt, a
// correction, a hold, a release, a shipment, a return — which is what lets
// one list be an audit trail and an event feed at the same time: these are
// the rows the `stock_movement.created` event carries, so a consumer that
// missed an event catches up by paging here. Append-only: the ledger has no
// update and no delete, because a correction is another booking.
// `order=created_at.desc` is the feed order.
func (srv *InventoriesStock) InventoriesMovementsList(optionalSetters ...InventoriesMovementsListOption)(*models.Error, error) {
	path := "/v1/inventories/movements"
	options := InventoriesMovementsListOptions{}.New()
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
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["LocationId"] {
		params["location_id"] = options.LocationId
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
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
	
// InventoriesMovementsGet a movement is one booking row in the ledger, and
// the ledger is append-only: there is no update and no delete, because a
// correction is another booking. `quantity` is SIGNED and its sign follows
// the `type` — a receipt books +5 and the reserve that promises those goods
// books −5, even though the reservation it created carries +5 as a positive
// hold. GET /inventories/vocabularies/movement-types is the list of types
// with the words for them. A booking says what changed, not what the balance
// became: it carries no running total, so the row's story is read by listing
// the ledger for that location and item rather than by fetching one id.
// `location_id` is a plain uuid and not a foreign key, so a booking outlives
// the location it was made at and this route will happily hand back one whose
// location no longer resolves — that is the audit trail doing its job, not
// a broken row. Fixing a wrong booking is another booking (POST
// /inventories/adjust); nothing here can be edited or removed.
func (srv *InventoriesStock) InventoriesMovementsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/movements/{id}")
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
type InventoriesReceiveOptions struct {
	Items []models.InventoryStockItem
	LocationCode string
	ProductId string
	Quantity float64
	Reason string
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesReceiveOptions) New() *InventoriesReceiveOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"LocationCode": false,
		"ProductId": false,
		"Quantity": false,
		"Reason": false,
		"Sku": false,
	}
	return &options
}
type InventoriesReceiveOption func(*InventoriesReceiveOptions)
func (srv *InventoriesStock) WithInventoriesReceiveItems(v []models.InventoryStockItem) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesReceiveLocationCode(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesReceiveProductId(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesReceiveQuantity(v float64) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesReceiveReason(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesReceiveSku(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
	
// InventoriesReceive books a delivery into the receiving location (the
// caller's location_code, else the default_location_code setting), creating
// the stock row if the item is new. A reason is optional unless
// movement_reason_required is 'all'. Takes a batch or one item inline.
func (srv *InventoriesStock) InventoriesReceive(optionalSetters ...InventoriesReceiveOption)(*models.Error, error) {
	path := "/v1/inventories/receive"
	options := InventoriesReceiveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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

// InventoriesReorderAlerts the replenishment worklist: the stock rows that
// have run down far enough that somebody has to order more, in one list
// rather than as a query a caller has to build. Computed on read, so it is
// never stale: a row alerts when available (on_hand − reserved) has fallen
// to or below its own reorder_point, or the reorder_point_default setting
// when it carries none. A point of 0 never alerts. Answers enabled:false with
// an empty list when reorder_alert_enabled is off — a tenant replenishing
// from an ERP should not be told twice.
func (srv *InventoriesStock) InventoriesReorderAlerts()(*models.ReorderAlerts, error) {
	path := "/v1/inventories/reorder-alerts"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ReorderAlerts{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReorderAlerts
	parsed, ok := resp.Result.(models.ReorderAlerts)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// InventoriesReorderScan publishes `stock_level.low` on the event bus for
// every row GET /inventories/reorder-alerts currently lists, so replenishment
// can be driven by a subscriber instead of by somebody refreshing that page.
// Also runs hourly as the `reorder-scan` schedule; this route is for driving
// it on demand. The event id is derived from the stock row and the day, so a
// re-run — a second click, a retried cron tick — publishes nothing new
// and returns the ids the first run produced. Nothing is written to the app's
// own data: this reads the same figures the alerts list computes and hands
// them to the bus. Answers enabled:false without publishing when
// reorder_alert_enabled is off.
func (srv *InventoriesStock) InventoriesReorderScan(Data interface{})(*models.ReorderScan, error) {
	path := "/v1/inventories/reorder-alerts/scan"
	params := map[string]interface{}{}
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

		parsed := models.ReorderScan{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReorderScan
	parsed, ok := resp.Result.(models.ReorderScan)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type InventoriesRestockOptions struct {
	Items []models.InventoryStockItem
	LocationCode string
	OrderRef string
	ProductId string
	Quantity float64
	Reason string
	Restock bool
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesRestockOptions) New() *InventoriesRestockOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"LocationCode": false,
		"OrderRef": false,
		"ProductId": false,
		"Quantity": false,
		"Reason": false,
		"Restock": false,
		"Sku": false,
	}
	return &options
}
type InventoriesRestockOption func(*InventoriesRestockOptions)
func (srv *InventoriesStock) WithInventoriesRestockItems(v []models.InventoryStockItem) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockLocationCode(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockOrderRef(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockProductId(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockQuantity(v float64) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockReason(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockRestock(v bool) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Restock = v
		o.enabledSetters["Restock"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesRestockSku(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
	
// InventoriesRestock whether a return rejoins sellable stock follows
// restock_on_return_default, overridable per call with 'restock'. When the
// answer is no the response says restocked:false and nothing moves — there
// is no movement to book, because no stock moved. That branch is why this
// route answers 200 and its sibling `receive` answers 201: a restock may
// legitimately create nothing.
func (srv *InventoriesStock) InventoriesRestock(optionalSetters ...InventoriesRestockOption)(*models.Error, error) {
	path := "/v1/inventories/restock"
	options := InventoriesRestockOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Restock"] {
		params["restock"] = options.Restock
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
type InventoriesStockListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	LocationId string
	ProductId string
	Sku string
	OnHand float64
	Reserved float64
	ReorderPoint float64
	Metadata string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options InventoriesStockListOptions) New() *InventoriesStockListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"LocationId": false,
		"ProductId": false,
		"Sku": false,
		"OnHand": false,
		"Reserved": false,
		"ReorderPoint": false,
		"Metadata": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type InventoriesStockListOption func(*InventoriesStockListOptions)
func (srv *InventoriesStock) WithInventoriesStockListLimit(v int) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListOffset(v int) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListOrder(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListId(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListLocationId(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.LocationId = v
		o.enabledSetters["LocationId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListProductId(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListSku(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListOnHand(v float64) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.OnHand = v
		o.enabledSetters["OnHand"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListReserved(v float64) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Reserved = v
		o.enabledSetters["Reserved"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListReorderPoint(v float64) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.ReorderPoint = v
		o.enabledSetters["ReorderPoint"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListMetadata(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListCreatedAt(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockListUpdatedAt(v string) InventoriesStockListOption {
	return func(o *InventoriesStockListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// InventoriesStockList a stock level is ONE item at ONE location, and it
// carries two numbers, neither of which is the sellable one: `on_hand` is
// what is physically there INCLUDING everything already promised, and
// `reserved` is what has been promised — it never reduces `on_hand`. What
// may still be sold is their difference, and it is derived on read and never
// stored, so there is no `available` column to read, filter or order by. This
// is the operator's view — the whole book, filtered by location or by item
// — not the shop's: a storefront asking "can I sell five of this" wants
// POST /inventories/availability, which sums an item across locations and
// answers `orderable` instead of leaving the caller to subtract. Two things
// this list will not do: it has no range filters, so "everything running low"
// is GET /inventories/reorder-alerts and not a query here; and it does not
// promise one row per item per location — no unique index enforces that.
// POST /inventories/stock refuses a duplicate with a 409, but that is a check
// and not a constraint, so a row written past it, or one that predates the
// guard, still splits an item's balance in two, and the write routes find and
// update whichever of them the database returns first.
func (srv *InventoriesStock) InventoriesStockList(optionalSetters ...InventoriesStockListOption)(*models.Error, error) {
	path := "/v1/inventories/stock"
	options := InventoriesStockListOptions{}.New()
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
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["LocationId"] {
		params["location_id"] = options.LocationId
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["OnHand"] {
		params["on_hand"] = options.OnHand
	}
	if options.enabledSetters["Reserved"] {
		params["reserved"] = options.Reserved
	}
	if options.enabledSetters["ReorderPoint"] {
		params["reorder_point"] = options.ReorderPoint
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
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
type InventoriesStockCreateOptions struct {
	Metadata interface{}
	ProductId string
	ReorderPoint float64
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesStockCreateOptions) New() *InventoriesStockCreateOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"ProductId": false,
		"ReorderPoint": false,
		"Sku": false,
	}
	return &options
}
type InventoriesStockCreateOption func(*InventoriesStockCreateOptions)
func (srv *InventoriesStock) WithInventoriesStockCreateMetadata(v interface{}) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockCreateProductId(v string) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockCreateReorderPoint(v float64) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.ReorderPoint = v
		o.enabledSetters["ReorderPoint"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockCreateSku(v string) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
			
// InventoriesStockCreate registers an item at a location. The row is born at
// ZERO and never gets a balance from this call: `on_hand` and `reserved` are
// NOT accepted, because they are the running total of the movements ledger,
// so an opening balance is a receipt (POST /inventories/receive) rather than
// a field here, and the only thing that ever moves either number afterwards
// is another booking. What this row carries is its identity (location +
// `product_id`/`sku`), its `reorder_point` and its metadata. `location_id` is
// the only field a create cannot omit; every other column is optional or
// defaulted by the database. The one rule that is a CHECK rather than a
// column is that a row has to identify its item, so `product_id` or `sku` has
// to be there as well. Mostly you do not need this route at all — every
// stock call creates the row it is missing — and a second row for an item
// this location already tracks is answered 409: no unique index enforces one
// row per item per location, so that row would split the item's balance
// across two rows the write routes cannot tell apart, each of them updating
// whichever the database returns first. That guard is a check before the
// insert and not a constraint, so it closes a double click or a re-run import
// and does not claim to close a race between two simultaneous creates.
func (srv *InventoriesStock) InventoriesStockCreate(LocationId string, optionalSetters ...InventoriesStockCreateOption)(*models.Error, error) {
	path := "/v1/inventories/stock"
	options := InventoriesStockCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["location_id"] = LocationId
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["ReorderPoint"] {
		params["reorder_point"] = options.ReorderPoint
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
	
// InventoriesStockDelete stops tracking one item at one location. A stock
// level is ONE item at ONE location, and it carries two numbers, neither of
// which is the sellable one: `on_hand` is what is physically there INCLUDING
// everything already promised, and `reserved` is what has been promised —
// it never reduces `on_hand`. What may still be sold is their difference, and
// it is derived on read and never stored, so there is no `available` column
// to read, filter or order by. A deleted balance is not recoverable: the
// ledger is the audit trail, not the source of truth, and nothing in this app
// ever replays it to rebuild a number — so the next receipt for the same
// item here creates a FRESH row at zero, standing next to movements that say
// otherwise. That used to be a trap a caller discovered afterwards. It is a
// stated property now, because the route REFUSES while the row still holds
// anything, and answers 409 with what it holds. The two things that block are
// the location delete's two, asked of one row. A reservation still `active`
// against this item at this location is the sharper one: /release and /commit
// look their stock row up by (location, item) on the very next call and would
// find nothing, so the hold would lower no `reserved` and /commit would book
// the whole quantity as a shortfall — orphaned immediately rather than
// eventually. `on_hand` above zero is the stronger one: deleting a LOCATION
// at least meant "close this warehouse" and took the balances as a side
// effect of the cascade, while this row IS the balance, so the delete can
// only ever mean "no longer tracked here" — true once the number is zero
// and a lie while it is not. POST /inventories/stock/{id}/adjust to zero is
// the operation that makes it true, and it BOOKS the movement, so the stock
// leaves through the ledger instead of vanishing with the row. Nothing points
// at it by foreign key, so the database takes nothing else with it. History
// therefore never blocks and is never deleted — the ledger is keyed on
// (location, item) and never on this id, so its bookings survive a row that
// is gone, BY DESIGN, exactly as they survive a location that is gone.
func (srv *InventoriesStock) InventoriesStockDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/stock/{id}")
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
	
// InventoriesStockGet a stock level is ONE item at ONE location, and it
// carries two numbers, neither of which is the sellable one: `on_hand` is
// what is physically there INCLUDING everything already promised, and
// `reserved` is what has been promised — it never reduces `on_hand`. What
// may still be sold is their difference, and it is derived on read and never
// stored, so there is no `available` column to read, filter or order by. Read
// it to see one item's position at one place, and to get the id the two
// row-scoped routes take: POST /inventories/stock/{id}/adjust corrects this
// balance, and GET /inventories/reorder-alerts reports it by this id. What it
// does not answer is how the balance got here — that is GET
// /inventories/movements filtered by the location and item on this row,
// because a movement points at (location, item) and never at a stock row id.
func (srv *InventoriesStock) InventoriesStockGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/stock/{id}")
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
type InventoriesStockUpdateOptions struct {
	LocationId string
	Metadata interface{}
	ProductId string
	ReorderPoint float64
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesStockUpdateOptions) New() *InventoriesStockUpdateOptions {
	options.enabledSetters = map[string]bool{
		"LocationId": false,
		"Metadata": false,
		"ProductId": false,
		"ReorderPoint": false,
		"Sku": false,
	}
	return &options
}
type InventoriesStockUpdateOption func(*InventoriesStockUpdateOptions)
func (srv *InventoriesStock) WithInventoriesStockUpdateLocationId(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.LocationId = v
		o.enabledSetters["LocationId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockUpdateMetadata(v interface{}) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockUpdateProductId(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockUpdateReorderPoint(v float64) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.ReorderPoint = v
		o.enabledSetters["ReorderPoint"] = true
	}
}
func (srv *InventoriesStock) WithInventoriesStockUpdateSku(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
			
// InventoriesStockUpdate partial update of everything on the row EXCEPT its
// balance: reorder_point, metadata, identity. on_hand and reserved are
// dropped from the body — every stock change is a movement, and a body
// carrying nothing else is answered 422 with the route that was meant (POST
// /inventories/stock/{id}/adjust).
func (srv *InventoriesStock) InventoriesStockUpdate(Id string, optionalSetters ...InventoriesStockUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/stock/{id}")
	options := InventoriesStockUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["LocationId"] {
		params["location_id"] = options.LocationId
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["ReorderPoint"] {
		params["reorder_point"] = options.ReorderPoint
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
type InventoriesStockAdjustOptions struct {
	Reason string
	enabledSetters map[string]bool
}
func (options InventoriesStockAdjustOptions) New() *InventoriesStockAdjustOptions {
	options.enabledSetters = map[string]bool{
		"Reason": false,
	}
	return &options
}
type InventoriesStockAdjustOption func(*InventoriesStockAdjustOptions)
func (srv *InventoriesStock) WithInventoriesStockAdjustReason(v string) InventoriesStockAdjustOption {
	return func(o *InventoriesStockAdjustOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
					
// InventoriesStockAdjust corrects the balance of ONE stock row, and only that
// one. It is the row-scoped twin of POST /inventories/adjust: the row already
// knows its location and item, so a caller owes nothing but a SIGNED delta on
// `on_hand` — positive to add, negative to take away — and a reason for
// it. The delta is not written onto the balance either; it is booked into the
// movements ledger as an `adjustment` and the balance follows, which is why
// the answer hands back the row at its new value instead of an
// acknowledgement. This is the route that replaced the Cockpit's editable
// on_hand field.
func (srv *InventoriesStock) InventoriesStockAdjust(Id string, Quantity float64, optionalSetters ...InventoriesStockAdjustOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/stock/{id}/adjust")
	options := InventoriesStockAdjustOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["quantity"] = Quantity
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

// InventoriesVocabulariesList discovery for the vocabulary routes: the enums
// this app publishes, each with its name, its title and its description and
// deliberately WITHOUT its values, so finding out what exists costs one small
// call and not one per vocabulary. Names: location-types, movement-types,
// reservation-statuses. Fetch one with GET /inventories/vocabularies/{name};
// a client holding the qualified pair 'inventories.<name>' builds that URL
// from the pair alone.
func (srv *InventoriesStock) InventoriesVocabulariesList()(*models.InventoryVocabularyIndex, error) {
	path := "/v1/inventories/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.InventoryVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.InventoryVocabularyIndex
	parsed, ok := resp.Result.(models.InventoryVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// InventoriesVocabulariesGet one vocabulary in full: every permitted value,
// each carrying the title and description a person reads for it and the badge
// tone a UI colours it with, so a client renders a status or a movement type
// without a hard-coded table of its own. The values are read out of the
// column's CHECK constraint, so the served set IS the enforced set and the
// two cannot drift — a value added to the constraint appears here even
// before anyone labels it, titled from its own key. Values come back in
// constraint order, which is lifecycle order for a status. 'closed' says the
// set is exhaustive, so a value outside it is stale data rather than a
// missing label. Names: location-types, movement-types, reservation-statuses.
func (srv *InventoriesStock) InventoriesVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/inventories/vocabularies/{name}")
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
