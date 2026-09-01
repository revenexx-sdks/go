package models

import (
    "encoding/json"
    "errors"
)

// AssetsFilter The exact-column filters this call was understood to carry,
// verbatim as they arrived. A query parameter that is not a column of
// `assets` — `?status=`, a typo, a filter another entity has — is DROPPED
// and does not appear here, and the list comes back unfiltered. This object
// is the only way to tell that apart from "nothing matched".
type AssetsFilter struct {
    // The literal `?asset_family_id=` value this call was understood to carry.
    AssetFamilyId string `json:"asset_family_id"`
    // The literal `?attribute_values=` value this call was understood to carry.
    AttributeValues string `json:"attribute_values"`
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?delivery_path=` value this call was understood to carry.
    DeliveryPath string `json:"delivery_path"`
    // The literal `?external_url=` value this call was understood to carry.
    ExternalUrl string `json:"external_url"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?source=` value this call was understood to carry.
    Source string `json:"source"`
    // The literal `?storage_asset_id=` value this call was understood to carry.
    StorageAssetId string `json:"storage_asset_id"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model AssetsFilter) New(data []byte) *AssetsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *AssetsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}