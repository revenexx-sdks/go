package models

import (
    "encoding/json"
    "errors"
)

// AssetFamiliesCreateRequest model.
type AssetFamiliesCreateRequest struct {
    // The asset family's stable identifier — a class of media with one shared
    // shape. Unique per tenant.
    Code string `json:"code"`
    // What the asset family is called, per language tag.
    Labels interface{} `json:"labels"`
    // How a file of this family is named, so an import can bind a file to a
    // product without a mapping table. `source` is the product value the file
    // name is built from, `pattern` how it is assembled, `allowed_extensions`
    // what may be uploaded.
    NamingConvention interface{} `json:"naming_convention"`

    // Used by Decode() method
    data []byte
}

func (model AssetFamiliesCreateRequest) New(data []byte) *AssetFamiliesCreateRequest {
    model.data = data
    return &model
}

func (model *AssetFamiliesCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}