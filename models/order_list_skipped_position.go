package models

import (
    "encoding/json"
    "errors"
)

// OrderListSkippedPosition A position left out of the conversion because the
// catalogue no longer knows its article (only ever non-empty when the
// tenant's 'on_missing_article' setting is 'skip').
type OrderListSkippedPosition struct {
    // The position that was left out, so a client can point at the row in the
    // list.
    Id string `json:"id"`
    // The saved article name, so the omission can be reported to the buyer in
    // words they recognise.
    Name string `json:"name"`
    // The catalogue product the position named, if it named one.
    ProductId string `json:"product_id"`
    // The article number the position named, if it named one.
    Sku string `json:"sku"`

    // Used by Decode() method
    data []byte
}

func (model OrderListSkippedPosition) New(data []byte) *OrderListSkippedPosition {
    model.data = data
    return &model
}

func (model *OrderListSkippedPosition) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}