package products

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Products service
type Products struct {
	client client.Client
}

func New(clt client.Client) *Products {
	return &Products{
		client: clt,
	}
}

type ProductsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Sku string
	Kind string
	ParentId string
	FamilyId string
	FamilyVariantId string
	Enabled bool
	TaxClass string
	AttributeValues string
	Label string
	QuantifiedAssociations string
	Completeness string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	enabledSetters map[string]bool
}
func (options ProductsListOptions) New() *ProductsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Sku": false,
		"Kind": false,
		"ParentId": false,
		"FamilyId": false,
		"FamilyVariantId": false,
		"Enabled": false,
		"TaxClass": false,
		"AttributeValues": false,
		"Label": false,
		"QuantifiedAssociations": false,
		"Completeness": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"DeletedAt": false,
	}
	return &options
}
type ProductsListOption func(*ProductsListOptions)
func (srv *Products) WithProductsListLimit(v int) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Products) WithProductsListOffset(v int) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Products) WithProductsListOrder(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Products) WithProductsListId(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Products) WithProductsListSku(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Products) WithProductsListKind(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Products) WithProductsListParentId(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *Products) WithProductsListFamilyId(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *Products) WithProductsListFamilyVariantId(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.FamilyVariantId = v
		o.enabledSetters["FamilyVariantId"] = true
	}
}
func (srv *Products) WithProductsListEnabled(v bool) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Products) WithProductsListTaxClass(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
func (srv *Products) WithProductsListAttributeValues(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsListLabel(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Products) WithProductsListQuantifiedAssociations(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.QuantifiedAssociations = v
		o.enabledSetters["QuantifiedAssociations"] = true
	}
}
func (srv *Products) WithProductsListCompleteness(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.Completeness = v
		o.enabledSetters["Completeness"] = true
	}
}
func (srv *Products) WithProductsListCreatedAt(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Products) WithProductsListUpdatedAt(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Products) WithProductsListDeletedAt(v string) ProductsListOption {
	return func(o *ProductsListOptions) {
		o.DeletedAt = v
		o.enabledSetters["DeletedAt"] = true
	}
}
	
