package models

import (
    "encoding/json"
    "errors"
)

// ContactEventKind One value of the activity types set. What kind of entry
// lands on a customer timeline. 'system' is the app's own decision trail and
// a caller may not file one, whatever the set says.
type ContactEventKind struct {
    // What `contact_events.kind` stores, and the only part of this row other data
    // depends on. Immutable once created: renaming it would orphan every record
    // carrying it.
    Code string `json:"code"`
    // When the value was added to this set.
    CreatedAt string `json:"created_at"`
    // One line of help for an operator choosing this value. Null when there is
    // nothing to add. A row seeded before 0.22.0 may hold a serialized locale map
    // here instead (PE-443).
    Description string `json:"description"`
    // Localized descriptions, keyed by language tag ({ "en": …, "de": … }).
    // Null when nobody translated this value — a client then falls back to
    // `description`.
    Descriptions interface{} `json:"descriptions"`
    // Primary key of this value. What the update and delete routes address it by
    // — the CODE is what records store.
    Id string `json:"id"`
    // The value a create falls back to when the caller names none. Exactly one
    // row of the set carries it; promoting another one demotes this.
    IsDefault bool `json:"is_default"`
    // True for a value this app seeded on install. Still renameable and still
    // removable — it only records where the value came from.
    IsSystem bool `json:"is_system"`
    // Localized titles, keyed by language tag ({ "en": …, "de": … }). Null
    // when nobody translated this value — a client then falls back to `title`.
    Labels interface{} `json:"labels"`
    // Where this value sits in the set, ascending. It is the order a select
    // should offer.
    Position int `json:"position"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // The fallback name — what a client shows when no locale in `labels`
    // matches. A row seeded before 0.22.0 may hold a serialized locale map here
    // instead (PE-443) — those rows were seeded with no `labels` at all.
    Title string `json:"title"`
    // Semantic badge colour. The palette stays fixed — it is a render concern,
    // not a merchant decision.
    Tone string `json:"tone"`
    // When it was last edited.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ContactEventKind) New(data []byte) *ContactEventKind {
    model.data = data
    return &model
}

func (model *ContactEventKind) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}