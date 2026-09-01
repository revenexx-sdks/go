package models

import (
    "encoding/json"
    "errors"
)

// CartItemCreateRequest An item needs an identity: 'name' or 'sku'.
type CartItemCreateRequest struct {
    // What was configured on this line, in the configurator's own vocabulary —
    // this app stores it and reads nothing out of it. Its mere PRESENCE is
    // behaviour: a line that carries a configuration never merges with another,
    // because two differently configured units of the same article are not one
    // line. Keys are the configurator's; the example is one shape, not the shape.
    Configuration interface{} `json:"configuration"`
    // ISO 4217 code. Defaults to the cart's currency.
    Currency string `json:"currency"`
    // Free-form data the storefront hangs on the line. Stored and returned
    // verbatim; no key in here is read by this app.
    Metadata interface{} `json:"metadata"`
    // What the line reads as on the cart page. Falls back to 'sku' when omitted,
    // so a line always has something to show.
    Name string `json:"name"`
    // Sort order within the cart, ascending. Default 0 when adding a line; in a
    // bulk replace the payload order fills it in.
    Position int `json:"position"`
    // The catalogue product, when the line comes from one. Part of the merge
    // identity: same product, same price, one line.
    ProductId string `json:"product_id"`
    // How much of it — default 1. Fractional is legal (2.5 m of cable); zero
    // and negative are not. On a plain product line that merges into an existing
    // one, this is ADDED to what is already there, and max_quantity_per_line is
    // checked on the result.
    Quantity float64 `json:"quantity"`
    // The article number, exactly as the merchant knows it. Free text — this
    // app does not resolve it against the catalogue — and part of the merge
    // identity together with product_id and unit_price. The example only shows
    // the shape of a real article number; nothing here enforces one.
    Sku string `json:"sku"`
    // The product as the buyer was shown it when this line was added — the
    // cart's own copy, so it stays honest when the catalogue moves underneath it.
    // Free-form apart from the price: conversion reads `unit_price` (or `price`
    // as a fallback) and nothing else. A snapshot without a readable price leaves
    // the line alone in both price modes, which is deliberate — a missing
    // snapshot must never be read as "free".
    Snapshot CartItemSnapshot `json:"snapshot"`
    // VAT percent for this line, as a number (19 means 19 %). Stored for the
    // order to use — no total in this app includes tax.
    TaxRate float64 `json:"tax_rate"`
    // Line type (default 'product'). Plain product lines merge by product+price;
    // configurations always stand alone.
    Type string `json:"type"`
    // The unit the quantity is counted in. Display and ERP hand-over only —
    // this app converts nothing.
    Unit string `json:"unit"`
    // Net price of one unit — line_total is always derived from it, never sent.
    // Part of the merge identity: the same article at a different price opens a
    // new line rather than averaging into the old one.
    UnitPrice float64 `json:"unit_price"`

    // Used by Decode() method
    data []byte
}

func (model CartItemCreateRequest) New(data []byte) *CartItemCreateRequest {
    model.data = data
    return &model
}

func (model *CartItemCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}