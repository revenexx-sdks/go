package models

import (
    "encoding/json"
    "errors"
)

// InventoryRestockRequest model.
type InventoryRestockRequest struct {
    // The goods that came back, at most 200 in one call. Whether they rejoin
    // sellable stock is `restock`, not this list.
    Items []InventoryStockItem `json:"items"`
    // Where the goods came back to — a returns warehouse is a location like any
    // other. Omitted, the `default_location_code` setting decides.
    LocationCode string `json:"location_code"`
    // The order the goods came back from. It is written onto the ledger booking,
    // so the return shows up in that order's stock history next to its reserve
    // and shipment — no reservation is touched by it.
    OrderRef string `json:"order_ref"`
    // Inline single-item form: the product to move, instead of a one-entry
    // `items` array. The two forms are equivalent — nothing downstream knows
    // which arrived.
    ProductId string `json:"product_id"`
    // Inline single-item form: how many came back. Positive.
    Quantity float64 `json:"quantity"`
    // Why the goods came back — 'wrong size', 'damaged on arrival'. Owed only
    // when `movement_reason_required` is 'all'.
    Reason string `json:"reason"`
    // Do these goods rejoin SELLABLE stock? A merchant decision, not a fact:
    // apparel usually restocks, hygiene articles never do, many merchants inspect
    // first. Omit it to follow the `restock_on_return_default` setting. `false`
    // answers `restocked: false`, moves nothing and books NOTHING — there is no
    // movement to write, because no stock moved, and that is the branch that
    // makes this route a 200 while its sibling `receive` is a 201.
    Restock bool `json:"restock"`
    // Inline single-item form: the article number to move (instead of
    // `product_id`).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryRestockRequest) New(data []byte) *InventoryRestockRequest {
    model.data = data
    return &model
}

func (model *InventoryRestockRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}