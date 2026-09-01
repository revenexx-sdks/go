package models

import (
    "encoding/json"
    "errors"
)

// FamilyVariantsCreateRequest model.
type FamilyVariantsCreateRequest struct {
    // The attribute codes a product model splits its variants on. Two shapes are
    // in the wild and both are read: a bare list of codes, or one entry per
    // level, outermost first — `[{"level": 1, "axes": ["colour"]}, {"level": 2,
    // "axes": ["size"]}]`. An attribute named here is READ-ONLY on the model and
    // set on each variant, which is what `AttributeField.readonly_reason`
    // reports.
    Axes interface{} `json:"axes"`
    // The variant structure's stable identifier — how this family splits, not
    // which product it splits. Unique per tenant.
    Code string `json:"code"`
    // The family this variant structure belongs to. A family may carry several,
    // and a product names the one it follows through `family_variant_id`.
    FamilyId string `json:"family_id"`
    // What the variant structure is called, per language tag.
    Labels interface{} `json:"labels"`

    // Used by Decode() method
    data []byte
}

func (model FamilyVariantsCreateRequest) New(data []byte) *FamilyVariantsCreateRequest {
    model.data = data
    return &model
}

func (model *FamilyVariantsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}