package apps

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Apps service
type Apps struct {
	client client.Client
}

func New(clt client.Client) *Apps {
	return &Apps{
		client: clt,
	}
}

type AppsListOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options AppsListOptions) New() *AppsListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type AppsListOption func(*AppsListOptions)
func (srv *Apps) WithAppsListQueries(v []string) AppsListOption {
	return func(o *AppsListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithAppsListSearch(v string) AppsListOption {
	return func(o *AppsListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Apps) WithAppsListTotal(v bool) AppsListOption {
	return func(o *AppsListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// AppsList list all Apps in the active project. Pass `search` to filter by
// name.
func (srv *Apps) AppsList(optionalSetters ...AppsListOption)(*models.FunctionList, error) {
	path := "/v1/apps"
	options := AppsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.FunctionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FunctionList
	parsed, ok := resp.Result.(models.FunctionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateOptions struct {
	Commands string
	Enabled bool
	Entrypoint string
	Events []string
	Execute []string
	InstallationId string
	Logging bool
	ProviderBranch string
	ProviderRepositoryId string
	ProviderRootDirectory string
	ProviderSilentMode bool
	Schedule string
	Scopes []string
	Specification string
	Timeout int
	enabledSetters map[string]bool
}
func (options AppsCreateOptions) New() *AppsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Commands": false,
		"Enabled": false,
		"Entrypoint": false,
		"Events": false,
		"Execute": false,
		"InstallationId": false,
		"Logging": false,
		"ProviderBranch": false,
		"ProviderRepositoryId": false,
		"ProviderRootDirectory": false,
		"ProviderSilentMode": false,
		"Schedule": false,
		"Scopes": false,
		"Specification": false,
		"Timeout": false,
	}
	return &options
}
type AppsCreateOption func(*AppsCreateOptions)
func (srv *Apps) WithAppsCreateCommands(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
func (srv *Apps) WithAppsCreateEnabled(v bool) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Apps) WithAppsCreateEntrypoint(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
func (srv *Apps) WithAppsCreateEvents(v []string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func (srv *Apps) WithAppsCreateExecute(v []string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func (srv *Apps) WithAppsCreateInstallationId(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Apps) WithAppsCreateLogging(v bool) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Apps) WithAppsCreateProviderBranch(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Apps) WithAppsCreateProviderRepositoryId(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Apps) WithAppsCreateProviderRootDirectory(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Apps) WithAppsCreateProviderSilentMode(v bool) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Apps) WithAppsCreateSchedule(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Apps) WithAppsCreateScopes(v []string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Scopes = v
		o.enabledSetters["Scopes"] = true
	}
}
func (srv *Apps) WithAppsCreateSpecification(v string) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Apps) WithAppsCreateTimeout(v int) AppsCreateOption {
	return func(o *AppsCreateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
							
// AppsCreate create a new revenexx App. An App is the deployment surface for
// code that runs on the platform — backend jobs, APIs, integrations. The
// created App owns subsequent deployments and executions.
// 
// Phase 1 mirrors the underlying Functions runtime 1:1; future phases will
// add manifest validation, registry coupling and schema migrations.
func (srv *Apps) AppsCreate(FunctionId string, Name string, Runtime string, optionalSetters ...AppsCreateOption)(*models.Function, error) {
	path := "/v1/apps"
	options := AppsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	params["runtime"] = Runtime
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["ProviderBranch"] {
		params["providerBranch"] = options.ProviderBranch
	}
	if options.enabledSetters["ProviderRepositoryId"] {
		params["providerRepositoryId"] = options.ProviderRepositoryId
	}
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
	}
	if options.enabledSetters["ProviderSilentMode"] {
		params["providerSilentMode"] = options.ProviderSilentMode
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Scopes"] {
		params["scopes"] = options.Scopes
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
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

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsListMarketplaceOptions struct {
	Search string
	PerPage int
	Page int
	enabledSetters map[string]bool
}
func (options AppsListMarketplaceOptions) New() *AppsListMarketplaceOptions {
	options.enabledSetters = map[string]bool{
		"Search": false,
		"PerPage": false,
		"Page": false,
	}
	return &options
}
type AppsListMarketplaceOption func(*AppsListMarketplaceOptions)
func (srv *Apps) WithAppsListMarketplaceSearch(v string) AppsListMarketplaceOption {
	return func(o *AppsListMarketplaceOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Apps) WithAppsListMarketplacePerPage(v int) AppsListMarketplaceOption {
	return func(o *AppsListMarketplaceOptions) {
		o.PerPage = v
		o.enabledSetters["PerPage"] = true
	}
}
func (srv *Apps) WithAppsListMarketplacePage(v int) AppsListMarketplaceOption {
	return func(o *AppsListMarketplaceOptions) {
		o.Page = v
		o.enabledSetters["Page"] = true
	}
}
	
// AppsListMarketplace list apps published to the Marketplace. Proxies the App
// Registry on Console with `?published=true` filter.
func (srv *Apps) AppsListMarketplace(optionalSetters ...AppsListMarketplaceOption)(*interface{}, error) {
	path := "/v1/apps/marketplace"
	options := AppsListMarketplaceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["PerPage"] {
		params["per_page"] = options.PerPage
	}
	if options.enabledSetters["Page"] {
		params["page"] = options.Page
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
			
// AppsInstallFromMarketplace install a Marketplace app on the calling
// project's tenant. Body: { owner, name }.
func (srv *Apps) AppsInstallFromMarketplace(Name string, Owner string)(*interface{}, error) {
	path := "/v1/apps/marketplace/install"
	params := map[string]interface{}{}
	params["name"] = Name
	params["owner"] = Owner
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

// AppsListRuntimes get a list of all runtimes available for an App. Identical
// content to `functions.listRuntimes()`.
func (srv *Apps) AppsListRuntimes()(*models.RuntimeList, error) {
	path := "/v1/apps/runtimes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.RuntimeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RuntimeList
	parsed, ok := resp.Result.(models.RuntimeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// AppsListSpecifications list the compute specifications (CPU + memory)
// available to Apps in this project.
func (srv *Apps) AppsListSpecifications()(*models.SpecificationList, error) {
	path := "/v1/apps/specifications"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SpecificationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SpecificationList
	parsed, ok := resp.Result.(models.SpecificationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsListTemplatesOptions struct {
	Runtimes []string
	UseCases []string
	Limit int
	Offset int
	Total bool
	enabledSetters map[string]bool
}
func (options AppsListTemplatesOptions) New() *AppsListTemplatesOptions {
	options.enabledSetters = map[string]bool{
		"Runtimes": false,
		"UseCases": false,
		"Limit": false,
		"Offset": false,
		"Total": false,
	}
	return &options
}
type AppsListTemplatesOption func(*AppsListTemplatesOptions)
func (srv *Apps) WithAppsListTemplatesRuntimes(v []string) AppsListTemplatesOption {
	return func(o *AppsListTemplatesOptions) {
		o.Runtimes = v
		o.enabledSetters["Runtimes"] = true
	}
}
func (srv *Apps) WithAppsListTemplatesUseCases(v []string) AppsListTemplatesOption {
	return func(o *AppsListTemplatesOptions) {
		o.UseCases = v
		o.enabledSetters["UseCases"] = true
	}
}
func (srv *Apps) WithAppsListTemplatesLimit(v int) AppsListTemplatesOption {
	return func(o *AppsListTemplatesOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Apps) WithAppsListTemplatesOffset(v int) AppsListTemplatesOption {
	return func(o *AppsListTemplatesOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Apps) WithAppsListTemplatesTotal(v bool) AppsListTemplatesOption {
	return func(o *AppsListTemplatesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// AppsListTemplates list the curated catalogue of App templates that can be
// used as starting points.
func (srv *Apps) AppsListTemplates(optionalSetters ...AppsListTemplatesOption)(*models.TemplateFunctionList, error) {
	path := "/v1/apps/templates"
	options := AppsListTemplatesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Runtimes"] {
		params["runtimes"] = options.Runtimes
	}
	if options.enabledSetters["UseCases"] {
		params["useCases"] = options.UseCases
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TemplateFunctionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TemplateFunctionList
	parsed, ok := resp.Result.(models.TemplateFunctionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// AppsGetTemplate get a single App template by its ID.
func (srv *Apps) AppsGetTemplate(TemplateId string)(*models.TemplateFunction, error) {
	r := strings.NewReplacer("{templateId}", TemplateId)
	path := r.Replace("/v1/apps/templates/{templateId}")
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TemplateFunction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TemplateFunction
	parsed, ok := resp.Result.(models.TemplateFunction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsListUsageOptions struct {
	Range string
	enabledSetters map[string]bool
}
func (options AppsListUsageOptions) New() *AppsListUsageOptions {
	options.enabledSetters = map[string]bool{
		"Range": false,
	}
	return &options
}
type AppsListUsageOption func(*AppsListUsageOptions)
func (srv *Apps) WithAppsListUsageRange(v string) AppsListUsageOption {
	return func(o *AppsListUsageOptions) {
		o.Range = v
		o.enabledSetters["Range"] = true
	}
}
	
// AppsListUsage get aggregated usage stats across all Apps in the project for
// the requested time range.
func (srv *Apps) AppsListUsage(optionalSetters ...AppsListUsageOption)(*models.UsageFunctions, error) {
	path := "/v1/apps/usage"
	options := AppsListUsageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Range"] {
		params["range"] = options.Range
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UsageFunctions{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageFunctions
	parsed, ok := resp.Result.(models.UsageFunctions)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// AppsDelete delete an App and all of its deployments. Cascades to the App
// Registry — Console removes the matching `RegisteredApp` row.
func (srv *Apps) AppsDelete(FunctionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
	
// AppsGet get an App by its unique ID.
func (srv *Apps) AppsGet(FunctionId string)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsUpdateOptions struct {
	Commands string
	Enabled bool
	Entrypoint string
	Events []string
	Execute []string
	InstallationId string
	Logging bool
	ProviderBranch string
	ProviderRepositoryId string
	ProviderRootDirectory string
	ProviderSilentMode bool
	Runtime string
	Schedule string
	Scopes []string
	Specification string
	Timeout int
	enabledSetters map[string]bool
}
func (options AppsUpdateOptions) New() *AppsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Commands": false,
		"Enabled": false,
		"Entrypoint": false,
		"Events": false,
		"Execute": false,
		"InstallationId": false,
		"Logging": false,
		"ProviderBranch": false,
		"ProviderRepositoryId": false,
		"ProviderRootDirectory": false,
		"ProviderSilentMode": false,
		"Runtime": false,
		"Schedule": false,
		"Scopes": false,
		"Specification": false,
		"Timeout": false,
	}
	return &options
}
type AppsUpdateOption func(*AppsUpdateOptions)
func (srv *Apps) WithAppsUpdateCommands(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
func (srv *Apps) WithAppsUpdateEnabled(v bool) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Apps) WithAppsUpdateEntrypoint(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
func (srv *Apps) WithAppsUpdateEvents(v []string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func (srv *Apps) WithAppsUpdateExecute(v []string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func (srv *Apps) WithAppsUpdateInstallationId(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Apps) WithAppsUpdateLogging(v bool) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Apps) WithAppsUpdateProviderBranch(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Apps) WithAppsUpdateProviderRepositoryId(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Apps) WithAppsUpdateProviderRootDirectory(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Apps) WithAppsUpdateProviderSilentMode(v bool) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Apps) WithAppsUpdateRuntime(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Runtime = v
		o.enabledSetters["Runtime"] = true
	}
}
func (srv *Apps) WithAppsUpdateSchedule(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Apps) WithAppsUpdateScopes(v []string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Scopes = v
		o.enabledSetters["Scopes"] = true
	}
}
func (srv *Apps) WithAppsUpdateSpecification(v string) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Apps) WithAppsUpdateTimeout(v int) AppsUpdateOption {
	return func(o *AppsUpdateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
					
// AppsUpdate update an App. Use this endpoint to rename, change runtime,
// schedule, environment variables and other configuration.
func (srv *Apps) AppsUpdate(FunctionId string, Name string, optionalSetters ...AppsUpdateOption)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}")
	options := AppsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["ProviderBranch"] {
		params["providerBranch"] = options.ProviderBranch
	}
	if options.enabledSetters["ProviderRepositoryId"] {
		params["providerRepositoryId"] = options.ProviderRepositoryId
	}
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
	}
	if options.enabledSetters["ProviderSilentMode"] {
		params["providerSilentMode"] = options.ProviderSilentMode
	}
	if options.enabledSetters["Runtime"] {
		params["runtime"] = options.Runtime
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Scopes"] {
		params["scopes"] = options.Scopes
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
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

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// AppsUpdateDeployment set the active deployment for an App. The chosen
// deployment must already be `ready`.
func (srv *Apps) AppsUpdateDeployment(FunctionId string, DeploymentId string)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployment")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsListDeploymentsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options AppsListDeploymentsOptions) New() *AppsListDeploymentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type AppsListDeploymentsOption func(*AppsListDeploymentsOptions)
func (srv *Apps) WithAppsListDeploymentsQueries(v []string) AppsListDeploymentsOption {
	return func(o *AppsListDeploymentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithAppsListDeploymentsSearch(v string) AppsListDeploymentsOption {
	return func(o *AppsListDeploymentsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Apps) WithAppsListDeploymentsTotal(v bool) AppsListDeploymentsOption {
	return func(o *AppsListDeploymentsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// AppsListDeployments list the deployment history of an App.
func (srv *Apps) AppsListDeployments(FunctionId string, optionalSetters ...AppsListDeploymentsOption)(*models.DeploymentList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployments")
	options := AppsListDeploymentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.DeploymentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DeploymentList
	parsed, ok := resp.Result.(models.DeploymentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateDeploymentOptions struct {
	Commands string
	Entrypoint string
	enabledSetters map[string]bool
}
func (options AppsCreateDeploymentOptions) New() *AppsCreateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Commands": false,
		"Entrypoint": false,
	}
	return &options
}
type AppsCreateDeploymentOption func(*AppsCreateDeploymentOptions)
func (srv *Apps) WithAppsCreateDeploymentCommands(v string) AppsCreateDeploymentOption {
	return func(o *AppsCreateDeploymentOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
func (srv *Apps) WithAppsCreateDeploymentEntrypoint(v string) AppsCreateDeploymentOption {
	return func(o *AppsCreateDeploymentOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
							
// AppsCreateDeployment upload a new code deployment for an App. Accepts a
// `.tar.gz`
// archive containing the App source. Phase 2 will extract the
// manifest from this archive and validate it against the App
// Registry before kicking off the build.
func (srv *Apps) AppsCreateDeployment(FunctionId string, Activate bool, Code string, optionalSetters ...AppsCreateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployments")
	options := AppsCreateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["activate"] = Activate
	params["code"] = Code
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}


    uploadId := ""

    resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
    if err != nil {
		return nil, err
	}
	var parsed models.Deployment
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil
}
type AppsCreateDuplicateDeploymentOptions struct {
	BuildId string
	enabledSetters map[string]bool
}
func (options AppsCreateDuplicateDeploymentOptions) New() *AppsCreateDuplicateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"BuildId": false,
	}
	return &options
}
type AppsCreateDuplicateDeploymentOption func(*AppsCreateDuplicateDeploymentOptions)
func (srv *Apps) WithAppsCreateDuplicateDeploymentBuildId(v string) AppsCreateDuplicateDeploymentOption {
	return func(o *AppsCreateDuplicateDeploymentOptions) {
		o.BuildId = v
		o.enabledSetters["BuildId"] = true
	}
}
					
// AppsCreateDuplicateDeployment re-deploy an existing build under a new
// deployment ID. Useful for promoting a known-good preview build to
// production without rebuilding.
func (srv *Apps) AppsCreateDuplicateDeployment(FunctionId string, DeploymentId string, optionalSetters ...AppsCreateDuplicateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployments/duplicate")
	options := AppsCreateDuplicateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	if options.enabledSetters["BuildId"] {
		params["buildId"] = options.BuildId
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

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateTemplateDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options AppsCreateTemplateDeploymentOptions) New() *AppsCreateTemplateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type AppsCreateTemplateDeploymentOption func(*AppsCreateTemplateDeploymentOptions)
func (srv *Apps) WithAppsCreateTemplateDeploymentActivate(v bool) AppsCreateTemplateDeploymentOption {
	return func(o *AppsCreateTemplateDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
													
// AppsCreateTemplateDeployment create a new App deployment from a template in
// the App Templates catalogue.
func (srv *Apps) AppsCreateTemplateDeployment(FunctionId string, Owner string, Reference string, Repository string, RootDirectory string, Type string, optionalSetters ...AppsCreateTemplateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployments/template")
	options := AppsCreateTemplateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["owner"] = Owner
	params["reference"] = Reference
	params["repository"] = Repository
	params["rootDirectory"] = RootDirectory
	params["type"] = Type
	if options.enabledSetters["Activate"] {
		params["activate"] = options.Activate
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

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateVcsDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options AppsCreateVcsDeploymentOptions) New() *AppsCreateVcsDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type AppsCreateVcsDeploymentOption func(*AppsCreateVcsDeploymentOptions)
func (srv *Apps) WithAppsCreateVcsDeploymentActivate(v bool) AppsCreateVcsDeploymentOption {
	return func(o *AppsCreateVcsDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
							
// AppsCreateVcsDeployment trigger a new deployment from the App's connected
// Git repository.
func (srv *Apps) AppsCreateVcsDeployment(FunctionId string, Reference string, Type string, optionalSetters ...AppsCreateVcsDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/deployments/vcs")
	options := AppsCreateVcsDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["reference"] = Reference
	params["type"] = Type
	if options.enabledSetters["Activate"] {
		params["activate"] = options.Activate
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

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// AppsDeleteDeployment delete a deployment. The active deployment cannot be
// deleted while it is active — switch first via the deployment-update
// endpoint.
func (srv *Apps) AppsDeleteDeployment(FunctionId string, DeploymentId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/apps/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
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
			
// AppsGetDeployment get a deployment by its unique ID.
func (srv *Apps) AppsGetDeployment(FunctionId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/apps/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsGetDeploymentDownloadOptions struct {
	Type string
	enabledSetters map[string]bool
}
func (options AppsGetDeploymentDownloadOptions) New() *AppsGetDeploymentDownloadOptions {
	options.enabledSetters = map[string]bool{
		"Type": false,
	}
	return &options
}
type AppsGetDeploymentDownloadOption func(*AppsGetDeploymentDownloadOptions)
func (srv *Apps) WithAppsGetDeploymentDownloadType(v string) AppsGetDeploymentDownloadOption {
	return func(o *AppsGetDeploymentDownloadOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// AppsGetDeploymentDownload get a redirect URL to download the source archive
// of an App deployment. Useful for re-running a build locally or auditing
// what was deployed.
func (srv *Apps) AppsGetDeploymentDownload(FunctionId string, DeploymentId string, optionalSetters ...AppsGetDeploymentDownloadOption)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/apps/{functionId}/deployments/{deploymentId}/download")
	options := AppsGetDeploymentDownloadOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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
			
// AppsUpdateDeploymentStatus cancel an in-progress deployment build. Used by
// the Cockpit "Cancel build" affordance.
func (srv *Apps) AppsUpdateDeploymentStatus(FunctionId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/apps/{functionId}/deployments/{deploymentId}/status")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsListExecutionsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options AppsListExecutionsOptions) New() *AppsListExecutionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type AppsListExecutionsOption func(*AppsListExecutionsOptions)
func (srv *Apps) WithAppsListExecutionsQueries(v []string) AppsListExecutionsOption {
	return func(o *AppsListExecutionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithAppsListExecutionsTotal(v bool) AppsListExecutionsOption {
	return func(o *AppsListExecutionsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// AppsListExecutions list the execution history of an App.
func (srv *Apps) AppsListExecutions(FunctionId string, optionalSetters ...AppsListExecutionsOption)(*models.ExecutionList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/executions")
	options := AppsListExecutionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ExecutionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ExecutionList
	parsed, ok := resp.Result.(models.ExecutionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateExecutionOptions struct {
	Async bool
	Body string
	Headers interface{}
	Method string
	Path string
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options AppsCreateExecutionOptions) New() *AppsCreateExecutionOptions {
	options.enabledSetters = map[string]bool{
		"Async": false,
		"Body": false,
		"Headers": false,
		"Method": false,
		"Path": false,
		"ScheduledAt": false,
	}
	return &options
}
type AppsCreateExecutionOption func(*AppsCreateExecutionOptions)
func (srv *Apps) WithAppsCreateExecutionAsync(v bool) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.Async = v
		o.enabledSetters["Async"] = true
	}
}
func (srv *Apps) WithAppsCreateExecutionBody(v string) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Apps) WithAppsCreateExecutionHeaders(v interface{}) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Apps) WithAppsCreateExecutionMethod(v string) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.Method = v
		o.enabledSetters["Method"] = true
	}
}
func (srv *Apps) WithAppsCreateExecutionPath(v string) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *Apps) WithAppsCreateExecutionScheduledAt(v string) AppsCreateExecutionOption {
	return func(o *AppsCreateExecutionOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
			
// AppsCreateExecution trigger an App execution. Use the optional `body`,
// `path`, `method` and `headers` parameters to invoke the App as if from an
// HTTP request.
func (srv *Apps) AppsCreateExecution(FunctionId string, optionalSetters ...AppsCreateExecutionOption)(*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/executions")
	options := AppsCreateExecutionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Async"] {
		params["async"] = options.Async
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["Method"] {
		params["method"] = options.Method
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
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

		parsed := models.Execution{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Execution
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// AppsDeleteExecution delete an App execution by its unique ID.
func (srv *Apps) AppsDeleteExecution(FunctionId string, ExecutionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/v1/apps/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["executionId"] = ExecutionId
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
			
// AppsGetExecution get an App execution by its unique ID.
func (srv *Apps) AppsGetExecution(FunctionId string, ExecutionId string)(*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/v1/apps/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["executionId"] = ExecutionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Execution{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Execution
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// AppsGetMarketplaceStatus read-through view of the App's App Registry row
// — visibility + Marketplace publish flag. Used by Cockpit to render the
// Publish/Unpublish button correctly on cold load.
func (srv *Apps) AppsGetMarketplaceStatus(FunctionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/marketplace-status")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
	
// AppsUnpublish remove this App from the Marketplace listing. Existing tenant
// installations are unaffected. Idempotent.
func (srv *Apps) AppsUnpublish(FunctionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/publish")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
	
// AppsPublish publish this App to the Marketplace. The App must have at
// least one `ready` deployment with a registered manifest,
// and its visibility (derived from `billing.json`) must be
// `public` or `included`. Idempotent.
func (srv *Apps) AppsPublish(FunctionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/publish")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
type AppsGetUsageOptions struct {
	Range string
	enabledSetters map[string]bool
}
func (options AppsGetUsageOptions) New() *AppsGetUsageOptions {
	options.enabledSetters = map[string]bool{
		"Range": false,
	}
	return &options
}
type AppsGetUsageOption func(*AppsGetUsageOptions)
func (srv *Apps) WithAppsGetUsageRange(v string) AppsGetUsageOption {
	return func(o *AppsGetUsageOptions) {
		o.Range = v
		o.enabledSetters["Range"] = true
	}
}
			
// AppsGetUsage get usage stats for a single App over the requested time
// range.
func (srv *Apps) AppsGetUsage(FunctionId string, optionalSetters ...AppsGetUsageOption)(*models.UsageFunction, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/usage")
	options := AppsGetUsageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Range"] {
		params["range"] = options.Range
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UsageFunction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageFunction
	parsed, ok := resp.Result.(models.UsageFunction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// AppsListVariables list all environment variables defined for the App.
func (srv *Apps) AppsListVariables(FunctionId string)(*models.VariableList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/variables")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.VariableList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VariableList
	parsed, ok := resp.Result.(models.VariableList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsCreateVariableOptions struct {
	Secret bool
	enabledSetters map[string]bool
}
func (options AppsCreateVariableOptions) New() *AppsCreateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Secret": false,
	}
	return &options
}
type AppsCreateVariableOption func(*AppsCreateVariableOptions)
func (srv *Apps) WithAppsCreateVariableSecret(v bool) AppsCreateVariableOption {
	return func(o *AppsCreateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
							
// AppsCreateVariable create a new App environment variable. These are passed
// into the App at runtime as `process.env.*`.
func (srv *Apps) AppsCreateVariable(FunctionId string, Key string, Value string, optionalSetters ...AppsCreateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/v1/apps/{functionId}/variables")
	options := AppsCreateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["key"] = Key
	params["value"] = Value
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
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

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// AppsDeleteVariable delete an App environment variable.
func (srv *Apps) AppsDeleteVariable(FunctionId string, VariableId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/v1/apps/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
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
			
// AppsGetVariable get an App variable by its unique ID.
func (srv *Apps) AppsGetVariable(FunctionId string, VariableId string)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/v1/apps/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type AppsUpdateVariableOptions struct {
	Secret bool
	Value string
	enabledSetters map[string]bool
}
func (options AppsUpdateVariableOptions) New() *AppsUpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Secret": false,
		"Value": false,
	}
	return &options
}
type AppsUpdateVariableOption func(*AppsUpdateVariableOptions)
func (srv *Apps) WithAppsUpdateVariableSecret(v bool) AppsUpdateVariableOption {
	return func(o *AppsUpdateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
func (srv *Apps) WithAppsUpdateVariableValue(v string) AppsUpdateVariableOption {
	return func(o *AppsUpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
							
// AppsUpdateVariable update an App environment variable.
func (srv *Apps) AppsUpdateVariable(FunctionId string, VariableId string, Key string, optionalSetters ...AppsUpdateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/v1/apps/{functionId}/variables/{variableId}")
	options := AppsUpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	params["key"] = Key
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
	}
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
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

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
