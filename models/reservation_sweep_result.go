package models

import (
    "encoding/json"
    "errors"
)

// ReservationSweepResult model.
type ReservationSweepResult struct {
    // How many active reservations were found past their hold: the ones with an
    // `expires_at` in the past, plus the undated ones older than their market's
    // TTL.
    Expired int `json:"expired"`
    // The market codes this run had to resolve a window for — every market that
    // had an undated active reservation. Empty when nothing is market-assigned,
    // which is the usual case.
    Markets []string `json:"markets"`
    // How many were actually given back — `reserved` lowered on the stock row
    // and a `release` booking written for each. It equals `expired` unless a row
    // vanished mid-run. Idempotent: a second run immediately after finds nothing
    // and answers 0.
    Released int `json:"released"`
    // The cut-off this run used — everything whose hold had run out by this
    // moment was released. It is the run's own clock, not a stored value.
    SweptAt string `json:"swept_at"`
    // The `reservation_ttl_minutes` that applied to reservations belonging to NO
    // market — the tenant baseline. A reservation assigned to a market is
    // judged against that market's own window instead, which is why this is
    // reported rather than assumed to be the only one.
    TtlMinutes float64 `json:"ttl_minutes"`

    // Used by Decode() method
    data []byte
}

func (model ReservationSweepResult) New(data []byte) *ReservationSweepResult {
    model.data = data
    return &model
}

func (model *ReservationSweepResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}