package models

import (
    "encoding/json"
    "errors"
)

// ChannelTypeUpdateRequest model.
type ChannelTypeUpdateRequest struct {
    // Replace the one-sentence description. Sent as null it is cleared; omitted
    // it is kept. `descriptions` carries the per-locale ones.
    Description string `json:"description"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Descriptions interface{} `json:"descriptions"`
    // Promote this type; the previous default is demoted. Only `true` does
    // anything — sending false does not demote this type, because some type
    // must hold the flag.
    IsDefault bool `json:"is_default"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Labels interface{} `json:"labels"`
    // Move the type in the order GET /channels/types answers in.
    Position int `json:"position"`
    // Rename the type. A blank or non-string title is ignored, not refused —
    // the stored one is kept.
    Title string `json:"title"`
    // Change the badge colour. A value outside the palette is ignored rather than
    // refused, and the stored tone is kept.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ChannelTypeUpdateRequest) New(data []byte) *ChannelTypeUpdateRequest {
    model.data = data
    return &model
}

func (model *ChannelTypeUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}