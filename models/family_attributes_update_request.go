package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type FamilyAttributesUpdateRequest struct {
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

func (model FamilyAttributesUpdateRequest) New(data []byte) *FamilyAttributesUpdateRequest {
    model.data = data
    return &model
}

func (model *FamilyAttributesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}