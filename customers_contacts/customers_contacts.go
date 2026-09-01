package customers_contacts

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CustomersContacts service
type CustomersContacts struct {
	client client.Client
}

func New(clt client.Client) *CustomersContacts {
	return &CustomersContacts{
		client: clt,
	}
}

type CustomersContactEventsListOptions struct {
	Id string
	ContactId string
	OrganizationId string
	Kind string
	Name string
	Subject string
	Actor string
	OccurredAt string
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersContactEventsListOptions) New() *CustomersContactEventsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"ContactId": false,
		"OrganizationId": false,
		"Kind": false,
		"Name": false,
		"Subject": false,
		"Actor": false,
		"OccurredAt": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersContactEventsListOption func(*CustomersContactEventsListOptions)
func (srv *CustomersContacts) WithCustomersContactEventsListId(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListContactId(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListOrganizationId(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListKind(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListName(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListSubject(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListActor(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Actor = v
		o.enabledSetters["Actor"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListOccurredAt(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.OccurredAt = v
		o.enabledSetters["OccurredAt"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListCreatedAt(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListLimit(v int) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListOffset(v int) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactEventsListOrder(v string) CustomersContactEventsListOption {
	return func(o *CustomersContactEventsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersContactEventsList a contact event is one entry on a customer's
// timeline: an activity somebody logged (a call, a visit, a meeting, a note)
// or a registration decision this app recorded itself. Every entry is keyed
// by a CONTACT and stamped with the organization derived from that contact,
// so a company's history is one indexed read rather than a join. Append-only
// — there is no update and no delete, which is what makes it usable as
// evidence. The activity feed, filtered by whichever column the question
// needs: `contact_id` for one person, `organization_id` for a whole company,
// `kind` for one type of activity. `kind: "system"` is this app's own
// registration decision trail (`registration.submitted` / `.approved` /
// `.rejected`), and no caller may file one of those. Paged with
// `limit`/`offset`/`order`; newest first is `order=occurred_at.desc`.
func (srv *CustomersContacts) CustomersContactEventsList(optionalSetters ...CustomersContactEventsListOption)(*interface{}, error) {
	path := "/v1/customers/contact_events"
	options := CustomersContactEventsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["Actor"] {
		params["actor"] = options.Actor
	}
	if options.enabledSetters["OccurredAt"] {
		params["occurred_at"] = options.OccurredAt
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
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
	
// CustomersContactEventsGet a contact event is one entry on a customer's
// timeline: an activity somebody logged (a call, a visit, a meeting, a note)
// or a registration decision this app recorded itself. Every entry is keyed
// by a CONTACT and stamped with the organization derived from that contact,
// so a company's history is one indexed read rather than a join. Append-only
// — there is no update and no delete, which is what makes it usable as
// evidence. One timeline entry by id, as it was written. Entries are never
// edited, so what this answers is what was recorded at the time.
func (srv *CustomersContacts) CustomersContactEventsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/contact_events/{id}")
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
type CustomersContactsListOptions struct {
	Id string
	OrganizationId string
	Email string
	FirstName string
	LastName string
	Phone string
	JobTitle string
	Role string
	Status string
	OrderApprovalLimit float64
	RegistrationStatus string
	RegistrationDecidedAt string
	RegistrationDecidedBy string
	RegistrationReason string
	Locale string
	IsPrimary bool
	ExternalUserId string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersContactsListOptions) New() *CustomersContactsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"OrganizationId": false,
		"Email": false,
		"FirstName": false,
		"LastName": false,
		"Phone": false,
		"JobTitle": false,
		"Role": false,
		"Status": false,
		"OrderApprovalLimit": false,
		"RegistrationStatus": false,
		"RegistrationDecidedAt": false,
		"RegistrationDecidedBy": false,
		"RegistrationReason": false,
		"Locale": false,
		"IsPrimary": false,
		"ExternalUserId": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersContactsListOption func(*CustomersContactsListOptions)
func (srv *CustomersContacts) WithCustomersContactsListId(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListOrganizationId(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListEmail(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListFirstName(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListLastName(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListPhone(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListJobTitle(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.JobTitle = v
		o.enabledSetters["JobTitle"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListRole(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Role = v
		o.enabledSetters["Role"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListStatus(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListOrderApprovalLimit(v float64) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.OrderApprovalLimit = v
		o.enabledSetters["OrderApprovalLimit"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListRegistrationStatus(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.RegistrationStatus = v
		o.enabledSetters["RegistrationStatus"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListRegistrationDecidedAt(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.RegistrationDecidedAt = v
		o.enabledSetters["RegistrationDecidedAt"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListRegistrationDecidedBy(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.RegistrationDecidedBy = v
		o.enabledSetters["RegistrationDecidedBy"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListRegistrationReason(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.RegistrationReason = v
		o.enabledSetters["RegistrationReason"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListLocale(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListIsPrimary(v bool) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.IsPrimary = v
		o.enabledSetters["IsPrimary"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListExternalUserId(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.ExternalUserId = v
		o.enabledSetters["ExternalUserId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListCreatedAt(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListUpdatedAt(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListLimit(v int) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListOffset(v int) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsListOrder(v string) CustomersContactsListOption {
	return func(o *CustomersContactsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersContactsList a contact is a PERSON, and the unit that logs in: one
// platform user, one email address, one role held inside its organization. A
// contact without an organization is a standalone buyer rather than an error,
// and two people at the same company are two contacts sharing an
// `organization_id`. The people list, and the read behind an approval queue:
// `registration_status=pending` is every application waiting for a decision.
// Every column is a filter — `external_user_id` in particular is how a
// storefront turns a platform auth id back into a customer — and the page
// is `limit`/`offset`/`order`.
func (srv *CustomersContacts) CustomersContactsList(optionalSetters ...CustomersContactsListOption)(*interface{}, error) {
	path := "/v1/customers/contacts"
	options := CustomersContactsListOptions{}.New()
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
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["FirstName"] {
		params["first_name"] = options.FirstName
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["JobTitle"] {
		params["job_title"] = options.JobTitle
	}
	if options.enabledSetters["Role"] {
		params["role"] = options.Role
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["OrderApprovalLimit"] {
		params["order_approval_limit"] = options.OrderApprovalLimit
	}
	if options.enabledSetters["RegistrationStatus"] {
		params["registration_status"] = options.RegistrationStatus
	}
	if options.enabledSetters["RegistrationDecidedAt"] {
		params["registration_decided_at"] = options.RegistrationDecidedAt
	}
	if options.enabledSetters["RegistrationDecidedBy"] {
		params["registration_decided_by"] = options.RegistrationDecidedBy
	}
	if options.enabledSetters["RegistrationReason"] {
		params["registration_reason"] = options.RegistrationReason
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["IsPrimary"] {
		params["is_primary"] = options.IsPrimary
	}
	if options.enabledSetters["ExternalUserId"] {
		params["external_user_id"] = options.ExternalUserId
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
type CustomersContactsCreateOptions struct {
	FirstName string
	IsPrimary bool
	JobTitle string
	LastName string
	Locale string
	OrderApprovalLimit float64
	OrganizationId string
	Phone string
	RegistrationStatus string
	Role string
	Status string
	enabledSetters map[string]bool
}
func (options CustomersContactsCreateOptions) New() *CustomersContactsCreateOptions {
	options.enabledSetters = map[string]bool{
		"FirstName": false,
		"IsPrimary": false,
		"JobTitle": false,
		"LastName": false,
		"Locale": false,
		"OrderApprovalLimit": false,
		"OrganizationId": false,
		"Phone": false,
		"RegistrationStatus": false,
		"Role": false,
		"Status": false,
	}
	return &options
}
type CustomersContactsCreateOption func(*CustomersContactsCreateOptions)
func (srv *CustomersContacts) WithCustomersContactsCreateFirstName(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateIsPrimary(v bool) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.IsPrimary = v
		o.enabledSetters["IsPrimary"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateJobTitle(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.JobTitle = v
		o.enabledSetters["JobTitle"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateLastName(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateLocale(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateOrderApprovalLimit(v float64) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.OrderApprovalLimit = v
		o.enabledSetters["OrderApprovalLimit"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateOrganizationId(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreatePhone(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateRegistrationStatus(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.RegistrationStatus = v
		o.enabledSetters["RegistrationStatus"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateRole(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Role = v
		o.enabledSetters["Role"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsCreateStatus(v string) CustomersContactsCreateOption {
	return func(o *CustomersContactsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// CustomersContactsCreate a contact is a PERSON, and the unit that logs in:
// one platform user, one email address, one role held inside its
// organization. A contact without an organization is a standalone buyer
// rather than an error, and two people at the same company are two contacts
// sharing an `organization_id`. Creates the person and their platform login
// together, so a contact that exists can always sign in. `role` names one of
// this tenant's own roles and decides what they may do; `registration_status`
// may only be set to `pending` or `approved` here, because a rejection has to
// carry a reason and that is the reject route's job. `email` is the only
// field a create cannot omit; everything else is optional or defaulted by the
// database. Two rows of this tenant may not share `email` or
// `external_user_id` (while external_user_id IS NOT NULL).
func (srv *CustomersContacts) CustomersContactsCreate(Email string, optionalSetters ...CustomersContactsCreateOption)(*models.Error, error) {
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
	if options.enabledSetters["JobTitle"] {
		params["job_title"] = options.JobTitle
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrderApprovalLimit"] {
		params["order_approval_limit"] = options.OrderApprovalLimit
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["RegistrationStatus"] {
		params["registration_status"] = options.RegistrationStatus
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
type CustomersContactsEventsCreateOptions struct {
	Actor string
	Kind string
	Note string
	OccurredAt string
	enabledSetters map[string]bool
}
func (options CustomersContactsEventsCreateOptions) New() *CustomersContactsEventsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Actor": false,
		"Kind": false,
		"Note": false,
		"OccurredAt": false,
	}
	return &options
}
type CustomersContactsEventsCreateOption func(*CustomersContactsEventsCreateOptions)
func (srv *CustomersContacts) WithCustomersContactsEventsCreateActor(v string) CustomersContactsEventsCreateOption {
	return func(o *CustomersContactsEventsCreateOptions) {
		o.Actor = v
		o.enabledSetters["Actor"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsEventsCreateKind(v string) CustomersContactsEventsCreateOption {
	return func(o *CustomersContactsEventsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsEventsCreateNote(v string) CustomersContactsEventsCreateOption {
	return func(o *CustomersContactsEventsCreateOptions) {
		o.Note = v
		o.enabledSetters["Note"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsEventsCreateOccurredAt(v string) CustomersContactsEventsCreateOption {
	return func(o *CustomersContactsEventsCreateOptions) {
		o.OccurredAt = v
		o.enabledSetters["OccurredAt"] = true
	}
}
					
// CustomersContactsEventsCreate this is how a call, a visit, a meeting, an
// email or a plain note reaches one person's timeline. It writes a
// contact_events row with kind != 'system' and emits contact_event.created,
// so an activity travels on the same bus as a registration decision and a
// timeline is one query rather than a union. organization_id is DERIVED from
// the contact, never taken from the body — an activity cannot be filed
// under a company the person does not belong to.
func (srv *CustomersContacts) CustomersContactsEventsCreate(ContactId string, Subject string, optionalSetters ...CustomersContactsEventsCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{contact_id}", ContactId)
	path := r.Replace("/v1/customers/contacts/{contact_id}/events")
	options := CustomersContactsEventsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	params["subject"] = Subject
	if options.enabledSetters["Actor"] {
		params["actor"] = options.Actor
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Note"] {
		params["note"] = options.Note
	}
	if options.enabledSetters["OccurredAt"] {
		params["occurred_at"] = options.OccurredAt
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
type CustomersContactsInviteOptions struct {
	InvitedBy string
	enabledSetters map[string]bool
}
func (options CustomersContactsInviteOptions) New() *CustomersContactsInviteOptions {
	options.enabledSetters = map[string]bool{
		"InvitedBy": false,
	}
	return &options
}
type CustomersContactsInviteOption func(*CustomersContactsInviteOptions)
func (srv *CustomersContacts) WithCustomersContactsInviteInvitedBy(v string) CustomersContactsInviteOption {
	return func(o *CustomersContactsInviteOptions) {
		o.InvitedBy = v
		o.enabledSetters["InvitedBy"] = true
	}
}
					
// CustomersContactsInvite tell somebody they were added to a company. A
// deliberate act rather than a side effect of creating the contact: a
// merchant entering a colleague from a business card is not always ready to
// mail them, and "added" and "told" are different decisions. No secret
// travels — the platform team membership is confirmed as it is created, so
// there is nothing to accept; the message says "you are in, here is the way
// in". Unlike the auth mails, a failure here IS a failure: the identity
// service sends nothing for this occasion, so this is the only message the
// person gets.
func (srv *CustomersContacts) CustomersContactsInvite(ContactId string, Url string, optionalSetters ...CustomersContactsInviteOption)(*models.Error, error) {
	r := strings.NewReplacer("{contact_id}", ContactId)
	path := r.Replace("/v1/customers/contacts/{contact_id}/invite")
	options := CustomersContactsInviteOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	params["url"] = Url
	if options.enabledSetters["InvitedBy"] {
		params["invited_by"] = options.InvitedBy
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
	
// CustomersContactsPermissions computed from contacts.role on every call —
// the grants are never persisted, so this always reflects the role the
// contact holds right now.
func (srv *CustomersContacts) CustomersContactsPermissions(ContactId string)(*models.Error, error) {
	r := strings.NewReplacer("{contact_id}", ContactId)
	path := r.Replace("/v1/customers/contacts/{contact_id}/permissions")
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
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
type CustomersRegistrationsApproveOptions struct {
	DecidedBy string
	enabledSetters map[string]bool
}
func (options CustomersRegistrationsApproveOptions) New() *CustomersRegistrationsApproveOptions {
	options.enabledSetters = map[string]bool{
		"DecidedBy": false,
	}
	return &options
}
type CustomersRegistrationsApproveOption func(*CustomersRegistrationsApproveOptions)
func (srv *CustomersContacts) WithCustomersRegistrationsApproveDecidedBy(v string) CustomersRegistrationsApproveOption {
	return func(o *CustomersRegistrationsApproveOptions) {
		o.DecidedBy = v
		o.enabledSetters["DecidedBy"] = true
	}
}
			
// CustomersRegistrationsApprove only reachable for a contact whose
// registration_status is 'pending' or 'rejected' (approving a rejection
// reinstates it). Enables the platform user FIRST — the password the
// applicant chose at submit time works immediately, no new credential is
// issued — then sets registration_status='approved' and status='active',
// and un-blocks the organization this registration itself founded. Approving
// an already-approved registration is a no-op that emits nothing, so a retry
// is safe. Writes a contact_events row named 'registration.approved'.
func (srv *CustomersContacts) CustomersRegistrationsApprove(ContactId string, optionalSetters ...CustomersRegistrationsApproveOption)(*models.Error, error) {
	r := strings.NewReplacer("{contact_id}", ContactId)
	path := r.Replace("/v1/customers/contacts/{contact_id}/registration/approve")
	options := CustomersRegistrationsApproveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	if options.enabledSetters["DecidedBy"] {
		params["decided_by"] = options.DecidedBy
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
type CustomersRegistrationsRejectOptions struct {
	DecidedBy string
	enabledSetters map[string]bool
}
func (options CustomersRegistrationsRejectOptions) New() *CustomersRegistrationsRejectOptions {
	options.enabledSetters = map[string]bool{
		"DecidedBy": false,
	}
	return &options
}
type CustomersRegistrationsRejectOption func(*CustomersRegistrationsRejectOptions)
func (srv *CustomersContacts) WithCustomersRegistrationsRejectDecidedBy(v string) CustomersRegistrationsRejectOption {
	return func(o *CustomersRegistrationsRejectOptions) {
		o.DecidedBy = v
		o.enabledSetters["DecidedBy"] = true
	}
}
					
// CustomersRegistrationsReject only reachable from 'pending'. Sets
// registration_status='rejected' and status='blocked', keeps the platform
// user in place but disabled — the email must not fall free for a silent
// second identity, and the merchant keeps the record. Delete the contact to
// remove both. 'reason' is mandatory and is stored on the contact plus
// carried in the event payload, so the applicant can be told why. Rejecting
// an already-rejected registration is a no-op. Writes a contact_events row
// named 'registration.rejected'.
func (srv *CustomersContacts) CustomersRegistrationsReject(ContactId string, Reason string, optionalSetters ...CustomersRegistrationsRejectOption)(*models.Error, error) {
	r := strings.NewReplacer("{contact_id}", ContactId)
	path := r.Replace("/v1/customers/contacts/{contact_id}/registration/reject")
	options := CustomersRegistrationsRejectOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	params["reason"] = Reason
	if options.enabledSetters["DecidedBy"] {
		params["decided_by"] = options.DecidedBy
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
	
// CustomersContactsDelete a contact is a PERSON, and the unit that logs in:
// one platform user, one email address, one role held inside its
// organization. A contact without an organization is a standalone buyer
// rather than an error, and two people at the same company are two contacts
// sharing an `organization_id`. Removes the person and their platform login,
// so they can no longer sign in anywhere. Their company keeps trading; use
// `status: "blocked"` instead when the intent is to stop one person without
// erasing what they did. Deleting one takes every `contact_events` and
// `addresses` row that points at it with it — the foreign keys decide, not
// this route.
func (srv *CustomersContacts) CustomersContactsDelete(Id string)(*models.Error, error) {
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
	
// CustomersContactsGet a contact is a PERSON, and the unit that logs in: one
// platform user, one email address, one role held inside its organization. A
// contact without an organization is a standalone buyer rather than an error,
// and two people at the same company are two contacts sharing an
// `organization_id`. One person by id. What they are ALLOWED to do is not in
// here: permissions are derived from `role` at read time and answered by `GET
// /customers/contacts/{contact_id}/permissions`.
func (srv *CustomersContacts) CustomersContactsGet(Id string)(*models.Error, error) {
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
type CustomersContactsUpdateOptions struct {
	Email string
	FirstName string
	IsPrimary bool
	JobTitle string
	LastName string
	Locale string
	OrderApprovalLimit float64
	OrganizationId string
	Phone string
	RegistrationStatus string
	Role string
	Status string
	enabledSetters map[string]bool
}
func (options CustomersContactsUpdateOptions) New() *CustomersContactsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Email": false,
		"FirstName": false,
		"IsPrimary": false,
		"JobTitle": false,
		"LastName": false,
		"Locale": false,
		"OrderApprovalLimit": false,
		"OrganizationId": false,
		"Phone": false,
		"RegistrationStatus": false,
		"Role": false,
		"Status": false,
	}
	return &options
}
type CustomersContactsUpdateOption func(*CustomersContactsUpdateOptions)
func (srv *CustomersContacts) WithCustomersContactsUpdateEmail(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateFirstName(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateIsPrimary(v bool) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.IsPrimary = v
		o.enabledSetters["IsPrimary"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateJobTitle(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.JobTitle = v
		o.enabledSetters["JobTitle"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateLastName(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateLocale(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateOrderApprovalLimit(v float64) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.OrderApprovalLimit = v
		o.enabledSetters["OrderApprovalLimit"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateOrganizationId(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdatePhone(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateRegistrationStatus(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.RegistrationStatus = v
		o.enabledSetters["RegistrationStatus"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateRole(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Role = v
		o.enabledSetters["Role"] = true
	}
}
func (srv *CustomersContacts) WithCustomersContactsUpdateStatus(v string) CustomersContactsUpdateOption {
	return func(o *CustomersContactsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// CustomersContactsUpdate a contact is a PERSON, and the unit that logs in:
// one platform user, one email address, one role held inside its
// organization. A contact without an organization is a standalone buyer
// rather than an error, and two people at the same company are two contacts
// sharing an `organization_id`. A partial update — send only what changes.
// `external_user_id` and every `registration_*` column are ignored: the link
// to platform auth is mirror-managed, and registration state is only ever
// moved by the approve and reject routes, which record why. Two rows of this
// tenant may not share `email` or `external_user_id` (while external_user_id
// IS NOT NULL).
func (srv *CustomersContacts) CustomersContactsUpdate(Id string, optionalSetters ...CustomersContactsUpdateOption)(*models.Error, error) {
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
	if options.enabledSetters["JobTitle"] {
		params["job_title"] = options.JobTitle
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrderApprovalLimit"] {
		params["order_approval_limit"] = options.OrderApprovalLimit
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["RegistrationStatus"] {
		params["registration_status"] = options.RegistrationStatus
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
type CustomersOrganizationsEventsCreateOptions struct {
	Actor string
	Kind string
	Note string
	OccurredAt string
	enabledSetters map[string]bool
}
func (options CustomersOrganizationsEventsCreateOptions) New() *CustomersOrganizationsEventsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Actor": false,
		"Kind": false,
		"Note": false,
		"OccurredAt": false,
	}
	return &options
}
type CustomersOrganizationsEventsCreateOption func(*CustomersOrganizationsEventsCreateOptions)
func (srv *CustomersContacts) WithCustomersOrganizationsEventsCreateActor(v string) CustomersOrganizationsEventsCreateOption {
	return func(o *CustomersOrganizationsEventsCreateOptions) {
		o.Actor = v
		o.enabledSetters["Actor"] = true
	}
}
func (srv *CustomersContacts) WithCustomersOrganizationsEventsCreateKind(v string) CustomersOrganizationsEventsCreateOption {
	return func(o *CustomersOrganizationsEventsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *CustomersContacts) WithCustomersOrganizationsEventsCreateNote(v string) CustomersOrganizationsEventsCreateOption {
	return func(o *CustomersOrganizationsEventsCreateOptions) {
		o.Note = v
		o.enabledSetters["Note"] = true
	}
}
func (srv *CustomersContacts) WithCustomersOrganizationsEventsCreateOccurredAt(v string) CustomersOrganizationsEventsCreateOption {
	return func(o *CustomersOrganizationsEventsCreateOptions) {
		o.OccurredAt = v
		o.enabledSetters["OccurredAt"] = true
	}
}
							
// CustomersOrganizationsEventsCreate same row as the contact route, reached
// from the organization. 'contact_id' is required and must belong to THIS
// organization — the picker offering the contacts is not filtered, so the
// membership check here is what stops a call with one company being filed
// under someone else's person.
func (srv *CustomersContacts) CustomersOrganizationsEventsCreate(OrganizationId string, ContactId string, Subject string, optionalSetters ...CustomersOrganizationsEventsCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{organization_id}", OrganizationId)
	path := r.Replace("/v1/customers/organizations/{organization_id}/events")
	options := CustomersOrganizationsEventsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organization_id"] = OrganizationId
	params["contact_id"] = ContactId
	params["subject"] = Subject
	if options.enabledSetters["Actor"] {
		params["actor"] = options.Actor
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Note"] {
		params["note"] = options.Note
	}
	if options.enabledSetters["OccurredAt"] {
		params["occurred_at"] = options.OccurredAt
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
