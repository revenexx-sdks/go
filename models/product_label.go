package models

import (
    "encoding/json"
    "errors"
)

// ProductLabel model.
type ProductLabel struct {
    // The attribute code the name was read from.
    Attribute string `json:"attribute"`
    // How that attribute was chosen: 'family' is the product's own
    // `families.label_attribute`, 'setting' the tenant's
    // `default_label_attribute`, 'convention' the built-in fallback to `name`
    // when neither says anything.
    AttributeFrom string `json:"attribute_from"`
    // The product's id.
    Id string `json:"id"`
    // The name to show. Never empty — read `source` before treating it as a
    // name, because `sku` there means this is the SKU standing in for one.
    Label string `json:"label"`
    // Which locale the value came out of, when it came from a locale bucket. Null
    // for a value in `common` and for the SKU fallback.
    Locale string `json:"locale"`
    // The SKU, which is also the fallback shown as `label` when the catalog holds
    // no name.
    Sku string `json:"sku"`
    // Which bucket of attribute_values the name came from. 'sku' means the
    // catalog holds no name for this product — show that as a missing name, not
    // as a name.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model ProductLabel) New(data []byte) *ProductLabel {
    model.data = data
    return &model
}

func (model *ProductLabel) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}