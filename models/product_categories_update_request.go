package models

import (
    "encoding/json"
    "errors"
)

// ProductCategoriesUpdateRequest Partial update — omitted fields keep their
// current value.
type ProductCategoriesUpdateRequest struct {
    // The category it is filed into. One row per (product, category), whichever
    // way it got there.
    CategoryId string `json:"category_id"`
    // Sort order of this product inside the category.
    Position int `json:"position"`
    // The product filed into the category. Deleting the product deletes the
    // membership with it.
    ProductId string `json:"product_id"`
    // How the membership came about: 'manual' is hand-picked, 'rule' was
    // materialized by a category rule. The two never touch each other — a
    // recompute only ever inserts and deletes `rule` rows, so a hand-picked
    // membership survives every pass.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model ProductCategoriesUpdateRequest) New(data []byte) *ProductCategoriesUpdateRequest {
    model.data = data
    return &model
}

func (model *ProductCategoriesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}