package models

import (
    "encoding/json"
    "errors"
)

// CartMaintenanceResult model.
type CartMaintenanceResult struct {
    // The first sweep: active carts nobody has touched since their market's
    // window become abandoned. Nothing else in the platform ever stamps
    // abandoned_at, so without this the abandonment funnel is empty by
    // construction rather than empty because nobody abandons carts.
    Abandon CartAbandonSweep `json:"abandon"`
    // This pass wrote nothing. The counts and cart ids are the same ones the wet
    // run would produce.
    DryRun bool `json:"dry_run"`
    // The second sweep, and the only destructive thing this app does: carts past
    // their retention window are deleted, their lines with them. An ordered cart
    // is never touched at any setting — it is the source record of a sale.
    Purge CartPurgeSweep `json:"purge"`
    // The instant this pass measured every window against. One clock for both
    // sweeps, so a cart cannot be judged idle by one and fresh by the other.
    SweptAt string `json:"swept_at"`

    // Used by Decode() method
    data []byte
}

func (model CartMaintenanceResult) New(data []byte) *CartMaintenanceResult {
    model.data = data
    return &model
}

func (model *CartMaintenanceResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}