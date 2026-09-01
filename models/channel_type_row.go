package models

import (
    "encoding/json"
    "errors"
)

// ChannelTypeRow model.
type ChannelTypeRow struct {
    // What `channels.type` stores. Immutable once created — renaming it would
    // orphan every channel that carries it, and there is no FK behind
    // `channels.type` to cascade. A fresh install seeds storefront, punchout,
    // marketplace, api, pos; a merchant may retire any of them and add their own.
    Code string `json:"code"`
    // When the row was inserted, set by the database.
    CreatedAt string `json:"created_at"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Descriptions interface{} `json:"descriptions"`
    // Row id, and the only handle GET/PUT/DELETE /channels/types/{id} accept. Not
    // the type `code`. No example is published because no id this app could
    // invent names a row a tenant holds.
    Id string `json:"id"`
    // The type a channel created without one gets. Exactly one row carries it.
    IsDefault bool `json:"is_default"`
    // Seeded on install rather than added by the merchant. A flag about origin
    // only — a system type is still renameable, reorderable and retirable.
    IsSystem bool `json:"is_system"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Labels interface{} `json:"labels"`
    // Sort position. GET /channels/types always answers in this order and takes
    // no `order` parameter. It is not unique and defaults to 0, so ties are
    // broken by `code` — the order is total, which is what makes paging the
    // list safe to walk.
    Position int `json:"position"`
    // The tenant that owns this row. Added by the data plane, not by this app: it
    // is not a column of schema.json, so it is read-only and `?tenant_id=` is not
    // a filter — the key is silently dropped and never reaches the `filter`
    // echo.
    TenantId string `json:"tenant_id"`
    // The fallback name. `labels` carries the per-locale ones. Rows seeded before
    // 0.7.0 hold a serialized locale map here instead (PE-452).
    Title interface{} `json:"title"`
    // Semantic badge colour for this type, for a client that renders the list.
    // The client owns what each tone looks like; the value only says what it
    // MEANS.
    Tone string `json:"tone"`
    // When the row was last written, set by the database.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ChannelTypeRow) New(data []byte) *ChannelTypeRow {
    model.data = data
    return &model
}

func (model *ChannelTypeRow) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}