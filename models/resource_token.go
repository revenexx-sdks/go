package models

import (
    "encoding/json"
    "errors"
)

// ResourceToken Model
type ResourceToken struct {
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Token ID.
    Id string `json:"$id"`
    // Most recent access date in ISO 8601 format. This attribute is only updated
    // again after 24 hours.
    AccessedAt string `json:"accessedAt"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource type.
    ResourceType string `json:"resourceType"`
    // JWT encoded string.
    Secret string `json:"secret"`

    // Used by Decode() method
    data []byte
}

func (model ResourceToken) New(data []byte) *ResourceToken {
    model.data = data
    return &model
}

func (model *ResourceToken) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}