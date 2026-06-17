package prices

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Prices service
type Prices struct {
	client client.Client
}

func New(clt client.Client) *Prices {
	return &Prices{
		client: clt,
	}
}


// PricesListsList
func (srv *Prices) PricesListsList()(*interface{}, error) {
	path := "/v1/prices/lists"
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
type PricesListsCreateOptions struct {
	ChannelId string
	ContactId string
	Currency string
	Description string
	IsDefault bool
	Labels interface{}
	MarketId string
	Metadata interface{}
	OrganizationId string
	Priority int
	Status string
	TaxIncluded bool
	ValidFrom string
	ValidUntil string
	enabledSetters map[string]bool
}
func (options PricesListsCreateOptions) New() *PricesListsCreateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"ContactId": false,
		"Currency": false,
		"Description": false,
		"IsDefault": false,
		"Labels": false,
		"MarketId": false,
		"Metadata": false,
		"OrganizationId": false,
		"Priority": false,
		"Status": false,
		"TaxIncluded": false,
		"ValidFrom": false,
		"ValidUntil": false,
	}
	return &options
}
type PricesListsCreateOption func(*PricesListsCreateOptions)
func (srv *Prices) WithPricesListsCreateChannelId(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Prices) WithPricesListsCreateContactId(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Prices) WithPricesListsCreateCurrency(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Prices) WithPricesListsCreateDescription(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Prices) WithPricesListsCreateIsDefault(v bool) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Prices) WithPricesListsCreateLabels(v interface{}) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Prices) WithPricesListsCreateMarketId(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Prices) WithPricesListsCreateMetadata(v interface{}) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Prices) WithPricesListsCreateOrganizationId(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Prices) WithPricesListsCreatePriority(v int) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Prices) WithPricesListsCreateStatus(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Prices) WithPricesListsCreateTaxIncluded(v bool) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.TaxIncluded = v
		o.enabledSetters["TaxIncluded"] = true
	}
}
func (srv *Prices) WithPricesListsCreateValidFrom(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesListsCreateValidUntil(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
					
// PricesListsCreate
func (srv *Prices) PricesListsCreate(Code string, Name string, optionalSetters ...PricesListsCreateOption)(*models.PriceList, error) {
	path := "/v1/prices/lists"
	options := PricesListsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MarketId"] {
		params["market_id"] = options.MarketId
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TaxIncluded"] {
		params["tax_included"] = options.TaxIncluded
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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

		parsed := models.PriceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceList
	parsed, ok := resp.Result.(models.PriceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PricesListsDefaults
func (srv *Prices) PricesListsDefaults()(*interface{}, error) {
	path := "/v1/prices/lists/defaults"
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
	
// PricesListsDelete
func (srv *Prices) PricesListsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/prices/lists/{id}")
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
	
// PricesListsGet
func (srv *Prices) PricesListsGet(Id string)(*models.PriceList, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/prices/lists/{id}")
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

		parsed := models.PriceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceList
	parsed, ok := resp.Result.(models.PriceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PricesListsUpdateOptions struct {
	ChannelId string
	Code string
	ContactId string
	Currency string
	Description string
	IsDefault bool
	Labels interface{}
	MarketId string
	Metadata interface{}
	Name string
	OrganizationId string
	Priority int
	Status string
	TaxIncluded bool
	ValidFrom string
	ValidUntil string
	enabledSetters map[string]bool
}
func (options PricesListsUpdateOptions) New() *PricesListsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"Code": false,
		"ContactId": false,
		"Currency": false,
		"Description": false,
		"IsDefault": false,
		"Labels": false,
		"MarketId": false,
		"Metadata": false,
		"Name": false,
		"OrganizationId": false,
		"Priority": false,
		"Status": false,
		"TaxIncluded": false,
		"ValidFrom": false,
		"ValidUntil": false,
	}
	return &options
}
type PricesListsUpdateOption func(*PricesListsUpdateOptions)
func (srv *Prices) WithPricesListsUpdateChannelId(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateCode(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateContactId(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateCurrency(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateDescription(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateIsDefault(v bool) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateLabels(v interface{}) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateMarketId(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateMetadata(v interface{}) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateName(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateOrganizationId(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Prices) WithPricesListsUpdatePriority(v int) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateStatus(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateTaxIncluded(v bool) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.TaxIncluded = v
		o.enabledSetters["TaxIncluded"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateValidFrom(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateValidUntil(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
			
// PricesListsUpdate
func (srv *Prices) PricesListsUpdate(Id string, optionalSetters ...PricesListsUpdateOption)(*models.PriceList, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/prices/lists/{id}")
	options := PricesListsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["MarketId"] {
		params["market_id"] = options.MarketId
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TaxIncluded"] {
		params["tax_included"] = options.TaxIncluded
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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

		parsed := models.PriceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceList
	parsed, ok := resp.Result.(models.PriceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PricesEntriesList
func (srv *Prices) PricesEntriesList(ListId string)(*interface{}, error) {
	r := strings.NewReplacer("{listId}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries")
	params := map[string]interface{}{}
	params["list_id"] = ListId
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
type PricesEntriesCreateOptions struct {
	Metadata interface{}
	PriceType string
	ProductId string
	QuantityMin float64
	Sku string
	Unit string
	UnitPrice float64
	ValidFrom string
	ValidUntil string
	enabledSetters map[string]bool
}
func (options PricesEntriesCreateOptions) New() *PricesEntriesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"PriceType": false,
		"ProductId": false,
		"QuantityMin": false,
		"Sku": false,
		"Unit": false,
		"UnitPrice": false,
		"ValidFrom": false,
		"ValidUntil": false,
	}
	return &options
}
type PricesEntriesCreateOption func(*PricesEntriesCreateOptions)
func (srv *Prices) WithPricesEntriesCreateMetadata(v interface{}) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreatePriceType(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.PriceType = v
		o.enabledSetters["PriceType"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateProductId(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateQuantityMin(v float64) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.QuantityMin = v
		o.enabledSetters["QuantityMin"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateSku(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateUnit(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateUnitPrice(v float64) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateValidFrom(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesEntriesCreateValidUntil(v string) PricesEntriesCreateOption {
	return func(o *PricesEntriesCreateOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
			
// PricesEntriesCreate
func (srv *Prices) PricesEntriesCreate(ListId string, optionalSetters ...PricesEntriesCreateOption)(*models.PriceEntry, error) {
	r := strings.NewReplacer("{listId}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries")
	options := PricesEntriesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["PriceType"] {
		params["price_type"] = options.PriceType
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["QuantityMin"] {
		params["quantity_min"] = options.QuantityMin
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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

		parsed := models.PriceEntry{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceEntry
	parsed, ok := resp.Result.(models.PriceEntry)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// PricesEntriesReplace
func (srv *Prices) PricesEntriesReplace(ListId string, Entries []models.PriceEntryReplaceItem)(*interface{}, error) {
	r := strings.NewReplacer("{listId}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries")
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["entries"] = Entries
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
			
// PricesEntriesBulk
func (srv *Prices) PricesEntriesBulk(ListId string, Entries []models.PriceEntryReplaceItem)(*interface{}, error) {
	r := strings.NewReplacer("{listId}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/bulk")
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["entries"] = Entries
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
			
// PricesEntriesDelete
func (srv *Prices) PricesEntriesDelete(ListId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{listId}", ListId, "{id}", Id)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/{id}")
	params := map[string]interface{}{}
	params["list_id"] = ListId
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
			
// PricesEntriesGet
func (srv *Prices) PricesEntriesGet(ListId string, Id string)(*models.PriceEntry, error) {
	r := strings.NewReplacer("{listId}", ListId, "{id}", Id)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/{id}")
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PriceEntry{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceEntry
	parsed, ok := resp.Result.(models.PriceEntry)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PricesEntriesUpdateOptions struct {
	Metadata interface{}
	PriceType string
	ProductId string
	QuantityMin float64
	Sku string
	Unit string
	UnitPrice float64
	ValidFrom string
	ValidUntil string
	enabledSetters map[string]bool
}
func (options PricesEntriesUpdateOptions) New() *PricesEntriesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Metadata": false,
		"PriceType": false,
		"ProductId": false,
		"QuantityMin": false,
		"Sku": false,
		"Unit": false,
		"UnitPrice": false,
		"ValidFrom": false,
		"ValidUntil": false,
	}
	return &options
}
type PricesEntriesUpdateOption func(*PricesEntriesUpdateOptions)
func (srv *Prices) WithPricesEntriesUpdateMetadata(v interface{}) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdatePriceType(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.PriceType = v
		o.enabledSetters["PriceType"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateProductId(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateQuantityMin(v float64) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.QuantityMin = v
		o.enabledSetters["QuantityMin"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateSku(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateUnit(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateUnitPrice(v float64) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateValidFrom(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesEntriesUpdateValidUntil(v string) PricesEntriesUpdateOption {
	return func(o *PricesEntriesUpdateOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
					
// PricesEntriesUpdate
func (srv *Prices) PricesEntriesUpdate(ListId string, Id string, optionalSetters ...PricesEntriesUpdateOption)(*models.PriceEntry, error) {
	r := strings.NewReplacer("{listId}", ListId, "{id}", Id)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/{id}")
	options := PricesEntriesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["id"] = Id
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["PriceType"] {
		params["price_type"] = options.PriceType
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["QuantityMin"] {
		params["quantity_min"] = options.QuantityMin
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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

		parsed := models.PriceEntry{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceEntry
	parsed, ok := resp.Result.(models.PriceEntry)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PricesResolveOptions struct {
	At string
	ChannelId string
	ContactId string
	Currency string
	MarketId string
	OrganizationId string
	enabledSetters map[string]bool
}
func (options PricesResolveOptions) New() *PricesResolveOptions {
	options.enabledSetters = map[string]bool{
		"At": false,
		"ChannelId": false,
		"ContactId": false,
		"Currency": false,
		"MarketId": false,
		"OrganizationId": false,
	}
	return &options
}
type PricesResolveOption func(*PricesResolveOptions)
func (srv *Prices) WithPricesResolveAt(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.At = v
		o.enabledSetters["At"] = true
	}
}
func (srv *Prices) WithPricesResolveChannelId(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Prices) WithPricesResolveContactId(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Prices) WithPricesResolveCurrency(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Prices) WithPricesResolveMarketId(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Prices) WithPricesResolveOrganizationId(v string) PricesResolveOption {
	return func(o *PricesResolveOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
			
// PricesResolve
func (srv *Prices) PricesResolve(Items []models.PriceResolveItem, optionalSetters ...PricesResolveOption)(*interface{}, error) {
	path := "/v1/prices/resolve"
	options := PricesResolveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["At"] {
		params["at"] = options.At
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["MarketId"] {
		params["market_id"] = options.MarketId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
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
