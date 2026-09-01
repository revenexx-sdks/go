package models

import (
    "encoding/json"
    "errors"
)

// PagesVocabularyRef One vocabulary, named but not unpacked.
type PagesVocabularyRef struct {
    // What the set is for, or null. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Description interface{} `json:"description"`
    // The name to fetch it by — the part after the dot in the qualified id.
    Name string `json:"name"`
    // What this set of values is called. A plain string, or a locale map keyed by
    // language tag ({ "en": …, "de": … }). Read the requested tag, fall back
    // to `en`.
    Title interface{} `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model PagesVocabularyRef) New(data []byte) *PagesVocabularyRef {
    model.data = data
    return &model
}

func (model *PagesVocabularyRef) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}