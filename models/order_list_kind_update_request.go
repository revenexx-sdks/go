package models

import (
    "encoding/json"
    "errors"
)

// OrderListKindUpdateRequest model.
type OrderListKindUpdateRequest struct {
    // What this kind is for, in one sentence. Explicit null clears it.
    Description string `json:"description"`
    // Localized descriptions, keyed by language tag. Replaces the whole map
    // rather than merging into it.
    Descriptions interface{} `json:"descriptions"`
    // True promotes this kind and demotes the previous default — the same move
    // POST /orderlists/kinds/{id}/make-default makes on its own.
    IsDefault bool `json:"is_default"`
    // Localized titles, keyed by language tag. Replaces the whole map rather than
    // merging into it.
    Labels interface{} `json:"labels"`
    // Where the kind sits in a select, ascending.
    Position int `json:"position"`
    // What a person reads. A blank title is ignored rather than stored — a kind
    // with no words is unreadable in every UI.
    Title string `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model OrderListKindUpdateRequest) New(data []byte) *OrderListKindUpdateRequest {
    model.data = data
    return &model
}

func (model *OrderListKindUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}