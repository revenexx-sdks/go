package models

import (
    "encoding/json"
    "errors"
)

// CartUpdateRequest Only safe columns are updatable — status moves through
// the lifecycle routes.
type CartUpdateRequest struct {
    // Move the cart to another sales channel.
    ChannelId string `json:"channel_id"`
    // ISO 4217 code. Changes what NEW lines inherit; lines already in the cart
    // keep the currency they were added with.
    Currency string `json:"currency"`
    // Free-form data the storefront hangs on the cart. Stored and returned
    // verbatim; no key in here is read by this app, and none is indexed.
    Metadata interface{} `json:"metadata"`
    // Rename the cart. Unlike on create, this is written verbatim — `null` and
    // `''` are refused by the database.
    Name string `json:"name"`

    // Used by Decode() method
    data []byte
}

func (model CartUpdateRequest) New(data []byte) *CartUpdateRequest {
    model.data = data
    return &model
}

func (model *CartUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}