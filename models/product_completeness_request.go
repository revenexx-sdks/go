package models

import (
    "encoding/json"
    "errors"
)

// ProductCompletenessRequest No body. Everything this needs is the path id
// and what the catalog already holds; send `{}`.
type ProductCompletenessRequest struct {

    // Used by Decode() method
    data []byte
}

func (model ProductCompletenessRequest) New(data []byte) *ProductCompletenessRequest {
    model.data = data
    return &model
}

func (model *ProductCompletenessRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}