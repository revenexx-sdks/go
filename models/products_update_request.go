package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type ProductsUpdateRequest struct {
    // 
    AttributeValues interface{} `json:"attribute_values"`
    // 
    Completeness interface{} `json:"completeness"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    FamilyId string `json:"family_id"`
    // 
    FamilyVariantId string `json:"family_variant_id"`
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

    // Used by Decode() method
    data []byte
}

func (model ProductsUpdateRequest) New(data []byte) *ProductsUpdateRequest {
    model.data = data
    return &model
}

func (model *ProductsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}