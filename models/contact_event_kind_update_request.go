package models

import (
    "encoding/json"
    "errors"
)

// ContactEventKindUpdateRequest Everything but `code`. Sending a different
// one is a 400 rather than a silent no-op, because records already store it.
type ContactEventKindUpdateRequest struct {
    // One line of help for whoever picks this value.
    Description string `json:"description"`
    // Localized descriptions, keyed by language tag ({ "en": …, "de": … }).
    // Null when nobody translated this value — a client then falls back to
    // `description`.
    Descriptions interface{} `json:"descriptions"`
    // Promote this value; the previous default is demoted.
    IsDefault bool `json:"is_default"`
    // Localized titles, keyed by language tag ({ "en": …, "de": … }). Null
    // when nobody translated this value — a client then falls back to `title`.
    Labels interface{} `json:"labels"`
    // Where it sits in the set, ascending.
    Position int `json:"position"`
    // The fallback name shown when no locale matches.
    Title string `json:"title"`
    // Semantic badge colour.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ContactEventKindUpdateRequest) New(data []byte) *ContactEventKindUpdateRequest {
    model.data = data
    return &model
}

func (model *ContactEventKindUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}