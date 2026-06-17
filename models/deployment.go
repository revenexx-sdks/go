package models

import (
    "encoding/json"
    "errors"
)

// Deployment Model
type Deployment struct {
    // Deployment creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Deployment ID.
    Id string `json:"$id"`
    // Deployment update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Whether the deployment should be automatically activated.
    Activate bool `json:"activate"`
    // Raw billing.json bytes captured from the source archive at deploy time.
    // Empty when no billing.json was shipped (private app).
    BillingJson string `json:"billingJson"`
    // The current build time in seconds.
    BuildDuration int `json:"buildDuration"`
    // The current build ID.
    BuildId string `json:"buildId"`
    // The build logs.
    BuildLogs string `json:"buildLogs"`
    // The build output size in bytes.
    BuildSize int `json:"buildSize"`
    // The entrypoint file to use to execute the deployment code.
    Entrypoint string `json:"entrypoint"`
    // Raw manifest.json bytes captured from the source archive at deploy time.
    // Empty for legacy Function/Site deployments without a manifest.
    ManifestJson string `json:"manifestJson"`
    // The branch of the vcs repository
    ProviderBranch string `json:"providerBranch"`
    // The branch of the vcs repository
    ProviderBranchUrl string `json:"providerBranchUrl"`
    // The name of vcs commit author
    ProviderCommitAuthor string `json:"providerCommitAuthor"`
    // The url of vcs commit author
    ProviderCommitAuthorUrl string `json:"providerCommitAuthorUrl"`
    // The commit hash of the vcs commit
    ProviderCommitHash string `json:"providerCommitHash"`
    // The commit message
    ProviderCommitMessage string `json:"providerCommitMessage"`
    // The url of the vcs commit
    ProviderCommitUrl string `json:"providerCommitUrl"`
    // The name of the vcs provider repository
    ProviderRepositoryName string `json:"providerRepositoryName"`
    // The name of the vcs provider repository owner
    ProviderRepositoryOwner string `json:"providerRepositoryOwner"`
    // The url of the vcs provider repository
    ProviderRepositoryUrl string `json:"providerRepositoryUrl"`
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource type.
    ResourceType string `json:"resourceType"`
    // Screenshot with dark theme preference file ID.
    ScreenshotDark string `json:"screenshotDark"`
    // Screenshot with light theme preference file ID.
    ScreenshotLight string `json:"screenshotLight"`
    // The code size in bytes.
    SourceSize int `json:"sourceSize"`
    // The deployment status. Possible values are "waiting", "processing",
    // "building", "ready", "canceled" and "failed".
    Status string `json:"status"`
    // The total size in bytes (source and build output).
    TotalSize int `json:"totalSize"`
    // Type of deployment.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model Deployment) New(data []byte) *Deployment {
    model.data = data
    return &model
}

func (model *Deployment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}