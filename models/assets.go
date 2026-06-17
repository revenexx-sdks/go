package models

import (
    "encoding/json"
    "errors"
)

// Model
type Assets struct {
    // 
    AssetFamilyId string `json:"asset_family_id"`
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    MediaUuid string `json:"media_uuid"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Assets) New(data []byte) *Assets {
    model.data = data
    return &model
}

func (model *Assets) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}