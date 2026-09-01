package models

import (
    "encoding/json"
    "errors"
)

// ProductsFilter The exact-column filters this call was understood to carry,
// verbatim as they arrived. A query parameter that is not a column of
// `products` — `?status=`, a typo, a filter another entity has — is
// DROPPED and does not appear here, and the list comes back unfiltered. This
// object is the only way to tell that apart from "nothing matched".
type ProductsFilter struct {
    // The literal `?attribute_values=` value this call was understood to carry.
    AttributeValues string `json:"attribute_values"`
    // The literal `?completeness=` value this call was understood to carry.
    Completeness string `json:"completeness"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?deleted_at=` value this call was understood to carry.
    DeletedAt string `json:"deleted_at"`
    // The literal `?enabled=` value this call was understood to carry.
    Enabled string `json:"enabled"`
    // The literal `?family_id=` value this call was understood to carry.
    FamilyId string `json:"family_id"`
    // The literal `?family_variant_id=` value this call was understood to carry.
    FamilyVariantId string `json:"family_variant_id"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?kind=` value this call was understood to carry.
    Kind string `json:"kind"`
    // The literal `?label=` value this call was understood to carry.
    Label string `json:"label"`
    // The literal `?parent_id=` value this call was understood to carry.
    ParentId string `json:"parent_id"`
    // The literal `?quantified_associations=` value this call was understood to
    // carry.
    QuantifiedAssociations string `json:"quantified_associations"`
    // The literal `?sku=` value this call was understood to carry.
    Sku string `json:"sku"`
    // The literal `?tax_class=` value this call was understood to carry.
    TaxClass string `json:"tax_class"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ProductsFilter) New(data []byte) *ProductsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *ProductsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}