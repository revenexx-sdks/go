package models

import (
    "encoding/json"
    "errors"
)

// OrderListItemsReplaceRequest Replace ALL positions of the list (set
// semantics).
type OrderListItemsReplaceRequest struct {
    // The new full set of positions, in the order they should carry. An empty
    // array empties the list. Every existing position is deleted and rewritten,
    // so ids are NOT preserved. The array order is the DEFAULT and not an
    // override: an entry that names no `position` takes its index, one that names
    // its own keeps it — so a replace does not by itself renumber the list from
    // zero.
    Items []OrderListItemInput `json:"items"`

    // Used by Decode() method
    data []byte
}

func (model OrderListItemsReplaceRequest) New(data []byte) *OrderListItemsReplaceRequest {
    model.data = data
    return &model
}

func (model *OrderListItemsReplaceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}