package models

import (
    "encoding/json"
    "errors"
)

// FrameworkAdapter Framework Adapter
type FrameworkAdapter struct {
    // Default command to build site into output directory.
    BuildCommand string `json:"buildCommand"`
    // Name of the fallback file to serve instead of a 404 page. If null, the site
    // runtime's built-in 404 page is served.
    FallbackFile string `json:"fallbackFile"`
    // Default command to download dependencies.
    InstallCommand string `json:"installCommand"`
    // Adapter key.
    Key string `json:"key"`
    // Default output directory of build.
    OutputDirectory string `json:"outputDirectory"`

    // Used by Decode() method
    data []byte
}

func (model FrameworkAdapter) New(data []byte) *FrameworkAdapter {
    model.data = data
    return &model
}

func (model *FrameworkAdapter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}