package models

import (
    "encoding/json"
    "errors"
)

// CartPurgeSweep The second sweep, and the only destructive thing this app
// does: carts past their retention window are deleted, their lines with them.
// An ordered cart is never touched at any setting — it is the source record
// of a sale.
type CartPurgeSweep struct {
    // More carts were available to examine than one pass examines; the rest go
    // next tick, oldest first.
    Capped bool `json:"capped"`
    // The carts this sweep touched, so a merchant can look at them before or
    // after.
    CartIds []string `json:"cart_ids"`
    // The tenant baseline's window for CUSTOMER carts, in days. 0 is 'never
    // delete' — the default, and also where an unparsable value lands, so no
    // settings outage can start a purge.
    CartTtlDays float64 `json:"cart_ttl_days"`
    // The baseline cutoff, for carts belonging to no market. Null when the
    // baseline keeps everything.
    Cutoff string `json:"cutoff"`
    // Carts actually deleted. 0 on a dry run — see `found`.
    Deleted int `json:"deleted"`
    // Retention was in force for at least one cart this pass looked at — the
    // baseline, or some market that sets a window while the baseline leaves it
    // off. False means nothing could have been deleted.
    Enabled bool `json:"enabled"`
    // Carts past their retention window. On a dry run this is what the wet run
    // would remove.
    Found int `json:"found"`
    // The same for GUEST carts — a cart with a session key and no contact
    // behind it. Kept separate because the two are worth different amounts: a
    // named B2B cart may be a quote somebody is still thinking about.
    GuestCartTtlDays float64 `json:"guest_cart_ttl_days"`
    // Lines actually deleted with them. 0 on a dry run.
    ItemsDeleted int `json:"items_deleted"`
    // The market codes this pass came across. Each cart was held against ITS
    // market's window, not the baseline's.
    Markets []string `json:"markets"`
    // Lines the wet run would remove. Always present, on a wet run too, so a
    // client never has to tell "nothing to delete" apart from "this build did not
    // report it".
    WouldDeleteItems int `json:"would_delete_items"`

    // Used by Decode() method
    data []byte
}

func (model CartPurgeSweep) New(data []byte) *CartPurgeSweep {
    model.data = data
    return &model
}

func (model *CartPurgeSweep) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}