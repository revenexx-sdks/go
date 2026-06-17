package models

import (
    "encoding/json"
    "errors"
)

// Model
type FamilyAttributes struct {
    // 
    AttributeId string `json:"attribute_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    FamilyId string `json:"family_id"`
    // 
    Id string `json:"id"`
    // 
    IsRequired bool `json:"is_required"`
    // 
    Position int `json:"position"`
    // 
    RequiredChannels interface{} `json:"required_channels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyAttributes) New(data []byte) *FamilyAttributes {
    model.data = data
    return &model
}

func (model *FamilyAttributes) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}