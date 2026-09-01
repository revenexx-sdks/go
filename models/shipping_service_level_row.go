package models

import (
    "encoding/json"
    "errors"
)

// ShippingServiceLevelRow model.
type ShippingServiceLevelRow struct {
    // What `shipping_carriers.service_level` stores. Immutable once created —
    // renaming it would orphan every row carrying it.
    Code string `json:"code"`
    // When the row was created (UTC).
    CreatedAt string `json:"created_at"`
    // The sentence under the title, explaining when to pick this service level.
    // Null when the title says enough.
    Description string `json:"description"`
    // Localized descriptions. A flat map keyed by locale — the Cockpit falls
    // back to `en`. Null means the row has no translations and every client shows
    // the untranslated column instead.
    Descriptions interface{} `json:"descriptions"`
    // Row id, assigned by the database on insert.
    Id string `json:"id"`
    // The service level a fallback lands on. Exactly one row carries it, and POST
    // …/make-default is what moves it.
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

func (model ShippingServiceLevelRow) New(data []byte) *ShippingServiceLevelRow {
    model.data = data
    return &model
}

func (model *ShippingServiceLevelRow) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}