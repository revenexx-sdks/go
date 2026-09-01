package inventories_reservations

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// InventoriesReservations service
type InventoriesReservations struct {
	client client.Client
}

func New(clt client.Client) *InventoriesReservations {
	return &InventoriesReservations{
		client: clt,
	}
}

	
// InventoriesCommit call this when the goods leave the building, and not
// before. Reserving only promised them — `reserved` went up and `on_hand`
// did not move, because the stock was still on the shelf; committing is the
// moment they are gone, so it lowers BOTH on each stock row and writes one
// `shipment` booking per hold, with a SIGNED negative quantity, as the
// ledger's record that they left. It takes the whole `order_ref` and every
// hold still active on it: there is no partial commit and no per-line id, so
// a part shipment means reserving the parts separately in the first place. It
// is also final — 'committed' ends the lifecycle and nothing moves a hold
// out of it, so goods coming back are POST /inventories/restock (a new
// receipt), never an undo of this. An order with nothing active is a 422
// rather than a quiet zero, because it means the hold was already released or
// already shipped; /release answers the same situation with a 200 on purpose,
// since cancelling twice is harmless and shipping twice is not.
func (srv *InventoriesReservations) InventoriesCommit(OrderRef string)(*models.Error, error) {
	path := "/v1/inventories/commit"
	params := map[string]interface{}{}
	params["order_ref"] = OrderRef
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
	
// InventoriesRelease the cancellation end of the reserve → commit | release
// lifecycle: it takes an `order_ref`, ends every hold still active on it,
// gives the stock back and writes a 'release' booking for each one, exactly
// like the expiry sweeper. Idempotent: an order with nothing active answers
// released:0 — which is why it is a 200 and not the 422 commit answers.
func (srv *InventoriesReservations) InventoriesRelease(OrderRef string)(*models.Error, error) {
	path := "/v1/inventories/release"
	params := map[string]interface{}{}
	params["order_ref"] = OrderRef
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
type InventoriesReservationsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	LocationId string
	ProductId string
	Sku string
	Quantity float64
	OrderRef string
	Status string
	ExpiresAt string
	Metadata string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options InventoriesReservationsListOptions) New() *InventoriesReservationsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"LocationId": false,
		"ProductId": false,
		"Sku": false,
		"Quantity": false,
		"OrderRef": false,
		"Status": false,
		"ExpiresAt": false,
		"Metadata": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type InventoriesReservationsListOption func(*InventoriesReservationsListOptions)
func (srv *InventoriesReservations) WithInventoriesReservationsListLimit(v int) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListOffset(v int) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListOrder(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListId(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListLocationId(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.LocationId = v
		o.enabledSetters["LocationId"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListProductId(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListSku(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListQuantity(v float64) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListOrderRef(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListStatus(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListExpiresAt(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListMetadata(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListCreatedAt(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReservationsListUpdatedAt(v string) InventoriesReservationsListOption {
	return func(o *InventoriesReservationsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// InventoriesReservationsList a reservation is stock promised to an
// `order_ref`. It is created only by POST /inventories/reserve and moved only
// by /commit, /release and the expiry sweep — there is no create, update or
// delete route, because the lifecycle IS the API. Only an 'active' hold
// counts towards a stock row's `reserved`; 'released' and 'committed' rows
// stay for the audit trail and hold nothing. This is the answer to "what is
// this order actually holding" (`?order_ref=…`) and to "what is holding
// this stock" (`?status=active&location_id=…`) — the second is the only
// way to see WHY a row's `reserved` is what it is, since a stock row reports
// the total and never who asked for it. `expires_at` filters on an exact
// timestamp and not a range, so this cannot answer "what expires today"; the
// deadline is acted on by POST /inventories/reservations/sweep, not by
// reading it here.
func (srv *InventoriesReservations) InventoriesReservationsList(optionalSetters ...InventoriesReservationsListOption)(*models.Error, error) {
	path := "/v1/inventories/reservations"
	options := InventoriesReservationsListOptions{}.New()
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
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["ExpiresAt"] {
		params["expires_at"] = options.ExpiresAt
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
	
// InventoriesReservationsSweep the expiry sweeper, also run by the
// 'expire-reservations' schedule every 15 minutes. Releases reservations past
// their own expires_at and — once reservation_ttl_minutes is above 0 —
// reservations older than that lifetime which never carried a deadline. Each
// release gives the stock back and writes a 'release' booking, exactly like a
// cancellation. Idempotent: a second run finds nothing.
func (srv *InventoriesReservations) InventoriesReservationsSweep(Data interface{})(*models.ReservationSweepResult, error) {
	path := "/v1/inventories/reservations/sweep"
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

		parsed := models.ReservationSweepResult{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReservationSweepResult
	parsed, ok := resp.Result.(models.ReservationSweepResult)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// InventoriesReservationsGet a reservation is stock promised to an
// `order_ref`. It is created only by POST /inventories/reserve and moved only
// by /commit, /release and the expiry sweep — there is no create, update or
// delete route, because the lifecycle IS the API. Only an 'active' hold
// counts towards a stock row's `reserved`; 'released' and 'committed' rows
// stay for the audit trail and hold nothing. One hold, with the three facts
// that are not on the order it belongs to: which location it was allocated
// to, when it expires, and — in `metadata.backordered` — how much of it
// was never covered by stock, which is how a promise made under a permissive
// backorder policy stays visible afterwards. The id is for reading only.
// Every transition acts on the whole `order_ref` (/commit, /release, the
// sweep), so there is no route that takes this id and no way to release one
// line of an order on its own.
func (srv *InventoriesReservations) InventoriesReservationsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/reservations/{id}")
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
type InventoriesReserveOptions struct {
	ExpiresAt string
	Items []models.InventoryStockItem
	LocationCode string
	ProductId string
	Quantity float64
	ShipTo interface{}
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesReserveOptions) New() *InventoriesReserveOptions {
	options.enabledSetters = map[string]bool{
		"ExpiresAt": false,
		"Items": false,
		"LocationCode": false,
		"ProductId": false,
		"Quantity": false,
		"ShipTo": false,
		"Sku": false,
	}
	return &options
}
type InventoriesReserveOption func(*InventoriesReserveOptions)
func (srv *InventoriesReservations) WithInventoriesReserveExpiresAt(v string) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveItems(v []models.InventoryStockItem) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveLocationCode(v string) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveProductId(v string) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveQuantity(v float64) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveShipTo(v interface{}) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.ShipTo = v
		o.enabledSetters["ShipTo"] = true
	}
}
func (srv *InventoriesReservations) WithInventoriesReserveSku(v string) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
			
// InventoriesReserve takes a hold against an `order_ref`, and plans the whole
// call before writing anything, so a reservation that cannot be satisfied
// changes nothing. WHICH location serves an item is not the caller's to
// choose: the tenant's allocation_strategy decides it ('priority', walking
// the enabled locations by their priority; 'nearest', matching ship_to
// against a location's country; or 'single_location' for the whole order);
// backorder_policy decides what happens when none can — refuse (422), or
// reserve anyway and let availability go negative. expires_at defaults from
// reservation_ttl_minutes and the sweeper enforces it.
func (srv *InventoriesReservations) InventoriesReserve(OrderRef string, optionalSetters ...InventoriesReserveOption)(*models.Error, error) {
	path := "/v1/inventories/reserve"
	options := InventoriesReserveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["order_ref"] = OrderRef
	if options.enabledSetters["ExpiresAt"] {
		params["expires_at"] = options.ExpiresAt
	}
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
	if options.enabledSetters["ShipTo"] {
		params["ship_to"] = options.ShipTo
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
