package models

import (
    "encoding/json"
    "errors"
)

// OrderComment A note on an order, either internal between operators or meant
// for the customer to see.
type OrderComment struct {
    // Who wrote it, as the caller reported it. Free text; not resolved against a
    // user directory.
    Author string `json:"author"`
    // The comment itself. Plain text; this app neither renders nor sanitizes it.
    Body string `json:"body"`
    // When the comment was written. Comments come back oldest first.
    CreatedAt string `json:"created_at"`
    // Primary key of the comment.
    Id string `json:"id"`
    // The order the comment hangs on.
    OrderId string `json:"order_id"`
    // Who may see it: 'internal' is a note between operators, 'customer' is meant
    // to be shown in the customer's order view. Nothing here enforces that —
    // this app labels the comment and the client showing it decides. Defaults to
    // the tenant's default_comment_visibility.
    Visibility string `json:"visibility"`

    // Used by Decode() method
    data []byte
}

func (model OrderComment) New(data []byte) *OrderComment {
    model.data = data
    return &model
}

func (model *OrderComment) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}