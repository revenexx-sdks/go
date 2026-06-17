package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type IoProfileUpdateRequest struct {
    // Default 'insert'.
    ApplyMode string `json:"apply_mode"`
    // 
    Direction string `json:"direction"`
    // Default 'carts'.
    Entity string `json:"entity"`
    // Default 'json'.
    Format string `json:"format"`
    // 
    IsTemplate bool `json:"is_template"`
    // Column mapping (Baseline-IO-compatible).
    Mapping interface{} `json:"mapping"`
    // 
    Name string `json:"name"`
    // 
    Options interface{} `json:"options"`

    // Used by Decode() method
    data []byte
}

func (model IoProfileUpdateRequest) New(data []byte) *IoProfileUpdateRequest {
    model.data = data
    return &model
}

func (model *IoProfileUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}