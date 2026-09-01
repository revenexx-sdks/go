package models

import (
    "encoding/json"
    "errors"
)

// TenantLocaleKeys One locale somewhere in this tenant, its read and write
// keys, and the markets that asked for it.
type TenantLocaleKeys struct {
    // The locale this entry is about, as some market registered it.
    Code string `json:"code"`
    // Its language part, which is also the key under language granularity.
    Language string `json:"language"`
    // Codes of the markets that registered this locale, sorted — who a baseline
    // translation written here is actually for. An editor that lists six inputs
    // without saying who needs them invites translations nobody will ever read.
    Markets []string `json:"markets"`
    // Keys to try in order until one holds text — the same resolved order the
    // per-market answer gives, so a baseline value and a market value can never
    // be keyed differently.
    Read []string `json:"read"`
    // A key inside a labels bag: a full locale ('de-DE') under regional
    // granularity, a bare language ('de') under language granularity.
    Write string `json:"write"`

    // Used by Decode() method
    data []byte
}

func (model TenantLocaleKeys) New(data []byte) *TenantLocaleKeys {
    model.data = data
    return &model
}

func (model *TenantLocaleKeys) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}