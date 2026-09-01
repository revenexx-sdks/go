package models

import (
    "encoding/json"
    "errors"
)

// ChannelTypeCreateRequest model.
type ChannelTypeCreateRequest struct {
    // What `channels.type` will store. Lowercased and trimmed before it is
    // written, and fixed from then on — a rename would orphan every channel
    // carrying it.
    Code string `json:"code"`
    // One sentence on what kind of place this type of channel is, for the
    // merchant choosing between them. Plain text, in the tenant's primary
    // language; `descriptions` carries the per-locale ones.
    Description string `json:"description"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Descriptions interface{} `json:"descriptions"`
    // Promote this type; the previous default is demoted. The default is the type
    // a channel created without one gets.
    IsDefault bool `json:"is_default"`
    // A locale map keyed by language tag: {"en": …, "de": …}. Read the
    // requested tag and fall back to the plain column beside it.
    Labels interface{} `json:"labels"`
    // Sort position (default 0). GET /channels/types answers in this order; ties
    // fall back to the code.
    Position int `json:"position"`
    // The fallback name. `labels` carries the per-locale ones.
    Title string `json:"title"`
    // Badge colour (default 'neutral'). A value outside the palette is ignored
    // rather than refused.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ChannelTypeCreateRequest) New(data []byte) *ChannelTypeCreateRequest {
    model.data = data
    return &model
}

func (model *ChannelTypeCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}