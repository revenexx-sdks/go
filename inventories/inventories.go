package inventories

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Inventories service
type Inventories struct {
	client client.Client
}

func New(clt client.Client) *Inventories {
	return &Inventories{
		client: clt,
	}
}

type InventoriesAdjustOptions struct {
	LocationCode string
	enabledSetters map[string]bool
}
func (options InventoriesAdjustOptions) New() *InventoriesAdjustOptions {
	options.enabledSetters = map[string]bool{
		"LocationCode": false,
	}
	return &options
}
type InventoriesAdjustOption func(*InventoriesAdjustOptions)
func (srv *Inventories) WithInventoriesAdjustLocationCode(v string) InventoriesAdjustOption {
	return func(o *InventoriesAdjustOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
					
// InventoriesAdjust
func (srv *Inventories) InventoriesAdjust(Items []models.InventoryAdjustItem, Reason string, optionalSetters ...InventoriesAdjustOption)(*interface{}, error) {
	path := "/v1/inventories/adjust"
	options := InventoriesAdjustOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	params["reason"] = Reason
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
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
type InventoriesAvailabilityOptions struct {
	LocationCode string
	enabledSetters map[string]bool
}
func (options InventoriesAvailabilityOptions) New() *InventoriesAvailabilityOptions {
	options.enabledSetters = map[string]bool{
		"LocationCode": false,
	}
	return &options
}
type InventoriesAvailabilityOption func(*InventoriesAvailabilityOptions)
func (srv *Inventories) WithInventoriesAvailabilityLocationCode(v string) InventoriesAvailabilityOption {
	return func(o *InventoriesAvailabilityOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
			
// InventoriesAvailability
func (srv *Inventories) InventoriesAvailability(Items []models.InventoryAvailabilityItem, optionalSetters ...InventoriesAvailabilityOption)(*interface{}, error) {
	path := "/v1/inventories/availability"
	options := InventoriesAvailabilityOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
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
	
// InventoriesCommit
func (srv *Inventories) InventoriesCommit(OrderRef string)(*interface{}, error) {
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

// InventoriesLocationsList
func (srv *Inventories) InventoriesLocationsList()(*interface{}, error) {
	path := "/v1/inventories/locations"
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
func (srv *Inventories) WithInventoriesLocationsCreateAddress(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsCreateEnabled(v bool) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsCreateLabels(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsCreateMetadata(v interface{}) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsCreatePriority(v int) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsCreateType(v string) InventoriesLocationsCreateOption {
	return func(o *InventoriesLocationsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// InventoriesLocationsCreate
func (srv *Inventories) InventoriesLocationsCreate(Code string, Name string, optionalSetters ...InventoriesLocationsCreateOption)(*models.Location, error) {
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

		parsed := models.Location{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Location
	parsed, ok := resp.Result.(models.Location)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// InventoriesLocationsDefaults
func (srv *Inventories) InventoriesLocationsDefaults()(*interface{}, error) {
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
	
// InventoriesLocationsDelete
func (srv *Inventories) InventoriesLocationsDelete(Id string)(*interface{}, error) {
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
	
// InventoriesLocationsGet
func (srv *Inventories) InventoriesLocationsGet(Id string)(*models.Location, error) {
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

		parsed := models.Location{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Location
	parsed, ok := resp.Result.(models.Location)
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
func (srv *Inventories) WithInventoriesLocationsUpdateAddress(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateCode(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateEnabled(v bool) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateLabels(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateMetadata(v interface{}) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateName(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdatePriority(v int) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Inventories) WithInventoriesLocationsUpdateType(v string) InventoriesLocationsUpdateOption {
	return func(o *InventoriesLocationsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
			
// InventoriesLocationsUpdate
func (srv *Inventories) InventoriesLocationsUpdate(Id string, optionalSetters ...InventoriesLocationsUpdateOption)(*models.Location, error) {
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

		parsed := models.Location{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Location
	parsed, ok := resp.Result.(models.Location)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// InventoriesMovementsList
func (srv *Inventories) InventoriesMovementsList()(*interface{}, error) {
	path := "/v1/inventories/movements"
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
	
// InventoriesMovementsGet
func (srv *Inventories) InventoriesMovementsGet(Id string)(*models.StockMovement, error) {
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

		parsed := models.StockMovement{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.StockMovement
	parsed, ok := resp.Result.(models.StockMovement)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type InventoriesReceiveOptions struct {
	LocationCode string
	Reason string
	enabledSetters map[string]bool
}
func (options InventoriesReceiveOptions) New() *InventoriesReceiveOptions {
	options.enabledSetters = map[string]bool{
		"LocationCode": false,
		"Reason": false,
	}
	return &options
}
type InventoriesReceiveOption func(*InventoriesReceiveOptions)
func (srv *Inventories) WithInventoriesReceiveLocationCode(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *Inventories) WithInventoriesReceiveReason(v string) InventoriesReceiveOption {
	return func(o *InventoriesReceiveOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// InventoriesReceive
func (srv *Inventories) InventoriesReceive(Items []models.InventoryStockItem, optionalSetters ...InventoriesReceiveOption)(*interface{}, error) {
	path := "/v1/inventories/receive"
	options := InventoriesReceiveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
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
	
// InventoriesRelease
func (srv *Inventories) InventoriesRelease(OrderRef string)(*interface{}, error) {
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

// InventoriesReservationsList
func (srv *Inventories) InventoriesReservationsList()(*interface{}, error) {
	path := "/v1/inventories/reservations"
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
	
// InventoriesReservationsGet
func (srv *Inventories) InventoriesReservationsGet(Id string)(*models.Reservation, error) {
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

		parsed := models.Reservation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Reservation
	parsed, ok := resp.Result.(models.Reservation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type InventoriesReserveOptions struct {
	ExpiresAt string
	enabledSetters map[string]bool
}
func (options InventoriesReserveOptions) New() *InventoriesReserveOptions {
	options.enabledSetters = map[string]bool{
		"ExpiresAt": false,
	}
	return &options
}
type InventoriesReserveOption func(*InventoriesReserveOptions)
func (srv *Inventories) WithInventoriesReserveExpiresAt(v string) InventoriesReserveOption {
	return func(o *InventoriesReserveOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
					
// InventoriesReserve
func (srv *Inventories) InventoriesReserve(Items []models.InventoryStockItem, OrderRef string, optionalSetters ...InventoriesReserveOption)(*interface{}, error) {
	path := "/v1/inventories/reserve"
	options := InventoriesReserveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	params["order_ref"] = OrderRef
	if options.enabledSetters["ExpiresAt"] {
		params["expires_at"] = options.ExpiresAt
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
type InventoriesRestockOptions struct {
	LocationCode string
	OrderRef string
	Reason string
	enabledSetters map[string]bool
}
func (options InventoriesRestockOptions) New() *InventoriesRestockOptions {
	options.enabledSetters = map[string]bool{
		"LocationCode": false,
		"OrderRef": false,
		"Reason": false,
	}
	return &options
}
type InventoriesRestockOption func(*InventoriesRestockOptions)
func (srv *Inventories) WithInventoriesRestockLocationCode(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.LocationCode = v
		o.enabledSetters["LocationCode"] = true
	}
}
func (srv *Inventories) WithInventoriesRestockOrderRef(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *Inventories) WithInventoriesRestockReason(v string) InventoriesRestockOption {
	return func(o *InventoriesRestockOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// InventoriesRestock
func (srv *Inventories) InventoriesRestock(Items []models.InventoryStockItem, optionalSetters ...InventoriesRestockOption)(*interface{}, error) {
	path := "/v1/inventories/restock"
	options := InventoriesRestockOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["LocationCode"] {
		params["location_code"] = options.LocationCode
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
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

// InventoriesStockList
func (srv *Inventories) InventoriesStockList()(*interface{}, error) {
	path := "/v1/inventories/stock"
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
type InventoriesStockCreateOptions struct {
	Metadata interface{}
	OnHand float64
	ProductId string
	ReorderPoint float64
	Reserved float64
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesStockCreateOptions) New() *InventoriesStockCreateOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"OnHand": false,
		"ProductId": false,
		"ReorderPoint": false,
		"Reserved": false,
		"Sku": false,
	}
	return &options
}
type InventoriesStockCreateOption func(*InventoriesStockCreateOptions)
func (srv *Inventories) WithInventoriesStockCreateMetadata(v interface{}) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Inventories) WithInventoriesStockCreateOnHand(v float64) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.OnHand = v
		o.enabledSetters["OnHand"] = true
	}
}
func (srv *Inventories) WithInventoriesStockCreateProductId(v string) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Inventories) WithInventoriesStockCreateReorderPoint(v float64) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.ReorderPoint = v
		o.enabledSetters["ReorderPoint"] = true
	}
}
func (srv *Inventories) WithInventoriesStockCreateReserved(v float64) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.Reserved = v
		o.enabledSetters["Reserved"] = true
	}
}
func (srv *Inventories) WithInventoriesStockCreateSku(v string) InventoriesStockCreateOption {
	return func(o *InventoriesStockCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
			
// InventoriesStockCreate
func (srv *Inventories) InventoriesStockCreate(LocationId string, optionalSetters ...InventoriesStockCreateOption)(*models.StockLevel, error) {
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
	if options.enabledSetters["OnHand"] {
		params["on_hand"] = options.OnHand
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["ReorderPoint"] {
		params["reorder_point"] = options.ReorderPoint
	}
	if options.enabledSetters["Reserved"] {
		params["reserved"] = options.Reserved
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

		parsed := models.StockLevel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.StockLevel
	parsed, ok := resp.Result.(models.StockLevel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// InventoriesStockDelete
func (srv *Inventories) InventoriesStockDelete(Id string)(*interface{}, error) {
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
	
// InventoriesStockGet
func (srv *Inventories) InventoriesStockGet(Id string)(*models.StockLevel, error) {
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

		parsed := models.StockLevel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.StockLevel
	parsed, ok := resp.Result.(models.StockLevel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type InventoriesStockUpdateOptions struct {
	LocationId string
	Metadata interface{}
	OnHand float64
	ProductId string
	ReorderPoint float64
	Reserved float64
	Sku string
	enabledSetters map[string]bool
}
func (options InventoriesStockUpdateOptions) New() *InventoriesStockUpdateOptions {
	options.enabledSetters = map[string]bool{
		"LocationId": false,
		"Metadata": false,
		"OnHand": false,
		"ProductId": false,
		"ReorderPoint": false,
		"Reserved": false,
		"Sku": false,
	}
	return &options
}
type InventoriesStockUpdateOption func(*InventoriesStockUpdateOptions)
func (srv *Inventories) WithInventoriesStockUpdateLocationId(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.LocationId = v
		o.enabledSetters["LocationId"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateMetadata(v interface{}) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateOnHand(v float64) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.OnHand = v
		o.enabledSetters["OnHand"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateProductId(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateReorderPoint(v float64) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.ReorderPoint = v
		o.enabledSetters["ReorderPoint"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateReserved(v float64) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.Reserved = v
		o.enabledSetters["Reserved"] = true
	}
}
func (srv *Inventories) WithInventoriesStockUpdateSku(v string) InventoriesStockUpdateOption {
	return func(o *InventoriesStockUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
			
// InventoriesStockUpdate
func (srv *Inventories) InventoriesStockUpdate(Id string, optionalSetters ...InventoriesStockUpdateOption)(*models.StockLevel, error) {
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
	if options.enabledSetters["OnHand"] {
		params["on_hand"] = options.OnHand
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["ReorderPoint"] {
		params["reorder_point"] = options.ReorderPoint
	}
	if options.enabledSetters["Reserved"] {
		params["reserved"] = options.Reserved
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

		parsed := models.StockLevel{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.StockLevel
	parsed, ok := resp.Result.(models.StockLevel)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
