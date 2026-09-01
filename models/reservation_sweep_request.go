package models

import (
    "encoding/json"
    "errors"
)

// ReservationSweepRequest No fields — send `{}`. The cut-off is always now,
// and what counts as expired follows each reservation's own `expires_at` plus
// the `reservation_ttl_minutes` setting of the market it belongs to.
type ReservationSweepRequest struct {

    // Used by Decode() method
    data []byte
}

func (model ReservationSweepRequest) New(data []byte) *ReservationSweepRequest {
    model.data = data
    return &model
}

func (model *ReservationSweepRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}