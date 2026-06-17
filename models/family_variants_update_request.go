package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type FamilyVariantsUpdateRequest struct {
    // 
    Axes interface{} `json:"axes"`
    // 
    Code string `json:"code"`
    // 
    FamilyId string `json:"family_id"`
    // 
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyVariantsUpdateRequest) New(data []byte) *FamilyVariantsUpdateRequest {
    model.data = data
    return &model
}

func (model *FamilyVariantsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}