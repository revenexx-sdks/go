package models

import (
    "encoding/json"
    "errors"
)

// ChannelVocabularyValue model.
type ChannelVocabularyValue struct {
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // Table-backed vocabularies only: the localized descriptions. A locale map
    // keyed by language tag: {"en": …, "de": …}. Read the requested tag and
    // fall back to the plain column beside it.
    Descriptions interface{} `json:"descriptions"`
    // The value ends the lifecycle.
    Final bool `json:"final"`
    // Table-backed vocabularies only: the value a create falls back to.
    IsDefault bool `json:"is_default"`
    // Table-backed vocabularies only: seeded on install rather than added by the
    // tenant. Still renameable and retirable.
    IsSystem bool `json:"is_system"`
    // The value as the database stores and enforces it.
    Key string `json:"key"`
    // Table-backed vocabularies only: the localized titles. `title` stays the
    // fallback. A locale map keyed by language tag: {"en": …, "de": …}. Read
    // the requested tag and fall back to the plain column beside it.
    Labels interface{} `json:"labels"`
    // A plain string, or a locale map keyed by language tag ({"en": …, "de":
    // …}). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model ChannelVocabularyValue) New(data []byte) *ChannelVocabularyValue {
    model.data = data
    return &model
}

func (model *ChannelVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}