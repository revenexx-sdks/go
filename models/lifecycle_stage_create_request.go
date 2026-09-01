package models

import (
    "encoding/json"
    "errors"
)

// LifecycleStageCreateRequest Add one value to the lifecycle stages set. It
// is available to `organizations.lifecycle_stage` immediately.
type LifecycleStageCreateRequest struct {
    // What `organizations.lifecycle_stage` will store. Lowercase, starting with a
    // letter; immutable afterwards.
    Code string `json:"code"`
    // One line of help for whoever picks this value.
    Description string `json:"description"`
    // Localized descriptions, keyed by language tag ({ "en": …, "de": … }).
    // Null when nobody translated this value — a client then falls back to
    // `description`.
    Descriptions interface{} `json:"descriptions"`
    // Promote this value; the previous default is demoted in the same call.
    IsDefault bool `json:"is_default"`
    // Localized titles, keyed by language tag ({ "en": …, "de": … }). Null
    // when nobody translated this value — a client then falls back to `title`.
    Labels interface{} `json:"labels"`
    // Where it sits in the set, ascending. Default 0.
    Position int `json:"position"`
    // The fallback name shown when no locale matches.
    Title string `json:"title"`
    // Semantic badge colour.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model LifecycleStageCreateRequest) New(data []byte) *LifecycleStageCreateRequest {
    model.data = data
    return &model
}

func (model *LifecycleStageCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}