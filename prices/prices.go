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

type PricesListsListOptions struct {
	Id string
	Code string
	Name string
	Description string
	Currency string
	Status string
	Priority int
	IsDefault bool
	TaxBasis string
	TaxIncluded bool
	RequiresAuth bool
	ContactId string
	OrganizationId string
	ChannelId string
	ValidFrom string
	ValidUntil string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options PricesListsListOptions) New() *PricesListsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Name": false,
		"Description": false,
		"Currency": false,
		"Status": false,
		"Priority": false,
		"IsDefault": false,
		"TaxBasis": false,
		"TaxIncluded": false,
		"RequiresAuth": false,
		"ContactId": false,
		"OrganizationId": false,
		"ChannelId": false,
		"ValidFrom": false,
		"ValidUntil": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type PricesListsListOption func(*PricesListsListOptions)
func (srv *Prices) WithPricesListsListId(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Prices) WithPricesListsListCode(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Prices) WithPricesListsListName(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Prices) WithPricesListsListDescription(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Prices) WithPricesListsListCurrency(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Prices) WithPricesListsListStatus(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Prices) WithPricesListsListPriority(v int) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Prices) WithPricesListsListIsDefault(v bool) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Prices) WithPricesListsListTaxBasis(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.TaxBasis = v
		o.enabledSetters["TaxBasis"] = true
	}
}
func (srv *Prices) WithPricesListsListTaxIncluded(v bool) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.TaxIncluded = v
		o.enabledSetters["TaxIncluded"] = true
	}
}
func (srv *Prices) WithPricesListsListRequiresAuth(v bool) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.RequiresAuth = v
		o.enabledSetters["RequiresAuth"] = true
	}
}
func (srv *Prices) WithPricesListsListContactId(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Prices) WithPricesListsListOrganizationId(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Prices) WithPricesListsListChannelId(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Prices) WithPricesListsListValidFrom(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesListsListValidUntil(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
func (srv *Prices) WithPricesListsListCreatedAt(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Prices) WithPricesListsListUpdatedAt(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Prices) WithPricesListsListLimit(v int) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Prices) WithPricesListsListOffset(v int) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Prices) WithPricesListsListOrder(v string) PricesListsListOption {
	return func(o *PricesListsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// PricesListsList one page of the tenant's price list HEADERS — code,
// currency, tax basis, status, priority, validity window, buyer scope and the
// default flag. Never the prices themselves: those are a separate page per
// list (`GET /prices/lists/{list_id}/entries`).
// 
// Every filter is an EXACT match on a column, ANDed together; a query key
// that is not a column is dropped in silence, which is why the answer echoes
// `filter`. The scope, currency and status filters are the useful ones,
// because between them they narrow the set to the candidates a resolve call
// in a given currency for a given buyer can draw on at all.
// 
// Market is deliberately not among them: a list is scoped to a market by an
// assignment, not a column, and the `X-Revenexx-Market` header is what
// narrows the set — this admin listing shows the tenant's lists whatever
// their market.
func (srv *Prices) PricesListsList(optionalSetters ...PricesListsListOption)(*models.Error, error) {
	path := "/v1/prices/lists"
	options := PricesListsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["TaxBasis"] {
		params["tax_basis"] = options.TaxBasis
	}
	if options.enabledSetters["TaxIncluded"] {
		params["tax_included"] = options.TaxIncluded
	}
	if options.enabledSetters["RequiresAuth"] {
		params["requires_auth"] = options.RequiresAuth
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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
type PricesListsCreateOptions struct {
	ChannelId string
	ContactId string
	Currency string
	Description string
	IsDefault bool
	Labels interface{}
	Metadata interface{}
	OrganizationId string
	Priority int
	RequiresAuth bool
	Status string
	TaxBasis string
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
		"Metadata": false,
		"OrganizationId": false,
		"Priority": false,
		"RequiresAuth": false,
		"Status": false,
		"TaxBasis": false,
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
func (srv *Prices) WithPricesListsCreateRequiresAuth(v bool) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.RequiresAuth = v
		o.enabledSetters["RequiresAuth"] = true
	}
}
func (srv *Prices) WithPricesListsCreateStatus(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Prices) WithPricesListsCreateTaxBasis(v string) PricesListsCreateOption {
	return func(o *PricesListsCreateOptions) {
		o.TaxBasis = v
		o.enabledSetters["TaxBasis"] = true
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
					
// PricesListsCreate opens an empty book, and states in one row the four
// things that decide whether it will ever price anything: its currency, its
// priority within a specificity group, its validity window, and its buyer
// scope (contact, organization or channel — leave all three empty for a
// list open to everyone).
// 
// `code` and `name` are the only fields required — they are the two columns
// with no default — and `code` is unique per tenant, so a code already in
// use is a 409 rather than an overwrite of prices somebody is selling on.
// 
// Everything else has a default, and two of them are worth choosing rather
// than accepting. `currency` defaults to EUR and is the currency of every
// amount in the list, since entries carry none; a resolve call only considers
// lists in the currency it is asked about, and nothing is ever converted.
// `tax_basis` defaults to NOTHING, which means the amounts inherit the
// tenant's `tax_inclusive_default` — state net or gross here and the answer
// stops depending on a tenant setting somebody may change later.
// 
// `is_default: true` here does NOT demote the list that currently holds the
// flag: you end up with two defaults, and which of them prices an item is
// left to the tenant's tie-break. Create the list, then move the flag with
// `POST /prices/lists/{list_id}/make-default`.
// 
// A new list prices nothing at all until it has entries, so it is inert until
// you add them — which makes it safe to create one ahead of the prices that
// will fill it.
func (srv *Prices) PricesListsCreate(Code string, Name string, optionalSetters ...PricesListsCreateOption)(*models.Error, error) {
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
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["RequiresAuth"] {
		params["requires_auth"] = options.RequiresAuth
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TaxBasis"] {
		params["tax_basis"] = options.TaxBasis
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

// PricesListsDefaults gives a tenant the one open list every tenant needs, so
// nothing has to exist before the first price can be written. Almost nobody
// calls it: the app runs it by itself on `app.installed`, and the route is
// the manual re-run — for a tenant installed before that hook existed, or
// one whose standard list was deleted. Because it is idempotent it is also
// safe to call from a provisioning script that cannot know which of the two
// is the case.
// 
// What it writes comes from settings, not from constants: the code is the
// tenant's `default_price_list_code`, the currency its `default_currency`,
// and the seeded list STATES its tax basis from `tax_inclusive_default`
// instead of inheriting it, because the one list every tenant gets should not
// be the ambiguous one.
// 
// Idempotent twice over — by that code, and by the existence of ANY default
// list. So calling it repeatedly is free, changing `default_price_list_code`
// later never produces a second list, and a tenant that has made some other
// list the default is left exactly as it is (the answer names that list under
// `existing`). It writes nothing else: it never demotes, never touches
// entries, and never repairs a list that is already there.
func (srv *Prices) PricesListsDefaults()(*models.PriceListDefaultsResponse, error) {
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

		parsed := models.PriceListDefaultsResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceListDefaultsResponse
	parsed, ok := resp.Result.(models.PriceListDefaultsResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PricesListsDelete deletes the list AND every price in it.
// `price_entries.price_list_id` references this row ON DELETE CASCADE, so the
// entries go in the same statement: nothing asks, nothing blocks, a book of
// 40 000 prices deletes exactly as fast as an empty one, and the answer is a
// bare `{deleted, id}` that never says how many prices went with it.
// 
// What that means while a storefront is quoting: from the next resolve call
// the items this list priced fall through to the next candidate list, and
// where there is none the answer is `on_request` — "price on request" for
// something that had a price a second ago, never €0. If the deleted list
// held the default flag the tenant has no default until one is moved onto
// another list; re-running `POST /prices/lists/defaults` recreates the
// standard list only while no other default exists.
// 
// This is not the way to take a list out of circulation. `status: "inactive"`
// does that immediately and reversibly and keeps the prices; deleting is for
// a list whose contents you are prepared to import again, because nothing
// here is recoverable.
func (srv *Prices) PricesListsDelete(Id string)(*models.Error, error) {
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
	
// PricesListsGet the list HEADER, never its prices: currency, tax basis,
// buyer scope, priority, validity window and the default flag — the
// settings that decide WHETHER this list prices a given buyer, before any
// amount is looked at. Its entries are a separate page (`GET
// /prices/lists/{list_id}/entries`), because a price book runs to thousands
// of rows and no read of a list should carry them. This is the admin view and
// it reads the base table rather than the market-scoped one the resolve call
// uses, so a list that is invisible in the active market is still returned
// here.
func (srv *Prices) PricesListsGet(Id string)(*models.Error, error) {
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
type PricesListsUpdateOptions struct {
	ChannelId string
	Code string
	ContactId string
	Currency string
	Description string
	IsDefault bool
	Labels interface{}
	Metadata interface{}
	Name string
	OrganizationId string
	Priority int
	RequiresAuth bool
	Status string
	TaxBasis string
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
		"Metadata": false,
		"Name": false,
		"OrganizationId": false,
		"Priority": false,
		"RequiresAuth": false,
		"Status": false,
		"TaxBasis": false,
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
func (srv *Prices) WithPricesListsUpdateRequiresAuth(v bool) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.RequiresAuth = v
		o.enabledSetters["RequiresAuth"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateStatus(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Prices) WithPricesListsUpdateTaxBasis(v string) PricesListsUpdateOption {
	return func(o *PricesListsUpdateOptions) {
		o.TaxBasis = v
		o.enabledSetters["TaxBasis"] = true
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
			
// PricesListsUpdate a partial update: send only what changes, omitted fields
// keep their value, and a payload with no updatable column at all is refused
// rather than answered with an unchanged row. There is no draft and no
// publish step — the next resolve call reads what this one wrote.
// 
// Three edits do more than their field names suggest. `currency`
// re-denominates without converting: entries carry no currency of their own,
// so 19.90 EUR becomes 19.90 CHF and the whole book is re-priced by one edit.
// `status: "inactive"` takes the list out of every quote immediately while
// keeping its prices — the reversible way to stop selling on a list, and
// the one to reach for instead of deleting it. `code` is the handle imports
// and integrations address the list by, and a code another list already holds
// is a 409.
// 
// `is_default` behaves here exactly as it does on create: setting it true
// leaves the incumbent default in place, so use `POST
// /prices/lists/{list_id}/make-default`, which demotes in the same call.
func (srv *Prices) PricesListsUpdate(Id string, optionalSetters ...PricesListsUpdateOption)(*models.Error, error) {
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
	if options.enabledSetters["RequiresAuth"] {
		params["requires_auth"] = options.RequiresAuth
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TaxBasis"] {
		params["tax_basis"] = options.TaxBasis
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
type PricesEntriesListOptions struct {
	Id string
	ProductId string
	Sku string
	PriceType string
	QuantityMin float64
	UnitPrice float64
	Unit string
	ValidFrom string
	ValidUntil string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options PricesEntriesListOptions) New() *PricesEntriesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"ProductId": false,
		"Sku": false,
		"PriceType": false,
		"QuantityMin": false,
		"UnitPrice": false,
		"Unit": false,
		"ValidFrom": false,
		"ValidUntil": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type PricesEntriesListOption func(*PricesEntriesListOptions)
func (srv *Prices) WithPricesEntriesListId(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Prices) WithPricesEntriesListProductId(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Prices) WithPricesEntriesListSku(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Prices) WithPricesEntriesListPriceType(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.PriceType = v
		o.enabledSetters["PriceType"] = true
	}
}
func (srv *Prices) WithPricesEntriesListQuantityMin(v float64) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.QuantityMin = v
		o.enabledSetters["QuantityMin"] = true
	}
}
func (srv *Prices) WithPricesEntriesListUnitPrice(v float64) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
func (srv *Prices) WithPricesEntriesListUnit(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Prices) WithPricesEntriesListValidFrom(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Prices) WithPricesEntriesListValidUntil(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
func (srv *Prices) WithPricesEntriesListCreatedAt(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Prices) WithPricesEntriesListUpdatedAt(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Prices) WithPricesEntriesListLimit(v int) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Prices) WithPricesEntriesListOffset(v int) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Prices) WithPricesEntriesListOrder(v string) PricesEntriesListOption {
	return func(o *PricesEntriesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// PricesEntriesList the prices inside one list, a page at a time. An entry is
// a rung rather than "the price of a product": it carries a quantity
// threshold, an amount and a unit, its own validity window, and — where the
// answer is deliberately no number at all — an `on_request` marker instead
// of one. So this page is where the quantity tiers, the promo windows and the
// "ask us" markers of a book are read.
// 
// The ladder of one item is the set of entries sharing an identity, so
// `?product_id=…` (or `?sku=…`) is how a caller reads the Staffel a
// resolve answer was built from, and `?price_type=on_request` is how the
// markers are audited. The response also carries `page` and `filter` like
// every other list, and an unknown list_id answers 404 instead of an empty
// page.
func (srv *Prices) PricesEntriesList(ListId string, optionalSetters ...PricesEntriesListOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries")
	options := PricesEntriesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["PriceType"] {
		params["price_type"] = options.PriceType
	}
	if options.enabledSetters["QuantityMin"] {
		params["quantity_min"] = options.QuantityMin
	}
	if options.enabledSetters["UnitPrice"] {
		params["unit_price"] = options.UnitPrice
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
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
			
// PricesEntriesCreate adds ONE rung to one item's quantity ladder in this
// list. The only thing an entry must have is an identity — `product_id` or
// `sku`, which the row CHECK enforces; everything else defaults, and one of
// those defaults deserves a warning.
// 
// `unit_price` defaults to **0**. That is the one door through which a zero
// price enters an app whose whole doctrine is that a missing price is
// `on_request` and never €0: a create that forgets the amount publishes a
// free item, and the storefront shows 0.00 instead of "price on request".
// Send the amount, or send `price_type: "on_request"` where there genuinely
// is none. The amount is per ONE unit of `unit`, in the LIST's currency
// (entries carry none) and on the LIST's tax basis, as a decimal in major
// units — 19.90, never 1990.
// 
// Nothing enforces one rung per (item, quantity): create the same
// `quantity_min` twice and both rows come back in the resolved `tiers`, with
// the last of them setting the price — an ambiguous ladder no error ever
// mentions. `quantity_min` defaults to 1 and `price_type` to `standard`.
// 
// This route is for a rung at a time. A whole ladder in one call is `POST
// …/entries/ladder`, an import is `POST …/entries/bulk`, and a complete
// rewrite of the book is `PUT …/entries`. An unknown `list_id` answers 404
// rather than attaching a price to nothing.
func (srv *Prices) PricesEntriesCreate(ListId string, optionalSetters ...PricesEntriesCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
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
			
// PricesEntriesReplace set semantics over the WHOLE list, not over one item:
// every entry of the list is deleted and the payload becomes the complete new
// book. It exists for the two callers that genuinely hold the whole book in
// hand — the Cockpit's table editor, whose save is this call, and a small
// import. `entries: []` is a legal payload and empties the list — the items
// it priced then resolve from the next candidate list, or come back
// `on_request`.
// 
// Two consequences of "delete, then insert". Every row is inserted fresh, so
// all entry ids change and anything holding one is stale afterwards. And it
// is not a transaction: the deletes go out before the inserts, so a payload
// that fails part-way through leaves the list holding the rows that landed
// and none of the ones it had. What protects you is that the whole payload is
// normalized and validated BEFORE the first delete — a malformed row is a
// 400 with the list untouched.
// 
// For a book of any size, or for adding to one you want to keep, use `POST
// …/entries/bulk`: it upserts in chunks and never wipes.
func (srv *Prices) PricesEntriesReplace(ListId string, Entries []models.PriceEntryReplaceItem)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
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
type PricesEntriesAdjustOptions struct {
	Amount float64
	DryRun bool
	Percent float64
	Rounding string
	SkuPrefix string
	enabledSetters map[string]bool
}
func (options PricesEntriesAdjustOptions) New() *PricesEntriesAdjustOptions {
	options.enabledSetters = map[string]bool{
		"Amount": false,
		"DryRun": false,
		"Percent": false,
		"Rounding": false,
		"SkuPrefix": false,
	}
	return &options
}
type PricesEntriesAdjustOption func(*PricesEntriesAdjustOptions)
func (srv *Prices) WithPricesEntriesAdjustAmount(v float64) PricesEntriesAdjustOption {
	return func(o *PricesEntriesAdjustOptions) {
		o.Amount = v
		o.enabledSetters["Amount"] = true
	}
}
func (srv *Prices) WithPricesEntriesAdjustDryRun(v bool) PricesEntriesAdjustOption {
	return func(o *PricesEntriesAdjustOptions) {
		o.DryRun = v
		o.enabledSetters["DryRun"] = true
	}
}
func (srv *Prices) WithPricesEntriesAdjustPercent(v float64) PricesEntriesAdjustOption {
	return func(o *PricesEntriesAdjustOptions) {
		o.Percent = v
		o.enabledSetters["Percent"] = true
	}
}
func (srv *Prices) WithPricesEntriesAdjustRounding(v string) PricesEntriesAdjustOption {
	return func(o *PricesEntriesAdjustOptions) {
		o.Rounding = v
		o.enabledSetters["Rounding"] = true
	}
}
func (srv *Prices) WithPricesEntriesAdjustSkuPrefix(v string) PricesEntriesAdjustOption {
	return func(o *PricesEntriesAdjustOptions) {
		o.SkuPrefix = v
		o.enabledSetters["SkuPrefix"] = true
	}
}
			
// PricesEntriesAdjust moves every priced entry of the list at once, in
// whichever of the two ways a merchant thinks about a price change: `percent`
// for a relative one (5 raises everything by 5 %) or `amount` for a flat one
// added to every unit price. One or the other, never both, and `sku_prefix`
// narrows the change to part of the book. On-request entries are never
// touched, because a percentage of "ask us" is not a number.
// 
// The other half of a bulk change is what the arithmetic leaves behind: a 7 %
// increase turns 19.90 into 21.293, which no merchant prints. Results are
// therefore rounded to the tenant's price_precision/rounding_mode and then
// snapped to a declared merchant price ending — x.99, x.95, a whole number
// — either the one this call names or the tenant's `bulk_adjust_rounding`.
// dry_run answers the same preview and writes nothing, which is what the
// Cockpit dialog shows before it commits.
func (srv *Prices) PricesEntriesAdjust(ListId string, optionalSetters ...PricesEntriesAdjustOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/adjust")
	options := PricesEntriesAdjustOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	if options.enabledSetters["Amount"] {
		params["amount"] = options.Amount
	}
	if options.enabledSetters["DryRun"] {
		params["dry_run"] = options.DryRun
	}
	if options.enabledSetters["Percent"] {
		params["percent"] = options.Percent
	}
	if options.enabledSetters["Rounding"] {
		params["rounding"] = options.Rounding
	}
	if options.enabledSetters["SkuPrefix"] {
		params["sku_prefix"] = options.SkuPrefix
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
type PricesEntriesBulkOptions struct {
	Mode string
	enabledSetters map[string]bool
}
func (options PricesEntriesBulkOptions) New() *PricesEntriesBulkOptions {
	options.enabledSetters = map[string]bool{
		"Mode": false,
	}
	return &options
}
type PricesEntriesBulkOption func(*PricesEntriesBulkOptions)
func (srv *Prices) WithPricesEntriesBulkMode(v string) PricesEntriesBulkOption {
	return func(o *PricesEntriesBulkOptions) {
		o.Mode = v
		o.enabledSetters["Mode"] = true
	}
}
					
// PricesEntriesBulk adds entries to a list without wiping it, and UPSERTS
// rather than inserts: a row naming a rung the list already has (same
// product_id/sku AND quantity_min) updates that rung, so re-running an import
// corrects prices instead of duplicating the ladder. `mode: 'append'` keeps
// the old insert-everything behaviour. Inserts go out as one PostgREST bulk
// write per 1000 rows.
// 
// This is the route for a large price book, and a large book arrives in
// chunks: a call carries at most 5000 entries and a longer payload is refused
// with 400 rather than truncated, so an importer of 200 000 prices sends
// forty calls. Because the upsert is keyed on the rung rather than on a row
// id, the chunks may be re-sent and re-ordered freely — a chunk that lands
// twice writes the same prices twice.
func (srv *Prices) PricesEntriesBulk(ListId string, Entries []models.PriceEntryReplaceItem, optionalSetters ...PricesEntriesBulkOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/bulk")
	options := PricesEntriesBulkOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["entries"] = Entries
	if options.enabledSetters["Mode"] {
		params["mode"] = options.Mode
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
type PricesEntriesLadderOptions struct {
	DiscountPercent float64
	ProductId string
	Quantities []float64
	Replace bool
	Rounding string
	Sku string
	Unit string
	enabledSetters map[string]bool
}
func (options PricesEntriesLadderOptions) New() *PricesEntriesLadderOptions {
	options.enabledSetters = map[string]bool{
		"DiscountPercent": false,
		"ProductId": false,
		"Quantities": false,
		"Replace": false,
		"Rounding": false,
		"Sku": false,
		"Unit": false,
	}
	return &options
}
type PricesEntriesLadderOption func(*PricesEntriesLadderOptions)
func (srv *Prices) WithPricesEntriesLadderDiscountPercent(v float64) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.DiscountPercent = v
		o.enabledSetters["DiscountPercent"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderProductId(v string) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderQuantities(v []float64) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.Quantities = v
		o.enabledSetters["Quantities"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderReplace(v bool) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.Replace = v
		o.enabledSetters["Replace"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderRounding(v string) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.Rounding = v
		o.enabledSetters["Rounding"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderSku(v string) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Prices) WithPricesEntriesLadderUnit(v string) PricesEntriesLadderOption {
	return func(o *PricesEntriesLadderOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
					
// PricesEntriesLadder writes a whole quantity-tier ladder (Staffelpreise) for
// ONE item in one call, instead of typing a rung at a time. Tiers are a flat
// quantity_min column on purpose — the ladder IS the set of entries sharing
// an identity, and resolve returns it sorted as one array. What was missing
// was the gesture: "19.90 from 1, 5 % off per tier at 10 and 50". Prices are
// rounded and snapped exactly as a bulk adjust is.
func (srv *Prices) PricesEntriesLadder(ListId string, BasePrice float64, optionalSetters ...PricesEntriesLadderOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/entries/ladder")
	options := PricesEntriesLadderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["base_price"] = BasePrice
	if options.enabledSetters["DiscountPercent"] {
		params["discount_percent"] = options.DiscountPercent
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantities"] {
		params["quantities"] = options.Quantities
	}
	if options.enabledSetters["Replace"] {
		params["replace"] = options.Replace
	}
	if options.enabledSetters["Rounding"] {
		params["rounding"] = options.Rounding
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
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
			
// PricesEntriesDelete removes ONE rung. The item keeps its other rungs and
// stays priced — which is exactly what makes the lowest rung the dangerous
// one to delete.
// 
// Below the first threshold the FIRST rung's price applies (a minimum
// quantity belongs to the catalog, not to the price ladder). So deleting the
// "from 1" rung of a 1/10/50 ladder does not make single units unpriced: it
// sells them at the 10-up volume price, silently, from the next resolve call
// onwards. Nothing in the answer marks that the ladder no longer starts where
// it used to.
// 
// Delete an item's LAST rung and this list stops pricing it altogether: the
// item falls through to the next candidate list, or comes back `on_request`
// — never €0. To retire a price without losing it, set the rung's
// `price_type` to `on_request` instead, or deactivate the list. An entry
// belonging to another list answers 404 rather than being deleted through the
// wrong parent.
func (srv *Prices) PricesEntriesDelete(ListId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
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
			
// PricesEntriesGet one rung of one ladder, exactly as stored — nothing is
// rounded, converted or taxed on the way out. `unit_price` is per ONE unit of
// `unit`, in the LIST's currency and on the LIST's tax basis; the entry
// itself carries neither, which is why a rung read on its own is not yet a
// price you can show a buyer. `POST /prices/resolve` is what turns it into
// one: it picks the rung that applies to a quantity, names the basis, and
// adds the net/gross pair and the tax rate. The id is checked against the
// list in the path, so an entry belonging to another list answers 404 rather
// than being read through the wrong parent.
func (srv *Prices) PricesEntriesGet(ListId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
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
					
// PricesEntriesUpdate a partial update of one rung: send only what changes, a
// payload with no updatable column at all is refused, and the next resolve
// call reads what this one wrote.
// 
// Two edits reach further than the field they touch. Moving `quantity_min`
// moves the rung within the ladder and may land on a threshold the item
// already has — nothing stops it, and both rows then sit in the resolved
// `tiers`. Setting `price_type: "on_request"` on ONE rung takes the WHOLE
// item off price in this list: resolution stops there and answers "price on
// request" even though the other rungs still carry amounts, and even where a
// less specific list would have priced it. That is the intended way to say
// "ask us" for an item, and a surprise if you meant to retire a single tier.
// 
// What this route cannot change is what the amount MEANS: currency and tax
// basis belong to the list, so re-denominating or switching net/gross is a
// list edit, not an entry edit. An entry of another list answers 404.
func (srv *Prices) PricesEntriesUpdate(ListId string, Id string, optionalSetters ...PricesEntriesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
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
			
// PricesListsMakeDefault promotes this list AND demotes whoever held the
// flag, in one call. The flag is a single answer, not a per-row opinion:
// resolution uses it as the last tie-break, so two defaults leave the winner
// to row order and none leaves a tie unsettled. Promote-then-demote as two
// PATCHes from a client produces exactly those two states whenever the second
// call does not land.
// 
// The write is as small as the change: exactly one write per row whose flag
// was wrong, and none at all for the rows that were already right. A tenant
// already in this state is therefore not written to, which is what makes
// repeating the call free. The answer is this list as it now stands plus the
// codes it demoted — empty when it already held the flag.
func (srv *Prices) PricesListsMakeDefault(ListId string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/prices/lists/{list_id}/make-default")
	params := map[string]interface{}{}
	params["list_id"] = ListId
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
			
// PricesResolve the live price call. Everything else in this app configures
// prices; this is the one route that ANSWERS them, and a storefront reaches
// it on every listing, every product page and every cart. Send up to 200
// items and the buyer context they are for — contact, organization, market
// and channel — and get back, per item, the unit price this buyer pays, the
// net/gross pair, the tax rate, the list that decided it and that item's full
// quantity ladder.
// 
// Which price wins when several match is the whole value of this app, and it
// is not guessable from the field types. The order, in full:
// 
// 1. **Candidates.** A list is a candidate when it is `active`, its currency
// EQUALS the currency of the call (nothing is ever converted — a list in
// another currency simply does not price the item), the instant `at` falls
// inside its validity window, it is visible in the buyer’s market (the
// `X-Revenexx-Market` header scopes the list view; lists assigned to no
// market are global and always visible), and its buyer scope matches or is
// open. A `requires_auth` list is dropped for a buyer with neither
// `contact_id` nor `organization_id`.
// 2. **Specificity decides first, and priority never overrules it.**
// contact-scoped (4) beats organization-scoped (3) beats channel-scoped (2)
// beats open (0). An organization list at `priority: 0` therefore wins over
// an open list at `priority: 100`.
// 3. **Within one specificity level:** `priority` descending, then
// non-default before default — the default list is deliberately last, so it
// prices only what nothing else did.
// 4. **A genuine tie** (same specificity, same priority, same default flag)
// is settled by the tenant’s `price_list_priority_tiebreak` setting —
// `lowest_price`, `highest_price`, `newest` or `code` — never by the order
// the database happened to return rows in. The setting in force is echoed in
// `basis.price_list_priority_tiebreak`.
// 5. **The first list that prices the item wins, and the search stops there**
// — even if a later, less specific list is cheaper. Its FULL tier ladder
// comes back in `tiers`; the rung with the highest `quantity_min` at or below
// the requested `quantity` sets `unit_price`, and below the first rung the
// first rung applies.
// 6. **An `on_request` entry stops the search too**, and inside a tie it
// outranks every price: a list that says "ask us" for this buyer is
// authoritative, and cannot be undercut by a list that happens to sort after
// it.
// 7. **Nothing found → `on_request`, never 0**, with a reason
// (`not_priced`, `on_request_entry`, `anonymous_denied`, `no_identity`). A
// storefront shows "price on request"; it must never show €0.
// 
// Amounts: `unit_price` is per ONE unit of the entry’s `unit`, in
// `currency`, as a decimal in MAJOR units (19.90) — never minor units/cents
// — and on the basis `tax_basis` names. `tax_basis` comes from the list’s
// own column, else from a legacy `tax_included: true` on it, else from the
// tenant’s `tax_inclusive_default`; `tax_basis_source` says which of the
// three. Read `unit_price_net`/`unit_price_gross` where you need an
// unambiguous number.
// 
// Tax is never guessed. The market comes from the `X-Revenexx-Market` header
// (a market CODE) or from `market_id` in the body; with several markets whose
// rates differ and no signal, the answer is `tax.resolved: false`, `reason:
// market_required` rather than another market’s VAT. `tax_rate: null` means
// UNKNOWN, not 0 %.
// 
// An item that cannot be priced never fails the call: it comes back
// on_request with its reason, so one bad line in a cart does not cost the
// other lines their prices.
// 
// One last thing worth knowing before you build on it. This is the most
// customised surface this app has in the field: pricing is where a tenant's
// ERP usually has the last word, and a tenant whose prices are computed there
// does not want this app's resolution order at all. So the route is
// deliberately shaped to be REPLACED — one required field, no rejection of
// an item the caller got wrong, an answer that stands on its own — and it
// is designed to be swapped 1:1 for a custom app through the gateway's
// capability override. An ERP-priced tenant overrides `prices.resolve` alone:
// the same path, the same request and the same response, answered by their
// own service, while every configuration route here (lists, entries, ladders,
// bulk changes, vocabularies) stays standard and keeps working. That is why
// the contract below is smaller than the machinery behind it, and why it
// changes reluctantly.
func (srv *Prices) PricesResolve(Items []models.PriceResolveItem, optionalSetters ...PricesResolveOption)(*models.Error, error) {
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

// PricesVocabulariesList discovery for the vocabulary routes: the enums this
// app enforces, each with its name, its title and its description — and
// deliberately WITHOUT its values, so a UI can cache this one small answer
// and then fetch only the value sets it actually renders. Names:
// list-statuses, price-types, tax-bases. Fetch one with GET
// /prices/vocabularies/{name}; a client holding the qualified pair
// 'prices.<name>' builds that URL from the pair alone.
func (srv *Prices) PricesVocabulariesList()(*models.PriceVocabularyIndex, error) {
	path := "/v1/prices/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PriceVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PriceVocabularyIndex
	parsed, ok := resp.Result.(models.PriceVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PricesVocabulariesGet one vocabulary in full: every permitted value, each
// with the title and description a human reads for it and the badge tone a UI
// colours it with — enough to render a select or a status chip without
// keeping a private copy of an enum this app enforces. The values are read
// out of the column's CHECK constraint, so the served set IS the enforced set
// and the two cannot drift — a value added to the constraint appears here
// even before anyone labels it, titled from its own key. Values come back in
// constraint order, which is the order a select should offer. 'closed' says
// the set is exhaustive, so a value outside it is stale data rather than a
// missing label. Answers 404 for an unknown name. Names: list-statuses,
// price-types, tax-bases.
func (srv *Prices) PricesVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/prices/vocabularies/{name}")
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
