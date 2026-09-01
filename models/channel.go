package models

import (
    "encoding/json"
    "errors"
)

// Channel model.
type Channel struct {
    // The scope slug Baseline matches channel assignments on
    // (manifest.provides_scopes[].slug_source). Unique per tenant and, in
    // practice, immutable — changing it orphans every assignment made against
    // it.
    Code string `json:"code"`
    // When the row was inserted, set by the database.
    CreatedAt string `json:"created_at"`
    // Row id, and the only handle GET/PUT/DELETE /channels/{id} accept. Not the
    // scope slug — that is `code`. No example is published because no id this
    // app could invent names a row a tenant holds.
    Id string `json:"id"`
    // The channel a request that names none falls back to. At most one channel
    // carries it.
    IsDefault bool `json:"is_default"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Labels interface{} `json:"labels"`
    // Display name. `labels` carries the per-locale ones.
    Name string `json:"name"`
    // Sort position — ascending, and the tiebreak when two channels both claim
    // is_default.
    Position int `json:"position"`
    // Whether the channel is in service. What 'inactive' DOES is the tenant's
    // inactive_channel_behavior setting: on 'serve' it is a label and the channel
    // still resolves, on 'block' /channels/context answers resolved:false with
    // reason 'channel_inactive'. Served as the 'channels.statuses' vocabulary.
    Status string `json:"status"`
    // The tenant that owns this row. Added by the data plane, not by this app: it
    // is not a column of schema.json, so it is read-only and `?tenant_id=` is not
    // a filter — the key is silently dropped and never reaches the `filter`
    // echo.
    TenantId string `json:"tenant_id"`
    // One of the codes the tenant keeps under GET /channels/types — served with
    // labels as the 'channels.types' vocabulary. Deliberately NOT an enum: the
    // set is the tenant's own rows, not a CHECK constraint this repo could quote.
    // A fresh install starts with storefront, punchout, marketplace, api, pos,
    // which is why 'storefront' is the example here, but a merchant may rename or
    // retire any of them and add their own (a feed or a print channel), so read
    // the list rather than assuming it.
    Type string `json:"type"`
    // What it means, IN THIS CHANNEL, that a row carries no channel assignment at
    // all — the per-channel override of the tenant-wide
    // unassigned_channel_visibility setting. 'inherit' (the default) takes the
    // tenant's answer and changes nothing. 'all' shows unassigned rows:
    // everything is on sale unless somebody carved it out, which is what an open
    // storefront wants and what Baseline's is_visible() does today.
    // 'assigned_only' hides them until they are explicitly assigned — the
    // negotiated assortment a punchout contract describes, and the one answer the
    // generated _scoped view has no way to express, which is why POST
    // /channels/visibility exists to apply it. Rows that DO carry assignments are
    // unaffected either way. Served with its labels as the
    // 'channels.unassigned-visibility' vocabulary.
    UnassignedVisibility string `json:"unassigned_visibility"`
    // When the row was last written, set by the database.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Channel) New(data []byte) *Channel {
    model.data = data
    return &model
}

func (model *Channel) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}