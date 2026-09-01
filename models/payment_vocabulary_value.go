package models

import (
    "encoding/json"
    "errors"
)

// PaymentVocabularyValue One permitted value, with the words and the colour a
// human reads for it.
type PaymentVocabularyValue struct {
    // One sentence on what the value means, or null where the key speaks for
    // itself. A plain string, or a locale map keyed by language tag ({ "en": …,
    // "de": … }). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // This value ends the lifecycle — the honest way to ask "is this still
    // open?" instead of matching status names.
    Final bool `json:"final"`
    // The value exactly as the database stores it — what a filter sends and
    // what a row carries.
    Key string `json:"key"`
    // The label to show for this value. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Title interface{} `json:"title"`
    // What the state MEANS, semantically: neutral, info, success, warning or
    // danger. The client decides what each one looks like in its own design
    // system.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model PaymentVocabularyValue) New(data []byte) *PaymentVocabularyValue {
    model.data = data
    return &model
}

func (model *PaymentVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}