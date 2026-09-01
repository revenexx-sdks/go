package carts_items

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CartsItems service
type CartsItems struct {
	client client.Client
}

func New(clt client.Client) *CartsItems {
	return &CartsItems{
		client: clt,
	}
}

type CartsItemsListOptions struct {
	Id string
	Type string
	ProductId string
	Sku string
	Name string
	Quantity float64
	Unit string
	UnitPrice float64
	Currency string
	TaxRate float64
	LineTotal float64
	Position int
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CartsItemsListOptions) New() *CartsItemsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Type": false,
		"ProductId": false,
		"Sku": false,
		"Name": false,
		"Quantity": false,
		"Unit": false,
		"UnitPrice": false,
		"Currency": false,
		"TaxRate": false,
		"LineTotal": false,
		"Position": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CartsItemsListOption func(*CartsItemsListOptions)
func (srv *CartsItems) WithCartsItemsListId(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListType(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListProductId(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListSku(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListName(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListQuantity(v float64) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListUnit(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListUnitPrice(v float64) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListCurrency(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListTaxRate(v float64) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListLineTotal(v float64) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.LineTotal = v
		o.enabledSetters["LineTotal"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListPosition(v int) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListCreatedAt(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListUpdatedAt(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListLimit(v int) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListOffset(v int) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CartsItems) WithCartsItemsListOrder(v string) CartsItemsListOption {
	return func(o *CartsItemsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// CartsItemsList the array is still called 'items'; the response also carries
// 'page' and 'filter' like every other list, and an unknown cart_id answers
// 404 instead of an empty page. A cart with more lines than the page size is
// not silently truncated — 'page.hasMore' says so. Lines come back in
// position order unless 'order' says otherwise.
func (srv *CartsItems) CartsItemsList(CartId string, optionalSetters ...CartsItemsListOption)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId)
	path := r.Replace("/v1/carts/{cart_id}/items")
	options := CartsItemsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["cart_id"] = CartId
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["LineTotal"] {
		params["line_total"] = options.LineTotal
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type CartsItemsCreateOptions struct {
	Configuration interface{}
	Currency string
	Metadata interface{}
	Name string
	Position int
	ProductId string
	Quantity float64
	Sku string
	Snapshot interface{}
	TaxRate float64
	Type string
	Unit string
	UnitPrice float64
	enabledSetters map[string]bool
}
func (options CartsItemsCreateOptions) New() *CartsItemsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Configuration": false,
		"Currency": false,
		"Metadata": false,
		"Name": false,
		"Position": false,
		"ProductId": false,
		"Quantity": false,
		"Sku": false,
		"Snapshot": false,
		"TaxRate": false,
		"Type": false,
		"Unit": false,
		"UnitPrice": false,
	}
	return &options
}
type CartsItemsCreateOption func(*CartsItemsCreateOptions)
func (srv *CartsItems) WithCartsItemsCreateConfiguration(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Configuration = v
		o.enabledSetters["Configuration"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateCurrency(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateMetadata(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateName(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreatePosition(v int) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateProductId(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateQuantity(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateSku(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateSnapshot(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Snapshot = v
		o.enabledSetters["Snapshot"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateTaxRate(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateType(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateUnit(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *CartsItems) WithCartsItemsCreateUnitPrice(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
			
// CartsItemsCreate adds one line to an ACTIVE cart — the add-to-basket
// call. `name` or `sku` is required (a line sent with only a SKU takes the
// SKU as its name, so a line always has something to show) and `quantity`
// must be greater than zero; everything else defaults, including the
// currency, which falls back to the cart's. The one thing that surprises a
// caller: a plain product line with the same product/sku AND the same
// `unit_price` as a line already in the cart does not open a second row —
// its quantity is added to that line, and the 201 names a row that already
// existed. Price is part of that identity on purpose, so a changed price
// never averages into an old line. A configured or custom line always stands
// alone. The cart's `item_count` (the sum of QUANTITIES) and `subtotal` are
// recomputed before the answer, and `max_items_per_cart` /
// `max_quantity_per_line` are checked on the RESULT of the merge (422), so
// ten calls of one piece cannot walk past a limit one call of ten would hit.
func (srv *CartsItems) CartsItemsCreate(CartId string, optionalSetters ...CartsItemsCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId)
	path := r.Replace("/v1/carts/{cart_id}/items")
	options := CartsItemsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["cart_id"] = CartId
	if options.enabledSetters["Configuration"] {
		params["configuration"] = options.Configuration
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Snapshot"] {
		params["snapshot"] = options.Snapshot
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
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
			
// CartsItemsReplace set semantics: the payload IS the cart. Every existing
// line is dropped and the payload is written in its place, so a line left out
// of the array is a line removed — this is the storefront sync, not a bulk
// add, and carts.items.create is what adds. Lines are numbered by their place
// in the array unless they carry their own `position`, and nothing merges:
// two identical lines in one payload stay two rows. The limits are checked
// against the payload BEFORE a single existing line is destroyed, so a sync
// refused with 422 leaves the cart exactly as it was. The cart must be
// active, and its totals are recomputed before the answer.
func (srv *CartsItems) CartsItemsReplace(CartId string, Items []models.CartItemCreateRequest)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId)
	path := r.Replace("/v1/carts/{cart_id}/items")
	params := map[string]interface{}{}
	params["cart_id"] = CartId
	params["items"] = Items
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
			
// CartsItemsDelete removes one line from an ACTIVE cart and recomputes the
// owning cart's `item_count` and `subtotal` before answering. This is how a
// quantity reaches zero: `quantity` is constrained to be greater than zero,
// so "none of it" is a DELETE and never an update to 0. The cart in the path
// is part of the address — a line belonging to a different cart answers 404
// and is left where it is. Deleting the last line leaves an empty cart, not a
// deleted one; the cart itself goes through carts.delete, which takes every
// line with it in one call.
func (srv *CartsItems) CartsItemsDelete(CartId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId, "{id}", Id)
	path := r.Replace("/v1/carts/{cart_id}/items/{id}")
	params := map[string]interface{}{}
	params["cart_id"] = CartId
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
			
// CartsItemsGet one line, addressed through the cart that owns it. Both ids
// are checked, not just the line's: a line that exists but belongs to a
// different cart answers 404 rather than the row, so an id copied out of
// another cart never resolves here and a caller can trust that what came back
// is a line of the cart they asked about. The line carries both of its prices
// — the working `unit_price`, which a resync or a repricing job may have
// moved, and the `snapshot` the buyer was shown when the line was added —
// and its own `line_total`, which is always quantity × unit_price and never
// what a payload claimed. To read a whole cart's lines, list them: this route
// is for one known line.
func (srv *CartsItems) CartsItemsGet(CartId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId, "{id}", Id)
	path := r.Replace("/v1/carts/{cart_id}/items/{id}")
	params := map[string]interface{}{}
	params["cart_id"] = CartId
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
type CartsItemsUpdateOptions struct {
	Configuration interface{}
	Currency string
	Metadata interface{}
	Name string
	Position int
	ProductId string
	Quantity float64
	Sku string
	Snapshot interface{}
	TaxRate float64
	Type string
	Unit string
	UnitPrice float64
	enabledSetters map[string]bool
}
func (options CartsItemsUpdateOptions) New() *CartsItemsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Configuration": false,
		"Currency": false,
		"Metadata": false,
		"Name": false,
		"Position": false,
		"ProductId": false,
		"Quantity": false,
		"Sku": false,
		"Snapshot": false,
		"TaxRate": false,
		"Type": false,
		"Unit": false,
		"UnitPrice": false,
	}
	return &options
}
type CartsItemsUpdateOption func(*CartsItemsUpdateOptions)
func (srv *CartsItems) WithCartsItemsUpdateConfiguration(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Configuration = v
		o.enabledSetters["Configuration"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateCurrency(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateMetadata(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateName(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdatePosition(v int) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateProductId(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateQuantity(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateSku(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateSnapshot(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Snapshot = v
		o.enabledSetters["Snapshot"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateTaxRate(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateType(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateUnit(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *CartsItems) WithCartsItemsUpdateUnitPrice(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
					
// CartsItemsUpdate changes one line of an ACTIVE cart — the quantity
// stepper on the cart page, and the route a repricing job writes through. The
// fields sent are merged onto the stored line and the whole line is validated
// again, so `quantity` must still be greater than zero and `type` still one
// of the three. `line_total` is not settable: it is recomputed as quantity ×
// unit_price, and the cart's `item_count` and `subtotal` follow before the
// answer. What it will NOT do is merge — only carts.items.create folds one
// line into another, so giving this line the same product and price as a
// sibling leaves two rows standing, and the next add joins whichever it
// matches. `max_quantity_per_line` is enforced on the result (422). A
// quantity of zero is not the way to remove a line; the delete is.
func (srv *CartsItems) CartsItemsUpdate(CartId string, Id string, optionalSetters ...CartsItemsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{cart_id}", CartId, "{id}", Id)
	path := r.Replace("/v1/carts/{cart_id}/items/{id}")
	options := CartsItemsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["cart_id"] = CartId
	params["id"] = Id
	if options.enabledSetters["Configuration"] {
		params["configuration"] = options.Configuration
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Snapshot"] {
		params["snapshot"] = options.Snapshot
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
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
