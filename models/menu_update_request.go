package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type MenuUpdateRequest struct {
    // 
    Items []interface{} `json:"items"`
    // 
    Label string `json:"label"`

    // Used by Decode() method
    data []byte
}

func (model MenuUpdateRequest) New(data []byte) *MenuUpdateRequest {
    model.data = data
    return &model
}

func (model *MenuUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}