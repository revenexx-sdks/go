package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type FamiliesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    ImageAttribute string `json:"image_attribute"`
    // 
    LabelAttribute string `json:"label_attribute"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model FamiliesUpdateRequest) New(data []byte) *FamiliesUpdateRequest {
    model.data = data
    return &model
}

func (model *FamiliesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}