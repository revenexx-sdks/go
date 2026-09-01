package models

import (
    "encoding/json"
    "errors"
)

// ItemAvailability model.
type ItemAvailability struct {
    // on_hand − reserved across the locations in scope: available-to-promise,
    // and the number a storefront shows. It can be NEGATIVE once backorders have
    // been reserved beyond stock — nothing floors it, because "sold more than
    // we hold" is a real state a merchant needs to see.
    Available float64 `json:"available"`
    // The per-location breakdown behind the summed figures — which place could
    // actually ship it.
    Locations []LocationAvailability `json:"locations"`
    // Physically in stock, summed across the locations in scope (every enabled
    // location, or the one `location_code` named). Promised units are included,
    // so this is NOT what may be sold.
    OnHand float64 `json:"on_hand"`
    // True when the item is tracked and `available >= requested` at this moment.
    // A SNAPSHOT, not a hold: nothing is set aside until POST
    // /inventories/reserve, and two checkouts can both read true for the last
    // unit.
    Orderable bool `json:"orderable"`
    // The product id as it was asked for, echoed. Null when the item was named by
    // SKU.
    ProductId string `json:"product_id"`
    // The quantity the check was made against — the item's own `quantity`, or 1
    // when none was sent. `orderable` answers "can I have this many?", so it is
    // only as strict as this number.
    Requested float64 `json:"requested"`
    // Already promised to orders, summed across the same locations — the part
    // of `on_hand` that is spoken for.
    Reserved float64 `json:"reserved"`
    // The SKU as it was asked for, echoed. Null when the item was named by
    // product id.
    Sku string `json:"sku"`
    // False when this app has never seen the item: no stock row anywhere in
    // scope. It is not an error and not a zero — the storefront decides whether
    // an untracked item sells freely (a service, a made-to-order piece) or not at
    // all. `on_hand`, `reserved` and `available` are 0 in that case, and
    // `orderable` is false.
    Tracked bool `json:"tracked"`

    // Used by Decode() method
    data []byte
}

func (model ItemAvailability) New(data []byte) *ItemAvailability {
    model.data = data
    return &model
}

func (model *ItemAvailability) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}