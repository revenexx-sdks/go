package carts

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Carts service
type Carts struct {
	client client.Client
}

func New(clt client.Client) *Carts {
	return &Carts{
		client: clt,
	}
}


// CartsList
func (srv *Carts) CartsList()(*interface{}, error) {
	path := "/v1/carts"
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
type CartsCreateOptions struct {
	ChannelId string
	ContactId string
	Currency string
	IsCurrent bool
	MarketId string
	Metadata interface{}
	Name string
	SessionKey string
	enabledSetters map[string]bool
}
func (options CartsCreateOptions) New() *CartsCreateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"ContactId": false,
		"Currency": false,
		"IsCurrent": false,
		"MarketId": false,
		"Metadata": false,
		"Name": false,
		"SessionKey": false,
	}
	return &options
}
type CartsCreateOption func(*CartsCreateOptions)
func (srv *Carts) WithCartsCreateChannelId(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Carts) WithCartsCreateContactId(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Carts) WithCartsCreateCurrency(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsCreateIsCurrent(v bool) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.IsCurrent = v
		o.enabledSetters["IsCurrent"] = true
	}
}
func (srv *Carts) WithCartsCreateMarketId(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Carts) WithCartsCreateMetadata(v interface{}) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsCreateName(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsCreateSessionKey(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.SessionKey = v
		o.enabledSetters["SessionKey"] = true
	}
}
	
// CartsCreate
func (srv *Carts) CartsCreate(optionalSetters ...CartsCreateOption)(*models.Cart, error) {
	path := "/v1/carts"
	options := CartsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["IsCurrent"] {
		params["is_current"] = options.IsCurrent
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
	if options.enabledSetters["SessionKey"] {
		params["session_key"] = options.SessionKey
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

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CartsClaimOptions struct {
	TargetCartId string
	enabledSetters map[string]bool
}
func (options CartsClaimOptions) New() *CartsClaimOptions {
	options.enabledSetters = map[string]bool{
		"TargetCartId": false,
	}
	return &options
}
type CartsClaimOption func(*CartsClaimOptions)
func (srv *Carts) WithCartsClaimTargetCartId(v string) CartsClaimOption {
	return func(o *CartsClaimOptions) {
		o.TargetCartId = v
		o.enabledSetters["TargetCartId"] = true
	}
}
					
// CartsClaim
func (srv *Carts) CartsClaim(ContactId string, SessionKey string, optionalSetters ...CartsClaimOption)(*interface{}, error) {
	path := "/v1/carts/claim"
	options := CartsClaimOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	params["session_key"] = SessionKey
	if options.enabledSetters["TargetCartId"] {
		params["target_cart_id"] = options.TargetCartId
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
type CartsImportOptions struct {
	ContactId string
	Csv string
	Name string
	Payload interface{}
	ProfileId string
	SessionKey string
	TargetCartId string
	enabledSetters map[string]bool
}
func (options CartsImportOptions) New() *CartsImportOptions {
	options.enabledSetters = map[string]bool{
		"ContactId": false,
		"Csv": false,
		"Name": false,
		"Payload": false,
		"ProfileId": false,
		"SessionKey": false,
		"TargetCartId": false,
	}
	return &options
}
type CartsImportOption func(*CartsImportOptions)
func (srv *Carts) WithCartsImportContactId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Carts) WithCartsImportCsv(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Csv = v
		o.enabledSetters["Csv"] = true
	}
}
func (srv *Carts) WithCartsImportName(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsImportPayload(v interface{}) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Payload = v
		o.enabledSetters["Payload"] = true
	}
}
func (srv *Carts) WithCartsImportProfileId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
func (srv *Carts) WithCartsImportSessionKey(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.SessionKey = v
		o.enabledSetters["SessionKey"] = true
	}
}
func (srv *Carts) WithCartsImportTargetCartId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.TargetCartId = v
		o.enabledSetters["TargetCartId"] = true
	}
}
	
