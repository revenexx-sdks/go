package models

import (
    "encoding/json"
    "errors"
)

// InventoryVocabulary model.
type InventoryVocabulary struct {
    // This app's name — the part before the dot in the qualified id.
    App string `json:"app"`
    // True when these values are the complete permitted set, because they were
    // read out of a CHECK constraint. A value outside a closed set is therefore
    // stale data, not a missing label — which is what lets a client show it as
    // an error instead of inventing a title for it.
    Closed bool `json:"closed"`
    // The tone a value gets when nobody has labelled it — a value added to the
    // CHECK constraint is served with its key humanized and this tone, rather
    // than not being served at all.
    DefaultTone string `json:"default_tone"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Description interface{} `json:"description"`
    // The vocabulary name, echoed — the part after the dot in the qualified id.
    Name string `json:"name"`
    // Where the words come from: 'schema' — the app's own, read from the
    // constraint. Nothing here is renameable per tenant, so a client may cache it
    // per app version.
    Source string `json:"source"`
    // A plain string, or a locale map keyed by language tag ({ "en": …, "de":
    // … }). Read the requested tag, fall back to `en`.
    Title interface{} `json:"title"`
    // Every permitted value, IN CONSTRAINT ORDER — which is lifecycle order for
    // a status, so a UI can render the steps in the order they happen.
    Values []interface{} `json:"values"`

    // Used by Decode() method
    data []byte
}

func (model InventoryVocabulary) New(data []byte) *InventoryVocabulary {
    model.data = data
    return &model
}

func (model *InventoryVocabulary) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}