package models

import (
    "encoding/json"
    "errors"
)

// OrderListToCartResult model.
type OrderListToCartResult struct {
    // Positions written to the cart. Equal to the list's position count minus
    // `skipped`.
    Added int `json:"added"`
    // True when this call created the cart. A created cart is the owner's CURRENT
    // cart, because a cart the buyer cannot see is not "added to cart".
    CartCreated bool `json:"cart_created"`
    // The cart the positions landed in: the one that was passed in, or the one
    // this call created.
    CartId string `json:"cart_id"`
    // The list that was converted. Unchanged by the call — a conversion reads
    // the list, it never empties it.
    ListId string `json:"list_id"`
    // The mode that was actually applied — the one that was asked for, or the
    // tenant's 'cart_merge_mode' default when the call named none.
    Mode string `json:"mode"`
    // Positions left out because the catalogue no longer knows their article.
    // Only ever non-empty when 'on_missing_article' is 'skip' — 'include'
    // converts them anyway and 'fail' answers 400 instead.
    Skipped []OrderListSkippedPosition `json:"skipped"`

    // Used by Decode() method
    data []byte
}

func (model OrderListToCartResult) New(data []byte) *OrderListToCartResult {
    model.data = data
    return &model
}

func (model *OrderListToCartResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}