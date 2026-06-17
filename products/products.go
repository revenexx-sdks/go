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


// ProductsList
func (srv *Products) ProductsList()(*interface{}, error) {
	path := "/v1/products"
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
			
// ProductsCreate
func (srv *Products) ProductsCreate(Sku string, optionalSetters ...ProductsCreateOption)(*models.Products, error) {
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

		parsed := models.Products{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Products
	parsed, ok := resp.Result.(models.Products)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAssetFamiliesList
func (srv *Products) ProductsAssetFamiliesList()(*interface{}, error) {
	path := "/v1/products/asset_families"
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
func (srv *Products) WithProductsAssetFamiliesCreateLabels(v interface{}) ProductsAssetFamiliesCreateOption {
	return func(o *ProductsAssetFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAssetFamiliesCreateNamingConvention(v interface{}) ProductsAssetFamiliesCreateOption {
	return func(o *ProductsAssetFamiliesCreateOptions) {
		o.NamingConvention = v
		o.enabledSetters["NamingConvention"] = true
	}
}
			
// ProductsAssetFamiliesCreate
func (srv *Products) ProductsAssetFamiliesCreate(Code string, optionalSetters ...ProductsAssetFamiliesCreateOption)(*models.AssetFamilies, error) {
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

		parsed := models.AssetFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssetFamilies
	parsed, ok := resp.Result.(models.AssetFamilies)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAssetFamiliesDelete
func (srv *Products) ProductsAssetFamiliesDelete(Id string)(*interface{}, error) {
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
	
// ProductsAssetFamiliesGet
func (srv *Products) ProductsAssetFamiliesGet(Id string)(*models.AssetFamilies, error) {
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

		parsed := models.AssetFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssetFamilies
	parsed, ok := resp.Result.(models.AssetFamilies)
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
func (srv *Products) WithProductsAssetFamiliesUpdateCode(v string) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAssetFamiliesUpdateLabels(v interface{}) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAssetFamiliesUpdateNamingConvention(v interface{}) ProductsAssetFamiliesUpdateOption {
	return func(o *ProductsAssetFamiliesUpdateOptions) {
		o.NamingConvention = v
		o.enabledSetters["NamingConvention"] = true
	}
}
			
// ProductsAssetFamiliesUpdate
func (srv *Products) ProductsAssetFamiliesUpdate(Id string, optionalSetters ...ProductsAssetFamiliesUpdateOption)(*models.AssetFamilies, error) {
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

		parsed := models.AssetFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssetFamilies
	parsed, ok := resp.Result.(models.AssetFamilies)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAssetsList
func (srv *Products) ProductsAssetsList()(*interface{}, error) {
	path := "/v1/products/assets"
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
type ProductsAssetsCreateOptions struct {
	AttributeValues interface{}
	MediaUuid string
	enabledSetters map[string]bool
}
func (options ProductsAssetsCreateOptions) New() *ProductsAssetsCreateOptions {
	options.enabledSetters = map[string]bool{
		"AttributeValues": false,
		"MediaUuid": false,
	}
	return &options
}
type ProductsAssetsCreateOption func(*ProductsAssetsCreateOptions)
func (srv *Products) WithProductsAssetsCreateAttributeValues(v interface{}) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsAssetsCreateMediaUuid(v string) ProductsAssetsCreateOption {
	return func(o *ProductsAssetsCreateOptions) {
		o.MediaUuid = v
		o.enabledSetters["MediaUuid"] = true
	}
}
					
// ProductsAssetsCreate
func (srv *Products) ProductsAssetsCreate(AssetFamilyId string, Code string, optionalSetters ...ProductsAssetsCreateOption)(*models.Assets, error) {
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
	if options.enabledSetters["MediaUuid"] {
		params["media_uuid"] = options.MediaUuid
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

		parsed := models.Assets{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Assets
	parsed, ok := resp.Result.(models.Assets)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAssetsDelete
func (srv *Products) ProductsAssetsDelete(Id string)(*interface{}, error) {
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
	
// ProductsAssetsGet
func (srv *Products) ProductsAssetsGet(Id string)(*models.Assets, error) {
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

		parsed := models.Assets{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Assets
	parsed, ok := resp.Result.(models.Assets)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ProductsAssetsUpdateOptions struct {
	AssetFamilyId string
	AttributeValues interface{}
	Code string
	MediaUuid string
	enabledSetters map[string]bool
}
func (options ProductsAssetsUpdateOptions) New() *ProductsAssetsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AssetFamilyId": false,
		"AttributeValues": false,
		"Code": false,
		"MediaUuid": false,
	}
	return &options
}
type ProductsAssetsUpdateOption func(*ProductsAssetsUpdateOptions)
func (srv *Products) WithProductsAssetsUpdateAssetFamilyId(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.AssetFamilyId = v
		o.enabledSetters["AssetFamilyId"] = true
	}
}
func (srv *Products) WithProductsAssetsUpdateAttributeValues(v interface{}) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsAssetsUpdateCode(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAssetsUpdateMediaUuid(v string) ProductsAssetsUpdateOption {
	return func(o *ProductsAssetsUpdateOptions) {
		o.MediaUuid = v
		o.enabledSetters["MediaUuid"] = true
	}
}
			
// ProductsAssetsUpdate
func (srv *Products) ProductsAssetsUpdate(Id string, optionalSetters ...ProductsAssetsUpdateOption)(*models.Assets, error) {
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
	if options.enabledSetters["MediaUuid"] {
		params["media_uuid"] = options.MediaUuid
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

		parsed := models.Assets{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Assets
	parsed, ok := resp.Result.(models.Assets)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAssociationTypesList
func (srv *Products) ProductsAssociationTypesList()(*interface{}, error) {
	path := "/v1/products/association_types"
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
func (srv *Products) WithProductsAssociationTypesCreateIsQuantified(v bool) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.IsQuantified = v
		o.enabledSetters["IsQuantified"] = true
	}
}
func (srv *Products) WithProductsAssociationTypesCreateIsTwoWay(v bool) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.IsTwoWay = v
		o.enabledSetters["IsTwoWay"] = true
	}
}
func (srv *Products) WithProductsAssociationTypesCreateLabels(v interface{}) ProductsAssociationTypesCreateOption {
	return func(o *ProductsAssociationTypesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsAssociationTypesCreate
func (srv *Products) ProductsAssociationTypesCreate(Code string, optionalSetters ...ProductsAssociationTypesCreateOption)(*models.AssociationTypes, error) {
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

		parsed := models.AssociationTypes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssociationTypes
	parsed, ok := resp.Result.(models.AssociationTypes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAssociationTypesDelete
func (srv *Products) ProductsAssociationTypesDelete(Id string)(*interface{}, error) {
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
	
// ProductsAssociationTypesGet
func (srv *Products) ProductsAssociationTypesGet(Id string)(*models.AssociationTypes, error) {
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

		parsed := models.AssociationTypes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssociationTypes
	parsed, ok := resp.Result.(models.AssociationTypes)
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
func (srv *Products) WithProductsAssociationTypesUpdateCode(v string) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAssociationTypesUpdateIsQuantified(v bool) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.IsQuantified = v
		o.enabledSetters["IsQuantified"] = true
	}
}
func (srv *Products) WithProductsAssociationTypesUpdateIsTwoWay(v bool) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.IsTwoWay = v
		o.enabledSetters["IsTwoWay"] = true
	}
}
func (srv *Products) WithProductsAssociationTypesUpdateLabels(v interface{}) ProductsAssociationTypesUpdateOption {
	return func(o *ProductsAssociationTypesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsAssociationTypesUpdate
func (srv *Products) ProductsAssociationTypesUpdate(Id string, optionalSetters ...ProductsAssociationTypesUpdateOption)(*models.AssociationTypes, error) {
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

		parsed := models.AssociationTypes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AssociationTypes
	parsed, ok := resp.Result.(models.AssociationTypes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAttributeGroupsList
func (srv *Products) ProductsAttributeGroupsList()(*interface{}, error) {
	path := "/v1/products/attribute_groups"
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
func (srv *Products) WithProductsAttributeGroupsCreateLabels(v interface{}) ProductsAttributeGroupsCreateOption {
	return func(o *ProductsAttributeGroupsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributeGroupsCreatePosition(v int) ProductsAttributeGroupsCreateOption {
	return func(o *ProductsAttributeGroupsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
			
// ProductsAttributeGroupsCreate
func (srv *Products) ProductsAttributeGroupsCreate(Code string, optionalSetters ...ProductsAttributeGroupsCreateOption)(*models.AttributeGroups, error) {
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

		parsed := models.AttributeGroups{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeGroups
	parsed, ok := resp.Result.(models.AttributeGroups)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAttributeGroupsDelete
func (srv *Products) ProductsAttributeGroupsDelete(Id string)(*interface{}, error) {
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
	
// ProductsAttributeGroupsGet
func (srv *Products) ProductsAttributeGroupsGet(Id string)(*models.AttributeGroups, error) {
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

		parsed := models.AttributeGroups{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeGroups
	parsed, ok := resp.Result.(models.AttributeGroups)
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
func (srv *Products) WithProductsAttributeGroupsUpdateCode(v string) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAttributeGroupsUpdateLabels(v interface{}) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributeGroupsUpdatePosition(v int) ProductsAttributeGroupsUpdateOption {
	return func(o *ProductsAttributeGroupsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
			
// ProductsAttributeGroupsUpdate
func (srv *Products) ProductsAttributeGroupsUpdate(Id string, optionalSetters ...ProductsAttributeGroupsUpdateOption)(*models.AttributeGroups, error) {
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

		parsed := models.AttributeGroups{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeGroups
	parsed, ok := resp.Result.(models.AttributeGroups)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAttributeOptionsList
func (srv *Products) ProductsAttributeOptionsList()(*interface{}, error) {
	path := "/v1/products/attribute_options"
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
func (srv *Products) WithProductsAttributeOptionsCreateLabels(v interface{}) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsCreatePosition(v int) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsCreateSwatch(v interface{}) ProductsAttributeOptionsCreateOption {
	return func(o *ProductsAttributeOptionsCreateOptions) {
		o.Swatch = v
		o.enabledSetters["Swatch"] = true
	}
}
					
// ProductsAttributeOptionsCreate
func (srv *Products) ProductsAttributeOptionsCreate(AttributeId string, Code string, optionalSetters ...ProductsAttributeOptionsCreateOption)(*models.AttributeOptions, error) {
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

		parsed := models.AttributeOptions{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeOptions
	parsed, ok := resp.Result.(models.AttributeOptions)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAttributeOptionsDelete
func (srv *Products) ProductsAttributeOptionsDelete(Id string)(*interface{}, error) {
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
	
// ProductsAttributeOptionsGet
func (srv *Products) ProductsAttributeOptionsGet(Id string)(*models.AttributeOptions, error) {
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

		parsed := models.AttributeOptions{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeOptions
	parsed, ok := resp.Result.(models.AttributeOptions)
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
func (srv *Products) WithProductsAttributeOptionsUpdateAttributeId(v string) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsUpdateCode(v string) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsUpdateLabels(v interface{}) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsUpdatePosition(v int) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsAttributeOptionsUpdateSwatch(v interface{}) ProductsAttributeOptionsUpdateOption {
	return func(o *ProductsAttributeOptionsUpdateOptions) {
		o.Swatch = v
		o.enabledSetters["Swatch"] = true
	}
}
			
// ProductsAttributeOptionsUpdate
func (srv *Products) ProductsAttributeOptionsUpdate(Id string, optionalSetters ...ProductsAttributeOptionsUpdateOption)(*models.AttributeOptions, error) {
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

		parsed := models.AttributeOptions{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeOptions
	parsed, ok := resp.Result.(models.AttributeOptions)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsAttributesList
func (srv *Products) ProductsAttributesList()(*interface{}, error) {
	path := "/v1/products/attributes"
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
func (srv *Products) WithProductsAttributesCreateConfig(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Config = v
		o.enabledSetters["Config"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateEntityRef(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateEntityType(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateGroupId(v string) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.GroupId = v
		o.enabledSetters["GroupId"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateIsFilterable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.IsFilterable = v
		o.enabledSetters["IsFilterable"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateIsUnique(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.IsUnique = v
		o.enabledSetters["IsUnique"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateLabels(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateLocalizable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Localizable = v
		o.enabledSetters["Localizable"] = true
	}
}
func (srv *Products) WithProductsAttributesCreatePosition(v int) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateScopable(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Scopable = v
		o.enabledSetters["Scopable"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateUsableInGrid(v bool) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.UsableInGrid = v
		o.enabledSetters["UsableInGrid"] = true
	}
}
func (srv *Products) WithProductsAttributesCreateValidation(v interface{}) ProductsAttributesCreateOption {
	return func(o *ProductsAttributesCreateOptions) {
		o.Validation = v
		o.enabledSetters["Validation"] = true
	}
}
					
// ProductsAttributesCreate
func (srv *Products) ProductsAttributesCreate(Code string, Type string, optionalSetters ...ProductsAttributesCreateOption)(*models.Attributes, error) {
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

		parsed := models.Attributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Attributes
	parsed, ok := resp.Result.(models.Attributes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsAttributesDelete
func (srv *Products) ProductsAttributesDelete(Id string)(*interface{}, error) {
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
	
// ProductsAttributesGet
func (srv *Products) ProductsAttributesGet(Id string)(*models.Attributes, error) {
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

		parsed := models.Attributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Attributes
	parsed, ok := resp.Result.(models.Attributes)
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
func (srv *Products) WithProductsAttributesUpdateCode(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateConfig(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Config = v
		o.enabledSetters["Config"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateEntityRef(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.EntityRef = v
		o.enabledSetters["EntityRef"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateEntityType(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.EntityType = v
		o.enabledSetters["EntityType"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateGroupId(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.GroupId = v
		o.enabledSetters["GroupId"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateIsFilterable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.IsFilterable = v
		o.enabledSetters["IsFilterable"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateIsUnique(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.IsUnique = v
		o.enabledSetters["IsUnique"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateLabels(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateLocalizable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Localizable = v
		o.enabledSetters["Localizable"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdatePosition(v int) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateScopable(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Scopable = v
		o.enabledSetters["Scopable"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateType(v string) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateUsableInGrid(v bool) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.UsableInGrid = v
		o.enabledSetters["UsableInGrid"] = true
	}
}
func (srv *Products) WithProductsAttributesUpdateValidation(v interface{}) ProductsAttributesUpdateOption {
	return func(o *ProductsAttributesUpdateOptions) {
		o.Validation = v
		o.enabledSetters["Validation"] = true
	}
}
			
// ProductsAttributesUpdate
func (srv *Products) ProductsAttributesUpdate(Id string, optionalSetters ...ProductsAttributesUpdateOption)(*models.Attributes, error) {
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

		parsed := models.Attributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Attributes
	parsed, ok := resp.Result.(models.Attributes)
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
	
// ProductsBatch
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

// ProductsCategoriesList
func (srv *Products) ProductsCategoriesList()(*interface{}, error) {
	path := "/v1/products/categories"
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
type ProductsCategoriesCreateOptions struct {
	Labels interface{}
	ParentId string
	Path string
	Position int
	Values interface{}
	enabledSetters map[string]bool
}
func (options ProductsCategoriesCreateOptions) New() *ProductsCategoriesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"ParentId": false,
		"Path": false,
		"Position": false,
		"Values": false,
	}
	return &options
}
type ProductsCategoriesCreateOption func(*ProductsCategoriesCreateOptions)
func (srv *Products) WithProductsCategoriesCreateLabels(v interface{}) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsCategoriesCreateParentId(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *Products) WithProductsCategoriesCreatePath(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *Products) WithProductsCategoriesCreatePosition(v int) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsCategoriesCreateValues(v interface{}) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Values = v
		o.enabledSetters["Values"] = true
	}
}
			
// ProductsCategoriesCreate
func (srv *Products) ProductsCategoriesCreate(Code string, optionalSetters ...ProductsCategoriesCreateOption)(*models.Categories, error) {
	path := "/v1/products/categories"
	options := ProductsCategoriesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Values"] {
		params["values"] = options.Values
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

		parsed := models.Categories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Categories
	parsed, ok := resp.Result.(models.Categories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsCategoriesDelete
func (srv *Products) ProductsCategoriesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
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
	
// ProductsCategoriesGet
func (srv *Products) ProductsCategoriesGet(Id string)(*models.Categories, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
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

		parsed := models.Categories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Categories
	parsed, ok := resp.Result.(models.Categories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ProductsCategoriesUpdateOptions struct {
	Code string
	Labels interface{}
	ParentId string
	Path string
	Position int
	Values interface{}
	enabledSetters map[string]bool
}
func (options ProductsCategoriesUpdateOptions) New() *ProductsCategoriesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"ParentId": false,
		"Path": false,
		"Position": false,
		"Values": false,
	}
	return &options
}
type ProductsCategoriesUpdateOption func(*ProductsCategoriesUpdateOptions)
func (srv *Products) WithProductsCategoriesUpdateCode(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsCategoriesUpdateLabels(v interface{}) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsCategoriesUpdateParentId(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *Products) WithProductsCategoriesUpdatePath(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *Products) WithProductsCategoriesUpdatePosition(v int) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsCategoriesUpdateValues(v interface{}) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Values = v
		o.enabledSetters["Values"] = true
	}
}
			
// ProductsCategoriesUpdate
func (srv *Products) ProductsCategoriesUpdate(Id string, optionalSetters ...ProductsCategoriesUpdateOption)(*models.Categories, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
	options := ProductsCategoriesUpdateOptions{}.New()
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
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Values"] {
		params["values"] = options.Values
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

		parsed := models.Categories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Categories
	parsed, ok := resp.Result.(models.Categories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsFamiliesList
func (srv *Products) ProductsFamiliesList()(*interface{}, error) {
	path := "/v1/products/families"
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
func (srv *Products) WithProductsFamiliesCreateImageAttribute(v string) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.ImageAttribute = v
		o.enabledSetters["ImageAttribute"] = true
	}
}
func (srv *Products) WithProductsFamiliesCreateLabelAttribute(v string) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.LabelAttribute = v
		o.enabledSetters["LabelAttribute"] = true
	}
}
func (srv *Products) WithProductsFamiliesCreateLabels(v interface{}) ProductsFamiliesCreateOption {
	return func(o *ProductsFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamiliesCreate
func (srv *Products) ProductsFamiliesCreate(Code string, optionalSetters ...ProductsFamiliesCreateOption)(*models.Families, error) {
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

		parsed := models.Families{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Families
	parsed, ok := resp.Result.(models.Families)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsFamiliesDelete
func (srv *Products) ProductsFamiliesDelete(Id string)(*interface{}, error) {
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
	
// ProductsFamiliesGet
func (srv *Products) ProductsFamiliesGet(Id string)(*models.Families, error) {
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

		parsed := models.Families{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Families
	parsed, ok := resp.Result.(models.Families)
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
func (srv *Products) WithProductsFamiliesUpdateCode(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsFamiliesUpdateImageAttribute(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.ImageAttribute = v
		o.enabledSetters["ImageAttribute"] = true
	}
}
func (srv *Products) WithProductsFamiliesUpdateLabelAttribute(v string) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.LabelAttribute = v
		o.enabledSetters["LabelAttribute"] = true
	}
}
func (srv *Products) WithProductsFamiliesUpdateLabels(v interface{}) ProductsFamiliesUpdateOption {
	return func(o *ProductsFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamiliesUpdate
func (srv *Products) ProductsFamiliesUpdate(Id string, optionalSetters ...ProductsFamiliesUpdateOption)(*models.Families, error) {
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

		parsed := models.Families{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Families
	parsed, ok := resp.Result.(models.Families)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsFamilyAttributesList
func (srv *Products) ProductsFamilyAttributesList()(*interface{}, error) {
	path := "/v1/products/family_attributes"
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
func (srv *Products) WithProductsFamilyAttributesCreateIsRequired(v bool) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.IsRequired = v
		o.enabledSetters["IsRequired"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesCreatePosition(v int) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesCreateRequiredChannels(v interface{}) ProductsFamilyAttributesCreateOption {
	return func(o *ProductsFamilyAttributesCreateOptions) {
		o.RequiredChannels = v
		o.enabledSetters["RequiredChannels"] = true
	}
}
					
// ProductsFamilyAttributesCreate
func (srv *Products) ProductsFamilyAttributesCreate(AttributeId string, FamilyId string, optionalSetters ...ProductsFamilyAttributesCreateOption)(*models.FamilyAttributes, error) {
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

		parsed := models.FamilyAttributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyAttributes
	parsed, ok := resp.Result.(models.FamilyAttributes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsFamilyAttributesDelete
func (srv *Products) ProductsFamilyAttributesDelete(Id string)(*interface{}, error) {
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
	
// ProductsFamilyAttributesGet
func (srv *Products) ProductsFamilyAttributesGet(Id string)(*models.FamilyAttributes, error) {
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

		parsed := models.FamilyAttributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyAttributes
	parsed, ok := resp.Result.(models.FamilyAttributes)
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
func (srv *Products) WithProductsFamilyAttributesUpdateAttributeId(v string) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.AttributeId = v
		o.enabledSetters["AttributeId"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesUpdateFamilyId(v string) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesUpdateIsRequired(v bool) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.IsRequired = v
		o.enabledSetters["IsRequired"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesUpdatePosition(v int) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsFamilyAttributesUpdateRequiredChannels(v interface{}) ProductsFamilyAttributesUpdateOption {
	return func(o *ProductsFamilyAttributesUpdateOptions) {
		o.RequiredChannels = v
		o.enabledSetters["RequiredChannels"] = true
	}
}
			
// ProductsFamilyAttributesUpdate
func (srv *Products) ProductsFamilyAttributesUpdate(Id string, optionalSetters ...ProductsFamilyAttributesUpdateOption)(*models.FamilyAttributes, error) {
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

		parsed := models.FamilyAttributes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyAttributes
	parsed, ok := resp.Result.(models.FamilyAttributes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsFamilyVariantsList
func (srv *Products) ProductsFamilyVariantsList()(*interface{}, error) {
	path := "/v1/products/family_variants"
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
func (srv *Products) WithProductsFamilyVariantsCreateAxes(v interface{}) ProductsFamilyVariantsCreateOption {
	return func(o *ProductsFamilyVariantsCreateOptions) {
		o.Axes = v
		o.enabledSetters["Axes"] = true
	}
}
func (srv *Products) WithProductsFamilyVariantsCreateLabels(v interface{}) ProductsFamilyVariantsCreateOption {
	return func(o *ProductsFamilyVariantsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
					
// ProductsFamilyVariantsCreate
func (srv *Products) ProductsFamilyVariantsCreate(Code string, FamilyId string, optionalSetters ...ProductsFamilyVariantsCreateOption)(*models.FamilyVariants, error) {
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

		parsed := models.FamilyVariants{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyVariants
	parsed, ok := resp.Result.(models.FamilyVariants)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsFamilyVariantsDelete
func (srv *Products) ProductsFamilyVariantsDelete(Id string)(*interface{}, error) {
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
	
// ProductsFamilyVariantsGet
func (srv *Products) ProductsFamilyVariantsGet(Id string)(*models.FamilyVariants, error) {
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

		parsed := models.FamilyVariants{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyVariants
	parsed, ok := resp.Result.(models.FamilyVariants)
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
func (srv *Products) WithProductsFamilyVariantsUpdateAxes(v interface{}) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Axes = v
		o.enabledSetters["Axes"] = true
	}
}
func (srv *Products) WithProductsFamilyVariantsUpdateCode(v string) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsFamilyVariantsUpdateFamilyId(v string) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.FamilyId = v
		o.enabledSetters["FamilyId"] = true
	}
}
func (srv *Products) WithProductsFamilyVariantsUpdateLabels(v interface{}) ProductsFamilyVariantsUpdateOption {
	return func(o *ProductsFamilyVariantsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsFamilyVariantsUpdate
func (srv *Products) ProductsFamilyVariantsUpdate(Id string, optionalSetters ...ProductsFamilyVariantsUpdateOption)(*models.FamilyVariants, error) {
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

		parsed := models.FamilyVariants{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FamilyVariants
	parsed, ok := resp.Result.(models.FamilyVariants)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsMeasurementFamiliesList
func (srv *Products) ProductsMeasurementFamiliesList()(*interface{}, error) {
	path := "/v1/products/measurement_families"
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
func (srv *Products) WithProductsMeasurementFamiliesCreateLabels(v interface{}) ProductsMeasurementFamiliesCreateOption {
	return func(o *ProductsMeasurementFamiliesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsMeasurementFamiliesCreateUnits(v interface{}) ProductsMeasurementFamiliesCreateOption {
	return func(o *ProductsMeasurementFamiliesCreateOptions) {
		o.Units = v
		o.enabledSetters["Units"] = true
	}
}
					
// ProductsMeasurementFamiliesCreate
func (srv *Products) ProductsMeasurementFamiliesCreate(Code string, StandardUnit string, optionalSetters ...ProductsMeasurementFamiliesCreateOption)(*models.MeasurementFamilies, error) {
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

		parsed := models.MeasurementFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MeasurementFamilies
	parsed, ok := resp.Result.(models.MeasurementFamilies)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsMeasurementFamiliesDelete
func (srv *Products) ProductsMeasurementFamiliesDelete(Id string)(*interface{}, error) {
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
	
// ProductsMeasurementFamiliesGet
func (srv *Products) ProductsMeasurementFamiliesGet(Id string)(*models.MeasurementFamilies, error) {
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

		parsed := models.MeasurementFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MeasurementFamilies
	parsed, ok := resp.Result.(models.MeasurementFamilies)
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
func (srv *Products) WithProductsMeasurementFamiliesUpdateCode(v string) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsMeasurementFamiliesUpdateLabels(v interface{}) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsMeasurementFamiliesUpdateStandardUnit(v string) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.StandardUnit = v
		o.enabledSetters["StandardUnit"] = true
	}
}
func (srv *Products) WithProductsMeasurementFamiliesUpdateUnits(v interface{}) ProductsMeasurementFamiliesUpdateOption {
	return func(o *ProductsMeasurementFamiliesUpdateOptions) {
		o.Units = v
		o.enabledSetters["Units"] = true
	}
}
			
// ProductsMeasurementFamiliesUpdate
func (srv *Products) ProductsMeasurementFamiliesUpdate(Id string, optionalSetters ...ProductsMeasurementFamiliesUpdateOption)(*models.MeasurementFamilies, error) {
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

		parsed := models.MeasurementFamilies{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MeasurementFamilies
	parsed, ok := resp.Result.(models.MeasurementFamilies)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsProductAssociationsList
func (srv *Products) ProductsProductAssociationsList()(*interface{}, error) {
	path := "/v1/products/product_associations"
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
							
// ProductsProductAssociationsCreate
func (srv *Products) ProductsProductAssociationsCreate(AssociationTypeId string, ProductId string, TargetProductId string, optionalSetters ...ProductsProductAssociationsCreateOption)(*models.ProductAssociations, error) {
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

		parsed := models.ProductAssociations{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductAssociations
	parsed, ok := resp.Result.(models.ProductAssociations)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsProductAssociationsDelete
func (srv *Products) ProductsProductAssociationsDelete(Id string)(*interface{}, error) {
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
	
// ProductsProductAssociationsGet
func (srv *Products) ProductsProductAssociationsGet(Id string)(*models.ProductAssociations, error) {
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

		parsed := models.ProductAssociations{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductAssociations
	parsed, ok := resp.Result.(models.ProductAssociations)
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
			
// ProductsProductAssociationsUpdate
func (srv *Products) ProductsProductAssociationsUpdate(Id string, optionalSetters ...ProductsProductAssociationsUpdateOption)(*models.ProductAssociations, error) {
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

		parsed := models.ProductAssociations{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductAssociations
	parsed, ok := resp.Result.(models.ProductAssociations)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsProductCategoriesList
func (srv *Products) ProductsProductCategoriesList()(*interface{}, error) {
	path := "/v1/products/product_categories"
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
type ProductsProductCategoriesCreateOptions struct {
	Position int
	enabledSetters map[string]bool
}
func (options ProductsProductCategoriesCreateOptions) New() *ProductsProductCategoriesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Position": false,
	}
	return &options
}
type ProductsProductCategoriesCreateOption func(*ProductsProductCategoriesCreateOptions)
func (srv *Products) WithProductsProductCategoriesCreatePosition(v int) ProductsProductCategoriesCreateOption {
	return func(o *ProductsProductCategoriesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// ProductsProductCategoriesCreate
func (srv *Products) ProductsProductCategoriesCreate(CategoryId string, ProductId string, optionalSetters ...ProductsProductCategoriesCreateOption)(*models.ProductCategories, error) {
	path := "/v1/products/product_categories"
	options := ProductsProductCategoriesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["category_id"] = CategoryId
	params["product_id"] = ProductId
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

		parsed := models.ProductCategories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductCategories
	parsed, ok := resp.Result.(models.ProductCategories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsProductCategoriesDelete
func (srv *Products) ProductsProductCategoriesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
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
	
// ProductsProductCategoriesGet
func (srv *Products) ProductsProductCategoriesGet(Id string)(*models.ProductCategories, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
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

		parsed := models.ProductCategories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductCategories
	parsed, ok := resp.Result.(models.ProductCategories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ProductsProductCategoriesUpdateOptions struct {
	CategoryId string
	Position int
	ProductId string
	enabledSetters map[string]bool
}
func (options ProductsProductCategoriesUpdateOptions) New() *ProductsProductCategoriesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"CategoryId": false,
		"Position": false,
		"ProductId": false,
	}
	return &options
}
type ProductsProductCategoriesUpdateOption func(*ProductsProductCategoriesUpdateOptions)
func (srv *Products) WithProductsProductCategoriesUpdateCategoryId(v string) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.CategoryId = v
		o.enabledSetters["CategoryId"] = true
	}
}
func (srv *Products) WithProductsProductCategoriesUpdatePosition(v int) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Products) WithProductsProductCategoriesUpdateProductId(v string) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
			
// ProductsProductCategoriesUpdate
func (srv *Products) ProductsProductCategoriesUpdate(Id string, optionalSetters ...ProductsProductCategoriesUpdateOption)(*models.ProductCategories, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
	options := ProductsProductCategoriesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["CategoryId"] {
		params["category_id"] = options.CategoryId
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
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

		parsed := models.ProductCategories{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProductCategories
	parsed, ok := resp.Result.(models.ProductCategories)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsReferenceEntitiesList
func (srv *Products) ProductsReferenceEntitiesList()(*interface{}, error) {
	path := "/v1/products/reference_entities"
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
func (srv *Products) WithProductsReferenceEntitiesCreateImage(v string) ProductsReferenceEntitiesCreateOption {
	return func(o *ProductsReferenceEntitiesCreateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Products) WithProductsReferenceEntitiesCreateLabels(v interface{}) ProductsReferenceEntitiesCreateOption {
	return func(o *ProductsReferenceEntitiesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsReferenceEntitiesCreate
func (srv *Products) ProductsReferenceEntitiesCreate(Code string, optionalSetters ...ProductsReferenceEntitiesCreateOption)(*models.ReferenceEntities, error) {
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

		parsed := models.ReferenceEntities{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntities
	parsed, ok := resp.Result.(models.ReferenceEntities)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsReferenceEntitiesDelete
func (srv *Products) ProductsReferenceEntitiesDelete(Id string)(*interface{}, error) {
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
	
// ProductsReferenceEntitiesGet
func (srv *Products) ProductsReferenceEntitiesGet(Id string)(*models.ReferenceEntities, error) {
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

		parsed := models.ReferenceEntities{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntities
	parsed, ok := resp.Result.(models.ReferenceEntities)
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
func (srv *Products) WithProductsReferenceEntitiesUpdateCode(v string) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsReferenceEntitiesUpdateImage(v string) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Products) WithProductsReferenceEntitiesUpdateLabels(v interface{}) ProductsReferenceEntitiesUpdateOption {
	return func(o *ProductsReferenceEntitiesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
			
// ProductsReferenceEntitiesUpdate
func (srv *Products) ProductsReferenceEntitiesUpdate(Id string, optionalSetters ...ProductsReferenceEntitiesUpdateOption)(*models.ReferenceEntities, error) {
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

		parsed := models.ReferenceEntities{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntities
	parsed, ok := resp.Result.(models.ReferenceEntities)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ProductsReferenceEntityRecordsList
func (srv *Products) ProductsReferenceEntityRecordsList()(*interface{}, error) {
	path := "/v1/products/reference_entity_records"
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
func (srv *Products) WithProductsReferenceEntityRecordsCreateAttributeValues(v interface{}) ProductsReferenceEntityRecordsCreateOption {
	return func(o *ProductsReferenceEntityRecordsCreateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsReferenceEntityRecordsCreateLabels(v interface{}) ProductsReferenceEntityRecordsCreateOption {
	return func(o *ProductsReferenceEntityRecordsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
					
// ProductsReferenceEntityRecordsCreate
func (srv *Products) ProductsReferenceEntityRecordsCreate(Code string, ReferenceEntityId string, optionalSetters ...ProductsReferenceEntityRecordsCreateOption)(*models.ReferenceEntityRecords, error) {
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

		parsed := models.ReferenceEntityRecords{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntityRecords
	parsed, ok := resp.Result.(models.ReferenceEntityRecords)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsReferenceEntityRecordsDelete
func (srv *Products) ProductsReferenceEntityRecordsDelete(Id string)(*interface{}, error) {
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
	
// ProductsReferenceEntityRecordsGet
func (srv *Products) ProductsReferenceEntityRecordsGet(Id string)(*models.ReferenceEntityRecords, error) {
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

		parsed := models.ReferenceEntityRecords{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntityRecords
	parsed, ok := resp.Result.(models.ReferenceEntityRecords)
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
func (srv *Products) WithProductsReferenceEntityRecordsUpdateAttributeValues(v interface{}) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.AttributeValues = v
		o.enabledSetters["AttributeValues"] = true
	}
}
func (srv *Products) WithProductsReferenceEntityRecordsUpdateCode(v string) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Products) WithProductsReferenceEntityRecordsUpdateLabels(v interface{}) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Products) WithProductsReferenceEntityRecordsUpdateReferenceEntityId(v string) ProductsReferenceEntityRecordsUpdateOption {
	return func(o *ProductsReferenceEntityRecordsUpdateOptions) {
		o.ReferenceEntityId = v
		o.enabledSetters["ReferenceEntityId"] = true
	}
}
			
// ProductsReferenceEntityRecordsUpdate
func (srv *Products) ProductsReferenceEntityRecordsUpdate(Id string, optionalSetters ...ProductsReferenceEntityRecordsUpdateOption)(*models.ReferenceEntityRecords, error) {
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

		parsed := models.ReferenceEntityRecords{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReferenceEntityRecords
	parsed, ok := resp.Result.(models.ReferenceEntityRecords)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ProductsDelete
func (srv *Products) ProductsDelete(Id string)(*interface{}, error) {
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
	
// ProductsGet
func (srv *Products) ProductsGet(Id string)(*models.Products, error) {
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

		parsed := models.Products{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Products
	parsed, ok := resp.Result.(models.Products)
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
			
// ProductsUpdate
func (srv *Products) ProductsUpdate(Id string, optionalSetters ...ProductsUpdateOption)(*models.Products, error) {
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

		parsed := models.Products{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Products
	parsed, ok := resp.Result.(models.Products)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
