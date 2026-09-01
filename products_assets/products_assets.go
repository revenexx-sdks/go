package products_assets

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ProductsAssets service
type ProductsAssets struct {
	client client.Client
}

func New(clt client.Client) *ProductsAssets {
	return &ProductsAssets{
		client: clt,
	}
}

type ProductsAssetsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	AssetFamilyId string
	Code string
	Source string
	StorageAssetId string
	DeliveryPath string
	ExternalUrl string
	AttributeValues string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAssetsListOptions) New() *ProductsAssetsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"AssetFamilyId": false,
		"Code": false,
		"Source": false,
		"StorageAssetId": false,
		"DeliveryPath": false,
		"ExternalUrl": false,
		"AttributeValues": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsAssetsListOption func(*ProductsAssetsListOptions)
func (srv *ProductsAssets) WithProductsAssetsListLimit(v int) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListOffset(v int) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListOrder(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListId(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListAssetFamilyId(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.AssetFamilyId = v
		o.enabledSetters["AssetFamilyId"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListCode(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListSource(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListStorageAssetId(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.StorageAssetId = v
		o.enabledSetters["StorageAssetId"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListDeliveryPath(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.DeliveryPath = v
		o.enabledSetters["DeliveryPath"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListExternalUrl(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.ExternalUrl = v
		o.enabledSetters["ExternalUrl"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListAttributeValues(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListCreatedAt(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsListUpdatedAt(v string) ProductsAssetsListOption {
	return func(o *ProductsAssetsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsAssetsList one piece of media in the decoupled asset domain. The
// bytes live either in this platform's object store (`source: "storage"` with
// a `storage_asset_id` that survives a rename) or on somebody else's host
// (`source: "external"` with an `external_url`), and the database enforces
// the pair so neither half can be stored alone. A product points at an asset
// by its code through a media attribute; there is no product-to-asset link
// table in this app.
// 
// Every column of `assets` is an exact-match query parameter, `order` sorts
// by one column, and `limit`/`offset` page through `page.total`. A query key
// that is NOT a column is dropped rather than refused, and the `filter`
// object echoes the ones that were understood — that echo is the only way
// to tell an unfiltered answer from an empty one. It reads rows exactly as
// they are stored: no join is resolved, no jsonb value is unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsAssets) ProductsAssetsList(optionalSetters ...ProductsAssetsListOption)(*interface{}, error) {
	path := "/v1/products/assets"
	options := ProductsAssetsListOptions{}.New()
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
	if options.enabledSetters["AssetFamilyId"] {
		params["asset_family_id"] = options.AssetFamilyId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
	}
	if options.enabledSetters["StorageAssetId"] {
		params["storage_asset_id"] = options.StorageAssetId
	}
	if options.enabledSetters["DeliveryPath"] {
		params["delivery_path"] = options.DeliveryPath
	}
	if options.enabledSetters["ExternalUrl"] {
		params["external_url"] = options.ExternalUrl
	}
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
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
type ProductsAssetsCreateOptions struct {
	AttributeValues interface{}
	DeliveryPath string
	ExternalUrl string
	Source string
	StorageAssetId string
	enabledSetters map[string]bool
}
func (options ProductsAssetsCreateOptions) New() *ProductsAssetsCreateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"DeliveryPath": false,
		"ExternalUrl": false,
		"Source": false,
		"StorageAssetId": false,
	}
	return &options
}
type ProductsAssetsCreateOption func(*ProductsAssetsCreateOptions)
func (srv *ProductsAssets) WithProductsAssetsCreateAttributeValues(v interface{}) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsCreateDeliveryPath(v string) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.DeliveryPath = v
		o.enabledSetters["DeliveryPath"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsCreateExternalUrl(v string) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.ExternalUrl = v
		o.enabledSetters["ExternalUrl"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsCreateSource(v string) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsCreateStorageAssetId(v string) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.StorageAssetId = v
		o.enabledSetters["StorageAssetId"] = true
	}
}
					
// ProductsAssetsCreate creates one asset and answers 201 with the stored row,
// including the id and the timestamps the database filled in — a client
// never sends an id, it reads one back and uses it in the path of every later
// call.
// 
// One piece of media in the decoupled asset domain. The bytes live either in
// this platform's object store (`source: "storage"` with a `storage_asset_id`
// that survives a rename) or on somebody else's host (`source: "external"`
// with an `external_url`), and the database enforces the pair so neither half
// can be stored alone. A product points at an asset by its code through a
// media attribute; there is no product-to-asset link table in this app.
// 
// `asset_family_id` and `code` are the only columns the database refuses the
// row without; everything else has a default or is nullable. A second row
// with the same `asset_family_id` and `code` answers 409. This app owns the
// create, because it is the only place an external URL can enter the catalog:
// an asset with no family falls back to the `default_asset_family` setting,
// and an `external` one is refused unless the tenant allows external media
// and the URL's host is on its allow-list.
func (srv *ProductsAssets) ProductsAssetsCreate(AssetFamilyId string, Code string, optionalSetters ...ProductsAssetsCreateOption)(*models.Error, error) {
	path := "/v1/products/assets"
	options := ProductsAssetsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["asset_family_id"] = AssetFamilyId
	params["code"] = Code
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["DeliveryPath"] {
		params["delivery_path"] = options.DeliveryPath
	}
	if options.enabledSetters["ExternalUrl"] {
		params["external_url"] = options.ExternalUrl
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
	}
	if options.enabledSetters["StorageAssetId"] {
		params["storage_asset_id"] = options.StorageAssetId
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
	
// ProductsAssetsDelete deletes one asset by id. It is a hard delete — the
// row is gone, and the answer is a confirmation rather than a result to
// branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no asset of this tenant carries answers 404; there is no 409, because
// every foreign key pointing at this entity resolves itself on delete rather
// than blocking one.
func (srv *ProductsAssets) ProductsAssetsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/assets/{id}")
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
	
// ProductsAssetsGet reads one asset by its id — the whole row, every
// column, as it is stored.
// 
// One piece of media in the decoupled asset domain. The bytes live either in
// this platform's object store (`source: "storage"` with a `storage_asset_id`
// that survives a rename) or on somebody else's host (`source: "external"`
// with an `external_url`), and the database enforces the pair so neither half
// can be stored alone. A product points at an asset by its code through a
// media attribute; there is no product-to-asset link table in this app.
// 
// An id no asset of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsAssets) ProductsAssetsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/assets/{id}")
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
type ProductsAssetsUpdateOptions struct {
	AssetFamilyId string
	AttributeValues interface{}
	Code string
	DeliveryPath string
	ExternalUrl string
	Source string
	StorageAssetId string
	enabledSetters map[string]bool
}
func (options ProductsAssetsUpdateOptions) New() *ProductsAssetsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AssetFamilyId": false,
		"AttributeValues": false,
		"Code": false,
		"DeliveryPath": false,
		"ExternalUrl": false,
		"Source": false,
		"StorageAssetId": false,
	}
	return &options
}
type ProductsAssetsUpdateOption func(*ProductsAssetsUpdateOptions)
func (srv *ProductsAssets) WithProductsAssetsUpdateAssetFamilyId(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.AssetFamilyId = v
		o.enabledSetters["AssetFamilyId"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateAttributeValues(v interface{}) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateCode(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateDeliveryPath(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.DeliveryPath = v
		o.enabledSetters["DeliveryPath"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateExternalUrl(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.ExternalUrl = v
		o.enabledSetters["ExternalUrl"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateSource(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
func (srv *ProductsAssets) WithProductsAssetsUpdateStorageAssetId(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.StorageAssetId = v
		o.enabledSetters["StorageAssetId"] = true
	}
}
			
// ProductsAssetsUpdate updates one asset by id. A partial patch: the body
// names only the columns to change and every column it leaves out keeps its
// current value, so there is no read-modify-write and no way to blank a field
// by forgetting it.
// 
// One piece of media in the decoupled asset domain. The bytes live either in
// this platform's object store (`source: "storage"` with a `storage_asset_id`
// that survives a rename) or on somebody else's host (`source: "external"`
// with an `external_url`), and the database enforces the pair so neither half
// can be stored alone. A product points at an asset by its code through a
// media attribute; there is no product-to-asset link table in this app.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `asset_family_id` and `code` answers 409.
func (srv *ProductsAssets) ProductsAssetsUpdate(Id string, optionalSetters ...ProductsAssetsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/assets/{id}")
	options := ProductsAssetsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AssetFamilyId"] {
		params["asset_family_id"] = options.AssetFamilyId
	}
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["DeliveryPath"] {
		params["delivery_path"] = options.DeliveryPath
	}
	if options.enabledSetters["ExternalUrl"] {
		params["external_url"] = options.ExternalUrl
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
	}
	if options.enabledSetters["StorageAssetId"] {
		params["storage_asset_id"] = options.StorageAssetId
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
