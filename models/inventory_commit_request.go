package models

import (
    "encoding/json"
    "errors"
)

// InventoryCommitRequest model.
type InventoryCommitRequest struct {
    // The order this hold belongs to. The caller supplies it — this app mints
    // nothing — and it is the handle POST /inventories/release and POST
    // /inventories/commit act on, so it has to be the same string the order
    // carries elsewhere. At least one character (CHECK `length(order_ref) > 0`).
    // Not unique: an order holds one reservation per item, and they are released
    // or committed together. Every ACTIVE hold under this reference ships:
    // `on_hand` and `reserved` both fall and a `shipment` booking is written for
    // each. Unlike release, committing an order that has nothing active is a 422
    // — it means the hold was already released or already shipped, and shipping
    // twice is worth saying out loud.
    OrderRef string `json:"order_ref"`

    // Used by Decode() method
    data []byte
}

func (model InventoryCommitRequest) New(data []byte) *InventoryCommitRequest {
    model.data = data
    return &model
}

func (model *InventoryCommitRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}