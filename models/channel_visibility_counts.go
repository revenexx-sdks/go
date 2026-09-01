package models

import (
    "encoding/json"
    "errors"
)

// ChannelVisibilityCounts The three tallies, so a caller can log or alert on
// a batch without walking it.
type ChannelVisibilityCounts struct {
    // How many must not be. A batch where this equals `total` and the reason is
    // no_channel_context means the channel did not resolve, not that the
    // assortment is empty.
    Hidden int `json:"hidden"`
    // How many rows were decided — the length of the `items` sent.
    Total int `json:"total"`
    // How many may be shown.
    Visible int `json:"visible"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVisibilityCounts) New(data []byte) *ChannelVisibilityCounts {
    model.data = data
    return &model
}

func (model *ChannelVisibilityCounts) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}