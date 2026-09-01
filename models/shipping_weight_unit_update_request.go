package models

import (
    "encoding/json"
    "errors"
)

// ShippingWeightUnitUpdateRequest Everything but the code. Sending a
// different code is a 400 rather than a silent no-op: renaming it would
// orphan every row that carries it.
type ShippingWeightUnitUpdateRequest struct {
    // The sentence under the title, explaining when to pick this weight unit.
    // Null when the title says enough.
    Description string `json:"description"`
    // Localized descriptions. A flat map keyed by locale — the Cockpit falls
    // back to `en`. Null means the row has no translations and every client shows
    // the untranslated column instead.
    Descriptions interface{} `json:"descriptions"`
    // How many BASE units (kilograms) one of this unit weighs — a tonne is
    // 1000, a gram 0.001, a pound 0.45359237. This number prices parcels: every
    // weight matrix converts a request through it. Must be > 0; the base unit is
    // fixed at 1 and rejects a change.
    Factor float64 `json:"factor"`
    // Promote this value; the previous default is demoted. POST …/make-default
    // does the same thing without an edit.
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

func (model ShippingWeightUnitUpdateRequest) New(data []byte) *ShippingWeightUnitUpdateRequest {
    model.data = data
    return &model
}

func (model *ShippingWeightUnitUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}