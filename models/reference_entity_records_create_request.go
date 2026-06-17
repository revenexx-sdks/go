package models

import (
    "encoding/json"
    "errors"
)

// Model
type ReferenceEntityRecordsCreateRequest struct {
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    ReferenceEntityId string `json:"reference_entity_id"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntityRecordsCreateRequest) New(data []byte) *ReferenceEntityRecordsCreateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntityRecordsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}