package models

import (
    "encoding/json"
    "errors"
)

// ChannelVisibilityRequest model.
type ChannelVisibilityRequest struct {
    // The channel `code` (the scope slug) to evaluate against, trimmed and
    // lowercased before it is matched. Optional, and through api.revenexx.com it
    // is the ONLY way to name a channel explicitly: the x-revenexx-channel header
    // is not forwarded to the app, so without this the resolution falls through
    // to the scope_context.channel claim and then to the tenant's default
    // channel. A code no channel carries is not an error — the answer is
    // resolved:false with reason 'unknown_channel', so a caller can tell it from
    // an outage.
    Channel string `json:"channel"`
    // The rows to decide on, each with the channel assignments Baseline holds for
    // it. POST /api/v1/scopes/lookup?dimension=channel answers in exactly this
    // shape. At most 500 — Baseline's own lookup ceiling.
    Items []ChannelVisibilityItem `json:"items"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVisibilityRequest) New(data []byte) *ChannelVisibilityRequest {
    model.data = data
    return &model
}

func (model *ChannelVisibilityRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}