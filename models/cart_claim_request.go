package models

import (
    "encoding/json"
    "errors"
)

// Model
type CartClaimRequest struct {
    // Contact taking ownership.
    ContactId string `json:"contact_id"`
    // Guest session whose active carts are handed over.
    SessionKey string `json:"session_key"`
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