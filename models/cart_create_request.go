package models

import (
    "encoding/json"
    "errors"
)

// CartCreateRequest A cart needs an owner: 'contact_id' (customer) or
// 'session_key' (guest).
type CartCreateRequest struct {
    // The sales channel this cart is being opened in, as a channel of the
    // channels app. Stored for attribution; nothing in this app reads it.
    ChannelId string `json:"channel_id"`
    // The customer who owns this cart, as a contact of the customers app. Send
    // this OR session_key — a cart with neither owner is refused.
    ContactId string `json:"contact_id"`
    // ISO 4217 code the cart is priced in (default EUR). Lines added without a
    // currency inherit it.
    Currency string `json:"currency"`
    // Make this THE current cart of its owner as it is created — the same thing
    // carts.activate does later, and it clears the flag on every sibling cart of
    // the same owner.
    IsCurrent bool `json:"is_current"`
    // Free-form data the storefront hangs on the cart. Stored and returned
    // verbatim; no key in here is read by this app, and none is indexed.
    Metadata interface{} `json:"metadata"`
    // What the buyer calls this cart (default 'Cart'). An empty string is legal
    // and lands on the default.
    Name string `json:"name"`
    // The guest session that owns this cart — the key the storefront already
    // keeps in its own session or cookie. Any non-empty string is accepted; this
    // app issues none and parses none, so the example shows a shape and not a
    // format. Send this OR contact_id.
    SessionKey string `json:"session_key"`

    // Used by Decode() method
    data []byte
}

func (model CartCreateRequest) New(data []byte) *CartCreateRequest {
    model.data = data
    return &model
}

func (model *CartCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}