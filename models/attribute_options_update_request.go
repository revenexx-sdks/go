package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AttributeOptionsUpdateRequest struct {
    // 
    AttributeId string `json:"attribute_id"`
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    Position int `json:"position"`
    // 
    Swatch interface{} `json:"swatch"`

    // Used by Decode() method
    data []byte
}

func (model AttributeOptionsUpdateRequest) New(data []byte) *AttributeOptionsUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributeOptionsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}