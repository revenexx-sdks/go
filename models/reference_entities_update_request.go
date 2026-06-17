package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type ReferenceEntitiesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Image string `json:"image"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntitiesUpdateRequest) New(data []byte) *ReferenceEntitiesUpdateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntitiesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}