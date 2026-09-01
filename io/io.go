package io

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Io service
type Io struct {
	client client.Client
}

func New(clt client.Client) *Io {
	return &Io{
		client: clt,
	}
}

type ListBulkJobsOptions struct {
	Type interface{}
	Status interface{}
	Vendor string
	App string
	Entity string
	Limit int
	enabledSetters map[string]bool
}
func (options ListBulkJobsOptions) New() *ListBulkJobsOptions {
	options.enabledSetters = map[string]bool{
		"Type": false,
		"Status": false,
		"Vendor": false,
		"App": false,
		"Entity": false,
		"Limit": false,
	}
	return &options
}
type ListBulkJobsOption func(*ListBulkJobsOptions)
func (srv *Io) WithListBulkJobsType(v interface{}) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Io) WithListBulkJobsStatus(v interface{}) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Io) WithListBulkJobsVendor(v string) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.Vendor = v
		o.enabledSetters["Vendor"] = true
	}
}
func (srv *Io) WithListBulkJobsApp(v string) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.App = v
		o.enabledSetters["App"] = true
	}
}
func (srv *Io) WithListBulkJobsEntity(v string) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.Entity = v
		o.enabledSetters["Entity"] = true
	}
}
func (srv *Io) WithListBulkJobsLimit(v int) ListBulkJobsOption {
	return func(o *ListBulkJobsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
	
// ListBulkJobs the calling tenant's bulk jobs, newest first. Jobs are created
// by the
// feature blocks (import / export / A/B swap / tenant copy / sample) —
// never here; this surface is read-only.
func (srv *Io) ListBulkJobs(optionalSetters ...ListBulkJobsOption)(*models.ValidationFailedResponse, error) {
	path := "/v1/io/bulk-jobs"
	options := ListBulkJobsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Vendor"] {
		params["vendor"] = options.Vendor
	}
	if options.enabledSetters["App"] {
		params["app"] = options.App
	}
	if options.enabledSetters["Entity"] {
		params["entity"] = options.Entity
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetBulkJob status, row counts, and progress for one bulk job.
// 
// Tenant-scoped: an id belonging to another tenant is filtered out and
// is therefore indistinguishable from a non-existent one — which is the
// intent.
func (srv *Io) GetBulkJob(Id string)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/bulk-jobs/{id}")
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListIoEntities flat list of the entities the calling tenant's installed
// apps expose,
// sorted by vendor, app, entity. Feeds the entity pickers of the
// Integration Studio I/O nodes.
// 
// The app set comes from `baseline.tenant_app_versions`. Per app the
// entity list is resolved from the tenant's pinned schema version; when
// that pointer is stale (missing or not applied) it falls back to the
// latest applied version of `(vendor, app)`. Apps with no applied
// schema at all contribute no entities.
func (srv *Io) ListIoEntities()(*models.ValidationFailedResponse, error) {
	path := "/v1/io/entities"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateExportOptions struct {
	Format string
	ProfileId string
	enabledSetters map[string]bool
}
func (options CreateExportOptions) New() *CreateExportOptions {
	options.enabledSetters = map[string]bool{
		"Format": false,
		"ProfileId": false,
	}
	return &options
}
type CreateExportOption func(*CreateExportOptions)
func (srv *Io) WithCreateExportFormat(v string) CreateExportOption {
	return func(o *CreateExportOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *Io) WithCreateExportProfileId(v string) CreateExportOption {
	return func(o *CreateExportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
							
// CreateExport creates a `bulk_job` and dispatches the engine to export the
// tenant's
// rows for an entity. CSV/XML stream row-by-row into an S3 multipart
// upload (flat RAM); JSON/XLSX are buffered. The response carries the
// object key the result will be written to.
func (srv *Io) CreateExport(App string, Entity string, Vendor string, optionalSetters ...CreateExportOption)(*models.ValidationFailedResponse, error) {
	path := "/v1/io/exports"
	options := CreateExportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["app"] = App
	params["entity"] = Entity
	params["vendor"] = Vendor
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetExportUrl mints a short-TTL signed S3 `GET` URL for the object a
// completed
// export wrote. Tenant-scoped: an id belonging to another tenant — or
// to a job that is not an export — is indistinguishable from a
// non-existent one and answers `404`.
// 
// The job must have reached `completed` or `partial`; any earlier
// state answers `409` and carries the current `job_status`.
func (srv *Io) GetExportUrl(Id string)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/exports/{id}/url")
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateImportOptions struct {
	Format string
	Keys []string
	MaxRejects int
	Mode string
	ProfileId string
	Target string
	enabledSetters map[string]bool
}
func (options CreateImportOptions) New() *CreateImportOptions {
	options.enabledSetters = map[string]bool{
		"Format": false,
		"Keys": false,
		"MaxRejects": false,
		"Mode": false,
		"ProfileId": false,
		"Target": false,
	}
	return &options
}
type CreateImportOption func(*CreateImportOptions)
func (srv *Io) WithCreateImportFormat(v string) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.Format = v
		o.enabledSetters["Format"] = true
	}
}
func (srv *Io) WithCreateImportKeys(v []string) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.Keys = v
		o.enabledSetters["Keys"] = true
	}
}
func (srv *Io) WithCreateImportMaxRejects(v int) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.MaxRejects = v
		o.enabledSetters["MaxRejects"] = true
	}
}
func (srv *Io) WithCreateImportMode(v string) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.Mode = v
		o.enabledSetters["Mode"] = true
	}
}
func (srv *Io) WithCreateImportProfileId(v string) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.ProfileId = v
		o.enabledSetters["ProfileId"] = true
	}
}
func (srv *Io) WithCreateImportTarget(v string) CreateImportOption {
	return func(o *CreateImportOptions) {
		o.Target = v
		o.enabledSetters["Target"] = true
	}
}
									
