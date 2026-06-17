package customers

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Customers service
type Customers struct {
	client client.Client
}

func New(clt client.Client) *Customers {
	return &Customers{
		client: clt,
	}
}


// CustomersAddressesList
func (srv *Customers) CustomersAddressesList()(*interface{}, error) {
	path := "/v1/customers/addresses"
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
func (srv *Customers) WithCustomersAddressesCreateCompany(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Company = v
		o.enabledSetters["Company"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateContactId(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateIsDefault(v bool) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateName(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateOrganizationId(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreatePhone(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateRegion(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateStreet2(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Street2 = v
		o.enabledSetters["Street2"] = true
	}
}
func (srv *Customers) WithCustomersAddressesCreateType(v string) CustomersAddressesCreateOption {
	return func(o *CustomersAddressesCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
									
// CustomersAddressesCreate
func (srv *Customers) CustomersAddressesCreate(City string, Country string, Street string, Zip string, optionalSetters ...CustomersAddressesCreateOption)(*models.Address, error) {
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

		parsed := models.Address{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Address
	parsed, ok := resp.Result.(models.Address)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersAddressesDelete
func (srv *Customers) CustomersAddressesDelete(Id string)(*interface{}, error) {
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
	
// CustomersAddressesGet
func (srv *Customers) CustomersAddressesGet(Id string)(*models.Address, error) {
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

		parsed := models.Address{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Address
	parsed, ok := resp.Result.(models.Address)
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
func (srv *Customers) WithCustomersAddressesUpdateCity(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.City = v
		o.enabledSetters["City"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateCompany(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Company = v
		o.enabledSetters["Company"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateContactId(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateCountry(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateIsDefault(v bool) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateName(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateOrganizationId(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdatePhone(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateRegion(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateStreet(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Street = v
		o.enabledSetters["Street"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateStreet2(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Street2 = v
		o.enabledSetters["Street2"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateType(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Customers) WithCustomersAddressesUpdateZip(v string) CustomersAddressesUpdateOption {
	return func(o *CustomersAddressesUpdateOptions) {
		o.Zip = v
		o.enabledSetters["Zip"] = true
	}
}
			
// CustomersAddressesUpdate
func (srv *Customers) CustomersAddressesUpdate(Id string, optionalSetters ...CustomersAddressesUpdateOption)(*models.Address, error) {
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

		parsed := models.Address{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Address
	parsed, ok := resp.Result.(models.Address)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthLogin
func (srv *Customers) CustomersAuthLogin(Email string, Password string)(*models.AuthLoginResponse, error) {
	path := "/v1/customers/auth/login"
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.AuthLoginResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AuthLoginResponse
	parsed, ok := resp.Result.(models.AuthLoginResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthLogout
func (srv *Customers) CustomersAuthLogout(SessionId string, UserId string)(*interface{}, error) {
	path := "/v1/customers/auth/logout"
	params := map[string]interface{}{}
	params["session_id"] = SessionId
	params["user_id"] = UserId
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
type CustomersAuthMeOptions struct {
	SessionId string
	enabledSetters map[string]bool
}
func (options CustomersAuthMeOptions) New() *CustomersAuthMeOptions {
	options.enabledSetters = map[string]bool{
		"SessionId": false,
	}
	return &options
}
type CustomersAuthMeOption func(*CustomersAuthMeOptions)
func (srv *Customers) WithCustomersAuthMeSessionId(v string) CustomersAuthMeOption {
	return func(o *CustomersAuthMeOptions) {
		o.SessionId = v
		o.enabledSetters["SessionId"] = true
	}
}
			
// CustomersAuthMe
func (srv *Customers) CustomersAuthMe(UserId string, optionalSetters ...CustomersAuthMeOption)(*models.AuthMeResponse, error) {
	path := "/v1/customers/auth/me"
	options := CustomersAuthMeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["user_id"] = UserId
	if options.enabledSetters["SessionId"] {
		params["session_id"] = options.SessionId
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

		parsed := models.AuthMeResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AuthMeResponse
	parsed, ok := resp.Result.(models.AuthMeResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthRecovery
func (srv *Customers) CustomersAuthRecovery(Email string, Url string)(*interface{}, error) {
	path := "/v1/customers/auth/recovery"
	params := map[string]interface{}{}
	params["email"] = Email
	params["url"] = Url
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
					
// CustomersAuthRecoveryConfirm
func (srv *Customers) CustomersAuthRecoveryConfirm(Password string, Secret string, UserId string)(*interface{}, error) {
	path := "/v1/customers/auth/recovery"
	params := map[string]interface{}{}
	params["password"] = Password
	params["secret"] = Secret
	params["user_id"] = UserId
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
type CustomersAuthRegisterOptions struct {
	FirstName string
	LastName string
	Locale string
	OrganizationId string
	OrganizationName string
	enabledSetters map[string]bool
}
func (options CustomersAuthRegisterOptions) New() *CustomersAuthRegisterOptions {
	options.enabledSetters = map[string]bool{
		"FirstName": false,
		"LastName": false,
		"Locale": false,
		"OrganizationId": false,
		"OrganizationName": false,
	}
	return &options
}
type CustomersAuthRegisterOption func(*CustomersAuthRegisterOptions)
func (srv *Customers) WithCustomersAuthRegisterFirstName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterLastName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterLocale(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterOrganizationId(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterOrganizationName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.OrganizationName = v
		o.enabledSetters["OrganizationName"] = true
	}
}
					
// CustomersAuthRegister
func (srv *Customers) CustomersAuthRegister(Email string, Password string, optionalSetters ...CustomersAuthRegisterOption)(*models.AuthRegisterResponse, error) {
	path := "/v1/customers/auth/register"
	options := CustomersAuthRegisterOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["FirstName"] {
		params["first_name"] = options.FirstName
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["OrganizationName"] {
		params["organization_name"] = options.OrganizationName
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

		parsed := models.AuthRegisterResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AuthRegisterResponse
	parsed, ok := resp.Result.(models.AuthRegisterResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CustomersContactsList
func (srv *Customers) CustomersContactsList()(*interface{}, error) {
	path := "/v1/customers/contacts"
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
type CustomersContactsCreateOptions struct {
	FirstName string
	IsPrimary bool
	LastName string
	Locale string
	OrganizationId string
	Phone string
	Role string
	Status string
	enabledSetters map[string]bool
}
func (options CustomersContactsCreateOptions) New() *CustomersContactsCreateOptions {
	options.enabledSetters = map[string]bool{
		"FirstName": false,
		"IsPrimary": false,
		"LastName": false,
		"Locale": false,
		"OrganizationId": false,
		"Phone": false,
		"Role": false,
		"Status": false,
	}
	return &options
}
type CustomersContactsCreateOption func(*CustomersContactsCreateOptions)
func (srv *Customers) WithCustomersContactsCreateFirstName(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateIsPrimary(v bool) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.IsPrimary = v
		o.enabledSetters["IsPrimary"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateLastName(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateLocale(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateOrganizationId(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreatePhone(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateRole(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Role = v
		o.enabledSetters["Role"] = true
	}
}
func (srv *Customers) WithCustomersContactsCreateStatus(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// CustomersContactsCreate
func (srv *Customers) CustomersContactsCreate(Email string, optionalSetters ...CustomersContactsCreateOption)(*models.Contact, error) {
	path := "/v1/customers/contacts"
	options := CustomersContactsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["email"] = Email
	if options.enabledSetters["FirstName"] {
		params["first_name"] = options.FirstName
	}
	if options.enabledSetters["IsPrimary"] {
		params["is_primary"] = options.IsPrimary
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Role"] {
		params["role"] = options.Role
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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

		parsed := models.Contact{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Contact
	parsed, ok := resp.Result.(models.Contact)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersContactsDelete
func (srv *Customers) CustomersContactsDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contacts/{id}")
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
	
// CustomersContactsGet
func (srv *Customers) CustomersContactsGet(Id string)(*models.Contact, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contacts/{id}")
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

		parsed := models.Contact{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Contact
	parsed, ok := resp.Result.(models.Contact)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersContactsUpdateOptions struct {
	Email string
	FirstName string
	IsPrimary bool
	LastName string
	Locale string
	OrganizationId string
	Phone string
	Role string
	Status string
	enabledSetters map[string]bool
}
func (options CustomersContactsUpdateOptions) New() *CustomersContactsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Email": false,
		"FirstName": false,
		"IsPrimary": false,
		"LastName": false,
		"Locale": false,
		"OrganizationId": false,
		"Phone": false,
		"Role": false,
		"Status": false,
	}
	return &options
}
type CustomersContactsUpdateOption func(*CustomersContactsUpdateOptions)
func (srv *Customers) WithCustomersContactsUpdateEmail(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateFirstName(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateIsPrimary(v bool) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.IsPrimary = v
		o.enabledSetters["IsPrimary"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateLastName(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateLocale(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateOrganizationId(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdatePhone(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateRole(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Role = v
		o.enabledSetters["Role"] = true
	}
}
func (srv *Customers) WithCustomersContactsUpdateStatus(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// CustomersContactsUpdate
func (srv *Customers) CustomersContactsUpdate(Id string, optionalSetters ...CustomersContactsUpdateOption)(*models.Contact, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contacts/{id}")
	options := CustomersContactsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["FirstName"] {
		params["first_name"] = options.FirstName
	}
	if options.enabledSetters["IsPrimary"] {
		params["is_primary"] = options.IsPrimary
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Role"] {
		params["role"] = options.Role
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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

		parsed := models.Contact{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Contact
	parsed, ok := resp.Result.(models.Contact)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CustomersOrganizationsList
func (srv *Customers) CustomersOrganizationsList()(*interface{}, error) {
	path := "/v1/customers/organizations"
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
type CustomersOrganizationsCreateOptions struct {
	Settings interface{}
	Status string
	VatId string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsCreateOptions) New() *CustomersOrganizationsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Settings": false,
		"Status": false,
		"VatId": false,
	}
	return &options
}
type CustomersOrganizationsCreateOption func(*CustomersOrganizationsCreateOptions)
func (srv *Customers) WithCustomersOrganizationsCreateSettings(v interface{}) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
func (srv *Customers) WithCustomersOrganizationsCreateStatus(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Customers) WithCustomersOrganizationsCreateVatId(v string) CustomersOrganizationsCreateOption {
	return func(o *CustomersOrganizationsCreateOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
			
// CustomersOrganizationsCreate
func (srv *Customers) CustomersOrganizationsCreate(Name string, optionalSetters ...CustomersOrganizationsCreateOption)(*models.Organization, error) {
	path := "/v1/customers/organizations"
	options := CustomersOrganizationsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersOrganizationsDelete
func (srv *Customers) CustomersOrganizationsDelete(Id string)(*interface{}, error) {
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
	
// CustomersOrganizationsGet
func (srv *Customers) CustomersOrganizationsGet(Id string)(*models.Organization, error) {
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersOrganizationsUpdateOptions struct {
	Name string
	Settings interface{}
	Status string
	VatId string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsUpdateOptions) New() *CustomersOrganizationsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Settings": false,
		"Status": false,
		"VatId": false,
	}
	return &options
}
type CustomersOrganizationsUpdateOption func(*CustomersOrganizationsUpdateOptions)
func (srv *Customers) WithCustomersOrganizationsUpdateName(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Customers) WithCustomersOrganizationsUpdateSettings(v interface{}) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
func (srv *Customers) WithCustomersOrganizationsUpdateStatus(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Customers) WithCustomersOrganizationsUpdateVatId(v string) CustomersOrganizationsUpdateOption {
	return func(o *CustomersOrganizationsUpdateOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
			
// CustomersOrganizationsUpdate
func (srv *Customers) CustomersOrganizationsUpdate(Id string, optionalSetters ...CustomersOrganizationsUpdateOption)(*models.Organization, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/organizations/{id}")
	options := CustomersOrganizationsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
