package models

import (
    "encoding/json"
    "errors"
)

// Target Model
type Target struct {
    // Target creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Target ID.
    Id string `json:"$id"`
    // Target update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is the target expired.
    Expired bool `json:"expired"`
    // The target identifier.
    Identifier string `json:"identifier"`
    // Target Name.
    Name string `json:"name"`
    // Provider ID.
    ProviderId string `json:"providerId"`
    // The target provider type. Can be one of the following: `email`, `sms` or
    // `push`.
    ProviderType string `json:"providerType"`
    // User ID.
    UserId string `json:"userId"`

    // Used by Decode() method
    data []byte
}

func (model Target) New(data []byte) *Target {
    model.data = data
    return &model
}

func (model *Target) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}