package models

import (
    "encoding/json"
    "errors"
)

// Model
type FamilyVariants struct {
    // 
    Axes interface{} `json:"axes"`
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    FamilyId string `json:"family_id"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model FamilyVariants) New(data []byte) *FamilyVariants {
    model.data = data
    return &model
}

func (model *FamilyVariants) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}