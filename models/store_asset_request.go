package models

import (
    "encoding/json"
    "errors"
)

// Model
type StoreAssetRequest struct {
    // 
    AltText string `json:"alt_text"`
    // 
    Description string `json:"description"`
    // 
    DisplayName string `json:"display_name"`
    // 
    File string `json:"file"`
    // 
    FolderId string `json:"folder_id"`
    // 
    KeepArchive bool `json:"keep_archive"`
    // 
    Tags []string `json:"tags"`
    // Archives only: unpack the members after upload (see AssetController).
    Unpack bool `json:"unpack"`
    // 
    Visibility string `json:"visibility"`

    // Used by Decode() method
    data []byte
}

func (model StoreAssetRequest) New(data []byte) *StoreAssetRequest {
    model.data = data
    return &model
}

func (model *StoreAssetRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}