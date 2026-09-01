package models

import (
    "encoding/json"
    "errors"
)

// OrderListKindCreateRequest model.
type OrderListKindCreateRequest struct {
    // What `lists.kind` will store. Lowercased on the way in and immutable
    // afterwards — a merchant who wants a different code creates a new kind and
    // moves the lists over.
    Code string `json:"code"`
    // What this kind is for, in one sentence — the line a select shows under
    // the title.
    Description string `json:"description"`
    // Localized descriptions, keyed by language tag.
    Descriptions interface{} `json:"descriptions"`
    // Promote this kind; the previous default is demoted.
    IsDefault bool `json:"is_default"`
    // Localized titles, keyed by language tag.
    Labels interface{} `json:"labels"`
    // Where the kind sits in a select, ascending. Omitted means 0, which puts it
    // first among the unpositioned.
    Position int `json:"position"`
    // What a person reads. `labels` adds the localized forms on top; this one is
    // the fallback.
    Title string `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like; omitted
    // means `neutral`.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model OrderListKindCreateRequest) New(data []byte) *OrderListKindCreateRequest {
    model.data = data
    return &model
}

func (model *OrderListKindCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}