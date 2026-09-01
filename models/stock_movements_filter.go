package models

import (
    "encoding/json"
    "errors"
)

// StockMovementsFilter The exact-column filters this call was understood to
// carry, verbatim as they arrived. A query parameter that is not a column of
// `stock_movements` — a typo, a filter another entity has, `?q=` — is
// DROPPED and cannot appear here, and the list comes back unfiltered. This
// object is the only way to tell that apart from "nothing matched".
type StockMovementsFilter struct {
    // The literal `?created_at=` value this call was understood to carry.
    CreatedAt string `json:"created_at"`
    // The literal `?id=` value this call was understood to carry.
    Id string `json:"id"`
    // The literal `?location_id=` value this call was understood to carry.
    LocationId string `json:"location_id"`
    // The literal `?metadata=` value this call was understood to carry.
    Metadata string `json:"metadata"`
    // The literal `?order_ref=` value this call was understood to carry.
    OrderRef string `json:"order_ref"`
    // The literal `?product_id=` value this call was understood to carry.
    ProductId string `json:"product_id"`
    // The literal `?quantity=` value this call was understood to carry.
    Quantity string `json:"quantity"`
    // The literal `?reason=` value this call was understood to carry.
    Reason string `json:"reason"`
    // The literal `?sku=` value this call was understood to carry.
    Sku string `json:"sku"`
    // The literal `?type=` value this call was understood to carry.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model StockMovementsFilter) New(data []byte) *StockMovementsFilter {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *StockMovementsFilter) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}