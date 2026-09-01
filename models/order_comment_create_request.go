package models

import (
    "encoding/json"
    "errors"
)

// OrderCommentCreateRequest model.
type OrderCommentCreateRequest struct {
    // Who wrote it, as the caller reported it. Free text; not resolved against a
    // user directory.
    Author string `json:"author"`
    // The comment itself. Plain text; this app neither renders nor sanitizes it.
    Body string `json:"body"`
    // Who may see it: 'internal' is a note between operators, 'customer' is meant
    // to be shown in the customer's order view. Nothing here enforces that —
    // this app labels the comment and the client showing it decides. Defaults to
    // the tenant's default_comment_visibility. Defaults to the tenant's
    // default_comment_visibility setting, which is 'internal' out of the box.
    Visibility string `json:"visibility"`

    // Used by Decode() method
    data []byte
}

func (model OrderCommentCreateRequest) New(data []byte) *OrderCommentCreateRequest {
    model.data = data
    return &model
}

func (model *OrderCommentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}