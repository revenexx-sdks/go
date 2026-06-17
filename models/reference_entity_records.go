package models

import (
    "encoding/json"
    "errors"
)

// Model
type ReferenceEntityRecords struct {
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    ReferenceEntityId string `json:"reference_entity_id"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntityRecords) New(data []byte) *ReferenceEntityRecords {
    model.data = data
    return &model
}

func (model *ReferenceEntityRecords) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}