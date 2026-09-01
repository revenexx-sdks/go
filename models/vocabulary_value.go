package models

import (
    "encoding/json"
    "errors"
)

// VocabularyValue model.
type VocabularyValue struct {
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // A terminal state — nothing moves out of it. False or absent on a
    // vocabulary that is not a lifecycle.
    Final bool `json:"final"`
    // The value as it is STORED and as the CHECK admits it — what a filter or a
    // write sends.
    Key string `json:"key"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Which badge colour a UI should paint this value in.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model VocabularyValue) New(data []byte) *VocabularyValue {
    model.data = data
    return &model
}

func (model *VocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}