// CartsImport
func (srv *Carts) CartsImport(optionalSetters ...CartsImportOption)(*interface{}, error) {
	path := "/v1/carts/import"
	options := CartsImportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Csv"] {
		params["csv"] = options.Csv
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Payload"] {
		params["payload"] = options.Payload
	}
	if options.enabledSetters["ProfileId"] {
		params["profile_id"] = options.ProfileId
	}
	if options.enabledSetters["SessionKey"] {
		params["session_key"] = options.SessionKey
	}
	if options.enabledSetters["TargetCartId"] {
		params["target_cart_id"] = options.TargetCartId
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

// CartsIoProfilesList
func (srv *Carts) CartsIoProfilesList()(*interface{}, error) {
	path := "/v1/carts/io/profiles"
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
type CartsIoProfilesCreateOptions struct {
	ApplyMode string
	Entity string
	Format string
	IsTemplate bool
	Mapping interface{}
	Options interface{}
	enabledSetters map[string]bool
}
func (options CartsIoProfilesCreateOptions) New() *CartsIoProfilesCreateOptions {
	options.enabledSetters = map[string]bool{
		"ApplyMode": false,
		"Entity": false,
		"Format": false,
		"IsTemplate": false,
		"Mapping": false,
		"Options": false,
	}
	return &options
}
type CartsIoProfilesCreateOption func(*CartsIoProfilesCreateOptions)
func (srv *Carts) WithCartsIoProfilesCreateApplyMode(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesCreateEntity(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesCreateFormat(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesCreateIsTemplate(v bool) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.IsTemplate = v
		o.enabledSetters["IsTemplate"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesCreateMapping(v interface{}) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesCreateOptions(v interface{}) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
					
// CartsIoProfilesCreate
func (srv *Carts) CartsIoProfilesCreate(Direction string, Name string, optionalSetters ...CartsIoProfilesCreateOption)(*models.IoProfile, error) {
	path := "/v1/carts/io/profiles"
	options := CartsIoProfilesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["direction"] = Direction
	params["name"] = Name
	if options.enabledSetters["ApplyMode"] {
		params["apply_mode"] = options.ApplyMode
	}
	if options.enabledSetters["Entity"] {
		params["entity"] = options.Entity
	}
	if options.enabledSetters["Format"] {
		params["format"] = options.Format
	}
	if options.enabledSetters["IsTemplate"] {
		params["is_template"] = options.IsTemplate
	}
	if options.enabledSetters["Mapping"] {
		params["mapping"] = options.Mapping
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
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

		parsed := models.IoProfile{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.IoProfile
	parsed, ok := resp.Result.(models.IoProfile)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CartsIoProfilesDefaults
func (srv *Carts) CartsIoProfilesDefaults()(*interface{}, error) {
	path := "/v1/carts/io/profiles/defaults"
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
	
// CartsIoProfilesDelete
func (srv *Carts) CartsIoProfilesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/io/profiles/{id}")
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
	
// CartsIoProfilesGet
func (srv *Carts) CartsIoProfilesGet(Id string)(*models.IoProfile, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/io/profiles/{id}")
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

		parsed := models.IoProfile{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.IoProfile
	parsed, ok := resp.Result.(models.IoProfile)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CartsIoProfilesUpdateOptions struct {
	ApplyMode string
	Direction string
	Entity string
	Format string
	IsTemplate bool
	Mapping interface{}
	Name string
	Options interface{}
	enabledSetters map[string]bool
}
func (options CartsIoProfilesUpdateOptions) New() *CartsIoProfilesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"ApplyMode": false,
		"Direction": false,
		"Entity": false,
		"Format": false,
		"IsTemplate": false,
		"Mapping": false,
		"Name": false,
		"Options": false,
	}
	return &options
}
type CartsIoProfilesUpdateOption func(*CartsIoProfilesUpdateOptions)
func (srv *Carts) WithCartsIoProfilesUpdateApplyMode(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateDirection(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Direction = v
		o.enabledSetters["Direction"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateEntity(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateFormat(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateIsTemplate(v bool) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.IsTemplate = v
		o.enabledSetters["IsTemplate"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateMapping(v interface{}) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateName(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsIoProfilesUpdateOptions(v interface{}) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
			
// CartsIoProfilesUpdate
func (srv *Carts) CartsIoProfilesUpdate(Id string, optionalSetters ...CartsIoProfilesUpdateOption)(*models.IoProfile, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/io/profiles/{id}")
	options := CartsIoProfilesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ApplyMode"] {
		params["apply_mode"] = options.ApplyMode
	}
	if options.enabledSetters["Direction"] {
		params["direction"] = options.Direction
	}
	if options.enabledSetters["Entity"] {
		params["entity"] = options.Entity
	}
	if options.enabledSetters["Format"] {
		params["format"] = options.Format
	}
	if options.enabledSetters["IsTemplate"] {
		params["is_template"] = options.IsTemplate
	}
	if options.enabledSetters["Mapping"] {
		params["mapping"] = options.Mapping
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
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

		parsed := models.IoProfile{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.IoProfile
	parsed, ok := resp.Result.(models.IoProfile)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CartsMerge
func (srv *Carts) CartsMerge(SourceCartId string, TargetCartId string)(*interface{}, error) {
	path := "/v1/carts/merge"
	params := map[string]interface{}{}
	params["source_cart_id"] = SourceCartId
	params["target_cart_id"] = TargetCartId
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
	
// CartsItemsList
func (srv *Carts) CartsItemsList(CartId string)(*interface{}, error) {
	r := strings.NewReplacer("{cartId}", CartId)
	path := r.Replace("/v1/carts/{cart_id}/items")
	params := map[string]interface{}{}
	params["cart_id"] = CartId
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
func (srv *Carts) WithCartsItemsCreateConfiguration(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Configuration = v
		o.enabledSetters["Configuration"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateCurrency(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateMetadata(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateName(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsItemsCreatePosition(v int) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateProductId(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateQuantity(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateSku(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateSnapshot(v interface{}) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Snapshot = v
		o.enabledSetters["Snapshot"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateTaxRate(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateType(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateUnit(v string) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Carts) WithCartsItemsCreateUnitPrice(v float64) CartsItemsCreateOption {
	return func(o *CartsItemsCreateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
			
// CartsItemsCreate
func (srv *Carts) CartsItemsCreate(CartId string, optionalSetters ...CartsItemsCreateOption)(*models.CartItem, error) {
	r := strings.NewReplacer("{cartId}", CartId)
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

		parsed := models.CartItem{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CartItem
	parsed, ok := resp.Result.(models.CartItem)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CartsItemsReplace
func (srv *Carts) CartsItemsReplace(CartId string, Items []models.CartItemCreateRequest)(*interface{}, error) {
	r := strings.NewReplacer("{cartId}", CartId)
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
			
// CartsItemsDelete
func (srv *Carts) CartsItemsDelete(CartId string, Id string)(*interface{}, error) {
	r := strings.NewReplacer("{cartId}", CartId, "{id}", Id)
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
			
// CartsItemsGet
func (srv *Carts) CartsItemsGet(CartId string, Id string)(*models.CartItem, error) {
	r := strings.NewReplacer("{cartId}", CartId, "{id}", Id)
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

		parsed := models.CartItem{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CartItem
	parsed, ok := resp.Result.(models.CartItem)
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
func (srv *Carts) WithCartsItemsUpdateConfiguration(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Configuration = v
		o.enabledSetters["Configuration"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateCurrency(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateMetadata(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateName(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdatePosition(v int) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateProductId(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateQuantity(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateSku(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateSnapshot(v interface{}) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Snapshot = v
		o.enabledSetters["Snapshot"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateTaxRate(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateType(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateUnit(v string) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Carts) WithCartsItemsUpdateUnitPrice(v float64) CartsItemsUpdateOption {
	return func(o *CartsItemsUpdateOptions) {
		o.UnitPrice = v
		o.enabledSetters["UnitPrice"] = true
	}
}
					
// CartsItemsUpdate
func (srv *Carts) CartsItemsUpdate(CartId string, Id string, optionalSetters ...CartsItemsUpdateOption)(*models.CartItem, error) {
	r := strings.NewReplacer("{cartId}", CartId, "{id}", Id)
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

		parsed := models.CartItem{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CartItem
	parsed, ok := resp.Result.(models.CartItem)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CartsDelete
func (srv *Carts) CartsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
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
	
// CartsGet
func (srv *Carts) CartsGet(Id string)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
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

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CartsUpdateOptions struct {
	ChannelId string
	Currency string
	MarketId string
	Metadata interface{}
	Name string
	enabledSetters map[string]bool
}
func (options CartsUpdateOptions) New() *CartsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"Currency": false,
		"MarketId": false,
		"Metadata": false,
		"Name": false,
	}
	return &options
}
type CartsUpdateOption func(*CartsUpdateOptions)
func (srv *Carts) WithCartsUpdateChannelId(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Carts) WithCartsUpdateCurrency(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsUpdateMarketId(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.MarketId = v
		o.enabledSetters["MarketId"] = true
	}
}
func (srv *Carts) WithCartsUpdateMetadata(v interface{}) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsUpdateName(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// CartsUpdate
func (srv *Carts) CartsUpdate(Id string, optionalSetters ...CartsUpdateOption)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
	options := CartsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CartsAbandon
func (srv *Carts) CartsAbandon(Id string)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/abandon")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CartsActivate
func (srv *Carts) CartsActivate(Id string)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/activate")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CartsExportOptions struct {
	Format string
	ProfileId string
	enabledSetters map[string]bool
}
func (options CartsExportOptions) New() *CartsExportOptions {
	options.enabledSetters = map[string]bool{
		"Format": false,
		"ProfileId": false,
	}
	return &options
}
type CartsExportOption func(*CartsExportOptions)
func (srv *Carts) WithCartsExportFormat(v string) CartsExportOption {
	return func(o *CartsExportOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *Carts) WithCartsExportProfileId(v string) CartsExportOption {
	return func(o *CartsExportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
			
// CartsExport
func (srv *Carts) CartsExport(Id string, optionalSetters ...CartsExportOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/export")
	options := CartsExportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Format"] {
		params["format"] = options.Format
	}
	if options.enabledSetters["ProfileId"] {
		params["profile_id"] = options.ProfileId
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
type CartsOrderOptions struct {
	OrderRef string
	enabledSetters map[string]bool
}
func (options CartsOrderOptions) New() *CartsOrderOptions {
	options.enabledSetters = map[string]bool{
		"OrderRef": false,
	}
	return &options
}
type CartsOrderOption func(*CartsOrderOptions)
func (srv *Carts) WithCartsOrderOrderRef(v string) CartsOrderOption {
	return func(o *CartsOrderOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
			
// CartsOrder
func (srv *Carts) CartsOrder(Id string, optionalSetters ...CartsOrderOption)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/order")
	options := CartsOrderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
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

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CartsReopen
func (srv *Carts) CartsReopen(Id string)(*models.Cart, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/reopen")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Cart{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Cart
	parsed, ok := resp.Result.(models.Cart)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
