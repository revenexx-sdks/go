package carts_io

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CartsIo service
type CartsIo struct {
	client client.Client
}

func New(clt client.Client) *CartsIo {
	return &CartsIo{
		client: clt,
	}
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
func (srv *CartsIo) WithCartsImportContactId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *CartsIo) WithCartsImportCsv(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Csv = v
		o.enabledSetters["Csv"] = true
	}
}
func (srv *CartsIo) WithCartsImportName(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsIo) WithCartsImportPayload(v interface{}) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.Payload = v
		o.enabledSetters["Payload"] = true
	}
}
func (srv *CartsIo) WithCartsImportProfileId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
func (srv *CartsIo) WithCartsImportSessionKey(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.SessionKey = v
		o.enabledSetters["SessionKey"] = true
	}
}
func (srv *CartsIo) WithCartsImportTargetCartId(v string) CartsImportOption {
	return func(o *CartsImportOptions) {
		o.TargetCartId = v
		o.enabledSetters["TargetCartId"] = true
	}
}
	
// CartsImport reads a payload of lines into a cart — the bulk-order path a
// buyer pastes a spreadsheet into. With `target_cart_id` the lines land in
// that cart, which must be active, and the profile's `apply_mode` decides
// what happens to the lines already there: 'replace' clears them first,
// 'insert' and 'append' both add. Without a target a new cart is created, and
// an OWNER is then required — `contact_id` or `session_key` — because a
// cart with neither cannot exist. `profile_id` names an IMPORT profile;
// without one the payload is read ad hoc, as CSV when `csv` is present and as
// JSON otherwise. The lines fold into identical product lines exactly as
// carts.items.create does, so `imported_lines` counts the lines READ and the
// cart may have gained fewer rows than that. A payload that parses to no line
// at all is a 400 rather than a quiet no-op.
func (srv *CartsIo) CartsImport(optionalSetters ...CartsImportOption)(*models.Error, error) {
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
type CartsIoProfilesListOptions struct {
	Id string
	Name string
	Direction string
	Entity string
	Format string
	ApplyMode string
	IsTemplate bool
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CartsIoProfilesListOptions) New() *CartsIoProfilesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Name": false,
		"Direction": false,
		"Entity": false,
		"Format": false,
		"ApplyMode": false,
		"IsTemplate": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CartsIoProfilesListOption func(*CartsIoProfilesListOptions)
func (srv *CartsIo) WithCartsIoProfilesListId(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListName(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListDirection(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Direction = v
		o.enabledSetters["Direction"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListEntity(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListFormat(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListApplyMode(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListIsTemplate(v bool) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.IsTemplate = v
		o.enabledSetters["IsTemplate"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListCreatedAt(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListUpdatedAt(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListLimit(v int) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListOffset(v int) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesListOrder(v string) CartsIoProfilesListOption {
	return func(o *CartsIoProfilesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CartsIoProfilesList the filters are what make this list usable:
// `?direction=export` is how a client offers the profiles that carts.export
// will accept, and `?is_template=true` separates the four bundled templates
// from what a merchant wrote. An unknown column is dropped rather than
// refused — `filter` echoes what was understood.
func (srv *CartsIo) CartsIoProfilesList(optionalSetters ...CartsIoProfilesListOption)(*models.Error, error) {
	path := "/v1/carts/io/profiles"
	options := CartsIoProfilesListOptions{}.New()
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
	if options.enabledSetters["Direction"] {
		params["direction"] = options.Direction
	}
	if options.enabledSetters["Entity"] {
		params["entity"] = options.Entity
	}
	if options.enabledSetters["Format"] {
		params["format"] = options.Format
	}
	if options.enabledSetters["ApplyMode"] {
		params["apply_mode"] = options.ApplyMode
	}
	if options.enabledSetters["IsTemplate"] {
		params["is_template"] = options.IsTemplate
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
func (srv *CartsIo) WithCartsIoProfilesCreateApplyMode(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesCreateEntity(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesCreateFormat(v string) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesCreateIsTemplate(v bool) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.IsTemplate = v
		o.enabledSetters["IsTemplate"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesCreateMapping(v interface{}) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesCreateOptions(v interface{}) CartsIoProfilesCreateOption {
	return func(o *CartsIoProfilesCreateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
					
// CartsIoProfilesCreate defines a new import/export profile. Two fields are
// required and have no default — `name`, which must be unique within the
// tenant, and `direction`, which fixes the one way this profile will ever
// run. Everything else defaults to the common case: whole carts, JSON,
// `apply_mode` 'insert', not a template. The uniqueness of the name is a
// unique index rather than a check in this app, so a reused name is a 409 no
// matter which route wrote the other one, including the four bundled
// templates. The shape is Baseline-IO-compatible, so a mapping written for
// another app's import reads the same way here. Creating a profile does not
// move any data: carts.export and carts.import are what execute one, and each
// refuses a profile pointed the wrong way.
func (srv *CartsIo) CartsIoProfilesCreate(Direction string, Name string, optionalSetters ...CartsIoProfilesCreateOption)(*models.Error, error) {
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

// CartsIoProfilesDefaults seeds the 4 bundled templates and reports which of
// them it had to create — the call that gives a fresh tenant something to
// export through before anybody has written a profile. Idempotent and matched
// by NAME, so a second call answers with everything under 'existing' and
// writes nothing, and a template a merchant has edited is left exactly as
// they left it rather than reset. It also runs by itself on app.installed;
// call it by hand where that event cannot be relied on, and after deleting a
// template to get it back.
func (srv *CartsIo) CartsIoProfilesDefaults()(*interface{}, error) {
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
	
// CartsIoProfilesDelete removes a profile. Nothing in this app points at one
// — no cart and no line stores the profile it was imported through — so
// no foreign key holds the delete up and nothing is orphaned by it; what
// breaks is the caller still holding that `profile_id`, which answers 404 on
// its next run. Deleting one of the four bundled templates is not permanent
// either: the next carts.io.profiles.defaults, and the next install of this
// app, seeds it again by name, in the shape it ships with rather than the
// shape a merchant had edited it into.
func (srv *CartsIo) CartsIoProfilesDelete(Id string)(*models.Error, error) {
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
	
// CartsIoProfilesGet one profile by id — the id carts.export and
// carts.import name in `profile_id`. Read it to see what a run will do before
// starting one: `direction`, because a profile only ever runs the way it
// declares; `entity`, whole carts or bare lines; `format`, where json
// round-trips and csv carries line fields only; `mapping`, what the external
// columns are called; and `apply_mode`, which decides what an import does
// with the lines a target cart already has. `is_template` says whether this
// is one of the four the app ships with or something a merchant wrote.
// Reading a profile runs nothing and changes nothing.
func (srv *CartsIo) CartsIoProfilesGet(Id string)(*models.Error, error) {
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
func (srv *CartsIo) WithCartsIoProfilesUpdateApplyMode(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateDirection(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Direction = v
		o.enabledSetters["Direction"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateEntity(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateFormat(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateIsTemplate(v bool) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.IsTemplate = v
		o.enabledSetters["IsTemplate"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateMapping(v interface{}) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateName(v string) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *CartsIo) WithCartsIoProfilesUpdateOptions(v interface{}) CartsIoProfilesUpdateOption {
	return func(o *CartsIoProfilesUpdateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
			
// CartsIoProfilesUpdate edits a profile in place, the four bundled templates
// included — seeding matches on name and never rewrites what it finds, so
// an edit made here survives every later call to carts.io.profiles.defaults
// and every reinstall of the app. The name stays unique in the tenant, so
// renaming onto another profile's name is a 409, and a payload carrying no
// updatable field answers 400 rather than storing nothing quietly. Runs that
// already happened are unaffected: a profile is read at the moment
// carts.export or carts.import executes and nothing is kept pointing back at
// it, so changing a mapping changes the next run and no earlier one.
func (srv *CartsIo) CartsIoProfilesUpdate(Id string, optionalSetters ...CartsIoProfilesUpdateOption)(*models.Error, error) {
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
func (srv *CartsIo) WithCartsExportFormat(v string) CartsExportOption {
	return func(o *CartsExportOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *CartsIo) WithCartsExportProfileId(v string) CartsExportOption {
	return func(o *CartsExportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
			
// CartsExport renders one cart as a document somebody can take away. With
// `profile_id` the named EXPORT profile decides the format, the entity and
// the column names; handing it an import profile is a 400, because a profile
// only runs the way it declares. Without one the call runs ad hoc — JSON,
// unless `format: 'csv'` says otherwise. The JSON form is `{cart: {…},
// items: […]}` and is exactly what carts.import takes back, so an export
// round-trips; the CSV form is the lines only, header first, and drops
// everything that lives on the cart rather than on a line. Nothing is stored
// and nothing about the cart changes — `filename` is a suggestion for a
// browser download, not a file this app keeps — and a cart of any status
// can be exported, including one already ordered.
func (srv *CartsIo) CartsExport(Id string, optionalSetters ...CartsExportOption)(*models.Error, error) {
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
