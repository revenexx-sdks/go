package models

import (
    "encoding/json"
    "errors"
)

// TenantLocalePolicy How this tenant keys its translations, resolved rather
// than named: the key a client WRITES and the order it READS, per locale.
// Emitting the resolved answer is the point — a client handed only the
// setting names re-implements the policy and gets it subtly different, which
// is how a label editor came to ask for de-DE while the row held de.
type TenantLocalePolicy struct {
    // settings#locale_fallback — what a read tries after the exact key holds
    // nothing.
    Fallback string `json:"fallback"`
    // settings#locale_granularity — whether a value is keyed by the full locale
    // ('regional') or by its language alone.
    Granularity string `json:"granularity"`
    // The UNION of every market's locales, each one appearing once — the full
    // set of inputs a tenant-baseline editor has to offer. Empty when no market
    // registers a locale at all.
    Locales []TenantLocaleKeys `json:"locales"`

    // Used by Decode() method
    data []byte
}

func (model TenantLocalePolicy) New(data []byte) *TenantLocalePolicy {
    model.data = data
    return &model
}

func (model *TenantLocalePolicy) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}