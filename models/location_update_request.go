package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type LocationUpdateRequest struct {
    // 
    Address interface{} `json:"address"`
    // Unique location code (per tenant).
    Code string `json:"code"`
    // Disabled locations are skipped by availability and reserve (default true).
    Enabled bool `json:"enabled"`
    // Localised display names ({de, en, …}).
    Labels interface{} `json:"labels"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // Sourcing order — lower wins (default 0).
    Priority int `json:"priority"`
    // Default 'warehouse'.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model LocationUpdateRequest) New(data []byte) *LocationUpdateRequest {
    model.data = data
    return &model
}

func (model *LocationUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}