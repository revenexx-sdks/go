package models

import (
    "encoding/json"
    "errors"
)

// OrderListVocabularyValue model.
type OrderListVocabularyValue struct {
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Localized descriptions of a tenant-owned value, keyed by locale.
    Descriptions interface{} `json:"descriptions"`
    // The value ends the lifecycle. Always false for `kinds` — a list kind is
    // not a state.
    Final bool `json:"final"`
    // The value a create falls back to, so a client can mark it without reading
    // the settings as well.
    IsDefault bool `json:"is_default"`
    // Seeded on install rather than created by the tenant. Still renameable and
    // retirable.
    IsSystem bool `json:"is_system"`
    // The value as the database stores and enforces it — for `kinds`, the
    // `code` a list carries.
    Key string `json:"key"`
    // Localized titles of a tenant-owned value, keyed by locale.
    Labels interface{} `json:"labels"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model OrderListVocabularyValue) New(data []byte) *OrderListVocabularyValue {
    model.data = data
    return &model
}

func (model *OrderListVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}