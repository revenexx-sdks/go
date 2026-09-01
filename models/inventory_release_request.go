package models

import (
    "encoding/json"
    "errors"
)

// InventoryReleaseRequest model.
type InventoryReleaseRequest struct {
    // The order this hold belongs to. The caller supplies it — this app mints
    // nothing — and it is the handle POST /inventories/release and POST
    // /inventories/commit act on, so it has to be the same string the order
    // carries elsewhere. At least one character (CHECK `length(order_ref) > 0`).
    // Not unique: an order holds one reservation per item, and they are released
    // or committed together. Every ACTIVE hold under this reference is given
    // back; ones already committed or released are left alone. A reference no
    // reservation carries releases nothing and answers `released: 0` — not an
    // error, which is what makes a retried cancellation safe.
    OrderRef string `json:"order_ref"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReleaseRequest) New(data []byte) *InventoryReleaseRequest {
    model.data = data
    return &model
}

func (model *InventoryReleaseRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}