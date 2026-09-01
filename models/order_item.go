package models

import (
    "encoding/json"
    "errors"
)

// OrderItem One POSITION of an order, frozen at place-time: the article as it
// was, the price as it was, and three running quantities (shipped, cancelled,
// returned) that everything after placement books against. `quantity` itself
// never changes.
type OrderItem struct {
    // The chosen options of a configured line — what the configurator produced,
    // in whatever shape it produces. Only meaningful for type 'configuration';
    // null everywhere else.
    Configuration interface{} `json:"configuration"`
    // The buyer's own cost centre for this line — a B2B field: the same order
    // is split across several of them and the buyer's finance department needs
    // the split per line, not per order.
    CostCenter string `json:"cost_center"`
    // When the position was written — the moment the order was placed.
    CreatedAt string `json:"created_at"`
    // Primary key of the position. This is the id every positions[] payload
    // names: /ship, /items/cancel and /return all take order_item_id.
    Id string `json:"id"`
    // quantity × unit_price, NET, always COMPUTED here — a caller cannot set
    // it. The order's subtotal is the sum of these.
    LineTotal float64 `json:"line_total"`
    // Free-form data belonging to the integration side, per position. Stored and
    // returned untouched.
    Metadata interface{} `json:"metadata"`
    // The article name as it stood at place-time, frozen. Falls back to the sku
    // when the caller sent none — a position always reads as something.
    Name string `json:"name"`
    // The order this position belongs to. Deleting the order deletes its
    // positions.
    OrderId string `json:"order_id"`
    // The line number a human reads, and what the order is sorted by. Numbered in
    // steps of the range's position_step (10, 20, 30) unless the caller set it
    // explicitly — the gap is what lets a line be inserted later without
    // renumbering.
    Position int `json:"position"`
    // A free note the buyer attached to this line — an engraving, a delivery
    // instruction, the drawing number the line refers to. Printed on the
    // paperwork, read by nothing.
    PositionText string `json:"position_text"`
    // The product as it was at place-time, FROZEN: the copy that makes the order
    // still correct after the catalog changes its price, its name or its
    // attributes. The caller decides how much of the product to freeze; this app
    // stores it and reads nothing out of it.
    Product interface{} `json:"product"`
    // The catalog product this line was taken from (the products app). Null on a
    // custom line, and it stays a reference — the position keeps working after
    // the product is retired.
    ProductId string `json:"product_id"`
    // How much was ORDERED, in `unit`. Three decimal places, so 2.5 m of cable is
    // a real order line. Never changed afterwards — cancelling or returning
    // writes the quantity_* columns instead, which is what keeps the order a
    // truthful record of what was asked for.
    Quantity float64 `json:"quantity"`
    // How much of this position was cancelled and will never ship. Written by
    // /cancel (all of it) and /items/cancel (a named quantity). Cancelling
    // reduces the effective quantity, so an order whose every position is fully
    // cancelled becomes cancelled itself.
    QuantityCancelled float64 `json:"quantity_cancelled"`
    // How much of this position came BACK, booked when a return is completed —
    // not when it is registered or received. This is the goods accounting: it
    // never reduces quantity_shipped, so a position can be shipped 3 and returned
    // 3.
    QuantityReturned float64 `json:"quantity_returned"`
    // How much of this position has GONE OUT, summed over the shipments. Written
    // only by POST /orders/{id}/ship; it is what fulfillment_status is derived
    // from, and what a return is guarded against.
    QuantityShipped float64 `json:"quantity_shipped"`
    // The article number as it stood at place-time, frozen with the rest of the
    // line. The value an ERP and a warehouse both join on, and the one field a
    // picker reads. Null only on a line that never had one.
    Sku string `json:"sku"`
    // Tax on this line in `currency`. Derived from line_total × tax_rate/100
    // when the caller sent none, which is the normal case — but a caller may
    // send it, for a market whose rounding rules differ from ours.
    TaxAmount float64 `json:"tax_amount"`
    // Tax percentage for this line, as a number (19 means 19 %). Frozen at
    // place-time with everything else.
    TaxRate float64 `json:"tax_rate"`
    // What kind of line this is: 'product' is a catalog article, 'configuration'
    // a configured one carrying its configuration, 'custom' a line typed by hand
    // that no catalog knows.
    Type string `json:"type"`
    // The unit the quantity is counted in — piece, metre, kilogram, package.
    // Free text as the catalog carries it; this app does no conversion.
    Unit string `json:"unit"`
    // NET price per unit, FROZEN at place-time. A later price change in the
    // catalog does not reach this order.
    UnitPrice float64 `json:"unit_price"`
    // When the position last changed, which in practice means the last time a
    // quantity was booked onto it.
    UpdatedAt string `json:"updated_at"`
    // Free-form data belonging to the ordering side, per position — carried
    // through from the cart line and handed back untouched.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderItem) New(data []byte) *OrderItem {
    model.data = data
    return &model
}

func (model *OrderItem) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}