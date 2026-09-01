package models

import (
    "encoding/json"
    "errors"
)

// CartImportRequest Import into an existing cart ('target_cart_id') or a new
// cart (owner 'contact_id'/'session_key' required).
type CartImportRequest struct {
    // Owner of the cart this import creates. Ignored when target_cart_id is sent.
    ContactId string `json:"contact_id"`
    // The CSV rows, when that is easier than putting them in `payload`. First
    // line is the header, and its names are the ones the profile's mapping
    // expects (the bundled quick-order template reads sku, name, quantity,
    // unit_price). Numbers are coerced; a JSON column survives as a JSON string.
    Csv string `json:"csv"`
    // Name for the cart this import creates. A name in the payload's own `cart`
    // block wins over it; without either the cart is called 'Imported cart'.
    Name string `json:"name"`
    // The import itself. As an object: `{ "cart": { name, status, currency,
    // channel_id, metadata }, "items": [ … ] }` — the same document
    // carts.export produces, so an export round-trips. As a string: that document
    // as raw JSON, or CSV rows when the profile is a csv one. A line with neither
    // `name` nor `sku` is dropped, and a payload that leaves no line at all is a
    // 400.
    Payload interface{} `json:"payload"`
    // The import profile to run — one of the ids `GET
    // /carts/io/profiles?direction=import` lists. Omit it for an ad-hoc import:
    // the payload is then read in the canonical shape, and as CSV if `csv` is
    // what carried it.
    ProfileId string `json:"profile_id"`
    // Guest owner of the cart this import creates — the storefront's own
    // session key. Ignored when target_cart_id is sent.
    SessionKey string `json:"session_key"`
    // An existing ACTIVE cart to import into. The lines are added to it (merging
    // identical product lines), unless the profile says `apply_mode: replace`,
    // which clears it first. Without this a new cart is created and an owner is
    // required.
    TargetCartId string `json:"target_cart_id"`

    // Used by Decode() method
    data []byte
}

func (model CartImportRequest) New(data []byte) *CartImportRequest {
    model.data = data
    return &model
}

func (model *CartImportRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}