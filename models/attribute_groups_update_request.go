package models

import (
    "encoding/json"
    "errors"
)

// AttributeGroupsUpdateRequest Partial update — omitted fields keep their
// current value.
type AttributeGroupsUpdateRequest struct {
    // The group's stable identifier, and the value an `AttributeField` carries as
    // its `group` — a SECTION of the product form, not a label. Unique per
    // tenant and the key an import joins on.
    Code string `json:"code"`
    // The section heading a person sees, keyed by language tag. The code is never
    // shown to an operator; a tag nobody translated falls back to the next filled
    // one, then to English.
    Labels interface{} `json:"labels"`
    // Where this section sits in a form, ascending. Sections that tie keep the
    // order the database returns them in.
    Position int `json:"position"`

    // Used by Decode() method
    data []byte
}

func (model AttributeGroupsUpdateRequest) New(data []byte) *AttributeGroupsUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributeGroupsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}