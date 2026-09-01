package models

import (
    "encoding/json"
    "errors"
)

// CartClaimRequest model.
type CartClaimRequest struct {
    // The contact taking ownership. Every active cart of that session ends up
    // with this contact — adopted as it stands, or folded into
    // `target_cart_id`.
    ContactId string `json:"contact_id"`
    // The guest session whose active carts are handed over — the key the
    // storefront keeps in its own session or cookie and has been sending on every
    // anonymous call. This app neither issues nor parses it, so the example shows
    // the shape of an opaque token and not a format anything enforces.
    SessionKey string `json:"session_key"`
    // Override the tenant's cart_merge_strategy for this call: 'merge' keeps the
    // target cart's own lines, 'replace' clears them first. Omit to use the
    // setting.
    Strategy string `json:"strategy"`
    // Merge the session carts into this cart instead of adopting them.
    TargetCartId string `json:"target_cart_id"`

    // Used by Decode() method
    data []byte
}

func (model CartClaimRequest) New(data []byte) *CartClaimRequest {
    model.data = data
    return &model
}

func (model *CartClaimRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}