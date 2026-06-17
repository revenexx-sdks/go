package models

import (
    "encoding/json"
    "errors"
)

// Model
type Products struct {
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Completeness interface{} `json:"completeness"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    FamilyId string `json:"family_id"`
    // 
    FamilyVariantId string `json:"family_variant_id"`
    // 
    Id string `json:"id"`
    // 
    Kind string `json:"kind"`
    // 
    ParentId string `json:"parent_id"`
    // 
    QuantifiedAssociations interface{} `json:"quantified_associations"`
    // 
    Sku string `json:"sku"`
    // 
    TaxClass string `json:"tax_class"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Products) New(data []byte) *Products {
    model.data = data
    return &model
}

func (model *Products) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}