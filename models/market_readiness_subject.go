package models

import (
    "encoding/json"
    "errors"
)

// MarketReadinessSubject The market the verdict is about, identified rather
// than returned in full — the five columns a reader needs to know which
// market answered. Read GET /markets/{id} for the rest.
type MarketReadinessSubject struct {
    // Market code, unique per tenant, and the single most load-bearing string in
    // this app: it IS the market scope slug. The Entity Scoping Engine publishes
    // it as the `market` dimension (`scope_context.market` in the JWT), and every
    // other commerce app — products, prices, orders, customers — stores THIS
    // value to say which market a row belongs to. Renaming it re-keys that scope
    // for everyone, so treat it as permanent. Accepted in place of the uuid on
    // /readiness, /clone, /backfill and /make-default — but not on the item
    // routes or /context, which take a uuid only.
    Code string `json:"code"`
    // Base currency this market quotes in — ISO 4217, and schema.json's own
    // default is 'EUR'. This is the single currency prices are STATED in; the
    // currencies collection under the market is the wider set it accepts. A base
    // currency missing from that collection is a blocking readiness failure.
    Currency string `json:"currency"`
    // The market's primary key — resolved, so a call that named the market by
    // its code gets the uuid back.
    Id string `json:"id"`
    // Display name, in the operator's own language. Cockpit copy only — nothing
    // resolves a market by it.
    Name string `json:"name"`
    // Default 'active'. Only an active market serves a storefront; 'inactive'
    // keeps the market and all its configuration but takes it out of service.
    // Readiness reports an active market that cannot trade as `serving: true,
    // ready: false` — live and broken.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model MarketReadinessSubject) New(data []byte) *MarketReadinessSubject {
    model.data = data
    return &model
}

func (model *MarketReadinessSubject) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}