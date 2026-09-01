package customers_organizations

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CustomersOrganizations service
type CustomersOrganizations struct {
	client client.Client
}

func New(clt client.Client) *CustomersOrganizations {
	return &CustomersOrganizations{
		client: clt,
	}
}

type CustomersAddressesListOptions struct {
	Id string
	OrganizationId string
	ContactId string
	Type string
	Company string
	Name string
	Street string
	Street2 string
	Zip string
	City string
	Region string
	Country string
	Phone string
	IsDefault bool
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersAddressesListOptions) New() *CustomersAddressesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"OrganizationId": false,
		"ContactId": false,
		"Type": false,
		"Company": false,
		"Name": false,
		"Street": false,
		"Street2": false,
		"Zip": false,
		"City": false,
		"Region": false,
		"Country": false,
		"Phone": false,
		"IsDefault": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersAddressesListOption func(*CustomersAddressesListOptions)
func (srv *CustomersOrganizations) WithCustomersAddressesListId(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListOrganizationId(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListContactId(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListType(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListCompany(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Company = v
		o.enabledSetters["Company"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListName(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListStreet(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Street = v
		o.enabledSetters["Street"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListStreet2(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Street2 = v
		o.enabledSetters["Street2"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListZip(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Zip = v
		o.enabledSetters["Zip"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListCity(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.City = v
		o.enabledSetters["City"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListRegion(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListCountry(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListPhone(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListIsDefault(v bool) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListCreatedAt(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListUpdatedAt(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListLimit(v int) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListOffset(v int) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesListOrder(v string) CustomersAddressesListOption {
	return func(o *CustomersAddressesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersAddressesList a postal address used for billing or for shipping,
// owned by exactly one of the two parties: an organization (the company
// address everyone in it may use) or a contact (a private one only that
// person uses). Both owner columns are nullable and exactly one is set —
// sending both, or neither, is refused. Every address this tenant holds,
// filterable by owner (`organization_id`, `contact_id`), by `type` and by any
// other column. It is how the addresses tab of a company or a person is
// filled; the page is `limit`/`offset`/`order`.
func (srv *CustomersOrganizations) CustomersAddressesList(optionalSetters ...CustomersAddressesListOption)(*interface{}, error) {
	path := "/v1/customers/addresses"
	options := CustomersAddressesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Company"] {
		params["company"] = options.Company
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Street"] {
		params["street"] = options.Street
	}
	if options.enabledSetters["Street2"] {
		params["street2"] = options.Street2
	}
	if options.enabledSetters["Zip"] {
		params["zip"] = options.Zip
	}
	if options.enabledSetters["City"] {
		params["city"] = options.City
	}
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
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
type CustomersAddressesCreateOptions struct {
	Company string
	ContactId string
	IsDefault bool
	Name string
	OrganizationId string
	Phone string
	Region string
	Street2 string
	Type string
	enabledSetters map[string]bool
}
func (options CustomersAddressesCreateOptions) New() *CustomersAddressesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Company": false,
		"ContactId": false,
		"IsDefault": false,
		"Name": false,
		"OrganizationId": false,
		"Phone": false,
		"Region": false,
		"Street2": false,
		"Type": false,
	}
	return &options
}
type CustomersAddressesCreateOption func(*CustomersAddressesCreateOptions)
func (srv *CustomersOrganizations) WithCustomersAddressesCreateCompany(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Company = v
		o.enabledSetters["Company"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateContactId(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateIsDefault(v bool) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateName(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateOrganizationId(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreatePhone(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateRegion(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateStreet2(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Street2 = v
		o.enabledSetters["Street2"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesCreateType(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
									
// CustomersAddressesCreate a postal address used for billing or for shipping,
// owned by exactly one of the two parties: an organization (the company
// address everyone in it may use) or a contact (a private one only that
// person uses). Both owner columns are nullable and exactly one is set —
// sending both, or neither, is refused. `type` names one of this tenant's own
// address types — billing and shipping are seeded, and a merchant may add a
// works entrance or a central accounts office without a release of this app.
// `is_default` picks the one a checkout should preselect for that owner and
// that type. A create cannot omit `street`, `zip`, `city` and `country`;
// everything else is optional or defaulted by the database.
func (srv *CustomersOrganizations) CustomersAddressesCreate(City string, Country string, Street string, Zip string, optionalSetters ...CustomersAddressesCreateOption)(*models.Error, error) {
	path := "/v1/customers/addresses"
	options := CustomersAddressesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["city"] = City
	params["country"] = Country
	params["street"] = Street
	params["zip"] = Zip
	if options.enabledSetters["Company"] {
		params["company"] = options.Company
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	if options.enabledSetters["Street2"] {
		params["street2"] = options.Street2
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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
	
// CustomersAddressesDelete a postal address used for billing or for shipping,
// owned by exactly one of the two parties: an organization (the company
// address everyone in it may use) or a contact (a private one only that
// person uses). Both owner columns are nullable and exactly one is set —
// sending both, or neither, is refused. Removes the address. Orders already
// placed keep the address they were placed with; nothing in this app reaches
// back. Nothing else in this app points at it, so nothing else goes with it.
func (srv *CustomersOrganizations) CustomersAddressesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/addresses/{id}")
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
	
// CustomersAddressesGet a postal address used for billing or for shipping,
// owned by exactly one of the two parties: an organization (the company
// address everyone in it may use) or a contact (a private one only that
// person uses). Both owner columns are nullable and exactly one is set —
// sending both, or neither, is refused. One address by id, whichever of the
// two owners it hangs off.
func (srv *CustomersOrganizations) CustomersAddressesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/addresses/{id}")
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
type CustomersAddressesUpdateOptions struct {
	City string
	Company string
	ContactId string
	Country string
	IsDefault bool
	Name string
	OrganizationId string
	Phone string
	Region string
	Street string
	Street2 string
	Type string
	Zip string
	enabledSetters map[string]bool
}
func (options CustomersAddressesUpdateOptions) New() *CustomersAddressesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"City": false,
		"Company": false,
		"ContactId": false,
		"Country": false,
		"IsDefault": false,
		"Name": false,
		"OrganizationId": false,
		"Phone": false,
		"Region": false,
		"Street": false,
		"Street2": false,
		"Type": false,
		"Zip": false,
	}
	return &options
}
type CustomersAddressesUpdateOption func(*CustomersAddressesUpdateOptions)
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateCity(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.City = v
		o.enabledSetters["City"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateCompany(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Company = v
		o.enabledSetters["Company"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateContactId(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateCountry(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateIsDefault(v bool) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateName(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateOrganizationId(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdatePhone(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateRegion(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateStreet(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Street = v
		o.enabledSetters["Street"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateStreet2(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Street2 = v
		o.enabledSetters["Street2"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateType(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersAddressesUpdateZip(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Zip = v
		o.enabledSetters["Zip"] = true
	}
}
			
// CustomersAddressesUpdate a postal address used for billing or for shipping,
// owned by exactly one of the two parties: an organization (the company
// address everyone in it may use) or a contact (a private one only that
// person uses). Both owner columns are nullable and exactly one is set —
// sending both, or neither, is refused. A partial update — send only what
// changes. An empty body is refused rather than answered as a no-op, so a
// client that built the wrong patch finds out.
func (srv *CustomersOrganizations) CustomersAddressesUpdate(Id string, optionalSetters ...CustomersAddressesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/addresses/{id}")
	options := CustomersAddressesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["City"] {
		params["city"] = options.City
	}
	if options.enabledSetters["Company"] {
		params["company"] = options.Company
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	if options.enabledSetters["Street"] {
		params["street"] = options.Street
	}
	if options.enabledSetters["Street2"] {
		params["street2"] = options.Street2
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Zip"] {
		params["zip"] = options.Zip
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
type CustomersOrganizationMetricsListOptions struct {
	Id string
	OrganizationId string
	OrderCount int
	OrderCount30d int
	OrderCount90d int
	OrderCount365d int
	RevenueTotal float64
	Revenue30d float64
	Revenue90d float64
	Revenue365d float64
	AvgOrderValue float64
	AvgOrderValue365d float64
	FirstOrderAt string
	LastOrderAt string
	Currency string
	CurrencyMixed bool
	OrdersAsOf string
	ComputedAt string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationMetricsListOptions) New() *CustomersOrganizationMetricsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"OrganizationId": false,
		"OrderCount": false,
		"OrderCount30d": false,
		"OrderCount90d": false,
		"OrderCount365d": false,
		"RevenueTotal": false,
		"Revenue30d": false,
		"Revenue90d": false,
		"Revenue365d": false,
		"AvgOrderValue": false,
		"AvgOrderValue365d": false,
		"FirstOrderAt": false,
		"LastOrderAt": false,
		"Currency": false,
		"CurrencyMixed": false,
		"OrdersAsOf": false,
		"ComputedAt": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersOrganizationMetricsListOption func(*CustomersOrganizationMetricsListOptions)
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListId(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrganizationId(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrderCount(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrderCount = v
		o.enabledSetters["OrderCount"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrderCount30d(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrderCount30d = v
		o.enabledSetters["OrderCount30d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrderCount90d(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrderCount90d = v
		o.enabledSetters["OrderCount90d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrderCount365d(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrderCount365d = v
		o.enabledSetters["OrderCount365d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListRevenueTotal(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.RevenueTotal = v
		o.enabledSetters["RevenueTotal"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListRevenue30d(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Revenue30d = v
		o.enabledSetters["Revenue30d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListRevenue90d(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Revenue90d = v
		o.enabledSetters["Revenue90d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListRevenue365d(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Revenue365d = v
		o.enabledSetters["Revenue365d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListAvgOrderValue(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.AvgOrderValue = v
		o.enabledSetters["AvgOrderValue"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListAvgOrderValue365d(v float64) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.AvgOrderValue365d = v
		o.enabledSetters["AvgOrderValue365d"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListFirstOrderAt(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.FirstOrderAt = v
		o.enabledSetters["FirstOrderAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListLastOrderAt(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.LastOrderAt = v
		o.enabledSetters["LastOrderAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListCurrency(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListCurrencyMixed(v bool) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.CurrencyMixed = v
		o.enabledSetters["CurrencyMixed"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrdersAsOf(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.OrdersAsOf = v
		o.enabledSetters["OrdersAsOf"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListComputedAt(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.ComputedAt = v
		o.enabledSetters["ComputedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListCreatedAt(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListUpdatedAt(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListLimit(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOffset(v int) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsListOrder(v string) CustomersOrganizationMetricsListOption {
	return func(o *CustomersOrganizationMetricsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersOrganizationMetricsList what an organization has BOUGHT,
// materialized into this app from the orders app: lifetime revenue, revenue
// over the last 30/90/365 days, order count, average order value, and the
// first and last order dates. Revenue lives in orders and may not be joined
// (ADR-0055: no cross-app foreign key, grant or view), so it is pulled on a
// schedule and stored here — one row per organization, all-zero for a
// company that never ordered, so that a "never bought anything" rule has
// something to match. The customer-value list: sort by `revenue_365d` for the
// best customers, filter `last_order_at` for the dormant ones. Every row
// carries `computed_at`, and a row is only as current as the last refresh —
// `GET /customers/organization_metrics/freshness` says how stale the set is
// before a number is shown to anybody.
func (srv *CustomersOrganizations) CustomersOrganizationMetricsList(optionalSetters ...CustomersOrganizationMetricsListOption)(*interface{}, error) {
	path := "/v1/customers/organization_metrics"
	options := CustomersOrganizationMetricsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["OrderCount"] {
		params["order_count"] = options.OrderCount
	}
	if options.enabledSetters["OrderCount30d"] {
		params["order_count_30d"] = options.OrderCount30d
	}
	if options.enabledSetters["OrderCount90d"] {
		params["order_count_90d"] = options.OrderCount90d
	}
	if options.enabledSetters["OrderCount365d"] {
		params["order_count_365d"] = options.OrderCount365d
	}
	if options.enabledSetters["RevenueTotal"] {
		params["revenue_total"] = options.RevenueTotal
	}
	if options.enabledSetters["Revenue30d"] {
		params["revenue_30d"] = options.Revenue30d
	}
	if options.enabledSetters["Revenue90d"] {
		params["revenue_90d"] = options.Revenue90d
	}
	if options.enabledSetters["Revenue365d"] {
		params["revenue_365d"] = options.Revenue365d
	}
	if options.enabledSetters["AvgOrderValue"] {
		params["avg_order_value"] = options.AvgOrderValue
	}
	if options.enabledSetters["AvgOrderValue365d"] {
		params["avg_order_value_365d"] = options.AvgOrderValue365d
	}
	if options.enabledSetters["FirstOrderAt"] {
		params["first_order_at"] = options.FirstOrderAt
	}
	if options.enabledSetters["LastOrderAt"] {
		params["last_order_at"] = options.LastOrderAt
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["CurrencyMixed"] {
		params["currency_mixed"] = options.CurrencyMixed
	}
	if options.enabledSetters["OrdersAsOf"] {
		params["orders_as_of"] = options.OrdersAsOf
	}
	if options.enabledSetters["ComputedAt"] {
		params["computed_at"] = options.ComputedAt
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

// CustomersOrganizationMetricsFreshness the projection is materialized, so it
// is only as true as its last refresh. This is that fact as one answer: the
// OLDEST computed_at in the table (the floor, not an average), the anchor
// those numbers were measured from, and how many organizations are not
// covered at all yet.
func (srv *CustomersOrganizations) CustomersOrganizationMetricsFreshness()(*models.OrganizationMetricsFreshness, error) {
	path := "/v1/customers/organization_metrics/freshness"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.OrganizationMetricsFreshness{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrganizationMetricsFreshness
	parsed, ok := resp.Result.(models.OrganizationMetricsFreshness)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersOrganizationMetricsRefreshOptions struct {
	AsOf string
	Cursor string
	OrganizationIds []string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationMetricsRefreshOptions) New() *CustomersOrganizationMetricsRefreshOptions {
	options.enabledSetters = map[string]bool{
		"AsOf": false,
		"Cursor": false,
		"OrganizationIds": false,
	}
	return &options
}
type CustomersOrganizationMetricsRefreshOption func(*CustomersOrganizationMetricsRefreshOptions)
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsRefreshAsOf(v string) CustomersOrganizationMetricsRefreshOption {
	return func(o *CustomersOrganizationMetricsRefreshOptions) {
		o.AsOf = v
		o.enabledSetters["AsOf"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsRefreshCursor(v string) CustomersOrganizationMetricsRefreshOption {
	return func(o *CustomersOrganizationMetricsRefreshOptions) {
		o.Cursor = v
		o.enabledSetters["Cursor"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationMetricsRefreshOrganizationIds(v []string) CustomersOrganizationMetricsRefreshOption {
	return func(o *CustomersOrganizationMetricsRefreshOptions) {
		o.OrganizationIds = v
		o.enabledSetters["OrganizationIds"] = true
	}
}
	
// CustomersOrganizationMetricsRefresh revenue lives in the orders app and
// cannot be joined (ADR-0055: no cross-app FK, grant or view), so it is
// PULLED: this route walks organizations in id order, asks
// orders.reports.customer-rollup about a batch of them at a time and
// materializes the answer into organization_metrics — one row per
// organization, all-zero for those that never ordered, so that 'never bought'
// rules match something. Rows are only rewritten when a value actually
// changed, so a routine refresh costs almost no writes. Bounded by a
// wall-clock budget below the gateway's upstream timeout: while 'done' is
// false, POST again with the returned 'cursor' AND 'as_of' (pinning as_of is
// what stops the rolling windows sliding during a multi-call refresh).
// 'organization_ids' refreshes exactly those organizations in a single call
// — the targeted path after a customer ordered.
func (srv *CustomersOrganizations) CustomersOrganizationMetricsRefresh(optionalSetters ...CustomersOrganizationMetricsRefreshOption)(*models.Error, error) {
	path := "/v1/customers/organization_metrics/refresh"
	options := CustomersOrganizationMetricsRefreshOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["AsOf"] {
		params["as_of"] = options.AsOf
	}
	if options.enabledSetters["Cursor"] {
		params["cursor"] = options.Cursor
	}
	if options.enabledSetters["OrganizationIds"] {
		params["organization_ids"] = options.OrganizationIds
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
	
// CustomersOrganizationMetricsGet what an organization has BOUGHT,
// materialized into this app from the orders app: lifetime revenue, revenue
// over the last 30/90/365 days, order count, average order value, and the
// first and last order dates. Revenue lives in orders and may not be joined
// (ADR-0055: no cross-app foreign key, grant or view), so it is pulled on a
// schedule and stored here — one row per organization, all-zero for a
// company that never ordered, so that a "never bought anything" rule has
// something to match. One company's numbers by the metrics row id. All zeroes
// mean the company has never ordered, not that the projection is missing —
// a missing row means the refresh has not reached that company yet.
func (srv *CustomersOrganizations) CustomersOrganizationMetricsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/organization_metrics/{id}")
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
type CustomersOrganizationsListOptions struct {
	Id string
	Name string
	VatId string
	Branche string
	CustomerNumber string
	Status string
	LifecycleStage string
	PaymentTerms string
	CreditLimit float64
	PriceList string
	DeliveryBlock bool
	ExternalTeamId string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsListOptions) New() *CustomersOrganizationsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Name": false,
		"VatId": false,
		"Branche": false,
		"CustomerNumber": false,
		"Status": false,
		"LifecycleStage": false,
		"PaymentTerms": false,
		"CreditLimit": false,
		"PriceList": false,
		"DeliveryBlock": false,
		"ExternalTeamId": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersOrganizationsListOption func(*CustomersOrganizationsListOptions)
func (srv *CustomersOrganizations) WithCustomersOrganizationsListId(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListName(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListVatId(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListBranche(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Branche = v
		o.enabledSetters["Branche"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListCustomerNumber(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.CustomerNumber = v
		o.enabledSetters["CustomerNumber"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListStatus(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListLifecycleStage(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.LifecycleStage = v
		o.enabledSetters["LifecycleStage"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListPaymentTerms(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.PaymentTerms = v
		o.enabledSetters["PaymentTerms"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListCreditLimit(v float64) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.CreditLimit = v
		o.enabledSetters["CreditLimit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListPriceList(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.PriceList = v
		o.enabledSetters["PriceList"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListDeliveryBlock(v bool) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.DeliveryBlock = v
		o.enabledSetters["DeliveryBlock"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListExternalTeamId(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.ExternalTeamId = v
		o.enabledSetters["ExternalTeamId"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListCreatedAt(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListUpdatedAt(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListLimit(v int) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListOffset(v int) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsListOrder(v string) CustomersOrganizationsListOption {
	return func(o *CustomersOrganizationsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersOrganizationsList an organization is a buying COMPANY — the unit
// a contract, a credit limit, a price list and a payment term belong to, and
// the unit an order is placed on behalf of. It is not a household and not a
// person: the people are `contacts`, and a company with no contacts yet is a
// perfectly normal row. Every organization is mirrored into platform auth as
// a team, so a name written here is the name storefront authentication shows.
// The company list a sales or service desk works from, and the read a segment
// rule is written against. Every column of the table is a filter and the page
// is `limit`/`offset`/`order` — including the two that are constantly
// confused: `status` is ACCESS (active or blocked) and `lifecycle_stage` is
// the sales PIPELINE, so filtering the wrong one answers with the wrong
// companies rather than with an error.
func (srv *CustomersOrganizations) CustomersOrganizationsList(optionalSetters ...CustomersOrganizationsListOption)(*interface{}, error) {
	path := "/v1/customers/organizations"
	options := CustomersOrganizationsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["VatId"] {
		params["vat_id"] = options.VatId
	}
	if options.enabledSetters["Branche"] {
		params["branche"] = options.Branche
	}
	if options.enabledSetters["CustomerNumber"] {
		params["customer_number"] = options.CustomerNumber
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["LifecycleStage"] {
		params["lifecycle_stage"] = options.LifecycleStage
	}
	if options.enabledSetters["PaymentTerms"] {
		params["payment_terms"] = options.PaymentTerms
	}
	if options.enabledSetters["CreditLimit"] {
		params["credit_limit"] = options.CreditLimit
	}
	if options.enabledSetters["PriceList"] {
		params["price_list"] = options.PriceList
	}
	if options.enabledSetters["DeliveryBlock"] {
		params["delivery_block"] = options.DeliveryBlock
	}
	if options.enabledSetters["ExternalTeamId"] {
		params["external_team_id"] = options.ExternalTeamId
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
type CustomersOrganizationsCreateOptions struct {
	Branche string
	CreditLimit float64
	CustomerNumber string
	DeliveryBlock bool
	LifecycleStage string
	PaymentTerms string
	PriceList string
	Settings interface{}
	Status string
	VatId string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsCreateOptions) New() *CustomersOrganizationsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Branche": false,
		"CreditLimit": false,
		"CustomerNumber": false,
		"DeliveryBlock": false,
		"LifecycleStage": false,
		"PaymentTerms": false,
		"PriceList": false,
		"Settings": false,
		"Status": false,
		"VatId": false,
	}
	return &options
}
type CustomersOrganizationsCreateOption func(*CustomersOrganizationsCreateOptions)
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateBranche(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.Branche = v
		o.enabledSetters["Branche"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateCreditLimit(v float64) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.CreditLimit = v
		o.enabledSetters["CreditLimit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateCustomerNumber(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.CustomerNumber = v
		o.enabledSetters["CustomerNumber"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateDeliveryBlock(v bool) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.DeliveryBlock = v
		o.enabledSetters["DeliveryBlock"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateLifecycleStage(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.LifecycleStage = v
		o.enabledSetters["LifecycleStage"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreatePaymentTerms(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.PaymentTerms = v
		o.enabledSetters["PaymentTerms"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreatePriceList(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.PriceList = v
		o.enabledSetters["PriceList"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateSettings(v interface{}) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateStatus(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsCreateVatId(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
			
// CustomersOrganizationsCreate an organization is a buying COMPANY — the
// unit a contract, a credit limit, a price list and a payment term belong to,
// and the unit an order is placed on behalf of. It is not a household and not
// a person: the people are `contacts`, and a company with no contacts yet is
// a perfectly normal row. Every organization is mirrored into platform auth
// as a team, so a name written here is the name storefront authentication
// shows. Registers a company as a customer. It is mirrored into platform auth
// as a team in the same call, so a failure of the identity service fails the
// create rather than leaving half a company behind. `payment_terms` and
// `lifecycle_stage` name values from this tenant's own sets, and a newly
// founded company inherits the tenant's `default_payment_terms` /
// `default_credit_limit` where the merchant set them. `name` is the only
// field a create cannot omit; everything else is optional or defaulted by the
// database. Two rows of this tenant may not share `customer_number` (while
// customer_number IS NOT NULL) or `external_team_id` (while external_team_id
// IS NOT NULL).
func (srv *CustomersOrganizations) CustomersOrganizationsCreate(Name string, optionalSetters ...CustomersOrganizationsCreateOption)(*models.Error, error) {
	path := "/v1/customers/organizations"
	options := CustomersOrganizationsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Branche"] {
		params["branche"] = options.Branche
	}
	if options.enabledSetters["CreditLimit"] {
		params["credit_limit"] = options.CreditLimit
	}
	if options.enabledSetters["CustomerNumber"] {
		params["customer_number"] = options.CustomerNumber
	}
	if options.enabledSetters["DeliveryBlock"] {
		params["delivery_block"] = options.DeliveryBlock
	}
	if options.enabledSetters["LifecycleStage"] {
		params["lifecycle_stage"] = options.LifecycleStage
	}
	if options.enabledSetters["PaymentTerms"] {
		params["payment_terms"] = options.PaymentTerms
	}
	if options.enabledSetters["PriceList"] {
		params["price_list"] = options.PriceList
	}
	if options.enabledSetters["Settings"] {
		params["settings"] = options.Settings
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["VatId"] {
		params["vat_id"] = options.VatId
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
	
// CustomersOrganizationsDelete an organization is a buying COMPANY — the
// unit a contract, a credit limit, a price list and a payment term belong to,
// and the unit an order is placed on behalf of. It is not a household and not
// a person: the people are `contacts`, and a company with no contacts yet is
// a perfectly normal row. Every organization is mirrored into platform auth
// as a team, so a name written here is the name storefront authentication
// shows. Removes the company and its mirrored team. Its people are NOT
// deleted: they become standalone buyers who can still sign in and still
// order, which is the behaviour a merchant winding down a subsidiary wants.
// Deleting one takes every `contact_events`, `addresses`,
// `organization_metrics` and `segment_members` row that points at it with it
// and clears `contacts.organization_id` rather than deleting those rows —
// the foreign keys decide, not this route.
func (srv *CustomersOrganizations) CustomersOrganizationsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/organizations/{id}")
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
	
// CustomersOrganizationsGet an organization is a buying COMPANY — the unit
// a contract, a credit limit, a price list and a payment term belong to, and
// the unit an order is placed on behalf of. It is not a household and not a
// person: the people are `contacts`, and a company with no contacts yet is a
// perfectly normal row. Every organization is mirrored into platform auth as
// a team, so a name written here is the name storefront authentication shows.
// One company by id, with its commercial terms as stored. What it has BOUGHT
// is not in here — that is the `organization_metrics` row for the same id,
// refreshed on its own schedule.
func (srv *CustomersOrganizations) CustomersOrganizationsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/organizations/{id}")
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
type CustomersOrganizationsUpdateOptions struct {
	Branche string
	CreditLimit float64
	CustomerNumber string
	DeliveryBlock bool
	LifecycleStage string
	Name string
	PaymentTerms string
	PriceList string
	Settings interface{}
	Status string
	VatId string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsUpdateOptions) New() *CustomersOrganizationsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Branche": false,
		"CreditLimit": false,
		"CustomerNumber": false,
		"DeliveryBlock": false,
		"LifecycleStage": false,
		"Name": false,
		"PaymentTerms": false,
		"PriceList": false,
		"Settings": false,
		"Status": false,
		"VatId": false,
	}
	return &options
}
type CustomersOrganizationsUpdateOption func(*CustomersOrganizationsUpdateOptions)
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateBranche(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Branche = v
		o.enabledSetters["Branche"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateCreditLimit(v float64) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.CreditLimit = v
		o.enabledSetters["CreditLimit"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateCustomerNumber(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.CustomerNumber = v
		o.enabledSetters["CustomerNumber"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateDeliveryBlock(v bool) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.DeliveryBlock = v
		o.enabledSetters["DeliveryBlock"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateLifecycleStage(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.LifecycleStage = v
		o.enabledSetters["LifecycleStage"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateName(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdatePaymentTerms(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.PaymentTerms = v
		o.enabledSetters["PaymentTerms"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdatePriceList(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.PriceList = v
		o.enabledSetters["PriceList"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateSettings(v interface{}) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateStatus(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *CustomersOrganizations) WithCustomersOrganizationsUpdateVatId(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
			
// CustomersOrganizationsUpdate an organization is a buying COMPANY — the
// unit a contract, a credit limit, a price list and a payment term belong to,
// and the unit an order is placed on behalf of. It is not a household and not
// a person: the people are `contacts`, and a company with no contacts yet is
// a perfectly normal row. Every organization is mirrored into platform auth
// as a team, so a name written here is the name storefront authentication
// shows. A partial update — send only what changes. `external_team_id` is
// mirror-managed and ignored if sent. Blocking a company here is what stops
// it trading; moving it through the pipeline is `lifecycle_stage`, and the
// two are independent. Two rows of this tenant may not share
// `customer_number` (while customer_number IS NOT NULL) or `external_team_id`
// (while external_team_id IS NOT NULL).
func (srv *CustomersOrganizations) CustomersOrganizationsUpdate(Id string, optionalSetters ...CustomersOrganizationsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/organizations/{id}")
	options := CustomersOrganizationsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Branche"] {
		params["branche"] = options.Branche
	}
	if options.enabledSetters["CreditLimit"] {
		params["credit_limit"] = options.CreditLimit
	}
	if options.enabledSetters["CustomerNumber"] {
		params["customer_number"] = options.CustomerNumber
	}
	if options.enabledSetters["DeliveryBlock"] {
		params["delivery_block"] = options.DeliveryBlock
	}
	if options.enabledSetters["LifecycleStage"] {
		params["lifecycle_stage"] = options.LifecycleStage
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["PaymentTerms"] {
		params["payment_terms"] = options.PaymentTerms
	}
	if options.enabledSetters["PriceList"] {
		params["price_list"] = options.PriceList
	}
	if options.enabledSetters["Settings"] {
		params["settings"] = options.Settings
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["VatId"] {
		params["vat_id"] = options.VatId
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
