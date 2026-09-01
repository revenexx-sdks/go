package models

import (
    "encoding/json"
    "errors"
)

// IoProfileUpdateRequest Partial update — omitted fields keep their current
// value.
type IoProfileUpdateRequest struct {
    // What an import does with the lines the target cart already has: 'replace'
    // clears them first, 'insert' and 'append' both add and behave identically
    // today. Read only when the import names a target_cart_id. Default 'insert'.
    ApplyMode string `json:"apply_mode"`
    // Which way this profile runs. A profile only ever runs in the direction it
    // declares: handing an import profile to carts.export is a 400, and the other
    // way round.
    Direction string `json:"direction"`
    // What the profile carries: whole carts (the `{cart, items}` document) or
    // bare cart lines. Default 'carts'.
    Entity string `json:"entity"`
    // The wire format. 'json' is the canonical, re-importable document; 'csv' is
    // the spreadsheet form, and only line fields survive it. Default 'json'.
    Format string `json:"format"`
    // One of the bundled templates. Set by carts.io.profiles.defaults; a profile
    // a merchant writes is not one.
    IsTemplate bool `json:"is_template"`
    // Baseline-IO-compatible column mapping. An empty object (or null) is
    // identity: the full canonical shape, every field under its own name.
    Mapping CartIoMapping `json:"mapping"`
    // What a merchant picks this profile by. Unique within the tenant — reusing
    // a name is a 409.
    Name string `json:"name"`
    // Free-form options carried with the profile. The four bundled templates put
    // one human sentence under `description` and nothing else; no other key is
    // read by this app, so anything a merchant needs alongside a profile can live
    // here.
    Options interface{} `json:"options"`

    // Used by Decode() method
    data []byte
}

func (model IoProfileUpdateRequest) New(data []byte) *IoProfileUpdateRequest {
    model.data = data
    return &model
}

func (model *IoProfileUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}