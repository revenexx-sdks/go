package models

import (
    "encoding/json"
    "errors"
)

// Model
type FamilyVariantsCreateRequest struct {
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

func (model FamilyVariantsCreateRequest) New(data []byte) *FamilyVariantsCreateRequest {
    model.data = data
    return &model
}

func (model *FamilyVariantsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}