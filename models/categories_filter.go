package models

import (
    "encoding/json"
    "errors"
)

// CategoriesFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `categories` — `?status=`, a typo, a filter another entity has — is
// DROPPED and does not appear here, and the list comes back unfiltered. This
// object is the only way to tell that apart from "nothing matched".
type CategoriesFilter struct {
    // The literal `?code=` value this call was understood to carry.
    Code string `json:"code"`
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?labels=` value this call was understood to carry.
    Labels string `json:"labels"`
    // The literal `?parent_id=` value this call was understood to carry.
    ParentId string `json:"parent_id"`
    // The literal `?path=` value this call was understood to carry.
    Path string `json:"path"`
    // The literal `?position=` value this call was understood to carry.
    Position string `json:"position"`
    // The literal `?rule_match=` value this call was understood to carry.
    RuleMatch string `json:"rule_match"`
    // The literal `?rules=` value this call was understood to carry.
    Rules string `json:"rules"`
    // The literal `?rules_computed_at=` value this call was understood to carry.
    RulesComputedAt string `json:"rules_computed_at"`
    // The literal `?updated_at=` value this call was understood to carry.
    UpdatedAt string `json:"updated_at"`
    // The literal `?values=` value this call was understood to carry.
    Values string `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model CategoriesFilter) New(data []byte) *CategoriesFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *CategoriesFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}