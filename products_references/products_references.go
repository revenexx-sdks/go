package products_references

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ProductsReferences service
type ProductsReferences struct {
	client client.Client
}

func New(clt client.Client) *ProductsReferences {
	return &ProductsReferences{
		client: clt,
	}
}

type ProductsReferenceEntitiesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	Labels string
	Image string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntitiesListOptions) New() *ProductsReferenceEntitiesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"Labels": false,
		"Image": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsReferenceEntitiesListOption func(*ProductsReferenceEntitiesListOptions)
func (srv *ProductsReferences) WithProductsReferenceEntitiesListLimit(v int) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListOffset(v int) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListOrder(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListId(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListCode(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListLabels(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListImage(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListCreatedAt(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesListUpdatedAt(v string) ProductsReferenceEntitiesListOption {
	return func(o *ProductsReferenceEntitiesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsReferenceEntitiesList a domain of records the catalog POINTS AT
// instead of duplicating — brands, manufacturers, care instructions.
// Declaring one is how a brand comes to be edited in one place rather than on
// nine thousand products. A reference entity has attributes of its own
// (`attributes` rows with `entity_type: "reference_entity"` and this entity's
// code as `entity_ref`), which is what makes its records more than a label.
// 
// Every column of `reference_entities` is an exact-match query parameter,
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
func (srv *ProductsReferences) ProductsReferenceEntitiesList(optionalSetters ...ProductsReferenceEntitiesListOption)(*interface{}, error) {
	path := "/v1/products/reference_entities"
	options := ProductsReferenceEntitiesListOptions{}.New()
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
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
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
type ProductsReferenceEntitiesCreateOptions struct {
	Image string
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntitiesCreateOptions) New() *ProductsReferenceEntitiesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Image": false,
		"Labels": false,
	}
	return &options
}
type ProductsReferenceEntitiesCreateOption func(*ProductsReferenceEntitiesCreateOptions)
func (srv *ProductsReferences) WithProductsReferenceEntitiesCreateImage(v string) ProductsReferenceEntitiesCreateOption {
	return func(o *ProductsReferenceEntitiesCreateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesCreateLabels(v interface{}) ProductsReferenceEntitiesCreateOption {
	return func(o *ProductsReferenceEntitiesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsReferenceEntitiesCreate creates one reference entity and answers
// 201 with the stored row, including the id and the timestamps the database
// filled in — a client never sends an id, it reads one back and uses it in
// the path of every later call.
// 
// A domain of records the catalog POINTS AT instead of duplicating —
// brands, manufacturers, care instructions. Declaring one is how a brand
// comes to be edited in one place rather than on nine thousand products. A
// reference entity has attributes of its own (`attributes` rows with
// `entity_type: "reference_entity"` and this entity's code as `entity_ref`),
// which is what makes its records more than a label.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsReferences) ProductsReferenceEntitiesCreate(Code string, optionalSetters ...ProductsReferenceEntitiesCreateOption)(*models.Error, error) {
	path := "/v1/products/reference_entities"
	options := ProductsReferenceEntitiesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
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
	
// ProductsReferenceEntitiesDelete deletes one reference entity by id. It is a
// hard delete — the row is gone, and the answer is a confirmation rather
// than a result to branch on.
// 
// It takes what hangs off it: reference entity records
// (`reference_entity_id`) are deleted with it.
// 
// An id no reference entity of this tenant carries answers 404; there is no
// 409, because every foreign key pointing at this entity resolves itself on
// delete rather than blocking one.
func (srv *ProductsReferences) ProductsReferenceEntitiesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entities/{id}")
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
	
// ProductsReferenceEntitiesGet reads one reference entity by its id — the
// whole row, every column, as it is stored.
// 
// A domain of records the catalog POINTS AT instead of duplicating —
// brands, manufacturers, care instructions. Declaring one is how a brand
// comes to be edited in one place rather than on nine thousand products. A
// reference entity has attributes of its own (`attributes` rows with
// `entity_type: "reference_entity"` and this entity's code as `entity_ref`),
// which is what makes its records more than a label.
// 
// An id no reference entity of this tenant carries answers 404, and so does
// one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsReferences) ProductsReferenceEntitiesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entities/{id}")
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
type ProductsReferenceEntitiesUpdateOptions struct {
	Code string
	Image string
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntitiesUpdateOptions) New() *ProductsReferenceEntitiesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Image": false,
		"Labels": false,
	}
	return &options
}
type ProductsReferenceEntitiesUpdateOption func(*ProductsReferenceEntitiesUpdateOptions)
func (srv *ProductsReferences) WithProductsReferenceEntitiesUpdateCode(v string) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesUpdateImage(v string) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntitiesUpdateLabels(v interface{}) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsReferenceEntitiesUpdate updates one reference entity by id. A
// partial patch: the body names only the columns to change and every column
// it leaves out keeps its current value, so there is no read-modify-write and
// no way to blank a field by forgetting it.
// 
// A domain of records the catalog POINTS AT instead of duplicating —
// brands, manufacturers, care instructions. Declaring one is how a brand
// comes to be edited in one place rather than on nine thousand products. A
// reference entity has attributes of its own (`attributes` rows with
// `entity_type: "reference_entity"` and this entity's code as `entity_ref`),
// which is what makes its records more than a label.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsReferences) ProductsReferenceEntitiesUpdate(Id string, optionalSetters ...ProductsReferenceEntitiesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entities/{id}")
	options := ProductsReferenceEntitiesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
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
type ProductsReferenceEntityRecordsListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	ReferenceEntityId string
	Code string
	Labels string
	AttributeValues string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntityRecordsListOptions) New() *ProductsReferenceEntityRecordsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"ReferenceEntityId": false,
		"Code": false,
		"Labels": false,
		"AttributeValues": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsReferenceEntityRecordsListOption func(*ProductsReferenceEntityRecordsListOptions)
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListLimit(v int) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListOffset(v int) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListOrder(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListId(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListReferenceEntityId(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.ReferenceEntityId = v
		o.enabledSetters["ReferenceEntityId"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListCode(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListLabels(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListAttributeValues(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListCreatedAt(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsListUpdatedAt(v string) ProductsReferenceEntityRecordsListOption {
	return func(o *ProductsReferenceEntityRecordsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsReferenceEntityRecordsList one record of a reference entity — one
// brand, one manufacturer. A product that points at it stores this record's
// CODE, exactly the way a select stores an option code, and the record's own
// properties live in its scoped `attribute_values` document. `GET
// /products/attribute-schema` offers these records as the `options` of any
// attribute that points at their entity, so a picker needs no second call.
// 
// Every column of `reference_entity_records` is an exact-match query
// parameter, `order` sorts by one column, and `limit`/`offset` page through
// `page.total`. A query key that is NOT a column is dropped rather than
// refused, and the `filter` object echoes the ones that were understood —
// that echo is the only way to tell an unfiltered answer from an empty one.
// It reads rows exactly as they are stored: no join is resolved, no jsonb
// value is unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsReferences) ProductsReferenceEntityRecordsList(optionalSetters ...ProductsReferenceEntityRecordsListOption)(*interface{}, error) {
	path := "/v1/products/reference_entity_records"
	options := ProductsReferenceEntityRecordsListOptions{}.New()
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
	if options.enabledSetters["ReferenceEntityId"] {
		params["reference_entity_id"] = options.ReferenceEntityId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
type ProductsReferenceEntityRecordsCreateOptions struct {
	AttributeValues interface{}
	Labels interface{}
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntityRecordsCreateOptions) New() *ProductsReferenceEntityRecordsCreateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"Labels": false,
	}
	return &options
}
type ProductsReferenceEntityRecordsCreateOption func(*ProductsReferenceEntityRecordsCreateOptions)
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsCreateAttributeValues(v interface{}) ProductsReferenceEntityRecordsCreateOption {
	return func(o *ProductsReferenceEntityRecordsCreateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsCreateLabels(v interface{}) ProductsReferenceEntityRecordsCreateOption {
	return func(o *ProductsReferenceEntityRecordsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
					
// ProductsReferenceEntityRecordsCreate creates one reference entity record
// and answers 201 with the stored row, including the id and the timestamps
// the database filled in — a client never sends an id, it reads one back
// and uses it in the path of every later call.
// 
// One record of a reference entity — one brand, one manufacturer. A product
// that points at it stores this record's CODE, exactly the way a select
// stores an option code, and the record's own properties live in its scoped
// `attribute_values` document. `GET /products/attribute-schema` offers these
// records as the `options` of any attribute that points at their entity, so a
// picker needs no second call.
// 
// `reference_entity_id` and `code` are the only columns the database refuses
// the row without; everything else has a default or is nullable. A second row
// with the same `reference_entity_id` and `code` answers 409.
func (srv *ProductsReferences) ProductsReferenceEntityRecordsCreate(Code string, ReferenceEntityId string, optionalSetters ...ProductsReferenceEntityRecordsCreateOption)(*models.Error, error) {
	path := "/v1/products/reference_entity_records"
	options := ProductsReferenceEntityRecordsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["reference_entity_id"] = ReferenceEntityId
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
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
	
// ProductsReferenceEntityRecordsDelete deletes one reference entity record by
// id. It is a hard delete — the row is gone, and the answer is a
// confirmation rather than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no reference entity record of this tenant carries answers 404; there
// is no 409, because every foreign key pointing at this entity resolves
// itself on delete rather than blocking one.
func (srv *ProductsReferences) ProductsReferenceEntityRecordsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entity_records/{id}")
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
	
// ProductsReferenceEntityRecordsGet reads one reference entity record by its
// id — the whole row, every column, as it is stored.
// 
// One record of a reference entity — one brand, one manufacturer. A product
// that points at it stores this record's CODE, exactly the way a select
// stores an option code, and the record's own properties live in its scoped
// `attribute_values` document. `GET /products/attribute-schema` offers these
// records as the `options` of any attribute that points at their entity, so a
// picker needs no second call.
// 
// An id no reference entity record of this tenant carries answers 404, and so
// does one belonging to another tenant: row-level security makes that row
// invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsReferences) ProductsReferenceEntityRecordsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entity_records/{id}")
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
type ProductsReferenceEntityRecordsUpdateOptions struct {
	AttributeValues interface{}
	Code string
	Labels interface{}
	ReferenceEntityId string
	enabledSetters map[string]bool
}
func (options ProductsReferenceEntityRecordsUpdateOptions) New() *ProductsReferenceEntityRecordsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"Code": false,
		"Labels": false,
		"ReferenceEntityId": false,
	}
	return &options
}
type ProductsReferenceEntityRecordsUpdateOption func(*ProductsReferenceEntityRecordsUpdateOptions)
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsUpdateAttributeValues(v interface{}) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsUpdateCode(v string) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsUpdateLabels(v interface{}) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsReferences) WithProductsReferenceEntityRecordsUpdateReferenceEntityId(v string) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.ReferenceEntityId = v
		o.enabledSetters["ReferenceEntityId"] = true
	}
}
			
// ProductsReferenceEntityRecordsUpdate updates one reference entity record by
// id. A partial patch: the body names only the columns to change and every
// column it leaves out keeps its current value, so there is no
// read-modify-write and no way to blank a field by forgetting it.
// 
// One record of a reference entity — one brand, one manufacturer. A product
// that points at it stores this record's CODE, exactly the way a select
// stores an option code, and the record's own properties live in its scoped
// `attribute_values` document. `GET /products/attribute-schema` offers these
// records as the `options` of any attribute that points at their entity, so a
// picker needs no second call.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `reference_entity_id` and `code` answers 409.
func (srv *ProductsReferences) ProductsReferenceEntityRecordsUpdate(Id string, optionalSetters ...ProductsReferenceEntityRecordsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/reference_entity_records/{id}")
	options := ProductsReferenceEntityRecordsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AttributeValues"] {
		params["attribute_values"] = options.AttributeValues
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["ReferenceEntityId"] {
		params["reference_entity_id"] = options.ReferenceEntityId
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
