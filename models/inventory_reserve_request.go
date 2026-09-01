package models

import (
    "encoding/json"
    "errors"
)

// InventoryReserveRequest model.
type InventoryReserveRequest struct {
    // When this hold lapses. The sweeper — POST
    // /inventories/reservations/sweep, and the 'expire-reservations' schedule
    // that runs it every 15 minutes — releases everything past this moment
    // exactly as a cancellation would, so an abandoned checkout stops holding
    // stock on its own. Null means the row named no deadline: it is swept on its
    // AGE instead once `reservation_ttl_minutes` is above 0, which is what makes
    // turning that setting on retroactive. Omit it to let the
    // `reservation_ttl_minutes` setting stamp one (0 — its default — means no
    // deadline at all); send one to hold this order for a window of its own, e.g.
    // a quote that stands until Friday.
    ExpiresAt string `json:"expires_at"`
    // The items to hold, at most 200 in one call — a whole cart in one request.
    // The call is planned before anything is written, so either every item is
    // placed or nothing is.
    Items []InventoryStockItem `json:"items"`
    // Where a BACKORDERED item is booked when no location holds a stock row for
    // it at all — the last fallback, not the allocator: which location serves
    // an item that IS in stock comes from `allocation_strategy`. Omitted, the
    // `default_location_code` setting decides.
    LocationCode string `json:"location_code"`
    // The order this hold belongs to. The caller supplies it — this app mints
    // nothing — and it is the handle POST /inventories/release and POST
    // /inventories/commit act on, so it has to be the same string the order
    // carries elsewhere. At least one character (CHECK `length(order_ref) > 0`).
    // Not unique: an order holds one reservation per item, and they are released
    // or committed together. Reserving twice under the same reference ADDS holds
    // rather than replacing them — release first if you mean to replace.
    OrderRef string `json:"order_ref"`
    // Inline single-item form: the product to move, instead of a one-entry
    // `items` array. The two forms are equivalent — nothing downstream knows
    // which arrived.
    ProductId string `json:"product_id"`
    // Inline single-item form: how many to hold. Positive — the hold is
    // expressed as a positive reservation, while the ledger booking it writes
    // carries the negative.
    Quantity float64 `json:"quantity"`
    // Where the order is going. Read ONLY when the tenant's `allocation_strategy`
    // is 'nearest' — under 'priority' or 'single_location' it is accepted and
    // ignored, so sending it is never wrong, it is just not always heard.
    ShipTo InventoryShipTo `json:"ship_to"`
    // Inline single-item form: the article number to move (instead of
    // `product_id`).
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model InventoryReserveRequest) New(data []byte) *InventoryReserveRequest {
    model.data = data
    return &model
}

func (model *InventoryReserveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}