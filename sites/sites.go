package sites

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Sites service
type Sites struct {
	client client.Client
}

func New(clt client.Client) *Sites {
	return &Sites{
		client: clt,
	}
}

type SitesListOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options SitesListOptions) New() *SitesListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type SitesListOption func(*SitesListOptions)
func (srv *Sites) WithSitesListQueries(v []string) SitesListOption {
	return func(o *SitesListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithSitesListSearch(v string) SitesListOption {
	return func(o *SitesListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Sites) WithSitesListTotal(v bool) SitesListOption {
	return func(o *SitesListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// SitesList get a list of all the project's sites. You can use the query
// params to filter your results.
func (srv *Sites) SitesList(optionalSetters ...SitesListOption)(*models.SiteList, error) {
	path := "/v1/sites"
	options := SitesListOptions{}.New()
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

		parsed := models.SiteList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SiteList
	parsed, ok := resp.Result.(models.SiteList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SitesCreateOptions struct {
	Adapter string
	BuildCommand string
	Enabled bool
	FallbackFile string
	InstallCommand string
	InstallationId string
	Logging bool
	OutputDirectory string
	ProviderBranch string
	ProviderRepositoryId string
	ProviderRootDirectory string
	ProviderSilentMode bool
	Specification string
	Timeout int
	enabledSetters map[string]bool
}
func (options SitesCreateOptions) New() *SitesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Adapter": false,
		"BuildCommand": false,
		"Enabled": false,
		"FallbackFile": false,
		"InstallCommand": false,
		"InstallationId": false,
		"Logging": false,
		"OutputDirectory": false,
		"ProviderBranch": false,
		"ProviderRepositoryId": false,
		"ProviderRootDirectory": false,
		"ProviderSilentMode": false,
		"Specification": false,
		"Timeout": false,
	}
	return &options
}
type SitesCreateOption func(*SitesCreateOptions)
func (srv *Sites) WithSitesCreateAdapter(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.Adapter = v
		o.enabledSetters["Adapter"] = true
	}
}
func (srv *Sites) WithSitesCreateBuildCommand(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithSitesCreateEnabled(v bool) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Sites) WithSitesCreateFallbackFile(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.FallbackFile = v
		o.enabledSetters["FallbackFile"] = true
	}
}
func (srv *Sites) WithSitesCreateInstallCommand(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithSitesCreateInstallationId(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Sites) WithSitesCreateLogging(v bool) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Sites) WithSitesCreateOutputDirectory(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
func (srv *Sites) WithSitesCreateProviderBranch(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Sites) WithSitesCreateProviderRepositoryId(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Sites) WithSitesCreateProviderRootDirectory(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Sites) WithSitesCreateProviderSilentMode(v bool) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Sites) WithSitesCreateSpecification(v string) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Sites) WithSitesCreateTimeout(v int) SitesCreateOption {
	return func(o *SitesCreateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
									
// SitesCreate create a new site.
func (srv *Sites) SitesCreate(BuildRuntime string, Framework string, Name string, SiteId string, optionalSetters ...SitesCreateOption)(*models.Site, error) {
	path := "/v1/sites"
	options := SitesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["buildRuntime"] = BuildRuntime
	params["framework"] = Framework
	params["name"] = Name
	params["siteId"] = SiteId
	if options.enabledSetters["Adapter"] {
		params["adapter"] = options.Adapter
	}
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FallbackFile"] {
		params["fallbackFile"] = options.FallbackFile
	}
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
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

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// SitesListFrameworks get a list of all frameworks that are currently
// available on the server instance.
func (srv *Sites) SitesListFrameworks()(*models.FrameworkList, error) {
	path := "/v1/sites/frameworks"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.FrameworkList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FrameworkList
	parsed, ok := resp.Result.(models.FrameworkList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// SitesListSpecifications list allowed site specifications for this instance.
func (srv *Sites) SitesListSpecifications()(*models.SpecificationList, error) {
	path := "/v1/sites/specifications"
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
	
// SitesDelete delete a site by its unique ID.
func (srv *Sites) SitesDelete(SiteId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
	
// SitesGet get a site by its unique ID.
func (srv *Sites) SitesGet(SiteId string)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SitesUpdateOptions struct {
	Adapter string
	BuildCommand string
	BuildRuntime string
	Enabled bool
	FallbackFile string
	InstallCommand string
	InstallationId string
	Logging bool
	OutputDirectory string
	ProviderBranch string
	ProviderRepositoryId string
	ProviderRootDirectory string
	ProviderSilentMode bool
	Specification string
	Timeout int
	enabledSetters map[string]bool
}
func (options SitesUpdateOptions) New() *SitesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Adapter": false,
		"BuildCommand": false,
		"BuildRuntime": false,
		"Enabled": false,
		"FallbackFile": false,
		"InstallCommand": false,
		"InstallationId": false,
		"Logging": false,
		"OutputDirectory": false,
		"ProviderBranch": false,
		"ProviderRepositoryId": false,
		"ProviderRootDirectory": false,
		"ProviderSilentMode": false,
		"Specification": false,
		"Timeout": false,
	}
	return &options
}
type SitesUpdateOption func(*SitesUpdateOptions)
func (srv *Sites) WithSitesUpdateAdapter(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.Adapter = v
		o.enabledSetters["Adapter"] = true
	}
}
func (srv *Sites) WithSitesUpdateBuildCommand(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithSitesUpdateBuildRuntime(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.BuildRuntime = v
		o.enabledSetters["BuildRuntime"] = true
	}
}
func (srv *Sites) WithSitesUpdateEnabled(v bool) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Sites) WithSitesUpdateFallbackFile(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.FallbackFile = v
		o.enabledSetters["FallbackFile"] = true
	}
}
func (srv *Sites) WithSitesUpdateInstallCommand(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithSitesUpdateInstallationId(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Sites) WithSitesUpdateLogging(v bool) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Sites) WithSitesUpdateOutputDirectory(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
func (srv *Sites) WithSitesUpdateProviderBranch(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Sites) WithSitesUpdateProviderRepositoryId(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Sites) WithSitesUpdateProviderRootDirectory(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Sites) WithSitesUpdateProviderSilentMode(v bool) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Sites) WithSitesUpdateSpecification(v string) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Sites) WithSitesUpdateTimeout(v int) SitesUpdateOption {
	return func(o *SitesUpdateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
							
// SitesUpdate update site by its unique ID.
func (srv *Sites) SitesUpdate(SiteId string, Framework string, Name string, optionalSetters ...SitesUpdateOption)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}")
	options := SitesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["framework"] = Framework
	params["name"] = Name
	if options.enabledSetters["Adapter"] {
		params["adapter"] = options.Adapter
	}
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["BuildRuntime"] {
		params["buildRuntime"] = options.BuildRuntime
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FallbackFile"] {
		params["fallbackFile"] = options.FallbackFile
	}
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
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

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// SitesUpdateSiteDeployment update the site active deployment. Use this
// endpoint to switch the code deployment that should be used when visitor
// opens your site.
func (srv *Sites) SitesUpdateSiteDeployment(SiteId string, DeploymentId string)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployment")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SitesListDeploymentsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options SitesListDeploymentsOptions) New() *SitesListDeploymentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type SitesListDeploymentsOption func(*SitesListDeploymentsOptions)
func (srv *Sites) WithSitesListDeploymentsQueries(v []string) SitesListDeploymentsOption {
	return func(o *SitesListDeploymentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithSitesListDeploymentsSearch(v string) SitesListDeploymentsOption {
	return func(o *SitesListDeploymentsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Sites) WithSitesListDeploymentsTotal(v bool) SitesListDeploymentsOption {
	return func(o *SitesListDeploymentsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// SitesListDeployments get a list of all the site's code deployments. You can
// use the query params to filter your results.
func (srv *Sites) SitesListDeployments(SiteId string, optionalSetters ...SitesListDeploymentsOption)(*models.DeploymentList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployments")
	options := SitesListDeploymentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesCreateDeploymentOptions struct {
	BuildCommand string
	InstallCommand string
	OutputDirectory string
	enabledSetters map[string]bool
}
func (options SitesCreateDeploymentOptions) New() *SitesCreateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"BuildCommand": false,
		"InstallCommand": false,
		"OutputDirectory": false,
	}
	return &options
}
type SitesCreateDeploymentOption func(*SitesCreateDeploymentOptions)
func (srv *Sites) WithSitesCreateDeploymentBuildCommand(v string) SitesCreateDeploymentOption {
	return func(o *SitesCreateDeploymentOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithSitesCreateDeploymentInstallCommand(v string) SitesCreateDeploymentOption {
	return func(o *SitesCreateDeploymentOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithSitesCreateDeploymentOutputDirectory(v string) SitesCreateDeploymentOption {
	return func(o *SitesCreateDeploymentOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
							
// SitesCreateDeployment create a new site code deployment. Use this endpoint
// to upload a new version of your site code. To activate your newly uploaded
// code, you'll need to update the site's deployment to use your new
// deployment ID.
func (srv *Sites) SitesCreateDeployment(SiteId string, Activate bool, Code string, optionalSetters ...SitesCreateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployments")
	options := SitesCreateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["activate"] = Activate
	params["code"] = Code
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
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
			
// SitesCreateDuplicateDeployment create a new build for an existing site
// deployment. This endpoint allows you to rebuild a deployment with the
// updated site configuration, including its commands and output directory if
// they have been modified. The build process will be queued and executed
// asynchronously. The original deployment's code will be preserved and used
// for the new build.
func (srv *Sites) SitesCreateDuplicateDeployment(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployments/duplicate")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["deploymentId"] = DeploymentId
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
type SitesCreateTemplateDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options SitesCreateTemplateDeploymentOptions) New() *SitesCreateTemplateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type SitesCreateTemplateDeploymentOption func(*SitesCreateTemplateDeploymentOptions)
func (srv *Sites) WithSitesCreateTemplateDeploymentActivate(v bool) SitesCreateTemplateDeploymentOption {
	return func(o *SitesCreateTemplateDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
													
// SitesCreateTemplateDeployment create a deployment based on a template.
// 
// Use this endpoint with combination of
// [listTemplates](https://appwrite.io/docs/products/sites/templates) to find
// the template details.
func (srv *Sites) SitesCreateTemplateDeployment(SiteId string, Owner string, Reference string, Repository string, RootDirectory string, Type string, optionalSetters ...SitesCreateTemplateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployments/template")
	options := SitesCreateTemplateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesCreateVcsDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options SitesCreateVcsDeploymentOptions) New() *SitesCreateVcsDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type SitesCreateVcsDeploymentOption func(*SitesCreateVcsDeploymentOptions)
func (srv *Sites) WithSitesCreateVcsDeploymentActivate(v bool) SitesCreateVcsDeploymentOption {
	return func(o *SitesCreateVcsDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
							
// SitesCreateVcsDeployment create a deployment when a site is connected to
// VCS.
// 
// This endpoint lets you create deployment from a branch, commit, or a tag.
func (srv *Sites) SitesCreateVcsDeployment(SiteId string, Reference string, Type string, optionalSetters ...SitesCreateVcsDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/deployments/vcs")
	options := SitesCreateVcsDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesDeleteDeployment delete a site deployment by its unique ID.
func (srv *Sites) SitesDeleteDeployment(SiteId string, DeploymentId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/sites/{siteId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesGetDeployment get a site deployment by its unique ID.
func (srv *Sites) SitesGetDeployment(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/sites/{siteId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesGetDeploymentDownloadOptions struct {
	Type string
	enabledSetters map[string]bool
}
func (options SitesGetDeploymentDownloadOptions) New() *SitesGetDeploymentDownloadOptions {
	options.enabledSetters = map[string]bool{
		"Type": false,
	}
	return &options
}
type SitesGetDeploymentDownloadOption func(*SitesGetDeploymentDownloadOptions)
func (srv *Sites) WithSitesGetDeploymentDownloadType(v string) SitesGetDeploymentDownloadOption {
	return func(o *SitesGetDeploymentDownloadOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// SitesGetDeploymentDownload get a site deployment content by its unique ID.
// The endpoint response return with a 'Content-Disposition: attachment'
// header that tells the browser to start downloading the file to user
// downloads directory.
func (srv *Sites) SitesGetDeploymentDownload(SiteId string, DeploymentId string, optionalSetters ...SitesGetDeploymentDownloadOption)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/sites/{siteId}/deployments/{deploymentId}/download")
	options := SitesGetDeploymentDownloadOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesUpdateDeploymentStatus cancel an ongoing site deployment build. If the
// build is already in progress, it will be stopped and marked as canceled. If
// the build hasn't started yet, it will be marked as canceled without
// executing. You cannot cancel builds that have already completed (status
// 'ready') or failed. The response includes the final build status and
// details.
func (srv *Sites) SitesUpdateDeploymentStatus(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/v1/sites/{siteId}/deployments/{deploymentId}/status")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesListLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options SitesListLogsOptions) New() *SitesListLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type SitesListLogsOption func(*SitesListLogsOptions)
func (srv *Sites) WithSitesListLogsQueries(v []string) SitesListLogsOption {
	return func(o *SitesListLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithSitesListLogsTotal(v bool) SitesListLogsOption {
	return func(o *SitesListLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// SitesListLogs get a list of all site logs. You can use the query params to
// filter your results.
func (srv *Sites) SitesListLogs(SiteId string, optionalSetters ...SitesListLogsOption)(*models.ExecutionList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/logs")
	options := SitesListLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesDeleteLog delete a site log by its unique ID.
func (srv *Sites) SitesDeleteLog(SiteId string, LogId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{logId}", LogId)
	path := r.Replace("/v1/sites/{siteId}/logs/{logId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["logId"] = LogId
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
			
// SitesGetLog get a site request log by its unique ID.
func (srv *Sites) SitesGetLog(SiteId string, LogId string)(*models.Execution, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{logId}", LogId)
	path := r.Replace("/v1/sites/{siteId}/logs/{logId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["logId"] = LogId
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
	
// SitesListVariables get a list of all variables of a specific site.
func (srv *Sites) SitesListVariables(SiteId string)(*models.VariableList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/variables")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesCreateVariableOptions struct {
	Secret bool
	enabledSetters map[string]bool
}
func (options SitesCreateVariableOptions) New() *SitesCreateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Secret": false,
	}
	return &options
}
type SitesCreateVariableOption func(*SitesCreateVariableOptions)
func (srv *Sites) WithSitesCreateVariableSecret(v bool) SitesCreateVariableOption {
	return func(o *SitesCreateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
							
// SitesCreateVariable create a new site variable. These variables can be
// accessed during build and runtime (server-side rendering) as environment
// variables.
func (srv *Sites) SitesCreateVariable(SiteId string, Key string, Value string, optionalSetters ...SitesCreateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/v1/sites/{siteId}/variables")
	options := SitesCreateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesDeleteVariable delete a variable by its unique ID.
func (srv *Sites) SitesDeleteVariable(SiteId string, VariableId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/v1/sites/{siteId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// SitesGetVariable get a variable by its unique ID.
func (srv *Sites) SitesGetVariable(SiteId string, VariableId string)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/v1/sites/{siteId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type SitesUpdateVariableOptions struct {
	Secret bool
	Value string
	enabledSetters map[string]bool
}
func (options SitesUpdateVariableOptions) New() *SitesUpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Secret": false,
		"Value": false,
	}
	return &options
}
type SitesUpdateVariableOption func(*SitesUpdateVariableOptions)
func (srv *Sites) WithSitesUpdateVariableSecret(v bool) SitesUpdateVariableOption {
	return func(o *SitesUpdateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
func (srv *Sites) WithSitesUpdateVariableValue(v string) SitesUpdateVariableOption {
	return func(o *SitesUpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
							
// SitesUpdateVariable update variable by its unique ID.
func (srv *Sites) SitesUpdateVariable(SiteId string, VariableId string, Key string, optionalSetters ...SitesUpdateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/v1/sites/{siteId}/variables/{variableId}")
	options := SitesUpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
