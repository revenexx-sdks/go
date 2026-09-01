package models

import (
    "encoding/json"
    "errors"
)

// FamilyAttributesUpdateRequest Partial update — omitted fields keep their
// current value.
type FamilyAttributesUpdateRequest struct {
    // The attribute the family carries. One row per (family, attribute); deleting
    // either side deletes the link.
    AttributeId string `json:"attribute_id"`
    // The family this link belongs to — one side of the pair that makes an
    // attribute part of a family's form.
    FamilyId string `json:"family_id"`
    // The attribute has to carry a value for a product of this family to count as
    // complete. `POST /products/{id}/completeness` measures exactly these and
    // nothing else.
    IsRequired bool `json:"is_required"`
    // The family's own ordering of this attribute, which overrides the
    // attribute's default `position` in this family's form.
    Position int `json:"position"`
    // Narrows `is_required` to named channels. NULL or an empty list means
    // required EVERYWHERE, not nowhere — that is how every required link in the
    // wild is stored, and reading an empty list as "nowhere" reports a fully
    // configured family as demanding nothing.
    RequiredChannels interface{} `json:"required_channels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyAttributesUpdateRequest) New(data []byte) *FamilyAttributesUpdateRequest {
    model.data = data
    return &model
}

func (model *FamilyAttributesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}