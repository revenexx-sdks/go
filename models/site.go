package models

import (
    "encoding/json"
    "errors"
)

// Site Model
type Site struct {
    // Site creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Site ID.
    Id string `json:"$id"`
    // Site update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Site framework adapter.
    Adapter string `json:"adapter"`
    // The build command used to build the site.
    BuildCommand string `json:"buildCommand"`
    // Site build runtime.
    BuildRuntime string `json:"buildRuntime"`
    // Active deployment creation date in ISO 8601 format.
    DeploymentCreatedAt string `json:"deploymentCreatedAt"`
    // Site's active deployment ID.
    DeploymentId string `json:"deploymentId"`
    // Screenshot of active deployment with dark theme preference file ID.
    DeploymentScreenshotDark string `json:"deploymentScreenshotDark"`
    // Screenshot of active deployment with light theme preference file ID.
    DeploymentScreenshotLight string `json:"deploymentScreenshotLight"`
    // Site enabled.
    Enabled bool `json:"enabled"`
    // Name of fallback file to use instead of 404 page. If null, Appwrite 404
    // page will be displayed.
    FallbackFile string `json:"fallbackFile"`
    // Site framework.
    Framework string `json:"framework"`
    // The install command used to install the site dependencies.
    InstallCommand string `json:"installCommand"`
    // Site VCS (Version Control System) installation id.
    InstallationId string `json:"installationId"`
    // Latest deployment creation date in ISO 8601 format.
    LatestDeploymentCreatedAt string `json:"latestDeploymentCreatedAt"`
    // Site's latest deployment ID.
    LatestDeploymentId string `json:"latestDeploymentId"`
    // Status of latest deployment. Possible values are "waiting", "processing",
    // "building", "ready", and "failed".
    LatestDeploymentStatus string `json:"latestDeploymentStatus"`
    // Is the site deployed with the latest configuration? This is set to false if
    // you've changed an environment variables, entrypoint, commands, or other
    // settings that needs redeploy to be applied. When the value is false,
    // redeploy the site to update it with the latest configuration.
    Live bool `json:"live"`
    // When disabled, request logs will exclude logs and errors, and site
    // responses will be slightly faster.
    Logging bool `json:"logging"`
    // Site name.
    Name string `json:"name"`
    // The directory where the site build output is located.
    OutputDirectory string `json:"outputDirectory"`
    // VCS (Version Control System) branch name
    ProviderBranch string `json:"providerBranch"`
    // VCS (Version Control System) Repository ID
    ProviderRepositoryId string `json:"providerRepositoryId"`
    // Path to site in VCS (Version Control System) repository
    ProviderRootDirectory string `json:"providerRootDirectory"`
    // Is VCS (Version Control System) connection is in silent mode? When in
    // silence mode, no comments will be posted on the repository pull or merge
    // requests
    ProviderSilentMode bool `json:"providerSilentMode"`
    // Machine specification for builds and executions.
    Specification string `json:"specification"`
    // Site request timeout in seconds.
    Timeout int `json:"timeout"`
    // Site variables.
    Vars []Variable `json:"vars"`

    // Used by Decode() method
    data []byte
}

func (model Site) New(data []byte) *Site {
    model.data = data
    return &model
}

func (model *Site) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}