package models

import (
    "encoding/json"
    "errors"
)

// PagesVocabularyValue One permitted value of a vocabulary, with everything
// needed to render it.
type PagesVocabularyValue struct {
    // When to use this value, or null when nobody wrote one. A plain string, or a
    // locale map keyed by language tag ({ "en": …, "de": … }). Read the
    // requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // The value ends the lifecycle.
    Final bool `json:"final"`
    // The value as the database stores and enforces it.
    Key string `json:"key"`
    // What a person reads. Falls back to a humanized key. A plain string, or a
    // locale map keyed by language tag ({ "en": …, "de": … }). Read the
    // requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model PagesVocabularyValue) New(data []byte) *PagesVocabularyValue {
    model.data = data
    return &model
}

func (model *PagesVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}