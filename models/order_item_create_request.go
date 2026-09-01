package models

import (
    "encoding/json"
    "errors"
)

// OrderItemCreateRequest A position of the placed order — needs an
// identity: 'name' or 'sku'. Items are SNAPSHOTS: carry the product copy,
// prices are frozen at place-time.
type OrderItemCreateRequest struct {
    // The chosen options of a configured line — what the configurator produced,
    // in whatever shape it produces. Only meaningful for type 'configuration';
    // null everywhere else.
    Configuration interface{} `json:"configuration"`
    // The buyer's own cost centre for this line — a B2B field: the same order
    // is split across several of them and the buyer's finance department needs
    // the split per line, not per order.
    CostCenter string `json:"cost_center"`
    // Free-form data belonging to the integration side, per position. Stored and
    // returned untouched.
    Metadata interface{} `json:"metadata"`
    // The article name as it stood at place-time, frozen. Falls back to the sku
    // when the caller sent none — a position always reads as something. Falls
    // back to 'sku' when omitted; one of the two is required.
    Name string `json:"name"`
    // The line number a human reads, and what the order is sorted by. Numbered in
    // steps of the range's position_step (10, 20, 30) unless the caller set it
    // explicitly — the gap is what lets a line be inserted later without
    // renumbering. Omitted = numbered in steps of the order range's
    // position_step.
    Position int `json:"position"`
    // A free note the buyer attached to this line — an engraving, a delivery
    // instruction, the drawing number the line refers to. Printed on the
    // paperwork, read by nothing.
    PositionText string `json:"position_text"`
    // The product as it was at place-time, FROZEN: the copy that makes the order
    // still correct after the catalog changes its price, its name or its
    // attributes. The caller decides how much of the product to freeze; this app
    // stores it and reads nothing out of it. 'snapshot' is accepted as an alias
    // for this key.
    Product interface{} `json:"product"`
    // The catalog product this line was taken from (the products app). Null on a
    // custom line, and it stays a reference — the position keeps working after
    // the product is retired.
    ProductId string `json:"product_id"`
    // How much was ORDERED, in `unit`. Three decimal places, so 2.5 m of cable is
    // a real order line. Never changed afterwards — cancelling or returning
    // writes the quantity_* columns instead, which is what keeps the order a
    // truthful record of what was asked for. Defaults to 1.
    Quantity float64 `json:"quantity"`
    // The article number as it stood at place-time, frozen with the rest of the
    // line. The value an ERP and a warehouse both join on, and the one field a
    // picker reads. Null only on a line that never had one.
    Sku string `json:"sku"`
    // The product as it was at place-time, FROZEN: the copy that makes the order
    // still correct after the catalog changes its price, its name or its
    // attributes. The caller decides how much of the product to freeze; this app
    // stores it and reads nothing out of it. Alias for 'product' — send one or
    // the other, not both.
    Snapshot interface{} `json:"snapshot"`
    // Tax on this line in `currency`. Derived from line_total × tax_rate/100
    // when the caller sent none, which is the normal case — but a caller may
    // send it, for a market whose rounding rules differ from ours. Send it only
    // where your market rounds differently from line_total × tax_rate/100.
    TaxAmount float64 `json:"tax_amount"`
    // Tax percentage for this line, as a number (19 means 19 %). Frozen at
    // place-time with everything else. Defaults to 0.
    TaxRate float64 `json:"tax_rate"`
    // What kind of line this is: 'product' is a catalog article, 'configuration'
    // a configured one carrying its configuration, 'custom' a line typed by hand
    // that no catalog knows. Defaults to 'product'.
    Type string `json:"type"`
    // The unit the quantity is counted in — piece, metre, kilogram, package.
    // Free text as the catalog carries it; this app does no conversion.
    Unit string `json:"unit"`
    // NET price per unit, FROZEN at place-time. A later price change in the
    // catalog does not reach this order. Defaults to 0. line_total is always
    // derived from it and never taken from the body.
    UnitPrice float64 `json:"unit_price"`
    // Free-form data belonging to the ordering side, per position — carried
    // through from the cart line and handed back untouched.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderItemCreateRequest) New(data []byte) *OrderItemCreateRequest {
    model.data = data
    return &model
}

func (model *OrderItemCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}