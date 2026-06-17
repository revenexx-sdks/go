package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type ProductCategoriesUpdateRequest struct {
    // 
    CategoryId string `json:"category_id"`
    // 
    Position int `json:"position"`
    // 
    ProductId string `json:"product_id"`

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