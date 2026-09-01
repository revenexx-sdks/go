package models

import (
    "encoding/json"
    "errors"
)

// InventoryReceiveRequest model.
type InventoryReceiveRequest struct {
    // The goods that arrived, at most 200 in one call — a delivery, a
    // production batch, an opening balance.
    Items []InventoryStockItem `json:"items"`
    // Which location took the delivery. Omitted, the `default_location_code`
    // setting decides; a code no location carries is answered 400 rather than
    // booked somewhere else.
    LocationCode string `json:"location_code"`
    // Inline single-item form: the product to move, instead of a one-entry
    // `items` array. The two forms are equivalent — nothing downstream knows
    // which arrived.
    ProductId string `json:"product_id"`
    // Inline single-item form: how many arrived. Positive.
    Quantity float64 `json:"quantity"`
    // What the ledger should record about this receipt — a delivery note
    // number, a production order. Owed only when `movement_reason_required` is
    // 'all'; the contract does not require it, because whether it is owed is the
    // tenant's setting and not this route's rule.
    Reason string `json:"reason"`
    // Inline single-item form: the article number to move (instead of
    // `product_id`).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReceiveRequest) New(data []byte) *InventoryReceiveRequest {
    model.data = data
    return &model
}

func (model *InventoryReceiveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}