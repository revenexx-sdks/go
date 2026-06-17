package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AssociationTypesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    IsQuantified bool `json:"is_quantified"`
    // 
    IsTwoWay bool `json:"is_two_way"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model AssociationTypesUpdateRequest) New(data []byte) *AssociationTypesUpdateRequest {
    model.data = data
    return &model
}

func (model *AssociationTypesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}