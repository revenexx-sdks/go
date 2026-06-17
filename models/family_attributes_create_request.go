package models

import (
    "encoding/json"
    "errors"
)

// Model
type FamilyAttributesCreateRequest struct {
    // 
    AttributeId string `json:"attribute_id"`
    // 
    FamilyId string `json:"family_id"`
    // 
    IsRequired bool `json:"is_required"`
    // 
    Position int `json:"position"`
    // 
    RequiredChannels interface{} `json:"required_channels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyAttributesCreateRequest) New(data []byte) *FamilyAttributesCreateRequest {
    model.data = data
    return &model
}

func (model *FamilyAttributesCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}