package models

import (
    "encoding/json"
    "errors"
)

// OrderListToCartRequest Every field is optional: with an empty body the list
// goes into a NEW cart for its owner, on the tenant defaults.
type OrderListToCartRequest struct {
    // Add to this existing cart. Omit to create one for the list owner and make
    // it their current cart.
    CartId string `json:"cart_id"`
    // ISO 4217 code for the cart and its lines. Omit to let the carts app decide.
    Currency string `json:"currency"`
    // 'append' adds the positions (the carts app merges a line by product and
    // price, so quantities accumulate); 'replace' makes the list the cart's
    // entire contents. Defaults to the tenant's 'cart_merge_mode' setting.
    Mode string `json:"mode"`

    // Used by Decode() method
    data []byte
}

func (model OrderListToCartRequest) New(data []byte) *OrderListToCartRequest {
    model.data = data
    return &model
}

func (model *OrderListToCartRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}