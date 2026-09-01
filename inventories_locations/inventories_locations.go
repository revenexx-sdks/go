package inventories_locations

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// InventoriesLocations service
type InventoriesLocations struct {
	client client.Client
}

func New(clt client.Client) *InventoriesLocations {
	return &InventoriesLocations{
		client: clt,
	}
}

type InventoriesLocationsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	Name string
	Labels string
	Type string
	Priority int
	Enabled bool
	Address string
	Metadata string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options InventoriesLocationsListOptions) New() *InventoriesLocationsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"Name": false,
		"Labels": false,
		"Type": false,
		"Priority": false,
		"Enabled": false,
		"Address": false,
		"Metadata": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type InventoriesLocationsListOption func(*InventoriesLocationsListOptions)
func (srv *InventoriesLocations) WithInventoriesLocationsListLimit(v int) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListOffset(v int) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListOrder(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListId(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListCode(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListName(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListLabels(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListType(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListPriority(v int) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListEnabled(v bool) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListAddress(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListMetadata(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListCreatedAt(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsListUpdatedAt(v string) InventoriesLocationsListOption {
	return func(o *InventoriesLocationsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// InventoriesLocationsList a location is WHERE stock is kept — a warehouse,
// a shop floor, a supplier that dropships, or a virtual bucket for pre-orders
// and quarantine. It holds no quantity of its own: what is at it is a stock
// level. `type` is descriptive and nothing branches on it; `priority` is the
// number that decides which location a reservation is served from, and
// `enabled` decides whether it is offered at all. This is the list a
// `location_code` is resolved against on every stock call, so it is the first
// thing to read when a receipt answers "unknown location". It answers no
// quantities at all — how much is at a location is GET
// /inventories/stock?location_id=…, and what may still be sold is POST
// /inventories/availability. Filter `?enabled=true` for the operational
// subset: availability and reserve only ever look at enabled locations, so a
// disabled one is invisible to a shop while keeping every row that points at
// it.
func (srv *InventoriesLocations) InventoriesLocationsList(optionalSetters ...InventoriesLocationsListOption)(*models.Error, error) {
	path := "/v1/inventories/locations"
	options := InventoriesLocationsListOptions{}.New()
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
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Address"] {
		params["address"] = options.Address
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
type InventoriesLocationsCreateOptions struct {
	Address interface{}
	Enabled bool
	Labels interface{}
	Metadata interface{}
	Priority int
	Type string
	enabledSetters map[string]bool
}
func (options InventoriesLocationsCreateOptions) New() *InventoriesLocationsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Address": false,
		"Enabled": false,
		"Labels": false,
		"Metadata": false,
		"Priority": false,
		"Type": false,
	}
	return &options
}
type InventoriesLocationsCreateOption func(*InventoriesLocationsCreateOptions)
func (srv *InventoriesLocations) WithInventoriesLocationsCreateAddress(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsCreateEnabled(v bool) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsCreateLabels(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsCreateMetadata(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsCreatePriority(v int) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsCreateType(v string) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// InventoriesLocationsCreate registers a new place stock can be kept, and
// `type` says what kind of place it is: a warehouse of your own, a store
// whose shop floor a click-and-collect order draws on, a dropship supplier
// whose stock this row only tracks, or a virtual bucket that is not a
// building at all — pre-orders, consignment, a quarantine shelf. A create
// cannot omit `code` and `name`; every other column is optional or defaulted
// by the database. Two rows of this tenant may not share `code` — that is
// the 409, and it answers an update that moves a row onto a sibling's value
// exactly as it answers a second insert. A new location starts EMPTY and
// creating one moves nothing: stock arrives through POST
// /inventories/receive, or is transferred by two adjustments, one negative at
// the old location and one positive here. Mind the two columns that are not
// decoration — `priority` decides where a reservation is served from before
// `type` ever does (nothing branches on `type`), and `enabled` defaults to
// true, so a location created for a warehouse that has not opened yet starts
// being offered by availability and reserve immediately.
func (srv *InventoriesLocations) InventoriesLocationsCreate(Code string, Name string, optionalSetters ...InventoriesLocationsCreateOption)(*models.Error, error) {
	path := "/v1/inventories/locations"
	options := InventoriesLocationsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Address"] {
		params["address"] = options.Address
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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

// InventoriesLocationsDefaults gives a tenant its first location, `main`, so
// the stock calls have somewhere to book into: `receive`, `adjust` and
// `restock` fall back to the `default_location_code` setting when a caller
// names no `location_code`, and a tenant with no location at all answers 400
// on its first receipt. The platform already runs this on `app.installed`, so
// calling it by hand is the repair for an install that predates the event or
// a `main` somebody deleted. Idempotent by CODE, not by contents: a location
// already carrying that code is reported under `existing` and is NOT touched,
// so a renamed or disabled `main` stays renamed and disabled. It creates
// nothing else and never removes a location.
func (srv *InventoriesLocations) InventoriesLocationsDefaults()(*interface{}, error) {
	path := "/v1/inventories/locations/defaults"
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
	
// InventoriesLocationsDelete deleting one takes every `stock_levels` row that
// points at it with it — the foreign key decides that, not this route. What
// the database does NOT clean up is everything else carrying the same id:
// `stock_movements.location_id` and `reservations.location_id` are plain uuid
// columns and not foreign keys, so those rows stay exactly where they are,
// pointing at a row that no longer exists, and nothing nulls the pointer.
// That asymmetry destroys the balances and keeps everything that refers to
// them, so the route REFUSES while anything still depends on the location and
// answers 409 with the count — taken here rather than left to whoever is
// about to click delete, because a client that pre-counts asks a second
// question whose answer disagrees the moment a receipt lands between the two
// calls. Two things block it. A stock row still carrying `on_hand`: the
// cascade would destroy recorded inventory and nothing in this app ever
// replays the ledger to rebuild a balance, so there is no undo. And a
// reservation still `active`: a promise to a customer must not outlive the
// row backing it — such a hold used to survive its stock row, after which
// /release lowered no `reserved` and still wrote its `release` booking, and
// /commit booked the whole quantity as a shortfall, neither of them an error.
// A stock row at zero does not block: it records no quantity. HISTORY never
// blocks, and is never deleted either — a movement is an accounting record
// and removing one would falsify it, so the bookings stay, naming a location
// that no longer resolves, BY DESIGN. A location that once had traffic and
// now holds nothing is exactly what a merchant closes. To get past the 409,
// adjust the stock to zero and release or commit the holds; where the
// location is merely out of service, PUT `enabled: false` keeps every row and
// can be undone.
func (srv *InventoriesLocations) InventoriesLocationsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/locations/{id}")
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
	
// InventoriesLocationsGet a location is WHERE stock is kept — a warehouse,
// a shop floor, a supplier that dropships, or a virtual bucket for pre-orders
// and quarantine. It holds no quantity of its own: what is at it is a stock
// level. `type` is descriptive and nothing branches on it; `priority` is the
// number that decides which location a reservation is served from, and
// `enabled` decides whether it is offered at all. This is the route that
// turns an id back into a place: `location_id` is on every stock row, every
// ledger booking and every reservation, and none of them carries the code or
// the name. Reading it also answers the two questions those rows raise —
// whether the location is still `enabled` (a disabled one is skipped by
// availability and reserve while its stock stays exactly where it is) and
// where its `priority` puts it when the allocation strategy picks somewhere
// to reserve from.
func (srv *InventoriesLocations) InventoriesLocationsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/locations/{id}")
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
type InventoriesLocationsUpdateOptions struct {
	Address interface{}
	Code string
	Enabled bool
	Labels interface{}
	Metadata interface{}
	Name string
	Priority int
	Type string
	enabledSetters map[string]bool
}
func (options InventoriesLocationsUpdateOptions) New() *InventoriesLocationsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Address": false,
		"Code": false,
		"Enabled": false,
		"Labels": false,
		"Metadata": false,
		"Name": false,
		"Priority": false,
		"Type": false,
	}
	return &options
}
type InventoriesLocationsUpdateOption func(*InventoriesLocationsUpdateOptions)
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateAddress(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateCode(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateEnabled(v bool) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateLabels(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateMetadata(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateName(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdatePriority(v int) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *InventoriesLocations) WithInventoriesLocationsUpdateType(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
			
// InventoriesLocationsUpdate partial update: send the fields that change. The
// one with consequences is `enabled` — setting it to false is how a
// location is taken out of service WITHOUT losing anything. Availability and
// reserve stop looking at it, so its stock stops being sellable, while every
// stock row, ledger booking and reservation that points at it survives
// untouched and comes back the moment it is enabled again. That is the
// reversible alternative to DELETE, which is not reversible at all. Changing
// `code` is the other sharp edge: rows keep their `location_id` so nothing
// moves, but every caller that names the old code in `location_code` starts
// getting 400 "unknown location". Two rows of this tenant may not share
// `code` — that is the 409, and it answers an update that moves a row onto
// a sibling's value exactly as it answers a second insert.
func (srv *InventoriesLocations) InventoriesLocationsUpdate(Id string, optionalSetters ...InventoriesLocationsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/inventories/locations/{id}")
	options := InventoriesLocationsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Address"] {
		params["address"] = options.Address
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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
