package models

import (
    "encoding/json"
    "errors"
)

// ShippingWeightUnitRow model.
type ShippingWeightUnitRow struct {
    // What a rate request names in `weight_unit`, and what a market's
    // `weight_unit` setting stores. Immutable once created — renaming it would
    // orphan every row carrying it.
    Code string `json:"code"`
    // When the row was created (UTC).
    CreatedAt string `json:"created_at"`
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
    // Row id, assigned by the database on insert.
    Id string `json:"id"`
    // The anchor every other factor is expressed in. Exactly one row, fixed at
    // install, not writable and not deletable — moving it would silently
    // reprice every weight matrix.
    IsBase bool `json:"is_base"`
    // The unit a market whose `weight_unit` setting is unset keys its tiers in.
    // Exactly one row carries it.
    IsDefault bool `json:"is_default"`
    // Seeded on install rather than typed by the merchant. Still renameable and
    // still deletable; it only says where the row came from.
    IsSystem bool `json:"is_system"`
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
    // When the row was last written (UTC).
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ShippingWeightUnitRow) New(data []byte) *ShippingWeightUnitRow {
    model.data = data
    return &model
}

func (model *ShippingWeightUnitRow) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}