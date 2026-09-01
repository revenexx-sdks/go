package models

import (
    "encoding/json"
    "errors"
)

// ProductLabelsRequest model.
type ProductLabelsRequest struct {
    // Product ids to name. At most 500.
    Ids []string `json:"ids"`
    // Product SKUs to name. At most 500.
    Skus []string `json:"skus"`

    // Used by Decode() method
    data []byte
}

func (model ProductLabelsRequest) New(data []byte) *ProductLabelsRequest {
    model.data = data
    return &model
}

func (model *ProductLabelsRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}