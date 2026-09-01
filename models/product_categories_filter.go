package models

import (
    "encoding/json"
    "errors"
)

// ProductCategoriesFilter The exact-column filters this call was understood
// to carry, verbatim as they arrived. A query parameter that is not a column
// of `product_categories` — `?status=`, a typo, a filter another entity has
// — is DROPPED and does not appear here, and the list comes back
// unfiltered. This object is the only way to tell that apart from "nothing
// matched".
type ProductCategoriesFilter struct {
    // The literal `?category_id=` value this call was understood to carry.
    CategoryId string `json:"category_id"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?position=` value this call was understood to carry.
    Position string `json:"position"`
    // The literal `?product_id=` value this call was understood to carry.
    ProductId string `json:"product_id"`
    // The literal `?source=` value this call was understood to carry.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model ProductCategoriesFilter) New(data []byte) *ProductCategoriesFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *ProductCategoriesFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}