// CreateImport creates a `bulk_job` and dispatches the engine to import a
// previously
// uploaded object into the named entity. The engine streams CSV
// row-by-row (flat RAM at 1M+ rows) and COPYs into the entity's staging
// sibling before a merge / content-hash delta into the target.
func (srv *Io) CreateImport(App string, Entity string, ObjectKey string, Vendor string, optionalSetters ...CreateImportOption)(*models.ValidationFailedResponse, error) {
	path := "/v1/io/imports"
	options := CreateImportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["app"] = App
	params["entity"] = Entity
	params["object_key"] = ObjectKey
	params["vendor"] = Vendor
	if options.enabledSetters["Format"] {
		params["format"] = options.Format
	}
	if options.enabledSetters["Keys"] {
		params["keys"] = options.Keys
	}
	if options.enabledSetters["MaxRejects"] {
		params["max_rejects"] = options.MaxRejects
	}
	if options.enabledSetters["Mode"] {
		params["mode"] = options.Mode
	}
	if options.enabledSetters["ProfileId"] {
		params["profile_id"] = options.ProfileId
	}
	if options.enabledSetters["Target"] {
		params["target"] = options.Target
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListProfiles the calling tenant's saved profiles, ordered by name.
// 
// When `X-Revenexx-Market` is present the listing is filtered to the
// profiles offered for that market — global profiles (`markets: null`)
// plus those whose `markets` contain it. Omit the header to get every
// profile, which is what the management view wants.
func (srv *Io) ListProfiles()(*models.ValidationFailedResponse, error) {
	path := "/v1/io/profiles"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateProfileOptions struct {
	ApplyMode string
	Mapping interface{}
	Markets []string
	Options interface{}
	enabledSetters map[string]bool
}
func (options CreateProfileOptions) New() *CreateProfileOptions {
	options.enabledSetters = map[string]bool{
		"ApplyMode": false,
		"Mapping": false,
		"Markets": false,
		"Options": false,
	}
	return &options
}
type CreateProfileOption func(*CreateProfileOptions)
func (srv *Io) WithCreateProfileApplyMode(v string) CreateProfileOption {
	return func(o *CreateProfileOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *Io) WithCreateProfileMapping(v interface{}) CreateProfileOption {
	return func(o *CreateProfileOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *Io) WithCreateProfileMarkets(v []string) CreateProfileOption {
	return func(o *CreateProfileOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Io) WithCreateProfileOptions(v interface{}) CreateProfileOption {
	return func(o *CreateProfileOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
													
// CreateProfile a tenant-secured, reusable mapping (field rename + transforms
// + keys)
// for a direction (`import`/`export`), format, and entity. Runnable
// on-click via `/io/profiles/{id}/run`.
func (srv *Io) CreateProfile(App string, Direction string, Entity string, Format string, Name string, Vendor string, optionalSetters ...CreateProfileOption)(*models.ValidationFailedResponse, error) {
	path := "/v1/io/profiles"
	options := CreateProfileOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["app"] = App
	params["direction"] = Direction
	params["entity"] = Entity
	params["format"] = Format
	params["name"] = Name
	params["vendor"] = Vendor
	if options.enabledSetters["ApplyMode"] {
		params["apply_mode"] = options.ApplyMode
	}
	if options.enabledSetters["Mapping"] {
		params["mapping"] = options.Mapping
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteProfile permanently remove a saved profile owned by the calling
// tenant.
// 
// Idempotent, and deliberately not a `404` path: deleting an id that
// does not belong to the tenant still answers `200`, with `deleted: 0`.
func (srv *Io) DeleteProfile(Id string)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/profiles/{id}")
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ShowProfile a single saved profile. Tenant-scoped: an id owned by another
// tenant
// is indistinguishable from a non-existent one and answers `404`.
func (srv *Io) ShowProfile(Id string)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/profiles/{id}")
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateProfileOptions struct {
	ApplyMode string
	Mapping interface{}
	Markets []string
	Options interface{}
	enabledSetters map[string]bool
}
func (options UpdateProfileOptions) New() *UpdateProfileOptions {
	options.enabledSetters = map[string]bool{
		"ApplyMode": false,
		"Mapping": false,
		"Markets": false,
		"Options": false,
	}
	return &options
}
type UpdateProfileOption func(*UpdateProfileOptions)
func (srv *Io) WithUpdateProfileApplyMode(v string) UpdateProfileOption {
	return func(o *UpdateProfileOptions) {
		o.ApplyMode = v
		o.enabledSetters["ApplyMode"] = true
	}
}
func (srv *Io) WithUpdateProfileMapping(v interface{}) UpdateProfileOption {
	return func(o *UpdateProfileOptions) {
		o.Mapping = v
		o.enabledSetters["Mapping"] = true
	}
}
func (srv *Io) WithUpdateProfileMarkets(v []string) UpdateProfileOption {
	return func(o *UpdateProfileOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Io) WithUpdateProfileOptions(v interface{}) UpdateProfileOption {
	return func(o *UpdateProfileOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
															
// UpdateProfile replace a saved profile's mapping, format, or apply mode
// (tenant-scoped).
func (srv *Io) UpdateProfile(Id string, App string, Direction string, Entity string, Format string, Name string, Vendor string, optionalSetters ...UpdateProfileOption)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/profiles/{id}")
	options := UpdateProfileOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["app"] = App
	params["direction"] = Direction
	params["entity"] = Entity
	params["format"] = Format
	params["name"] = Name
	params["vendor"] = Vendor
	if options.enabledSetters["ApplyMode"] {
		params["apply_mode"] = options.ApplyMode
	}
	if options.enabledSetters["Mapping"] {
		params["mapping"] = options.Mapping
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type RunProfileOptions struct {
	Markets []string
	ObjectKey string
	enabledSetters map[string]bool
}
func (options RunProfileOptions) New() *RunProfileOptions {
	options.enabledSetters = map[string]bool{
		"Markets": false,
		"ObjectKey": false,
	}
	return &options
}
type RunProfileOption func(*RunProfileOptions)
func (srv *Io) WithRunProfileMarkets(v []string) RunProfileOption {
	return func(o *RunProfileOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Io) WithRunProfileObjectKey(v string) RunProfileOption {
	return func(o *RunProfileOptions) {
		o.ObjectKey = v
		o.enabledSetters["ObjectKey"] = true
	}
}
			
// RunProfile dispatches the engine using the saved profile. An import run
// requires
// `object_key` (upload first); an export run writes a generated key.
func (srv *Io) RunProfile(Id string, optionalSetters ...RunProfileOption)(*models.ValidationFailedResponse, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/io/profiles/{id}/run")
	options := RunProfileOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
	}
	if options.enabledSetters["ObjectKey"] {
		params["object_key"] = options.ObjectKey
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateUploadOptions struct {
	Extension string
	enabledSetters map[string]bool
}
func (options CreateUploadOptions) New() *CreateUploadOptions {
	options.enabledSetters = map[string]bool{
		"Extension": false,
	}
	return &options
}
type CreateUploadOption func(*CreateUploadOptions)
func (srv *Io) WithCreateUploadExtension(v string) CreateUploadOption {
	return func(o *CreateUploadOptions) {
		o.Extension = v
		o.enabledSetters["Extension"] = true
	}
}
	
// CreateUpload returns a short-lived signed S3 `PUT` URL (+ required headers)
// and
// the `object_key` to reference in a subsequent `/io/imports`. The
// client uploads bytes directly to object storage — never through
// Baseline.
func (srv *Io) CreateUpload(optionalSetters ...CreateUploadOption)(*models.ValidationFailedResponse, error) {
	path := "/v1/io/uploads"
	options := CreateUploadOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Extension"] {
		params["extension"] = options.Extension
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

		parsed := models.ValidationFailedResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ValidationFailedResponse
	parsed, ok := resp.Result.(models.ValidationFailedResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
