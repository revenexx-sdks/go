package models

import (
    "encoding/json"
    "errors"
)

// ShippingServiceLevelCreateRequest model.
type ShippingServiceLevelCreateRequest struct {
    // Lowercase letters, digits, - or _, starting with a letter. What
    // `shipping_carriers.service_level` stores. Immutable once created —
    // renaming it would orphan every row carrying it.
    Code string `json:"code"`
    // The sentence under the title, explaining when to pick this service level.
    // Null when the title says enough.
    Description string `json:"description"`
    // Localized descriptions. A flat map keyed by locale — the Cockpit falls
    // back to `en`. Null means the row has no translations and every client shows
    // the untranslated column instead.
    Descriptions interface{} `json:"descriptions"`
    // Promote this value on creation; the previous default is demoted.
    IsDefault bool `json:"is_default"`
    // Localized titles. A flat map keyed by locale — the Cockpit falls back to
    // `en`. Null means the row has no translations and every client shows the
    // untranslated column instead.
    Labels interface{} `json:"labels"`
    // Sort order in a select — the collection is returned in it.
    Position int `json:"position"`
    // What an operator reads in a select. The name a merchant renames; the code
    // underneath never moves.
    Title string `json:"title"`
    // Semantic badge colour for a UI listing the set. The client owns what each
    // tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ShippingServiceLevelCreateRequest) New(data []byte) *ShippingServiceLevelCreateRequest {
    model.data = data
    return &model
}

func (model *ShippingServiceLevelCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}