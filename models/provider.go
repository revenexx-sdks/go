package models

import (
    "encoding/json"
    "errors"
)

// Provider Model
type Provider struct {
    // Provider creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Provider ID.
    Id string `json:"$id"`
    // Provider update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Provider credentials.
    Credentials interface{} `json:"credentials"`
    // Is provider enabled?
    Enabled bool `json:"enabled"`
    // The name for the provider instance.
    Name string `json:"name"`
    // Provider options.
    Options interface{} `json:"options"`
    // The name of the provider service.
    Provider string `json:"provider"`
    // Type of provider.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model Provider) New(data []byte) *Provider {
    model.data = data
    return &model
}

func (model *Provider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}