// ProductsList the catalog itself. A product row carries only what every
// product has — SKU, kind, family, enabled, tax class — and everything
// the tenant modelled lives in the `attribute_values` jsonb document, keyed
// by attribute CODE inside one of four scope buckets (common, per locale, per
// channel, per channel and locale). `label` is a generated column, maintained
// by the database so a grid of twenty thousand rows can sort and filter on a
// name with no join. `kind` says where the row sits in the variant hierarchy:
// a `model` carries what its variants share and is never sold itself.
// 
// Every column of `products` is an exact-match query parameter, `order` sorts
// by one column, and `limit`/`offset` page through `page.total`. A query key
// that is NOT a column is dropped rather than refused, and the `filter`
// object echoes the ones that were understood — that echo is the only way
// to tell an unfiltered answer from an empty one. It reads rows exactly as
// they are stored: no join is resolved, no jsonb value is unpacked, and
// soft-deleted products are included — filter on `deleted_at` to read the
// live catalog, or use `GET /products/grid`, which excludes them.
func (srv *Products) ProductsList(optionalSetters ...ProductsListOption)(*interface{}, error) {
	path := "/v1/products"
	options := ProductsListOptions{}.New()
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
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["FamilyVariantId"] {
		params["family_variant_id"] = options.FamilyVariantId
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["TaxClass"] {
		params["tax_class"] = options.TaxClass
	}
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["QuantifiedAssociations"] {
		params["quantified_associations"] = options.QuantifiedAssociations
	}
	if options.enabledSetters["Completeness"] {
		params["completeness"] = options.Completeness
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
	}
	if options.enabledSetters["DeletedAt"] {
		params["deleted_at"] = options.DeletedAt
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
type ProductsCreateOptions struct {
	AttributeValues interface{}
	Completeness interface{}
	DeletedAt string
	Enabled bool
	FamilyId string
	FamilyVariantId string
	Kind string
	ParentId string
	QuantifiedAssociations interface{}
	TaxClass string
	enabledSetters map[string]bool
}
func (options ProductsCreateOptions) New() *ProductsCreateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"Completeness": false,
		"DeletedAt": false,
		"Enabled": false,
		"FamilyId": false,
		"FamilyVariantId": false,
		"Kind": false,
		"ParentId": false,
		"QuantifiedAssociations": false,
		"TaxClass": false,
	}
	return &options
}
type ProductsCreateOption func(*ProductsCreateOptions)
func (srv *Products) WithProductsCreateAttributeValues(v interface{}) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsCreateCompleteness(v interface{}) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.Completeness = v
		o.enabledSetters["Completeness"] = true
	}
}
func (srv *Products) WithProductsCreateDeletedAt(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.DeletedAt = v
		o.enabledSetters["DeletedAt"] = true
	}
}
func (srv *Products) WithProductsCreateEnabled(v bool) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Products) WithProductsCreateFamilyId(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *Products) WithProductsCreateFamilyVariantId(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.FamilyVariantId = v
		o.enabledSetters["FamilyVariantId"] = true
	}
}
func (srv *Products) WithProductsCreateKind(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Products) WithProductsCreateParentId(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *Products) WithProductsCreateQuantifiedAssociations(v interface{}) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.QuantifiedAssociations = v
		o.enabledSetters["QuantifiedAssociations"] = true
	}
}
func (srv *Products) WithProductsCreateTaxClass(v string) ProductsCreateOption {
	return func(o *ProductsCreateOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
			
// ProductsCreate creates one product and answers 201 with the stored row,
// including the id and the timestamps the database filled in — a client
// never sends an id, it reads one back and uses it in the path of every later
// call.
// 
// The catalog itself. A product row carries only what every product has —
// SKU, kind, family, enabled, tax class — and everything the tenant
// modelled lives in the `attribute_values` jsonb document, keyed by attribute
// CODE inside one of four scope buckets (common, per locale, per channel, per
// channel and locale). `label` is a generated column, maintained by the
// database so a grid of twenty thousand rows can sort and filter on a name
// with no join. `kind` says where the row sits in the variant hierarchy: a
// `model` carries what its variants share and is never sold itself.
// 
// `sku` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `sku` answers
// 409. This app owns the create: `enabled` defaults from the
// `new_products_enabled_by_default` tenant setting rather than blindly to
// true, so an import cannot publish twenty thousand unfinished products the
// moment it lands, and a product that names no family gets the
// `default_product_family` one. An explicit value in the body always wins
// over both.
func (srv *Products) ProductsCreate(Sku string, optionalSetters ...ProductsCreateOption)(*models.Error, error) {
	path := "/v1/products"
	options := ProductsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["sku"] = Sku
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["Completeness"] {
		params["completeness"] = options.Completeness
	}
	if options.enabledSetters["DeletedAt"] {
		params["deleted_at"] = options.DeletedAt
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["FamilyVariantId"] {
		params["family_variant_id"] = options.FamilyVariantId
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["QuantifiedAssociations"] {
		params["quantified_associations"] = options.QuantifiedAssociations
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
type ProductsBatchOptions struct {
	Ids []string
	Skus []string
	enabledSetters map[string]bool
}
func (options ProductsBatchOptions) New() *ProductsBatchOptions {
	options.enabledSetters = map[string]bool{
		"Ids": false,
		"Skus": false,
	}
	return &options
}
type ProductsBatchOption func(*ProductsBatchOptions)
func (srv *Products) WithProductsBatchIds(v []string) ProductsBatchOption {
	return func(o *ProductsBatchOptions) {
		o.Ids = v
		o.enabledSetters["Ids"] = true
	}
}
func (srv *Products) WithProductsBatchSkus(v []string) ProductsBatchOption {
	return func(o *ProductsBatchOptions) {
		o.Skus = v
		o.enabledSetters["Skus"] = true
	}
}
	
// ProductsBatch answers four fields — id, sku, tax_class and the resolved
// display name — for a list of ids and/or SKUs in ONE call. It exists for
// the app on the other side of a product reference: the prices app holds SKUs
// and needs a tax class, a feed builder holds ids and needs names, and
// neither should page through the catalog or fire a request per line. Ask by
// either identifier or both; the two are unioned and a product named twice
// comes back once.
// 
// It answers what it FOUND: an id or SKU that names nothing is simply absent
// from `items` rather than an error, so compare the length of what you sent
// with what came back if a miss matters. It is not a general product read —
// for the whole row use `GET /products/{id}`, and for a scannable list use
// `GET /products/grid`.
func (srv *Products) ProductsBatch(optionalSetters ...ProductsBatchOption)(*interface{}, error) {
	path := "/v1/products/batch"
	options := ProductsBatchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Ids"] {
		params["ids"] = options.Ids
	}
	if options.enabledSetters["Skus"] {
		params["skus"] = options.Skus
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
type ProductsGridOptions struct {
	Limit int
	Offset int
	Order string
	Q string
	Kind string
	Enabled bool
	FamilyId string
	enabledSetters map[string]bool
}
func (options ProductsGridOptions) New() *ProductsGridOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Q": false,
		"Kind": false,
		"Enabled": false,
		"FamilyId": false,
	}
	return &options
}
type ProductsGridOption func(*ProductsGridOptions)
func (srv *Products) WithProductsGridLimit(v int) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Products) WithProductsGridOffset(v int) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Products) WithProductsGridOrder(v string) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Products) WithProductsGridQ(v string) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Q = v
		o.enabledSetters["Q"] = true
	}
}
func (srv *Products) WithProductsGridKind(v string) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Products) WithProductsGridEnabled(v bool) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Products) WithProductsGridFamilyId(v string) ProductsGridOption {
	return func(o *ProductsGridOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
	
// ProductsGrid the list a merchant can actually scan, as opposed to `GET
// /products`, which answers SKUs and a jsonb blob. Every row arrives already
// flattened: its resolved display name and where that name came from, its
// family code, its stored completeness, and the value of every attribute the
// catalog marks `usable_in_grid` — no join, no second call. `q` is a
// case-insensitive substring of the stored `label` column, which falls back
// to the SKU, so one box finds a product by either. Soft-deleted products are
// excluded here, unlike `GET /products`.
// 
// It filters on `q`, `kind`, `enabled` and `family_id`, and on NOTHING ELSE
// — a query parameter it does not accept is refused with 400 rather than
// dropped. That matters because of `filters`: the array reports the
// attributes marked `is_filterable`, which is what a filter bar should OFFER,
// and it is not a query surface. Filtering on an attribute value is not
// offered by this API at all — the values live inside a four-bucket jsonb
// document and are read through a fallback chain, so it is a feature with a
// design of its own rather than a parameter that was forgotten.
func (srv *Products) ProductsGrid(optionalSetters ...ProductsGridOption)(*models.Error, error) {
	path := "/v1/products/grid"
	options := ProductsGridOptions{}.New()
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
	if options.enabledSetters["Q"] {
		params["q"] = options.Q
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
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
type ProductsLabelsOptions struct {
	Ids []string
	Skus []string
	enabledSetters map[string]bool
}
func (options ProductsLabelsOptions) New() *ProductsLabelsOptions {
	options.enabledSetters = map[string]bool{
		"Ids": false,
		"Skus": false,
	}
	return &options
}
type ProductsLabelsOption func(*ProductsLabelsOptions)
func (srv *Products) WithProductsLabelsIds(v []string) ProductsLabelsOption {
	return func(o *ProductsLabelsOptions) {
		o.Ids = v
		o.enabledSetters["Ids"] = true
	}
}
func (srv *Products) WithProductsLabelsSkus(v []string) ProductsLabelsOption {
	return func(o *ProductsLabelsOptions) {
		o.Skus = v
		o.enabledSetters["Skus"] = true
	}
}
	
// ProductsLabels what is this product CALLED? A product's name is an
// attribute rather than a column, and which attribute it is, is per family
// — so no plain read can answer it. This resolves up to 500 products at
// once, by id and/or SKU: it reads families.label_attribute (falling back to
// the default_label_attribute setting, then to the conventional `name`) and
// looks the value up through the scoped attribute_values document — common,
// then locale_specific in the label_locales order, then the channel buckets.
// 
// It reports WHERE the name was found, which is the half that matters:
// `source: "sku"` means the catalog holds no name for this product and the
// SKU is standing in for one, so show it as a missing name rather than as a
// name. Writes nothing, and answers only what it found.
func (srv *Products) ProductsLabels(optionalSetters ...ProductsLabelsOption)(*models.Error, error) {
	path := "/v1/products/labels"
	options := ProductsLabelsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Ids"] {
		params["ids"] = options.Ids
	}
	if options.enabledSetters["Skus"] {
		params["skus"] = options.Skus
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
type ProductsProductAssociationsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	ProductId string
	AssociationTypeId string
	TargetProductId string
	Quantity float64
	Position int
	CreatedAt string
	enabledSetters map[string]bool
}
func (options ProductsProductAssociationsListOptions) New() *ProductsProductAssociationsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"ProductId": false,
		"AssociationTypeId": false,
		"TargetProductId": false,
		"Quantity": false,
		"Position": false,
		"CreatedAt": false,
	}
	return &options
}
type ProductsProductAssociationsListOption func(*ProductsProductAssociationsListOptions)
func (srv *Products) WithProductsProductAssociationsListLimit(v int) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListOffset(v int) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListOrder(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListId(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListProductId(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListAssociationTypeId(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.AssociationTypeId = v
		o.enabledSetters["AssociationTypeId"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListTargetProductId(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.TargetProductId = v
		o.enabledSetters["TargetProductId"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListQuantity(v float64) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListPosition(v int) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsListCreatedAt(v string) ProductsProductAssociationsListOption {
	return func(o *ProductsProductAssociationsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// ProductsProductAssociationsList one relation from one product to another,
// of a declared type: this drill's accessories, this bundle's parts, this
// article's cross-sells. `quantity` is the number in "this bundle contains 4
// casters" and is meaningful only when the association type carries
// `is_quantified`. This relational surface is the one this app serves; the
// `products.quantified_associations` column is an importer's blob that no
// route here reads or writes.
// 
// Every column of `product_associations` is an exact-match query parameter,
// `order` sorts by one column, and `limit`/`offset` page through
// `page.total`. A query key that is NOT a column is dropped rather than
// refused, and the `filter` object echoes the ones that were understood —
// that echo is the only way to tell an unfiltered answer from an empty one.
// It reads rows exactly as they are stored: no join is resolved, no jsonb
// value is unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *Products) ProductsProductAssociationsList(optionalSetters ...ProductsProductAssociationsListOption)(*interface{}, error) {
	path := "/v1/products/product_associations"
	options := ProductsProductAssociationsListOptions{}.New()
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
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["AssociationTypeId"] {
		params["association_type_id"] = options.AssociationTypeId
	}
	if options.enabledSetters["TargetProductId"] {
		params["target_product_id"] = options.TargetProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type ProductsProductAssociationsCreateOptions struct {
	Position int
	Quantity float64
	enabledSetters map[string]bool
}
func (options ProductsProductAssociationsCreateOptions) New() *ProductsProductAssociationsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Position": false,
		"Quantity": false,
	}
	return &options
}
type ProductsProductAssociationsCreateOption func(*ProductsProductAssociationsCreateOptions)
func (srv *Products) WithProductsProductAssociationsCreatePosition(v int) ProductsProductAssociationsCreateOption {
	return func(o *ProductsProductAssociationsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsCreateQuantity(v float64) ProductsProductAssociationsCreateOption {
	return func(o *ProductsProductAssociationsCreateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
							
// ProductsProductAssociationsCreate creates one product association and
// answers 201 with the stored row, including the id and the timestamps the
// database filled in — a client never sends an id, it reads one back and
// uses it in the path of every later call.
// 
// One relation from one product to another, of a declared type: this drill's
// accessories, this bundle's parts, this article's cross-sells. `quantity` is
// the number in "this bundle contains 4 casters" and is meaningful only when
// the association type carries `is_quantified`. This relational surface is
// the one this app serves; the `products.quantified_associations` column is
// an importer's blob that no route here reads or writes.
// 
// `product_id`, `association_type_id`, `target_product_id` are the only
// columns the database refuses the row without; everything else has a default
// or is nullable. A second row with the same `product_id`,
// `association_type_id`, `target_product_id` answers 409.
func (srv *Products) ProductsProductAssociationsCreate(AssociationTypeId string, ProductId string, TargetProductId string, optionalSetters ...ProductsProductAssociationsCreateOption)(*models.Error, error) {
	path := "/v1/products/product_associations"
	options := ProductsProductAssociationsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["association_type_id"] = AssociationTypeId
	params["product_id"] = ProductId
	params["target_product_id"] = TargetProductId
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
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
	
// ProductsProductAssociationsDelete deletes one product association by id. It
// is a hard delete — the row is gone, and the answer is a confirmation
// rather than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no product association of this tenant carries answers 404; there is
// no 409, because every foreign key pointing at this entity resolves itself
// on delete rather than blocking one.
func (srv *Products) ProductsProductAssociationsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_associations/{id}")
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
	
// ProductsProductAssociationsGet reads one product association by its id —
// the whole row, every column, as it is stored.
// 
// One relation from one product to another, of a declared type: this drill's
// accessories, this bundle's parts, this article's cross-sells. `quantity` is
// the number in "this bundle contains 4 casters" and is meaningful only when
// the association type carries `is_quantified`. This relational surface is
// the one this app serves; the `products.quantified_associations` column is
// an importer's blob that no route here reads or writes.
// 
// An id no product association of this tenant carries answers 404, and so
// does one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *Products) ProductsProductAssociationsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_associations/{id}")
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
type ProductsProductAssociationsUpdateOptions struct {
	AssociationTypeId string
	Position int
	ProductId string
	Quantity float64
	TargetProductId string
	enabledSetters map[string]bool
}
func (options ProductsProductAssociationsUpdateOptions) New() *ProductsProductAssociationsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AssociationTypeId": false,
		"Position": false,
		"ProductId": false,
		"Quantity": false,
		"TargetProductId": false,
	}
	return &options
}
type ProductsProductAssociationsUpdateOption func(*ProductsProductAssociationsUpdateOptions)
func (srv *Products) WithProductsProductAssociationsUpdateAssociationTypeId(v string) ProductsProductAssociationsUpdateOption {
	return func(o *ProductsProductAssociationsUpdateOptions) {
		o.AssociationTypeId = v
		o.enabledSetters["AssociationTypeId"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsUpdatePosition(v int) ProductsProductAssociationsUpdateOption {
	return func(o *ProductsProductAssociationsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsUpdateProductId(v string) ProductsProductAssociationsUpdateOption {
	return func(o *ProductsProductAssociationsUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsUpdateQuantity(v float64) ProductsProductAssociationsUpdateOption {
	return func(o *ProductsProductAssociationsUpdateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Products) WithProductsProductAssociationsUpdateTargetProductId(v string) ProductsProductAssociationsUpdateOption {
	return func(o *ProductsProductAssociationsUpdateOptions) {
		o.TargetProductId = v
		o.enabledSetters["TargetProductId"] = true
	}
}
			
// ProductsProductAssociationsUpdate updates one product association by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// One relation from one product to another, of a declared type: this drill's
// accessories, this bundle's parts, this article's cross-sells. `quantity` is
// the number in "this bundle contains 4 casters" and is meaningful only when
// the association type carries `is_quantified`. This relational surface is
// the one this app serves; the `products.quantified_associations` column is
// an importer's blob that no route here reads or writes.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `product_id`, `association_type_id`, `target_product_id` answers 409.
func (srv *Products) ProductsProductAssociationsUpdate(Id string, optionalSetters ...ProductsProductAssociationsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_associations/{id}")
	options := ProductsProductAssociationsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AssociationTypeId"] {
		params["association_type_id"] = options.AssociationTypeId
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
	if options.enabledSetters["TargetProductId"] {
		params["target_product_id"] = options.TargetProductId
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

// ProductsVocabulariesList the index of the enums this app ENFORCES —
// `product-kinds`, `membership-sources`, `rule-matches`, `asset-sources` —
// served by the app that owns the CHECK constraint each one is parsed out of,
// so a UI never has to keep its own copy of a status map and watch it drift.
// Names and titles only: fetch one by name for its values, badge tones and
// descriptions.
// 
// The set is a fixed property of this app rather than tenant data, so it is
// the same list for every tenant. `attributes.type` is deliberately absent:
// it carries no CHECK, because the whole point of an attribute-driven PIM is
// that the type list is data an integrator extends.
func (srv *Products) ProductsVocabulariesList()(*interface{}, error) {
	path := "/v1/products/vocabularies"
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
	
// ProductsVocabulariesGet one vocabulary with every value it admits, each
// with a title, a description and the badge tone a UI should paint it in. The
// value set is parsed out of the CHECK constraint in schema.json, so what is
// served IS what is enforced. Labels are curated on top and can only add
// words and colour — a permitted value nobody labelled still appears,
// titled from its own key.
func (srv *Products) ProductsVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/products/vocabularies/{name}")
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
	
// ProductsDelete deletes one product by id. It is a hard delete — the row
// is gone, and the answer is a confirmation rather than a result to branch
// on.
// 
// It takes what hangs off it: product category memberships (`product_id`),
// product associations (`product_id` and `target_product_id`) are deleted
// with it. `products.parent_id` is set to null instead, so the rows that
// pointed at it survive the delete rather than going with it.
// 
// An id no product of this tenant carries answers 404; there is no 409,
// because every foreign key pointing at this entity resolves itself on delete
// rather than blocking one. `products.deleted_at` is a SOFT-delete marker
// that the grid and every category-rule evaluation honour, but no route in
// this app ever writes it — to soft-delete instead, `PUT /products/{id}`
// with a `deleted_at`.
func (srv *Products) ProductsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}")
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
	
// ProductsGet reads one product by its id — the whole row, every column, as
// it is stored.
// 
// The catalog itself. A product row carries only what every product has —
// SKU, kind, family, enabled, tax class — and everything the tenant
// modelled lives in the `attribute_values` jsonb document, keyed by attribute
// CODE inside one of four scope buckets (common, per locale, per channel, per
// channel and locale). `label` is a generated column, maintained by the
// database so a grid of twenty thousand rows can sort and filter on a name
// with no join. `kind` says where the row sits in the variant hierarchy: a
// `model` carries what its variants share and is never sold itself.
// 
// An id no product of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached. Nothing is resolved for you here — for the display name, the
// family code and the grid attributes already unpacked, use `GET
// /products/grid` or `POST /products/labels`.
func (srv *Products) ProductsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}")
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
type ProductsUpdateOptions struct {
	AttributeValues interface{}
	Completeness interface{}
	DeletedAt string
	Enabled bool
	FamilyId string
	FamilyVariantId string
	Kind string
	ParentId string
	QuantifiedAssociations interface{}
	Sku string
	TaxClass string
	enabledSetters map[string]bool
}
func (options ProductsUpdateOptions) New() *ProductsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"Completeness": false,
		"DeletedAt": false,
		"Enabled": false,
		"FamilyId": false,
		"FamilyVariantId": false,
		"Kind": false,
		"ParentId": false,
		"QuantifiedAssociations": false,
		"Sku": false,
		"TaxClass": false,
	}
	return &options
}
type ProductsUpdateOption func(*ProductsUpdateOptions)
func (srv *Products) WithProductsUpdateAttributeValues(v interface{}) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsUpdateCompleteness(v interface{}) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.Completeness = v
		o.enabledSetters["Completeness"] = true
	}
}
func (srv *Products) WithProductsUpdateDeletedAt(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.DeletedAt = v
		o.enabledSetters["DeletedAt"] = true
	}
}
func (srv *Products) WithProductsUpdateEnabled(v bool) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Products) WithProductsUpdateFamilyId(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *Products) WithProductsUpdateFamilyVariantId(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.FamilyVariantId = v
		o.enabledSetters["FamilyVariantId"] = true
	}
}
func (srv *Products) WithProductsUpdateKind(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Products) WithProductsUpdateParentId(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *Products) WithProductsUpdateQuantifiedAssociations(v interface{}) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.QuantifiedAssociations = v
		o.enabledSetters["QuantifiedAssociations"] = true
	}
}
func (srv *Products) WithProductsUpdateSku(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Products) WithProductsUpdateTaxClass(v string) ProductsUpdateOption {
	return func(o *ProductsUpdateOptions) {
		o.TaxClass = v
		o.enabledSetters["TaxClass"] = true
	}
}
			
// ProductsUpdate updates one product by id. A partial patch: the body names
// only the columns to change and every column it leaves out keeps its current
// value, so there is no read-modify-write and no way to blank a field by
// forgetting it.
// 
// The catalog itself. A product row carries only what every product has —
// SKU, kind, family, enabled, tax class — and everything the tenant
// modelled lives in the `attribute_values` jsonb document, keyed by attribute
// CODE inside one of four scope buckets (common, per locale, per channel, per
// channel and locale). `label` is a generated column, maintained by the
// database so a grid of twenty thousand rows can sort and filter on a name
// with no join. `kind` says where the row sits in the variant hierarchy: a
// `model` carries what its variants share and is never sold itself.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `sku` answers 409. `label` is a generated column: naming it is dropped
// rather than refused, and `completeness` is written by the two metadata
// routes, not here.
func (srv *Products) ProductsUpdate(Id string, optionalSetters ...ProductsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}")
	options := ProductsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["Completeness"] {
		params["completeness"] = options.Completeness
	}
	if options.enabledSetters["DeletedAt"] {
		params["deleted_at"] = options.DeletedAt
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["FamilyVariantId"] {
		params["family_variant_id"] = options.FamilyVariantId
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["QuantifiedAssociations"] {
		params["quantified_associations"] = options.QuantifiedAssociations
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
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
			
// ProductsCompleteness how much of what its family REQUIRES does this product
// actually carry — the number a merchandiser works down.
// products.completeness is jsonb that nothing had ever written. This computes
// it from family_attributes (is_required) against the product's own scoped
// attribute_values and stores the result. A product with no family answers
// 400 rather than an invented 0 % — it has nothing to be measured against.
func (srv *Products) ProductsCompleteness(Id string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}/completeness")
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
type ProductsFamilyAssignOptions struct {
	FamilyCode string
	FamilyId string
	enabledSetters map[string]bool
}
func (options ProductsFamilyAssignOptions) New() *ProductsFamilyAssignOptions {
	options.enabledSetters = map[string]bool{
		"FamilyCode": false,
		"FamilyId": false,
	}
	return &options
}
type ProductsFamilyAssignOption func(*ProductsFamilyAssignOptions)
func (srv *Products) WithProductsFamilyAssignFamilyCode(v string) ProductsFamilyAssignOption {
	return func(o *ProductsFamilyAssignOptions) {
		o.FamilyCode = v
		o.enabledSetters["FamilyCode"] = true
	}
}
func (srv *Products) WithProductsFamilyAssignFamilyId(v string) ProductsFamilyAssignOption {
	return func(o *ProductsFamilyAssignOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
			
// ProductsFamilyAssign names the family in the body — by `family_id` or by
// `family_code`, whichever the caller holds — and computes the product's
// completeness in the same call. The step every family-driven surface waits
// on: a product with no family has no required attributes, so its
// completeness cannot be computed and its family's label attribute never
// resolves. Assigning the family recomputes and STORES products.completeness
// immediately, so the metadata cannot go stale between the two operations.
func (srv *Products) ProductsFamilyAssign(Id string, optionalSetters ...ProductsFamilyAssignOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}/family")
	options := ProductsFamilyAssignOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["FamilyCode"] {
		params["family_code"] = options.FamilyCode
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
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
