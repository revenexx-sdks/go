package models

import (
    "encoding/json"
    "errors"
)

// IoProfile model.
type IoProfile struct {
    // What an import does with the lines the target cart already has. 'replace'
    // clears them first; 'insert' and 'append' both add, and behave identically
    // today. Read only by carts.import, and only when the call names a
    // target_cart_id — an import that creates its own cart has nothing to apply
    // a mode to.
    ApplyMode string `json:"apply_mode"`
    // When the profile was created — for the bundled templates, when the app
    // was installed.
    CreatedAt string `json:"created_at"`
    // Which way this profile runs. A profile only ever runs in the direction it
    // declares: handing an import profile to carts.export is a 400, and the other
    // way round.
    Direction string `json:"direction"`
    // What the profile carries: whole carts ('carts' — the `{cart, items}`
    // document) or bare cart lines ('cart_items' — the spreadsheet a buyer
    // quick-orders from).
    Entity string `json:"entity"`
    // The wire format. 'json' is the canonical, re-importable document; 'csv' is
    // the spreadsheet form, and only line fields survive it.
    Format string `json:"format"`
    // The profile, as carts.export and carts.import name it in `profile_id`.
    Id string `json:"id"`
    // One of the profiles this app ships with, seeded by
    // carts.io.profiles.defaults. A profile a merchant wrote is not one, so this
    // is how a UI separates "what came with the app" from "what we built".
    IsTemplate bool `json:"is_template"`
    // Baseline-IO-compatible column mapping. An empty object (or null) is
    // identity: the full canonical shape, every field under its own name.
    Mapping CartIoMapping `json:"mapping"`
    // What a merchant picks this profile by. Unique within the tenant — reusing
    // a name is a 409 — and the four bundled templates use it as their
    // identity, so seeding is idempotent by name.
    Name string `json:"name"`
    // Free-form options carried with the profile. The four bundled templates put
    // one human sentence under `description` and nothing else; no other key is
    // read by this app, so anything a merchant needs alongside a profile can live
    // here.
    Options interface{} `json:"options"`
    // The tenant this row belongs to, echoed by the data plane.
    TenantId string `json:"tenant_id"`
    // When the profile last changed.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model IoProfile) New(data []byte) *IoProfile {
    model.data = data
    return &model
}

func (model *IoProfile) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}