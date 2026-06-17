package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AttributesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Config interface{} `json:"config"`
    // 
    EntityRef string `json:"entity_ref"`
    // 
    EntityType string `json:"entity_type"`
    // 
    GroupId string `json:"group_id"`
    // 
    IsFilterable bool `json:"is_filterable"`
    // 
    IsUnique bool `json:"is_unique"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Localizable bool `json:"localizable"`
    // 
    Position int `json:"position"`
    // 
    Scopable bool `json:"scopable"`
    // 
    Type string `json:"type"`
    // 
    UsableInGrid bool `json:"usable_in_grid"`
    // 
    Validation interface{} `json:"validation"`

    // Used by Decode() method
    data []byte
}

func (model AttributesUpdateRequest) New(data []byte) *AttributesUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}