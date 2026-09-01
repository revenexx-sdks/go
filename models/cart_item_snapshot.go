package models

import (
    "encoding/json"
    "errors"
)

// CartItemSnapshot The product as the buyer was shown it when this line was
// added — the cart's own copy, so it stays honest when the catalogue moves
// underneath it. Free-form apart from the price: conversion reads
// `unit_price` (or `price` as a fallback) and nothing else. A snapshot
// without a readable price leaves the line alone in both price modes, which
// is deliberate — a missing snapshot must never be read as "free".
type CartItemSnapshot struct {
    // The older spelling of the same thing, read only when `unit_price` is
    // absent.
    Price float64 `json:"price"`
    // The net unit price the buyer was shown. This is what carts.order books the
    // line on under price_snapshot_mode = snapshot, and what it rewrites under =
    // live.
    UnitPrice float64 `json:"unit_price"`

    // Used by Decode() method
    data []byte
}

func (model CartItemSnapshot) New(data []byte) *CartItemSnapshot {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *CartItemSnapshot) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}