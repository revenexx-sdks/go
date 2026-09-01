package models

import (
    "encoding/json"
    "errors"
)

// FamiliesUpdateRequest Partial update — omitted fields keep their current
// value.
type FamiliesUpdateRequest struct {
    // The family's stable identifier — which set of attributes a product of
    // this family HAS. Unique per tenant, and the value `GET
    // /products/attribute-schema?family_code=` resolves.
    Code string `json:"code"`
    // Which attribute code carries the product's main image — the one a grid
    // thumbnail and a picker read.
    ImageAttribute string `json:"image_attribute"`
    // Which attribute CODE carries the display name of a product in this family.
    // A product's name is an attribute, not a column, and which attribute it is,
    // is per family. Null falls back to the `default_label_attribute` setting and
    // then to the conventional `name`.
    LabelAttribute string `json:"label_attribute"`
    // What the family is called, per language tag — the name an operator picks
    // from, while the code is what everything else joins on.
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model FamiliesUpdateRequest) New(data []byte) *FamiliesUpdateRequest {
    model.data = data
    return &model
}

func (model *FamiliesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}