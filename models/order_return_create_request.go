package models

import (
    "encoding/json"
    "errors"
)

// OrderReturnCreateRequest Register a return against the shipped quantities
// — the return number is drawn from the return range. Omitted positions =
// every position that still has a returnable quantity, in full ('the customer
// sent it all back').
type OrderReturnCreateRequest struct {
    // Free-form data for the caller — the returns portal's own reference.
    // Stored and returned untouched.
    Metadata interface{} `json:"metadata"`
    // What is coming back. Omitted = every position with a returnable (shipped,
    // not yet returned) quantity, in full.
    Positions []OrderReturnPosition `json:"positions"`
    // Why the goods are coming back, free text as the customer or the desk stated
    // it. Also what /reject stores when it is given no resolution out of the
    // published set.
    Reason string `json:"reason"`
    // The default restock flag for positions that carry none of their own — and
    // the only way to say "put it all back into stock" when the positions are
    // defaulted. It does not restock anything itself: it decides what the
    // completion REPORTS for the orchestrator's inventories.restock call.
    Restock bool `json:"restock"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnCreateRequest) New(data []byte) *OrderReturnCreateRequest {
    model.data = data
    return &model
}

func (model *OrderReturnCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}