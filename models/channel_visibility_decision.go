package models

import (
    "encoding/json"
    "errors"
)

// ChannelVisibilityDecision model.
type ChannelVisibilityDecision struct {
    // The id as it was sent, verbatim.
    Id string `json:"id"`
    // Why the row was shown or hidden — the answer is auditable, not a bare
    // boolean.
    Reason string `json:"reason"`
    // Whether this row may be shown in the resolved channel. The same answer as
    // membership in `visible`; `reason` says why.
    Visible bool `json:"visible"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVisibilityDecision) New(data []byte) *ChannelVisibilityDecision {
    model.data = data
    return &model
}

func (model *ChannelVisibilityDecision) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}