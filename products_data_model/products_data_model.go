package products_data_model

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ProductsDataModel service
type ProductsDataModel struct {
	client client.Client
}

func New(clt client.Client) *ProductsDataModel {
	return &ProductsDataModel{
		client: clt,
	}
}

type ProductsAssetFamiliesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	Labels string
	NamingConvention string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAssetFamiliesListOptions) New() *ProductsAssetFamiliesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"Labels": false,
		"NamingConvention": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsAssetFamiliesListOption func(*ProductsAssetFamiliesListOptions)
func (srv *ProductsDataModel) WithProductsAssetFamiliesListLimit(v int) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListOffset(v int) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListOrder(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListId(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListCode(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListLabels(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListNamingConvention(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.NamingConvention = v
		o.enabledSetters["NamingConvention"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListCreatedAt(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesListUpdatedAt(v string) ProductsAssetFamiliesListOption {
	return func(o *ProductsAssetFamiliesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsAssetFamiliesList a class of media with one shared shape —
// packshots, datasheets, line drawings. The family decides which attributes
// an asset of it carries (alt text, copyright, an expiry date) and, through
// `naming_convention`, how a file of it is named — which is what lets an
// import bind a file to a product with no mapping table.
// 
// Every column of `asset_families` is an exact-match query parameter, `order`
// sorts by one column, and `limit`/`offset` page through `page.total`. A
// query key that is NOT a column is dropped rather than refused, and the
// `filter` object echoes the ones that were understood — that echo is the
// only way to tell an unfiltered answer from an empty one. It reads rows
// exactly as they are stored: no join is resolved, no jsonb value is
// unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAssetFamiliesList(optionalSetters ...ProductsAssetFamiliesListOption)(*interface{}, error) {
	path := "/v1/products/asset_families"
	options := ProductsAssetFamiliesListOptions{}.New()
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
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["NamingConvention"] {
		params["naming_convention"] = options.NamingConvention
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
type ProductsAssetFamiliesCreateOptions struct {
	Labels interface{}
	NamingConvention interface{}
	enabledSetters map[string]bool
}
func (options ProductsAssetFamiliesCreateOptions) New() *ProductsAssetFamiliesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"NamingConvention": false,
	}
	return &options
}
type ProductsAssetFamiliesCreateOption func(*ProductsAssetFamiliesCreateOptions)
func (srv *ProductsDataModel) WithProductsAssetFamiliesCreateLabels(v interface{}) ProductsAssetFamiliesCreateOption {
	return func(o *ProductsAssetFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesCreateNamingConvention(v interface{}) ProductsAssetFamiliesCreateOption {
	return func(o *ProductsAssetFamiliesCreateOptions) {
		o.NamingConvention = v
		o.enabledSetters["NamingConvention"] = true
	}
}
			
// ProductsAssetFamiliesCreate creates one asset family and answers 201 with
// the stored row, including the id and the timestamps the database filled in
// — a client never sends an id, it reads one back and uses it in the path
// of every later call.
// 
// A class of media with one shared shape — packshots, datasheets, line
// drawings. The family decides which attributes an asset of it carries (alt
// text, copyright, an expiry date) and, through `naming_convention`, how a
// file of it is named — which is what lets an import bind a file to a
// product with no mapping table.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsDataModel) ProductsAssetFamiliesCreate(Code string, optionalSetters ...ProductsAssetFamiliesCreateOption)(*models.Error, error) {
	path := "/v1/products/asset_families"
	options := ProductsAssetFamiliesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["NamingConvention"] {
		params["naming_convention"] = options.NamingConvention
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
	
// ProductsAssetFamiliesDelete deletes one asset family by id. It is a hard
// delete — the row is gone, and the answer is a confirmation rather than a
// result to branch on.
// 
// It takes what hangs off it: assets (`asset_family_id`) are deleted with it.
// 
// An id no asset family of this tenant carries answers 404; there is no 409,
// because every foreign key pointing at this entity resolves itself on delete
// rather than blocking one.
func (srv *ProductsDataModel) ProductsAssetFamiliesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/asset_families/{id}")
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
	
// ProductsAssetFamiliesGet reads one asset family by its id — the whole
// row, every column, as it is stored.
// 
// A class of media with one shared shape — packshots, datasheets, line
// drawings. The family decides which attributes an asset of it carries (alt
// text, copyright, an expiry date) and, through `naming_convention`, how a
// file of it is named — which is what lets an import bind a file to a
// product with no mapping table.
// 
// An id no asset family of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAssetFamiliesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/asset_families/{id}")
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
type ProductsAssetFamiliesUpdateOptions struct {
	Code string
	Labels interface{}
	NamingConvention interface{}
	enabledSetters map[string]bool
}
func (options ProductsAssetFamiliesUpdateOptions) New() *ProductsAssetFamiliesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"NamingConvention": false,
	}
	return &options
}
type ProductsAssetFamiliesUpdateOption func(*ProductsAssetFamiliesUpdateOptions)
func (srv *ProductsDataModel) WithProductsAssetFamiliesUpdateCode(v string) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesUpdateLabels(v interface{}) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssetFamiliesUpdateNamingConvention(v interface{}) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.NamingConvention = v
		o.enabledSetters["NamingConvention"] = true
	}
}
			
// ProductsAssetFamiliesUpdate updates one asset family by id. A partial
// patch: the body names only the columns to change and every column it leaves
// out keeps its current value, so there is no read-modify-write and no way to
// blank a field by forgetting it.
// 
// A class of media with one shared shape — packshots, datasheets, line
// drawings. The family decides which attributes an asset of it carries (alt
// text, copyright, an expiry date) and, through `naming_convention`, how a
// file of it is named — which is what lets an import bind a file to a
// product with no mapping table.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsAssetFamiliesUpdate(Id string, optionalSetters ...ProductsAssetFamiliesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/asset_families/{id}")
	options := ProductsAssetFamiliesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["NamingConvention"] {
		params["naming_convention"] = options.NamingConvention
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
type ProductsAssociationTypesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	IsTwoWay bool
	IsQuantified bool
	Labels string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAssociationTypesListOptions) New() *ProductsAssociationTypesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"IsTwoWay": false,
		"IsQuantified": false,
		"Labels": false,
		"CreatedAt": false,
	}
	return &options
}
type ProductsAssociationTypesListOption func(*ProductsAssociationTypesListOptions)
func (srv *ProductsDataModel) WithProductsAssociationTypesListLimit(v int) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListOffset(v int) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListOrder(v string) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListId(v string) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListCode(v string) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListIsTwoWay(v bool) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.IsTwoWay = v
		o.enabledSetters["IsTwoWay"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListIsQuantified(v bool) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.IsQuantified = v
		o.enabledSetters["IsQuantified"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListLabels(v string) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesListCreatedAt(v string) ProductsAssociationTypesListOption {
	return func(o *ProductsAssociationTypesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// ProductsAssociationTypesList the KIND of relation two products can have —
// cross-sell, accessory, spare part, bill of materials. `is_two_way` declares
// the relation symmetric and `is_quantified` declares that it carries a
// quantity; both are declarations a client READS rather than behaviour this
// app performs — it stores one row per direction and never creates the
// mirror for you.
// 
// Every column of `association_types` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsAssociationTypesList(optionalSetters ...ProductsAssociationTypesListOption)(*interface{}, error) {
	path := "/v1/products/association_types"
	options := ProductsAssociationTypesListOptions{}.New()
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
	if options.enabledSetters["IsTwoWay"] {
		params["is_two_way"] = options.IsTwoWay
	}
	if options.enabledSetters["IsQuantified"] {
		params["is_quantified"] = options.IsQuantified
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsAssociationTypesCreateOptions struct {
	IsQuantified bool
	IsTwoWay bool
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsAssociationTypesCreateOptions) New() *ProductsAssociationTypesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsQuantified": false,
		"IsTwoWay": false,
		"Labels": false,
	}
	return &options
}
type ProductsAssociationTypesCreateOption func(*ProductsAssociationTypesCreateOptions)
func (srv *ProductsDataModel) WithProductsAssociationTypesCreateIsQuantified(v bool) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.IsQuantified = v
		o.enabledSetters["IsQuantified"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesCreateIsTwoWay(v bool) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.IsTwoWay = v
		o.enabledSetters["IsTwoWay"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesCreateLabels(v interface{}) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsAssociationTypesCreate creates one association type and answers 201
// with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// The KIND of relation two products can have — cross-sell, accessory, spare
// part, bill of materials. `is_two_way` declares the relation symmetric and
// `is_quantified` declares that it carries a quantity; both are declarations
// a client READS rather than behaviour this app performs — it stores one
// row per direction and never creates the mirror for you.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsDataModel) ProductsAssociationTypesCreate(Code string, optionalSetters ...ProductsAssociationTypesCreateOption)(*models.Error, error) {
	path := "/v1/products/association_types"
	options := ProductsAssociationTypesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["IsQuantified"] {
		params["is_quantified"] = options.IsQuantified
	}
	if options.enabledSetters["IsTwoWay"] {
		params["is_two_way"] = options.IsTwoWay
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
	
// ProductsAssociationTypesDelete deletes one association type by id. It is a
// hard delete — the row is gone, and the answer is a confirmation rather
// than a result to branch on.
// 
// It takes what hangs off it: product associations (`association_type_id`)
// are deleted with it.
// 
// An id no association type of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsAssociationTypesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/association_types/{id}")
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
	
// ProductsAssociationTypesGet reads one association type by its id — the
// whole row, every column, as it is stored.
// 
// The KIND of relation two products can have — cross-sell, accessory, spare
// part, bill of materials. `is_two_way` declares the relation symmetric and
// `is_quantified` declares that it carries a quantity; both are declarations
// a client READS rather than behaviour this app performs — it stores one
// row per direction and never creates the mirror for you.
// 
// An id no association type of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAssociationTypesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/association_types/{id}")
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
type ProductsAssociationTypesUpdateOptions struct {
	Code string
	IsQuantified bool
	IsTwoWay bool
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsAssociationTypesUpdateOptions) New() *ProductsAssociationTypesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsQuantified": false,
		"IsTwoWay": false,
		"Labels": false,
	}
	return &options
}
type ProductsAssociationTypesUpdateOption func(*ProductsAssociationTypesUpdateOptions)
func (srv *ProductsDataModel) WithProductsAssociationTypesUpdateCode(v string) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesUpdateIsQuantified(v bool) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.IsQuantified = v
		o.enabledSetters["IsQuantified"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesUpdateIsTwoWay(v bool) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.IsTwoWay = v
		o.enabledSetters["IsTwoWay"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAssociationTypesUpdateLabels(v interface{}) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsAssociationTypesUpdate updates one association type by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// The KIND of relation two products can have — cross-sell, accessory, spare
// part, bill of materials. `is_two_way` declares the relation symmetric and
// `is_quantified` declares that it carries a quantity; both are declarations
// a client READS rather than behaviour this app performs — it stores one
// row per direction and never creates the mirror for you.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsAssociationTypesUpdate(Id string, optionalSetters ...ProductsAssociationTypesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/association_types/{id}")
	options := ProductsAssociationTypesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["IsQuantified"] {
		params["is_quantified"] = options.IsQuantified
	}
	if options.enabledSetters["IsTwoWay"] {
		params["is_two_way"] = options.IsTwoWay
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsAttributeSchemaOptions struct {
	FamilyId string
	FamilyCode string
	EntityType string
	EntityRef string
	Locale string
	Channel string
	Kind string
	enabledSetters map[string]bool
}
func (options ProductsAttributeSchemaOptions) New() *ProductsAttributeSchemaOptions {
	options.enabledSetters = map[string]bool{
		"FamilyId": false,
		"FamilyCode": false,
		"EntityType": false,
		"EntityRef": false,
		"Locale": false,
		"Channel": false,
		"Kind": false,
	}
	return &options
}
type ProductsAttributeSchemaOption func(*ProductsAttributeSchemaOptions)
func (srv *ProductsDataModel) WithProductsAttributeSchemaFamilyId(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaFamilyCode(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.FamilyCode = v
		o.enabledSetters["FamilyCode"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaEntityType(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaEntityRef(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaLocale(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaChannel(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeSchemaKind(v string) ProductsAttributeSchemaOption {
	return func(o *ProductsAttributeSchemaOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
	
// ProductsAttributeSchema which fields does this family have — one
// ready-to-render list, not six joined tables. The catalog's SHAPE is tenant
// data: a product's properties are rows in `attributes`, grouped by
// `attribute_groups`, selected per family by `family_attributes`, with their
// permitted values in `attribute_options` and their variant axes in
// `family_variants`. Reading that shape used to mean five reads, a join, and
// a private `attributes.type` → input mapping in every client — and that
// mapping is the part that must live here, because the type list carries no
// CHECK by design and an integrator extends it. Answers one field list
// instead, ordered by group then by the family's own ordering. Without a
// family it answers every attribute declared for `entity_type`/`entity_ref`
// — the shape of a reference entity's records or an asset family, which
// have attributes but no family. Writes nothing.
func (srv *ProductsDataModel) ProductsAttributeSchema(optionalSetters ...ProductsAttributeSchemaOption)(*models.Error, error) {
	path := "/v1/products/attribute-schema"
	options := ProductsAttributeSchemaOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["FamilyCode"] {
		params["family_code"] = options.FamilyCode
	}
	if options.enabledSetters["EntityType"] {
		params["entity_type"] = options.EntityType
	}
	if options.enabledSetters["EntityRef"] {
		params["entity_ref"] = options.EntityRef
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
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
type ProductsAttributeGroupsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	Position int
	Labels string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAttributeGroupsListOptions) New() *ProductsAttributeGroupsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"Position": false,
		"Labels": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsAttributeGroupsListOption func(*ProductsAttributeGroupsListOptions)
func (srv *ProductsDataModel) WithProductsAttributeGroupsListLimit(v int) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListOffset(v int) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListOrder(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListId(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListCode(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListPosition(v int) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListLabels(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListCreatedAt(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsListUpdatedAt(v string) ProductsAttributeGroupsListOption {
	return func(o *ProductsAttributeGroupsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsAttributeGroupsList an attribute group is a SECTION of a product
// form — "Technical attributes", "Logistics" — and the thing every
// attribute is filed under. It carries a `position`, which is the order the
// sections appear in, and per-language `labels`, which is what an operator
// reads; the `code` is what an attribute joins on and is never shown. `GET
// /products/attribute-schema` already resolves a group's heading onto every
// field it returns, so these routes are for MANAGING the sections, not for
// rendering a form.
// 
// Every column of `attribute_groups` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsAttributeGroupsList(optionalSetters ...ProductsAttributeGroupsListOption)(*interface{}, error) {
	path := "/v1/products/attribute_groups"
	options := ProductsAttributeGroupsListOptions{}.New()
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
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsAttributeGroupsCreateOptions struct {
	Labels interface{}
	Position int
	enabledSetters map[string]bool
}
func (options ProductsAttributeGroupsCreateOptions) New() *ProductsAttributeGroupsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"Position": false,
	}
	return &options
}
type ProductsAttributeGroupsCreateOption func(*ProductsAttributeGroupsCreateOptions)
func (srv *ProductsDataModel) WithProductsAttributeGroupsCreateLabels(v interface{}) ProductsAttributeGroupsCreateOption {
	return func(o *ProductsAttributeGroupsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsCreatePosition(v int) ProductsAttributeGroupsCreateOption {
	return func(o *ProductsAttributeGroupsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
			
// ProductsAttributeGroupsCreate creates one attribute group and answers 201
// with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// An attribute group is a SECTION of a product form — "Technical
// attributes", "Logistics" — and the thing every attribute is filed under.
// It carries a `position`, which is the order the sections appear in, and
// per-language `labels`, which is what an operator reads; the `code` is what
// an attribute joins on and is never shown. `GET /products/attribute-schema`
// already resolves a group's heading onto every field it returns, so these
// routes are for MANAGING the sections, not for rendering a form.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsDataModel) ProductsAttributeGroupsCreate(Code string, optionalSetters ...ProductsAttributeGroupsCreateOption)(*models.Error, error) {
	path := "/v1/products/attribute_groups"
	options := ProductsAttributeGroupsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
	
// ProductsAttributeGroupsDelete deletes one attribute group by id. It is a
// hard delete — the row is gone, and the answer is a confirmation rather
// than a result to branch on.
// 
// `attributes.group_id` is set to null instead, so the rows that pointed at
// it survive the delete rather than going with it.
// 
// An id no attribute group of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsAttributeGroupsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_groups/{id}")
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
	
// ProductsAttributeGroupsGet reads one attribute group by its id — the
// whole row, every column, as it is stored.
// 
// An attribute group is a SECTION of a product form — "Technical
// attributes", "Logistics" — and the thing every attribute is filed under.
// It carries a `position`, which is the order the sections appear in, and
// per-language `labels`, which is what an operator reads; the `code` is what
// an attribute joins on and is never shown. `GET /products/attribute-schema`
// already resolves a group's heading onto every field it returns, so these
// routes are for MANAGING the sections, not for rendering a form.
// 
// An id no attribute group of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAttributeGroupsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_groups/{id}")
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
type ProductsAttributeGroupsUpdateOptions struct {
	Code string
	Labels interface{}
	Position int
	enabledSetters map[string]bool
}
func (options ProductsAttributeGroupsUpdateOptions) New() *ProductsAttributeGroupsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"Position": false,
	}
	return &options
}
type ProductsAttributeGroupsUpdateOption func(*ProductsAttributeGroupsUpdateOptions)
func (srv *ProductsDataModel) WithProductsAttributeGroupsUpdateCode(v string) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsUpdateLabels(v interface{}) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeGroupsUpdatePosition(v int) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
			
// ProductsAttributeGroupsUpdate updates one attribute group by id. A partial
// patch: the body names only the columns to change and every column it leaves
// out keeps its current value, so there is no read-modify-write and no way to
// blank a field by forgetting it.
// 
// An attribute group is a SECTION of a product form — "Technical
// attributes", "Logistics" — and the thing every attribute is filed under.
// It carries a `position`, which is the order the sections appear in, and
// per-language `labels`, which is what an operator reads; the `code` is what
// an attribute joins on and is never shown. `GET /products/attribute-schema`
// already resolves a group's heading onto every field it returns, so these
// routes are for MANAGING the sections, not for rendering a form.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsAttributeGroupsUpdate(Id string, optionalSetters ...ProductsAttributeGroupsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_groups/{id}")
	options := ProductsAttributeGroupsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type ProductsAttributeOptionsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	AttributeId string
	Code string
	Position int
	Swatch string
	Labels string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAttributeOptionsListOptions) New() *ProductsAttributeOptionsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"AttributeId": false,
		"Code": false,
		"Position": false,
		"Swatch": false,
		"Labels": false,
		"CreatedAt": false,
	}
	return &options
}
type ProductsAttributeOptionsListOption func(*ProductsAttributeOptionsListOptions)
func (srv *ProductsDataModel) WithProductsAttributeOptionsListLimit(v int) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListOffset(v int) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListOrder(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListId(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListAttributeId(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListCode(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListPosition(v int) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListSwatch(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Swatch = v
		o.enabledSetters["Swatch"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListLabels(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsListCreatedAt(v string) ProductsAttributeOptionsListOption {
	return func(o *ProductsAttributeOptionsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// ProductsAttributeOptionsList the permitted values of one select or
// multi-select attribute. A record stores the option's CODE and never its
// label, so renaming an option in every language leaves every product that
// picked it untouched, and `position` is the order it appears in the
// dropdown. `GET /products/attribute-schema` republishes these as a field's
// `options`, already resolved for a locale.
// 
// Every column of `attribute_options` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsAttributeOptionsList(optionalSetters ...ProductsAttributeOptionsListOption)(*interface{}, error) {
	path := "/v1/products/attribute_options"
	options := ProductsAttributeOptionsListOptions{}.New()
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
	if options.enabledSetters["AttributeId"] {
		params["attribute_id"] = options.AttributeId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Swatch"] {
		params["swatch"] = options.Swatch
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsAttributeOptionsCreateOptions struct {
	Labels interface{}
	Position int
	Swatch interface{}
	enabledSetters map[string]bool
}
func (options ProductsAttributeOptionsCreateOptions) New() *ProductsAttributeOptionsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"Position": false,
		"Swatch": false,
	}
	return &options
}
type ProductsAttributeOptionsCreateOption func(*ProductsAttributeOptionsCreateOptions)
func (srv *ProductsDataModel) WithProductsAttributeOptionsCreateLabels(v interface{}) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsCreatePosition(v int) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsCreateSwatch(v interface{}) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Swatch = v
		o.enabledSetters["Swatch"] = true
	}
}
					
// ProductsAttributeOptionsCreate creates one attribute option and answers 201
// with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// The permitted values of one select or multi-select attribute. A record
// stores the option's CODE and never its label, so renaming an option in
// every language leaves every product that picked it untouched, and
// `position` is the order it appears in the dropdown. `GET
// /products/attribute-schema` republishes these as a field's `options`,
// already resolved for a locale.
// 
// `attribute_id` and `code` are the only columns the database refuses the row
// without; everything else has a default or is nullable. A second row with
// the same `attribute_id` and `code` answers 409.
func (srv *ProductsDataModel) ProductsAttributeOptionsCreate(AttributeId string, Code string, optionalSetters ...ProductsAttributeOptionsCreateOption)(*models.Error, error) {
	path := "/v1/products/attribute_options"
	options := ProductsAttributeOptionsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["attribute_id"] = AttributeId
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Swatch"] {
		params["swatch"] = options.Swatch
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
	
// ProductsAttributeOptionsDelete deletes one attribute option by id. It is a
// hard delete — the row is gone, and the answer is a confirmation rather
// than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no attribute option of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsAttributeOptionsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_options/{id}")
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
	
// ProductsAttributeOptionsGet reads one attribute option by its id — the
// whole row, every column, as it is stored.
// 
// The permitted values of one select or multi-select attribute. A record
// stores the option's CODE and never its label, so renaming an option in
// every language leaves every product that picked it untouched, and
// `position` is the order it appears in the dropdown. `GET
// /products/attribute-schema` republishes these as a field's `options`,
// already resolved for a locale.
// 
// An id no attribute option of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAttributeOptionsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_options/{id}")
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
type ProductsAttributeOptionsUpdateOptions struct {
	AttributeId string
	Code string
	Labels interface{}
	Position int
	Swatch interface{}
	enabledSetters map[string]bool
}
func (options ProductsAttributeOptionsUpdateOptions) New() *ProductsAttributeOptionsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeId": false,
		"Code": false,
		"Labels": false,
		"Position": false,
		"Swatch": false,
	}
	return &options
}
type ProductsAttributeOptionsUpdateOption func(*ProductsAttributeOptionsUpdateOptions)
func (srv *ProductsDataModel) WithProductsAttributeOptionsUpdateAttributeId(v string) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsUpdateCode(v string) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsUpdateLabels(v interface{}) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsUpdatePosition(v int) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributeOptionsUpdateSwatch(v interface{}) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Swatch = v
		o.enabledSetters["Swatch"] = true
	}
}
			
// ProductsAttributeOptionsUpdate updates one attribute option by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// The permitted values of one select or multi-select attribute. A record
// stores the option's CODE and never its label, so renaming an option in
// every language leaves every product that picked it untouched, and
// `position` is the order it appears in the dropdown. `GET
// /products/attribute-schema` republishes these as a field's `options`,
// already resolved for a locale.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `attribute_id` and `code` answers 409.
func (srv *ProductsDataModel) ProductsAttributeOptionsUpdate(Id string, optionalSetters ...ProductsAttributeOptionsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attribute_options/{id}")
	options := ProductsAttributeOptionsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AttributeId"] {
		params["attribute_id"] = options.AttributeId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Swatch"] {
		params["swatch"] = options.Swatch
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
type ProductsAttributesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	EntityType string
	EntityRef string
	Type string
	GroupId string
	Localizable bool
	Scopable bool
	IsUnique bool
	IsFilterable bool
	UsableInGrid bool
	Validation string
	Config string
	Labels string
	Position int
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsAttributesListOptions) New() *ProductsAttributesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"EntityType": false,
		"EntityRef": false,
		"Type": false,
		"GroupId": false,
		"Localizable": false,
		"Scopable": false,
		"IsUnique": false,
		"IsFilterable": false,
		"UsableInGrid": false,
		"Validation": false,
		"Config": false,
		"Labels": false,
		"Position": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsAttributesListOption func(*ProductsAttributesListOptions)
func (srv *ProductsDataModel) WithProductsAttributesListLimit(v int) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListOffset(v int) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListOrder(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListId(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListCode(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListEntityType(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListEntityRef(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListType(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListGroupId(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.GroupId = v
		o.enabledSetters["GroupId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListLocalizable(v bool) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Localizable = v
		o.enabledSetters["Localizable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListScopable(v bool) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Scopable = v
		o.enabledSetters["Scopable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListIsUnique(v bool) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.IsUnique = v
		o.enabledSetters["IsUnique"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListIsFilterable(v bool) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.IsFilterable = v
		o.enabledSetters["IsFilterable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListUsableInGrid(v bool) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.UsableInGrid = v
		o.enabledSetters["UsableInGrid"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListValidation(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Validation = v
		o.enabledSetters["Validation"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListConfig(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Config = v
		o.enabledSetters["Config"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListLabels(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListPosition(v int) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListCreatedAt(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesListUpdatedAt(v string) ProductsAttributesListOption {
	return func(o *ProductsAttributesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsAttributesList an attribute is one property a record can carry, and
// in an attribute-driven PIM it is a ROW rather than a column: giving the
// catalog a "net weight" is a create here, not a migration. Its own flags
// decide everything downstream — `localizable` and `scopable` pick which of
// the four `attribute_values` buckets its values are written to, `type` picks
// the editor that renders it, `usable_in_grid` and `is_filterable` are what
// the product grid reads. `entity_type`/`entity_ref` say which kind of record
// carries it: a product, one reference entity's records, one asset family, or
// a category.
// 
// Every column of `attributes` is an exact-match query parameter, `order`
// sorts by one column, and `limit`/`offset` page through `page.total`. A
// query key that is NOT a column is dropped rather than refused, and the
// `filter` object echoes the ones that were understood — that echo is the
// only way to tell an unfiltered answer from an empty one. It reads rows
// exactly as they are stored: no join is resolved, no jsonb value is
// unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAttributesList(optionalSetters ...ProductsAttributesListOption)(*interface{}, error) {
	path := "/v1/products/attributes"
	options := ProductsAttributesListOptions{}.New()
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
	if options.enabledSetters["EntityType"] {
		params["entity_type"] = options.EntityType
	}
	if options.enabledSetters["EntityRef"] {
		params["entity_ref"] = options.EntityRef
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["GroupId"] {
		params["group_id"] = options.GroupId
	}
	if options.enabledSetters["Localizable"] {
		params["localizable"] = options.Localizable
	}
	if options.enabledSetters["Scopable"] {
		params["scopable"] = options.Scopable
	}
	if options.enabledSetters["IsUnique"] {
		params["is_unique"] = options.IsUnique
	}
	if options.enabledSetters["IsFilterable"] {
		params["is_filterable"] = options.IsFilterable
	}
	if options.enabledSetters["UsableInGrid"] {
		params["usable_in_grid"] = options.UsableInGrid
	}
	if options.enabledSetters["Validation"] {
		params["validation"] = options.Validation
	}
	if options.enabledSetters["Config"] {
		params["config"] = options.Config
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsAttributesCreateOptions struct {
	Config interface{}
	EntityRef string
	EntityType string
	GroupId string
	IsFilterable bool
	IsUnique bool
	Labels interface{}
	Localizable bool
	Position int
	Scopable bool
	UsableInGrid bool
	Validation interface{}
	enabledSetters map[string]bool
}
func (options ProductsAttributesCreateOptions) New() *ProductsAttributesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Config": false,
		"EntityRef": false,
		"EntityType": false,
		"GroupId": false,
		"IsFilterable": false,
		"IsUnique": false,
		"Labels": false,
		"Localizable": false,
		"Position": false,
		"Scopable": false,
		"UsableInGrid": false,
		"Validation": false,
	}
	return &options
}
type ProductsAttributesCreateOption func(*ProductsAttributesCreateOptions)
func (srv *ProductsDataModel) WithProductsAttributesCreateConfig(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Config = v
		o.enabledSetters["Config"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateEntityRef(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateEntityType(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateGroupId(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.GroupId = v
		o.enabledSetters["GroupId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateIsFilterable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.IsFilterable = v
		o.enabledSetters["IsFilterable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateIsUnique(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.IsUnique = v
		o.enabledSetters["IsUnique"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateLabels(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateLocalizable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Localizable = v
		o.enabledSetters["Localizable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreatePosition(v int) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateScopable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Scopable = v
		o.enabledSetters["Scopable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateUsableInGrid(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.UsableInGrid = v
		o.enabledSetters["UsableInGrid"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesCreateValidation(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Validation = v
		o.enabledSetters["Validation"] = true
	}
}
					
// ProductsAttributesCreate creates one attribute and answers 201 with the
// stored row, including the id and the timestamps the database filled in —
// a client never sends an id, it reads one back and uses it in the path of
// every later call.
// 
// An attribute is one property a record can carry, and in an attribute-driven
// PIM it is a ROW rather than a column: giving the catalog a "net weight" is
// a create here, not a migration. Its own flags decide everything downstream
// — `localizable` and `scopable` pick which of the four `attribute_values`
// buckets its values are written to, `type` picks the editor that renders it,
// `usable_in_grid` and `is_filterable` are what the product grid reads.
// `entity_type`/`entity_ref` say which kind of record carries it: a product,
// one reference entity's records, one asset family, or a category.
// 
// `code` and `type` are the only columns the database refuses the row
// without; everything else has a default or is nullable. A second row with
// the same `entity_type`, `entity_ref`, `code` answers 409.
func (srv *ProductsDataModel) ProductsAttributesCreate(Code string, Type string, optionalSetters ...ProductsAttributesCreateOption)(*models.Error, error) {
	path := "/v1/products/attributes"
	options := ProductsAttributesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["type"] = Type
	if options.enabledSetters["Config"] {
		params["config"] = options.Config
	}
	if options.enabledSetters["EntityRef"] {
		params["entity_ref"] = options.EntityRef
	}
	if options.enabledSetters["EntityType"] {
		params["entity_type"] = options.EntityType
	}
	if options.enabledSetters["GroupId"] {
		params["group_id"] = options.GroupId
	}
	if options.enabledSetters["IsFilterable"] {
		params["is_filterable"] = options.IsFilterable
	}
	if options.enabledSetters["IsUnique"] {
		params["is_unique"] = options.IsUnique
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Localizable"] {
		params["localizable"] = options.Localizable
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Scopable"] {
		params["scopable"] = options.Scopable
	}
	if options.enabledSetters["UsableInGrid"] {
		params["usable_in_grid"] = options.UsableInGrid
	}
	if options.enabledSetters["Validation"] {
		params["validation"] = options.Validation
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
	
// ProductsAttributesDelete deletes one attribute by id. It is a hard delete
// — the row is gone, and the answer is a confirmation rather than a result
// to branch on.
// 
// It takes what hangs off it: attribute options (`attribute_id`), family
// attributes (`attribute_id`) are deleted with it.
// 
// An id no attribute of this tenant carries answers 404; there is no 409,
// because every foreign key pointing at this entity resolves itself on delete
// rather than blocking one.
func (srv *ProductsDataModel) ProductsAttributesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attributes/{id}")
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
	
// ProductsAttributesGet reads one attribute by its id — the whole row,
// every column, as it is stored.
// 
// An attribute is one property a record can carry, and in an attribute-driven
// PIM it is a ROW rather than a column: giving the catalog a "net weight" is
// a create here, not a migration. Its own flags decide everything downstream
// — `localizable` and `scopable` pick which of the four `attribute_values`
// buckets its values are written to, `type` picks the editor that renders it,
// `usable_in_grid` and `is_filterable` are what the product grid reads.
// `entity_type`/`entity_ref` say which kind of record carries it: a product,
// one reference entity's records, one asset family, or a category.
// 
// An id no attribute of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsAttributesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attributes/{id}")
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
type ProductsAttributesUpdateOptions struct {
	Code string
	Config interface{}
	EntityRef string
	EntityType string
	GroupId string
	IsFilterable bool
	IsUnique bool
	Labels interface{}
	Localizable bool
	Position int
	Scopable bool
	Type string
	UsableInGrid bool
	Validation interface{}
	enabledSetters map[string]bool
}
func (options ProductsAttributesUpdateOptions) New() *ProductsAttributesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Config": false,
		"EntityRef": false,
		"EntityType": false,
		"GroupId": false,
		"IsFilterable": false,
		"IsUnique": false,
		"Labels": false,
		"Localizable": false,
		"Position": false,
		"Scopable": false,
		"Type": false,
		"UsableInGrid": false,
		"Validation": false,
	}
	return &options
}
type ProductsAttributesUpdateOption func(*ProductsAttributesUpdateOptions)
func (srv *ProductsDataModel) WithProductsAttributesUpdateCode(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateConfig(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Config = v
		o.enabledSetters["Config"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateEntityRef(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateEntityType(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateGroupId(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.GroupId = v
		o.enabledSetters["GroupId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateIsFilterable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.IsFilterable = v
		o.enabledSetters["IsFilterable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateIsUnique(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.IsUnique = v
		o.enabledSetters["IsUnique"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateLabels(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateLocalizable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Localizable = v
		o.enabledSetters["Localizable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdatePosition(v int) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateScopable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Scopable = v
		o.enabledSetters["Scopable"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateType(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateUsableInGrid(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.UsableInGrid = v
		o.enabledSetters["UsableInGrid"] = true
	}
}
func (srv *ProductsDataModel) WithProductsAttributesUpdateValidation(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Validation = v
		o.enabledSetters["Validation"] = true
	}
}
			
// ProductsAttributesUpdate updates one attribute by id. A partial patch: the
// body names only the columns to change and every column it leaves out keeps
// its current value, so there is no read-modify-write and no way to blank a
// field by forgetting it.
// 
// An attribute is one property a record can carry, and in an attribute-driven
// PIM it is a ROW rather than a column: giving the catalog a "net weight" is
// a create here, not a migration. Its own flags decide everything downstream
// — `localizable` and `scopable` pick which of the four `attribute_values`
// buckets its values are written to, `type` picks the editor that renders it,
// `usable_in_grid` and `is_filterable` are what the product grid reads.
// `entity_type`/`entity_ref` say which kind of record carries it: a product,
// one reference entity's records, one asset family, or a category.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `entity_type`, `entity_ref`, `code` answers 409.
func (srv *ProductsDataModel) ProductsAttributesUpdate(Id string, optionalSetters ...ProductsAttributesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/attributes/{id}")
	options := ProductsAttributesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Config"] {
		params["config"] = options.Config
	}
	if options.enabledSetters["EntityRef"] {
		params["entity_ref"] = options.EntityRef
	}
	if options.enabledSetters["EntityType"] {
		params["entity_type"] = options.EntityType
	}
	if options.enabledSetters["GroupId"] {
		params["group_id"] = options.GroupId
	}
	if options.enabledSetters["IsFilterable"] {
		params["is_filterable"] = options.IsFilterable
	}
	if options.enabledSetters["IsUnique"] {
		params["is_unique"] = options.IsUnique
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Localizable"] {
		params["localizable"] = options.Localizable
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Scopable"] {
		params["scopable"] = options.Scopable
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["UsableInGrid"] {
		params["usable_in_grid"] = options.UsableInGrid
	}
	if options.enabledSetters["Validation"] {
		params["validation"] = options.Validation
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
type ProductsFamiliesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	LabelAttribute string
	ImageAttribute string
	Labels string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsFamiliesListOptions) New() *ProductsFamiliesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"LabelAttribute": false,
		"ImageAttribute": false,
		"Labels": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsFamiliesListOption func(*ProductsFamiliesListOptions)
func (srv *ProductsDataModel) WithProductsFamiliesListLimit(v int) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListOffset(v int) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListOrder(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListId(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListCode(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListLabelAttribute(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.LabelAttribute = v
		o.enabledSetters["LabelAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListImageAttribute(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.ImageAttribute = v
		o.enabledSetters["ImageAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListLabels(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListCreatedAt(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesListUpdatedAt(v string) ProductsFamiliesListOption {
	return func(o *ProductsFamiliesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsFamiliesList a family decides WHICH attributes a product has —
// the set is `family_attributes`, and every family-driven surface follows
// from it. It also names which attribute carries the display name
// (`label_attribute`) and which carries the main image. A product with no
// family has no required attributes at all, so its completeness cannot be
// measured and its name never resolves past the SKU; `POST
// /products/{id}/family` is the call that ends that state.
// 
// Every column of `families` is an exact-match query parameter, `order` sorts
// by one column, and `limit`/`offset` page through `page.total`. A query key
// that is NOT a column is dropped rather than refused, and the `filter`
// object echoes the ones that were understood — that echo is the only way
// to tell an unfiltered answer from an empty one. It reads rows exactly as
// they are stored: no join is resolved, no jsonb value is unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsFamiliesList(optionalSetters ...ProductsFamiliesListOption)(*interface{}, error) {
	path := "/v1/products/families"
	options := ProductsFamiliesListOptions{}.New()
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
	if options.enabledSetters["LabelAttribute"] {
		params["label_attribute"] = options.LabelAttribute
	}
	if options.enabledSetters["ImageAttribute"] {
		params["image_attribute"] = options.ImageAttribute
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsFamiliesCreateOptions struct {
	ImageAttribute string
	LabelAttribute string
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamiliesCreateOptions) New() *ProductsFamiliesCreateOptions {
	options.enabledSetters = map[string]bool{
		"ImageAttribute": false,
		"LabelAttribute": false,
		"Labels": false,
	}
	return &options
}
type ProductsFamiliesCreateOption func(*ProductsFamiliesCreateOptions)
func (srv *ProductsDataModel) WithProductsFamiliesCreateImageAttribute(v string) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.ImageAttribute = v
		o.enabledSetters["ImageAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesCreateLabelAttribute(v string) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.LabelAttribute = v
		o.enabledSetters["LabelAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesCreateLabels(v interface{}) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamiliesCreate creates one family and answers 201 with the stored
// row, including the id and the timestamps the database filled in — a
// client never sends an id, it reads one back and uses it in the path of
// every later call.
// 
// A family decides WHICH attributes a product has — the set is
// `family_attributes`, and every family-driven surface follows from it. It
// also names which attribute carries the display name (`label_attribute`) and
// which carries the main image. A product with no family has no required
// attributes at all, so its completeness cannot be measured and its name
// never resolves past the SKU; `POST /products/{id}/family` is the call that
// ends that state.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsDataModel) ProductsFamiliesCreate(Code string, optionalSetters ...ProductsFamiliesCreateOption)(*models.Error, error) {
	path := "/v1/products/families"
	options := ProductsFamiliesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["ImageAttribute"] {
		params["image_attribute"] = options.ImageAttribute
	}
	if options.enabledSetters["LabelAttribute"] {
		params["label_attribute"] = options.LabelAttribute
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
	
// ProductsFamiliesDelete deletes one family by id. It is a hard delete —
// the row is gone, and the answer is a confirmation rather than a result to
// branch on.
// 
// It takes what hangs off it: family attributes (`family_id`), family
// variants (`family_id`) are deleted with it. `products.family_id` is set to
// null instead, so the rows that pointed at it survive the delete rather than
// going with it.
// 
// An id no family of this tenant carries answers 404; there is no 409,
// because every foreign key pointing at this entity resolves itself on delete
// rather than blocking one.
func (srv *ProductsDataModel) ProductsFamiliesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/families/{id}")
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
	
// ProductsFamiliesGet reads one family by its id — the whole row, every
// column, as it is stored.
// 
// A family decides WHICH attributes a product has — the set is
// `family_attributes`, and every family-driven surface follows from it. It
// also names which attribute carries the display name (`label_attribute`) and
// which carries the main image. A product with no family has no required
// attributes at all, so its completeness cannot be measured and its name
// never resolves past the SKU; `POST /products/{id}/family` is the call that
// ends that state.
// 
// An id no family of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsFamiliesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/families/{id}")
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
type ProductsFamiliesUpdateOptions struct {
	Code string
	ImageAttribute string
	LabelAttribute string
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamiliesUpdateOptions) New() *ProductsFamiliesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"ImageAttribute": false,
		"LabelAttribute": false,
		"Labels": false,
	}
	return &options
}
type ProductsFamiliesUpdateOption func(*ProductsFamiliesUpdateOptions)
func (srv *ProductsDataModel) WithProductsFamiliesUpdateCode(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesUpdateImageAttribute(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.ImageAttribute = v
		o.enabledSetters["ImageAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesUpdateLabelAttribute(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.LabelAttribute = v
		o.enabledSetters["LabelAttribute"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamiliesUpdateLabels(v interface{}) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamiliesUpdate updates one family by id. A partial patch: the body
// names only the columns to change and every column it leaves out keeps its
// current value, so there is no read-modify-write and no way to blank a field
// by forgetting it.
// 
// A family decides WHICH attributes a product has — the set is
// `family_attributes`, and every family-driven surface follows from it. It
// also names which attribute carries the display name (`label_attribute`) and
// which carries the main image. A product with no family has no required
// attributes at all, so its completeness cannot be measured and its name
// never resolves past the SKU; `POST /products/{id}/family` is the call that
// ends that state.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsFamiliesUpdate(Id string, optionalSetters ...ProductsFamiliesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/families/{id}")
	options := ProductsFamiliesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["ImageAttribute"] {
		params["image_attribute"] = options.ImageAttribute
	}
	if options.enabledSetters["LabelAttribute"] {
		params["label_attribute"] = options.LabelAttribute
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsFamilyAttributesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	FamilyId string
	AttributeId string
	Position int
	IsRequired bool
	RequiredChannels string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options ProductsFamilyAttributesListOptions) New() *ProductsFamilyAttributesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"FamilyId": false,
		"AttributeId": false,
		"Position": false,
		"IsRequired": false,
		"RequiredChannels": false,
		"CreatedAt": false,
	}
	return &options
}
type ProductsFamilyAttributesListOption func(*ProductsFamilyAttributesListOptions)
func (srv *ProductsDataModel) WithProductsFamilyAttributesListLimit(v int) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListOffset(v int) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListOrder(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListId(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListFamilyId(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListAttributeId(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListPosition(v int) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListIsRequired(v bool) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.IsRequired = v
		o.enabledSetters["IsRequired"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListRequiredChannels(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.RequiredChannels = v
		o.enabledSetters["RequiredChannels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesListCreatedAt(v string) ProductsFamilyAttributesListOption {
	return func(o *ProductsFamilyAttributesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// ProductsFamilyAttributesList one link between a family and an attribute —
// the row that puts an attribute INTO a family's form. It carries the
// family's own ordering of that attribute, which overrides the attribute's
// default position, and `is_required`, which is the flag `POST
// /products/{id}/completeness` measures and nothing else reads.
// `required_channels` narrows "required" to named channels; null or empty
// means required EVERYWHERE, not nowhere.
// 
// Every column of `family_attributes` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsFamilyAttributesList(optionalSetters ...ProductsFamilyAttributesListOption)(*interface{}, error) {
	path := "/v1/products/family_attributes"
	options := ProductsFamilyAttributesListOptions{}.New()
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
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["AttributeId"] {
		params["attribute_id"] = options.AttributeId
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["IsRequired"] {
		params["is_required"] = options.IsRequired
	}
	if options.enabledSetters["RequiredChannels"] {
		params["required_channels"] = options.RequiredChannels
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
type ProductsFamilyAttributesCreateOptions struct {
	IsRequired bool
	Position int
	RequiredChannels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamilyAttributesCreateOptions) New() *ProductsFamilyAttributesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsRequired": false,
		"Position": false,
		"RequiredChannels": false,
	}
	return &options
}
type ProductsFamilyAttributesCreateOption func(*ProductsFamilyAttributesCreateOptions)
func (srv *ProductsDataModel) WithProductsFamilyAttributesCreateIsRequired(v bool) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.IsRequired = v
		o.enabledSetters["IsRequired"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesCreatePosition(v int) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesCreateRequiredChannels(v interface{}) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.RequiredChannels = v
		o.enabledSetters["RequiredChannels"] = true
	}
}
					
// ProductsFamilyAttributesCreate creates one family attribute and answers 201
// with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// One link between a family and an attribute — the row that puts an
// attribute INTO a family's form. It carries the family's own ordering of
// that attribute, which overrides the attribute's default position, and
// `is_required`, which is the flag `POST /products/{id}/completeness`
// measures and nothing else reads. `required_channels` narrows "required" to
// named channels; null or empty means required EVERYWHERE, not nowhere.
// 
// `family_id` and `attribute_id` are the only columns the database refuses
// the row without; everything else has a default or is nullable. A second row
// with the same `family_id` and `attribute_id` answers 409.
func (srv *ProductsDataModel) ProductsFamilyAttributesCreate(AttributeId string, FamilyId string, optionalSetters ...ProductsFamilyAttributesCreateOption)(*models.Error, error) {
	path := "/v1/products/family_attributes"
	options := ProductsFamilyAttributesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["attribute_id"] = AttributeId
	params["family_id"] = FamilyId
	if options.enabledSetters["IsRequired"] {
		params["is_required"] = options.IsRequired
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["RequiredChannels"] {
		params["required_channels"] = options.RequiredChannels
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
	
// ProductsFamilyAttributesDelete deletes one family attribute by id. It is a
// hard delete — the row is gone, and the answer is a confirmation rather
// than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no family attribute of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsFamilyAttributesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_attributes/{id}")
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
	
// ProductsFamilyAttributesGet reads one family attribute by its id — the
// whole row, every column, as it is stored.
// 
// One link between a family and an attribute — the row that puts an
// attribute INTO a family's form. It carries the family's own ordering of
// that attribute, which overrides the attribute's default position, and
// `is_required`, which is the flag `POST /products/{id}/completeness`
// measures and nothing else reads. `required_channels` narrows "required" to
// named channels; null or empty means required EVERYWHERE, not nowhere.
// 
// An id no family attribute of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsFamilyAttributesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_attributes/{id}")
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
type ProductsFamilyAttributesUpdateOptions struct {
	AttributeId string
	FamilyId string
	IsRequired bool
	Position int
	RequiredChannels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamilyAttributesUpdateOptions) New() *ProductsFamilyAttributesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeId": false,
		"FamilyId": false,
		"IsRequired": false,
		"Position": false,
		"RequiredChannels": false,
	}
	return &options
}
type ProductsFamilyAttributesUpdateOption func(*ProductsFamilyAttributesUpdateOptions)
func (srv *ProductsDataModel) WithProductsFamilyAttributesUpdateAttributeId(v string) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesUpdateFamilyId(v string) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesUpdateIsRequired(v bool) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.IsRequired = v
		o.enabledSetters["IsRequired"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesUpdatePosition(v int) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyAttributesUpdateRequiredChannels(v interface{}) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.RequiredChannels = v
		o.enabledSetters["RequiredChannels"] = true
	}
}
			
// ProductsFamilyAttributesUpdate updates one family attribute by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// One link between a family and an attribute — the row that puts an
// attribute INTO a family's form. It carries the family's own ordering of
// that attribute, which overrides the attribute's default position, and
// `is_required`, which is the flag `POST /products/{id}/completeness`
// measures and nothing else reads. `required_channels` narrows "required" to
// named channels; null or empty means required EVERYWHERE, not nowhere.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `family_id` and `attribute_id` answers 409.
func (srv *ProductsDataModel) ProductsFamilyAttributesUpdate(Id string, optionalSetters ...ProductsFamilyAttributesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_attributes/{id}")
	options := ProductsFamilyAttributesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AttributeId"] {
		params["attribute_id"] = options.AttributeId
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["IsRequired"] {
		params["is_required"] = options.IsRequired
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["RequiredChannels"] {
		params["required_channels"] = options.RequiredChannels
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
type ProductsFamilyVariantsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	FamilyId string
	Code string
	Labels string
	Axes string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsFamilyVariantsListOptions) New() *ProductsFamilyVariantsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"FamilyId": false,
		"Code": false,
		"Labels": false,
		"Axes": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsFamilyVariantsListOption func(*ProductsFamilyVariantsListOptions)
func (srv *ProductsDataModel) WithProductsFamilyVariantsListLimit(v int) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListOffset(v int) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListOrder(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListId(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListFamilyId(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListCode(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListLabels(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListAxes(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.Axes = v
		o.enabledSetters["Axes"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListCreatedAt(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsListUpdatedAt(v string) ProductsFamilyVariantsListOption {
	return func(o *ProductsFamilyVariantsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsFamilyVariantsList a variant structure of a family: the attribute
// axes a product model splits its variants on — colour, then size. A
// product follows one through `family_variant_id`, and an attribute named as
// an axis becomes read-only on the model and is set on each variant instead,
// which is what `GET /products/attribute-schema` reports as
// `readonly_reason`. Two axis shapes are in the wild and both are read: a
// bare list of codes, or one entry per level.
// 
// Every column of `family_variants` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsFamilyVariantsList(optionalSetters ...ProductsFamilyVariantsListOption)(*interface{}, error) {
	path := "/v1/products/family_variants"
	options := ProductsFamilyVariantsListOptions{}.New()
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
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Axes"] {
		params["axes"] = options.Axes
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
type ProductsFamilyVariantsCreateOptions struct {
	Axes interface{}
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamilyVariantsCreateOptions) New() *ProductsFamilyVariantsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Axes": false,
		"Labels": false,
	}
	return &options
}
type ProductsFamilyVariantsCreateOption func(*ProductsFamilyVariantsCreateOptions)
func (srv *ProductsDataModel) WithProductsFamilyVariantsCreateAxes(v interface{}) ProductsFamilyVariantsCreateOption {
	return func(o *ProductsFamilyVariantsCreateOptions) {
		o.Axes = v
		o.enabledSetters["Axes"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsCreateLabels(v interface{}) ProductsFamilyVariantsCreateOption {
	return func(o *ProductsFamilyVariantsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
					
// ProductsFamilyVariantsCreate creates one family variant and answers 201
// with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// A variant structure of a family: the attribute axes a product model splits
// its variants on — colour, then size. A product follows one through
// `family_variant_id`, and an attribute named as an axis becomes read-only on
// the model and is set on each variant instead, which is what `GET
// /products/attribute-schema` reports as `readonly_reason`. Two axis shapes
// are in the wild and both are read: a bare list of codes, or one entry per
// level.
// 
// `family_id` and `code` are the only columns the database refuses the row
// without; everything else has a default or is nullable. A second row with
// the same `code` answers 409.
func (srv *ProductsDataModel) ProductsFamilyVariantsCreate(Code string, FamilyId string, optionalSetters ...ProductsFamilyVariantsCreateOption)(*models.Error, error) {
	path := "/v1/products/family_variants"
	options := ProductsFamilyVariantsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["family_id"] = FamilyId
	if options.enabledSetters["Axes"] {
		params["axes"] = options.Axes
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
	
// ProductsFamilyVariantsDelete deletes one family variant by id. It is a hard
// delete — the row is gone, and the answer is a confirmation rather than a
// result to branch on.
// 
// `products.family_variant_id` is set to null instead, so the rows that
// pointed at it survive the delete rather than going with it.
// 
// An id no family variant of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsFamilyVariantsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_variants/{id}")
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
	
// ProductsFamilyVariantsGet reads one family variant by its id — the whole
// row, every column, as it is stored.
// 
// A variant structure of a family: the attribute axes a product model splits
// its variants on — colour, then size. A product follows one through
// `family_variant_id`, and an attribute named as an axis becomes read-only on
// the model and is set on each variant instead, which is what `GET
// /products/attribute-schema` reports as `readonly_reason`. Two axis shapes
// are in the wild and both are read: a bare list of codes, or one entry per
// level.
// 
// An id no family variant of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsFamilyVariantsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_variants/{id}")
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
type ProductsFamilyVariantsUpdateOptions struct {
	Axes interface{}
	Code string
	FamilyId string
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsFamilyVariantsUpdateOptions) New() *ProductsFamilyVariantsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Axes": false,
		"Code": false,
		"FamilyId": false,
		"Labels": false,
	}
	return &options
}
type ProductsFamilyVariantsUpdateOption func(*ProductsFamilyVariantsUpdateOptions)
func (srv *ProductsDataModel) WithProductsFamilyVariantsUpdateAxes(v interface{}) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Axes = v
		o.enabledSetters["Axes"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsUpdateCode(v string) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsUpdateFamilyId(v string) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *ProductsDataModel) WithProductsFamilyVariantsUpdateLabels(v interface{}) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamilyVariantsUpdate updates one family variant by id. A partial
// patch: the body names only the columns to change and every column it leaves
// out keeps its current value, so there is no read-modify-write and no way to
// blank a field by forgetting it.
// 
// A variant structure of a family: the attribute axes a product model splits
// its variants on — colour, then size. A product follows one through
// `family_variant_id`, and an attribute named as an axis becomes read-only on
// the model and is set on each variant instead, which is what `GET
// /products/attribute-schema` reports as `readonly_reason`. Two axis shapes
// are in the wild and both are read: a bare list of codes, or one entry per
// level.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsFamilyVariantsUpdate(Id string, optionalSetters ...ProductsFamilyVariantsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/family_variants/{id}")
	options := ProductsFamilyVariantsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Axes"] {
		params["axes"] = options.Axes
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["FamilyId"] {
		params["family_id"] = options.FamilyId
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsMeasurementFamiliesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	StandardUnit string
	Units string
	Labels string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsMeasurementFamiliesListOptions) New() *ProductsMeasurementFamiliesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"StandardUnit": false,
		"Units": false,
		"Labels": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsMeasurementFamiliesListOption func(*ProductsMeasurementFamiliesListOptions)
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListLimit(v int) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListOffset(v int) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListOrder(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListId(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListCode(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListStandardUnit(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.StandardUnit = v
		o.enabledSetters["StandardUnit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListUnits(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Units = v
		o.enabledSetters["Units"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListLabels(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListCreatedAt(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesListUpdatedAt(v string) ProductsMeasurementFamiliesListOption {
	return func(o *ProductsMeasurementFamiliesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsMeasurementFamiliesList a family of units and the standard one they
// all convert to — weight in kilograms, length in metres. A `measure`
// attribute names one and then offers exactly that family's units, and each
// unit's `convert_factor` is what makes two values recorded in different
// units comparable at all.
// 
// Every column of `measurement_families` is an exact-match query parameter,
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
func (srv *ProductsDataModel) ProductsMeasurementFamiliesList(optionalSetters ...ProductsMeasurementFamiliesListOption)(*interface{}, error) {
	path := "/v1/products/measurement_families"
	options := ProductsMeasurementFamiliesListOptions{}.New()
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
	if options.enabledSetters["StandardUnit"] {
		params["standard_unit"] = options.StandardUnit
	}
	if options.enabledSetters["Units"] {
		params["units"] = options.Units
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsMeasurementFamiliesCreateOptions struct {
	Labels interface{}
	Units interface{}
	enabledSetters map[string]bool
}
func (options ProductsMeasurementFamiliesCreateOptions) New() *ProductsMeasurementFamiliesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"Units": false,
	}
	return &options
}
type ProductsMeasurementFamiliesCreateOption func(*ProductsMeasurementFamiliesCreateOptions)
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesCreateLabels(v interface{}) ProductsMeasurementFamiliesCreateOption {
	return func(o *ProductsMeasurementFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesCreateUnits(v interface{}) ProductsMeasurementFamiliesCreateOption {
	return func(o *ProductsMeasurementFamiliesCreateOptions) {
		o.Units = v
		o.enabledSetters["Units"] = true
	}
}
					
// ProductsMeasurementFamiliesCreate creates one measurement family and
// answers 201 with the stored row, including the id and the timestamps the
// database filled in — a client never sends an id, it reads one back and
// uses it in the path of every later call.
// 
// A family of units and the standard one they all convert to — weight in
// kilograms, length in metres. A `measure` attribute names one and then
// offers exactly that family's units, and each unit's `convert_factor` is
// what makes two values recorded in different units comparable at all.
// 
// `code` and `standard_unit` are the only columns the database refuses the
// row without; everything else has a default or is nullable. A second row
// with the same `code` answers 409.
func (srv *ProductsDataModel) ProductsMeasurementFamiliesCreate(Code string, StandardUnit string, optionalSetters ...ProductsMeasurementFamiliesCreateOption)(*models.Error, error) {
	path := "/v1/products/measurement_families"
	options := ProductsMeasurementFamiliesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["standard_unit"] = StandardUnit
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Units"] {
		params["units"] = options.Units
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
	
// ProductsMeasurementFamiliesDelete deletes one measurement family by id. It
// is a hard delete — the row is gone, and the answer is a confirmation
// rather than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no measurement family of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsDataModel) ProductsMeasurementFamiliesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/measurement_families/{id}")
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
	
// ProductsMeasurementFamiliesGet reads one measurement family by its id —
// the whole row, every column, as it is stored.
// 
// A family of units and the standard one they all convert to — weight in
// kilograms, length in metres. A `measure` attribute names one and then
// offers exactly that family's units, and each unit's `convert_factor` is
// what makes two values recorded in different units comparable at all.
// 
// An id no measurement family of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsDataModel) ProductsMeasurementFamiliesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/measurement_families/{id}")
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
type ProductsMeasurementFamiliesUpdateOptions struct {
	Code string
	Labels interface{}
	StandardUnit string
	Units interface{}
	enabledSetters map[string]bool
}
func (options ProductsMeasurementFamiliesUpdateOptions) New() *ProductsMeasurementFamiliesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"StandardUnit": false,
		"Units": false,
	}
	return &options
}
type ProductsMeasurementFamiliesUpdateOption func(*ProductsMeasurementFamiliesUpdateOptions)
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesUpdateCode(v string) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesUpdateLabels(v interface{}) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesUpdateStandardUnit(v string) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.StandardUnit = v
		o.enabledSetters["StandardUnit"] = true
	}
}
func (srv *ProductsDataModel) WithProductsMeasurementFamiliesUpdateUnits(v interface{}) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Units = v
		o.enabledSetters["Units"] = true
	}
}
			
// ProductsMeasurementFamiliesUpdate updates one measurement family by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// A family of units and the standard one they all convert to — weight in
// kilograms, length in metres. A `measure` attribute names one and then
// offers exactly that family's units, and each unit's `convert_factor` is
// what makes two values recorded in different units comparable at all.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsDataModel) ProductsMeasurementFamiliesUpdate(Id string, optionalSetters ...ProductsMeasurementFamiliesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/measurement_families/{id}")
	options := ProductsMeasurementFamiliesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["StandardUnit"] {
		params["standard_unit"] = options.StandardUnit
	}
	if options.enabledSetters["Units"] {
		params["units"] = options.Units
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
