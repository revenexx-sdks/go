package models

import (
    "encoding/json"
    "errors"
)

// ProductCategoryAssignRequest The category has to exist already; this route
// files a product into one, it does not create one.
type ProductCategoryAssignRequest struct {
    // The category to file the product into.
    CategoryId string `json:"category_id"`
    // Sort order inside the category. Default 0.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model ProductCategoryAssignRequest) New(data []byte) *ProductCategoryAssignRequest {
    model.data = data
    return &model
}

func (model *ProductCategoryAssignRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}