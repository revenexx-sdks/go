package models

import (
    "encoding/json"
    "errors"
)

// ProductGridRow model.
type ProductGridRow struct {
    // The grid cells: one key per attribute code that `columns` lists with
    // `source: "attribute"`, holding the value already resolved out of
    // `attribute_values` for the requested context. A code the product carries no
    // value for is null rather than absent, so a row is the same shape whatever
    // it holds. The keys are the tenant's own attribute codes, which is why this
    // object has no fixed properties — read `columns` for the set.
    Attributes interface{} `json:"attributes"`
    // The stored `products.completeness` document, verbatim. Null means it has
    // never been computed — not that the product is empty.
    Completeness interface{} `json:"completeness"`
    // Whether the product is offered.
    Enabled bool `json:"enabled"`
    // That family's code, resolved here so a grid can show and group by it
    // without a second read.
    FamilyCode string `json:"family_code"`
    // The product's family. Null is the state that makes completeness impossible.
    FamilyId string `json:"family_id"`
    // The product's id — what a row click navigates with.
    Id string `json:"id"`
    // 'simple', 'model' or 'variant' — a model is a row a person should not
    // price or sell.
    Kind string `json:"kind"`
    // The resolved display name. Never empty; read `label_source` before showing
    // it as a name.
    Label string `json:"label"`
    // Which attribute code the name was read from, per this product's family.
    LabelAttribute string `json:"label_attribute"`
    // Which bucket of attribute_values the name came from. 'sku' means the
    // catalog holds no name for this product — show that as a missing name, not
    // as a name.
    LabelSource string `json:"label_source"`
    // The merchant's article number.
    Sku string `json:"sku"`
    // When the product row was last written — the column a "recently changed"
    // sort uses.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ProductGridRow) New(data []byte) *ProductGridRow {
    model.data = data
    return &model
}

func (model *ProductGridRow) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}