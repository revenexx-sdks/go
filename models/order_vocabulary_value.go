package models

import (
    "encoding/json"
    "errors"
)

// OrderVocabularyValue One permitted value with the words and the badge tone
// a client should render for it.
type OrderVocabularyValue struct {
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Description string `json:"description"`
    // True when this value ENDS the lifecycle. Lets a reader ask "is this order
    // still open?" instead of matching status names it guessed.
    Final bool `json:"final"`
    // The value as stored — exactly what the CHECK constraint permits.
    Key string `json:"key"`
    // Only on 'return-resolutions': which return transition accepts this value. A
    // settlement word on the refusal dialog is how the two sets got mixed up.
    Stage string `json:"stage"`
    // Either one string, or a map of locale to string ({"en": …, "de": …}).
    Title string `json:"title"`
    // Semantic badge colour. The client owns what each tone looks like.
    Tone string `json:"tone"`

    // Used by Decode() method
    data []byte
}

func (model OrderVocabularyValue) New(data []byte) *OrderVocabularyValue {
    model.data = data
    return &model
}

func (model *OrderVocabularyValue) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}