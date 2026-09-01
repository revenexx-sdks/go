package models

import (
    "encoding/json"
    "errors"
)

// ChannelVisibilityItem model.
type ChannelVisibilityItem struct {
    // The row's channel scope slugs. Empty or absent means unassigned — the
    // case the policy decides.
    Channels []string `json:"channels"`
    // The row id, echoed back on the decision. Opaque to this app — it is never
    // looked up, so any non-empty string is accepted and nothing has to exist. In
    // practice it is the entity id POST /api/v1/scopes/lookup answered with,
    // which is what the example shows.
    Id string `json:"id"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVisibilityItem) New(data []byte) *ChannelVisibilityItem {
    model.data = data
    return &model
}

func (model *ChannelVisibilityItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}