package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type AssetsUpdateRequest struct {
    // 
    AssetFamilyId string `json:"asset_family_id"`
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Code string `json:"code"`
    // 
    MediaUuid string `json:"media_uuid"`

    // Used by Decode() method
    data []byte
}

func (model AssetsUpdateRequest) New(data []byte) *AssetsUpdateRequest {
    model.data = data
    return &model
}

func (model *AssetsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}