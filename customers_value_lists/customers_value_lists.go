package customers_value_lists

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CustomersValueLists service
type CustomersValueLists struct {
	client client.Client
}

func New(clt client.Client) *CustomersValueLists {
	return &CustomersValueLists{
		client: clt,
	}
}


// CustomersAddressTypesList what an address is used for. Billing and shipping
// are what a checkout needs; a works entrance or a central accounts office is
// the tenant's own. A fresh install is seeded with billing, shipping, and the
// set seeds on first read too, so the page is never empty and
// `addresses.type` always has a value it may carry. The whole set comes back
// in one page in the tenant's own order — this route takes no
// limit/offset/order and no column filters, so `page` describes the full set
// and `filter` is always empty.
func (srv *CustomersValueLists) CustomersAddressTypesList()(*interface{}, error) {
	path := "/v1/customers/address-types"
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
type CustomersAddressTypesCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersAddressTypesCreateOptions) New() *CustomersAddressTypesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Tone": false,
	}
	return &options
}
type CustomersAddressTypesCreateOption func(*CustomersAddressTypesCreateOptions)
func (srv *CustomersValueLists) WithCustomersAddressTypesCreateDescription(v string) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesCreateDescriptions(v interface{}) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesCreateIsDefault(v bool) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesCreateLabels(v interface{}) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesCreatePosition(v int) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesCreateTone(v string) CustomersAddressTypesCreateOption {
	return func(o *CustomersAddressTypesCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// CustomersAddressTypesCreate extends this tenant's address types set with a
// value of their own — the whole reason these four stopped being CHECK
// constraints. What an address is used for. Billing and shipping are what a
// checkout needs; a works entrance or a central accounts office is the
// tenant's own. The code is lowercase and becomes what `addresses.type`
// stores; it cannot be changed afterwards, because every record carrying it
// would be orphaned.
func (srv *CustomersValueLists) CustomersAddressTypesCreate(Code string, Title string, optionalSetters ...CustomersAddressTypesCreateOption)(*models.Error, error) {
	path := "/v1/customers/address-types"
	options := CustomersAddressTypesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["title"] = Title
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// CustomersAddressTypesDelete takes a value out of the address types set.
// There is no foreign key behind `addresses.type` — one added to a table
// that starts empty fails the migration of every existing tenant — so this
// route IS the integrity: it refuses while any record still carries the code,
// and it refuses to empty the set. Retiring a value that is in use is
// therefore a two-step job: move the records onto another value first, then
// remove it.
func (srv *CustomersValueLists) CustomersAddressTypesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/address-types/{id}")
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
	
// CustomersAddressTypesGet one value of the address types set, by its id —
// its code, its fallback title, the per-language `labels` an operator reads
// and the badge `tone` a client renders it with. What an address is used for.
// Billing and shipping are what a checkout needs; a works entrance or a
// central accounts office is the tenant's own. Reading one value is the rare
// path: `GET /customers/address-types` answers the whole set in a single
// page, which is what a select needs.
func (srv *CustomersValueLists) CustomersAddressTypesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/address-types/{id}")
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
type CustomersAddressTypesUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersAddressTypesUpdateOptions) New() *CustomersAddressTypesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type CustomersAddressTypesUpdateOption func(*CustomersAddressTypesUpdateOptions)
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateDescription(v string) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateDescriptions(v interface{}) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateIsDefault(v bool) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateLabels(v interface{}) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdatePosition(v int) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateTitle(v string) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersAddressTypesUpdateTone(v string) CustomersAddressTypesUpdateOption {
	return func(o *CustomersAddressTypesUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// CustomersAddressTypesUpdate everything about a value except the value
// itself: its titles, its help text, its badge tone, its `position` in the
// select, and which one of the set is the default. The `code` is immutable,
// so no record carrying it is ever orphaned by an edit here — a merchant
// who retitles `shipping` to wording of their own changes what people READ
// and nothing about what `addresses.type` stores. Seeded values (`is_system`)
// are renameable like any other, and re-seeding leaves the rename alone.
func (srv *CustomersValueLists) CustomersAddressTypesUpdate(Id string, optionalSetters ...CustomersAddressTypesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/address-types/{id}")
	options := CustomersAddressTypesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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

// CustomersContactEventKindsList what kind of entry lands on a customer
// timeline. 'system' is the app's own decision trail and a caller may not
// file one, whatever the set says. A fresh install is seeded with system,
// note, call, email, meeting, visit, task, and the set seeds on first read
// too, so the page is never empty and `contact_events.kind` always has a
// value it may carry. The whole set comes back in one page in the tenant's
// own order — this route takes no limit/offset/order and no column filters,
// so `page` describes the full set and `filter` is always empty.
func (srv *CustomersValueLists) CustomersContactEventKindsList()(*interface{}, error) {
	path := "/v1/customers/contact-event-kinds"
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
type CustomersContactEventKindsCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersContactEventKindsCreateOptions) New() *CustomersContactEventKindsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Tone": false,
	}
	return &options
}
type CustomersContactEventKindsCreateOption func(*CustomersContactEventKindsCreateOptions)
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreateDescription(v string) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreateDescriptions(v interface{}) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreateIsDefault(v bool) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreateLabels(v interface{}) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreatePosition(v int) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsCreateTone(v string) CustomersContactEventKindsCreateOption {
	return func(o *CustomersContactEventKindsCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// CustomersContactEventKindsCreate extends this tenant's activity types set
// with a value of their own — the whole reason these four stopped being
// CHECK constraints. What kind of entry lands on a customer timeline.
// 'system' is the app's own decision trail and a caller may not file one,
// whatever the set says. The code is lowercase and becomes what
// `contact_events.kind` stores; it cannot be changed afterwards, because
// every record carrying it would be orphaned.
func (srv *CustomersValueLists) CustomersContactEventKindsCreate(Code string, Title string, optionalSetters ...CustomersContactEventKindsCreateOption)(*models.Error, error) {
	path := "/v1/customers/contact-event-kinds"
	options := CustomersContactEventKindsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["title"] = Title
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// CustomersContactEventKindsDelete takes a value out of the activity types
// set. There is no foreign key behind `contact_events.kind` — one added to
// a table that starts empty fails the migration of every existing tenant —
// so this route IS the integrity: it refuses while any record still carries
// the code, and it refuses to empty the set. Retiring a value that is in use
// is therefore a two-step job: move the records onto another value first,
// then remove it.
func (srv *CustomersValueLists) CustomersContactEventKindsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contact-event-kinds/{id}")
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
	
// CustomersContactEventKindsGet one value of the activity types set, by its
// id — its code, its fallback title, the per-language `labels` an operator
// reads and the badge `tone` a client renders it with. What kind of entry
// lands on a customer timeline. 'system' is the app's own decision trail and
// a caller may not file one, whatever the set says. Reading one value is the
// rare path: `GET /customers/contact-event-kinds` answers the whole set in a
// single page, which is what a select needs.
func (srv *CustomersValueLists) CustomersContactEventKindsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contact-event-kinds/{id}")
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
type CustomersContactEventKindsUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersContactEventKindsUpdateOptions) New() *CustomersContactEventKindsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type CustomersContactEventKindsUpdateOption func(*CustomersContactEventKindsUpdateOptions)
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateDescription(v string) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateDescriptions(v interface{}) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateIsDefault(v bool) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateLabels(v interface{}) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdatePosition(v int) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateTitle(v string) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersContactEventKindsUpdateTone(v string) CustomersContactEventKindsUpdateOption {
	return func(o *CustomersContactEventKindsUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// CustomersContactEventKindsUpdate everything about a value except the value
// itself: its titles, its help text, its badge tone, its `position` in the
// select, and which one of the set is the default. The `code` is immutable,
// so no record carrying it is ever orphaned by an edit here — a merchant
// who retitles `call` to wording of their own changes what people READ and
// nothing about what `contact_events.kind` stores. Seeded values
// (`is_system`) are renameable like any other, and re-seeding leaves the
// rename alone.
func (srv *CustomersValueLists) CustomersContactEventKindsUpdate(Id string, optionalSetters ...CustomersContactEventKindsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contact-event-kinds/{id}")
	options := CustomersContactEventKindsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// CustomersDefaults what the app.installed event runs. It fills all four of
// the value sets a tenant needs before anything else works — the payment
// terms, the address types, the lifecycle stages and the activity types —
// in one call. Idempotent by code: a set that already has its rows is left
// completely alone, so a re-delivered event and a merchant's renames both
// survive. A tenant installed before these tables existed is seeded lazily
// instead, by the first read that finds one empty.
func (srv *CustomersValueLists) CustomersDefaults(Data interface{})(*models.Error, error) {
	path := "/v1/customers/defaults"
	params := map[string]interface{}{}
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

// CustomersLifecycleStagesList where a company stands in the sales pipeline
// — a separate axis from status, and one whose steps are a sales team's
// own. A fresh install is seeded with lead, prospect, customer, churned, and
// the set seeds on first read too, so the page is never empty and
// `organizations.lifecycle_stage` always has a value it may carry. The whole
// set comes back in one page in the tenant's own order — this route takes
// no limit/offset/order and no column filters, so `page` describes the full
// set and `filter` is always empty.
func (srv *CustomersValueLists) CustomersLifecycleStagesList()(*interface{}, error) {
	path := "/v1/customers/lifecycle-stages"
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
type CustomersLifecycleStagesCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersLifecycleStagesCreateOptions) New() *CustomersLifecycleStagesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Tone": false,
	}
	return &options
}
type CustomersLifecycleStagesCreateOption func(*CustomersLifecycleStagesCreateOptions)
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreateDescription(v string) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreateDescriptions(v interface{}) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreateIsDefault(v bool) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreateLabels(v interface{}) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreatePosition(v int) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesCreateTone(v string) CustomersLifecycleStagesCreateOption {
	return func(o *CustomersLifecycleStagesCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// CustomersLifecycleStagesCreate extends this tenant's lifecycle stages set
// with a value of their own — the whole reason these four stopped being
// CHECK constraints. Where a company stands in the sales pipeline — a
// separate axis from status, and one whose steps are a sales team's own. The
// code is lowercase and becomes what `organizations.lifecycle_stage` stores;
// it cannot be changed afterwards, because every record carrying it would be
// orphaned.
func (srv *CustomersValueLists) CustomersLifecycleStagesCreate(Code string, Title string, optionalSetters ...CustomersLifecycleStagesCreateOption)(*models.Error, error) {
	path := "/v1/customers/lifecycle-stages"
	options := CustomersLifecycleStagesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["title"] = Title
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// CustomersLifecycleStagesDelete takes a value out of the lifecycle stages
// set. There is no foreign key behind `organizations.lifecycle_stage` — one
// added to a table that starts empty fails the migration of every existing
// tenant — so this route IS the integrity: it refuses while any record
// still carries the code, and it refuses to empty the set. Retiring a value
// that is in use is therefore a two-step job: move the records onto another
// value first, then remove it.
func (srv *CustomersValueLists) CustomersLifecycleStagesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/lifecycle-stages/{id}")
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
	
// CustomersLifecycleStagesGet one value of the lifecycle stages set, by its
// id — its code, its fallback title, the per-language `labels` an operator
// reads and the badge `tone` a client renders it with. Where a company stands
// in the sales pipeline — a separate axis from status, and one whose steps
// are a sales team's own. Reading one value is the rare path: `GET
// /customers/lifecycle-stages` answers the whole set in a single page, which
// is what a select needs.
func (srv *CustomersValueLists) CustomersLifecycleStagesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/lifecycle-stages/{id}")
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
type CustomersLifecycleStagesUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersLifecycleStagesUpdateOptions) New() *CustomersLifecycleStagesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type CustomersLifecycleStagesUpdateOption func(*CustomersLifecycleStagesUpdateOptions)
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateDescription(v string) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateDescriptions(v interface{}) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateIsDefault(v bool) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateLabels(v interface{}) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdatePosition(v int) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateTitle(v string) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersLifecycleStagesUpdateTone(v string) CustomersLifecycleStagesUpdateOption {
	return func(o *CustomersLifecycleStagesUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// CustomersLifecycleStagesUpdate everything about a value except the value
// itself: its titles, its help text, its badge tone, its `position` in the
// select, and which one of the set is the default. The `code` is immutable,
// so no record carrying it is ever orphaned by an edit here — a merchant
// who retitles `customer` to wording of their own changes what people READ
// and nothing about what `organizations.lifecycle_stage` stores. Seeded
// values (`is_system`) are renameable like any other, and re-seeding leaves
// the rename alone.
func (srv *CustomersValueLists) CustomersLifecycleStagesUpdate(Id string, optionalSetters ...CustomersLifecycleStagesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/lifecycle-stages/{id}")
	options := CustomersLifecycleStagesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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

// CustomersPaymentTermsList when a company has to pay. A wholesaler who
// agrees net 45 with one customer used to need a release of this app to say
// so. A fresh install is seeded with prepayment, direct_debit, net_7, net_14,
// net_30, net_60, net_90, and the set seeds on first read too, so the page is
// never empty and `organizations.payment_terms` always has a value it may
// carry. The whole set comes back in one page in the tenant's own order —
// this route takes no limit/offset/order and no column filters, so `page`
// describes the full set and `filter` is always empty.
func (srv *CustomersValueLists) CustomersPaymentTermsList()(*interface{}, error) {
	path := "/v1/customers/payment-terms"
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
type CustomersPaymentTermsCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersPaymentTermsCreateOptions) New() *CustomersPaymentTermsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Tone": false,
	}
	return &options
}
type CustomersPaymentTermsCreateOption func(*CustomersPaymentTermsCreateOptions)
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreateDescription(v string) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreateDescriptions(v interface{}) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreateIsDefault(v bool) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreateLabels(v interface{}) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreatePosition(v int) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsCreateTone(v string) CustomersPaymentTermsCreateOption {
	return func(o *CustomersPaymentTermsCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// CustomersPaymentTermsCreate extends this tenant's payment terms set with a
// value of their own — the whole reason these four stopped being CHECK
// constraints. When a company has to pay. A wholesaler who agrees net 45 with
// one customer used to need a release of this app to say so. The code is
// lowercase and becomes what `organizations.payment_terms` stores; it cannot
// be changed afterwards, because every record carrying it would be orphaned.
func (srv *CustomersValueLists) CustomersPaymentTermsCreate(Code string, Title string, optionalSetters ...CustomersPaymentTermsCreateOption)(*models.Error, error) {
	path := "/v1/customers/payment-terms"
	options := CustomersPaymentTermsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["title"] = Title
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// CustomersPaymentTermsDelete takes a value out of the payment terms set.
// There is no foreign key behind `organizations.payment_terms` — one added
// to a table that starts empty fails the migration of every existing tenant
// — so this route IS the integrity: it refuses while any record still
// carries the code, and it refuses to empty the set. Retiring a value that is
// in use is therefore a two-step job: move the records onto another value
// first, then remove it.
func (srv *CustomersValueLists) CustomersPaymentTermsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/payment-terms/{id}")
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
	
// CustomersPaymentTermsGet one value of the payment terms set, by its id —
// its code, its fallback title, the per-language `labels` an operator reads
// and the badge `tone` a client renders it with. When a company has to pay. A
// wholesaler who agrees net 45 with one customer used to need a release of
// this app to say so. Reading one value is the rare path: `GET
// /customers/payment-terms` answers the whole set in a single page, which is
// what a select needs.
func (srv *CustomersValueLists) CustomersPaymentTermsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/payment-terms/{id}")
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
type CustomersPaymentTermsUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options CustomersPaymentTermsUpdateOptions) New() *CustomersPaymentTermsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type CustomersPaymentTermsUpdateOption func(*CustomersPaymentTermsUpdateOptions)
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateDescription(v string) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateDescriptions(v interface{}) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateIsDefault(v bool) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateLabels(v interface{}) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdatePosition(v int) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateTitle(v string) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *CustomersValueLists) WithCustomersPaymentTermsUpdateTone(v string) CustomersPaymentTermsUpdateOption {
	return func(o *CustomersPaymentTermsUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// CustomersPaymentTermsUpdate everything about a value except the value
// itself: its titles, its help text, its badge tone, its `position` in the
// select, and which one of the set is the default. The `code` is immutable,
// so no record carrying it is ever orphaned by an edit here — a merchant
// who retitles `net_30` to wording of their own changes what people READ and
// nothing about what `organizations.payment_terms` stores. Seeded values
// (`is_system`) are renameable like any other, and re-seeding leaves the
// rename alone.
func (srv *CustomersValueLists) CustomersPaymentTermsUpdate(Id string, optionalSetters ...CustomersPaymentTermsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/payment-terms/{id}")
	options := CustomersPaymentTermsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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

// CustomersVocabulariesList discovery for the vocabulary routes: every enum
// this app publishes, each as a name, a title and a description. The VALUES
// are deliberately left out — this is the call that says which vocabularies
// exist, and the detail route is the one that answers what is in them. Names:
// address-types, contact-event-kinds, contact-statuses, lifecycle-stages,
// locales, organization-statuses, payment-terms, registration-statuses,
// roles, rule-matches, segment-sources. Fetch one with GET
// /customers/vocabularies/{name}; a client holding the qualified pair
// 'customers.<name>' builds that URL from the pair alone.
func (srv *CustomersValueLists) CustomersVocabulariesList()(*models.VocabularyIndex, error) {
	path := "/v1/customers/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.VocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VocabularyIndex
	parsed, ok := resp.Result.(models.VocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersVocabulariesGet one vocabulary in full: every permitted value,
// each with its title, its description and the badge tone a client renders it
// with — enough to build a select without a second call. Two kinds of set,
// and 'source' says which one answered. 'schema' — the values are read out
// of the column's CHECK constraint, so the served set IS the enforced set and
// the two cannot drift; a value added to the constraint appears here even
// before anyone labels it, titled from its own key. 'table' — the values
// are the TENANT's own rows (payment terms, address types, lifecycle stages,
// activity types, roles), so they carry labels/descriptions per locale,
// is_system and is_default, and a merchant may add to them without a release
// of this app. 'tenant'/'defaults' are the two answers for a set the merchant
// configures but may not extend. Either way 'closed' is true: the set is
// exhaustive at this moment, so a value outside it is stale data rather than
// a missing label. Values come back in the order a select should offer them
// — lifecycle order for a status, the merchant's own position for a table.
// Names: address-types, contact-event-kinds, contact-statuses,
// lifecycle-stages, locales, organization-statuses, payment-terms,
// registration-statuses, roles, rule-matches, segment-sources.
func (srv *CustomersValueLists) CustomersVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/customers/vocabularies/{name}")
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
