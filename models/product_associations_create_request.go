package models

import (
    "encoding/json"
    "errors"
)

// Model
type ProductAssociationsCreateRequest struct {
    // 
    AssociationTypeId string `json:"association_type_id"`
    // 
    Position int `json:"position"`
    // 
    ProductId string `json:"product_id"`
    // 
    Quantity float64 `json:"quantity"`
    // 
    TargetProductId string `json:"target_product_id"`

    // Used by Decode() method
    data []byte
}

func (model ProductAssociationsCreateRequest) New(data []byte) *ProductAssociationsCreateRequest {
    model.data = data
    return &model
}

func (model *ProductAssociationsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}