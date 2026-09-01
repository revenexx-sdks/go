package models

import (
    "encoding/json"
    "errors"
)

// MarketLocalePolicy How this tenant keys its translations, resolved rather
// than named: the key a client WRITES and the order it READS, per locale.
// Emitting the resolved answer is the point — a client handed only the
// setting names re-implements the policy and gets it subtly different, which
// is how a label editor came to ask for de-DE while the row held de.
type MarketLocalePolicy struct {
    // settings#locale_fallback — what a read tries after the exact key holds
    // nothing.
    Fallback string `json:"fallback"`
    // settings#locale_granularity — whether a value is keyed by the full locale
    // ('regional') or by its language alone.
    Granularity string `json:"granularity"`
    // One entry per locale this market registers, in position order — the keys
    // to use for that locale. A market with no locale of its own has an empty
    // array here, not the fallback: the fallback answers `default_locale`, and
    // there is nothing to key against.
    Locales []MarketLocaleKeys `json:"locales"`

    // Used by Decode() method
    data []byte
}

func (model MarketLocalePolicy) New(data []byte) *MarketLocalePolicy {
    model.data = data
    return &model
}

func (model *MarketLocalePolicy) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}