package models

import (
    "encoding/json"
    "errors"
)

// MarketTaxClassCreateRequest The owning market comes from the route path
// ('market_id').
type MarketTaxClassCreateRequest struct {
    // Tax class code, unique per market — the rate bucket a product or a
    // shipping method is assigned to ('standard', 'reduced', 'zero'). Other apps
    // name a class by THIS and by nothing else: there is no foreign key behind it
    // and there cannot be (ADR-0055), which is why the delete route asks the
    // shipping app what still points at the code before removing it.
    Code string `json:"code"`
    // The class applied to a line that names none. At most one per market. A
    // market that stores GROSS prices and marks no default cannot break those
    // prices back down into net, which is why readiness turns that combination
    // from a warning into a blocking failure.
    IsDefault bool `json:"is_default"`
    // Localized display names for storefronts and invoices, keyed by locale: a
    // flat {locale: label} map, one level deep, string values. The key to write
    // is the `locale_policy.write` from GET /markets/{id}/context, exactly as for
    // a market's labels. Null means nothing is translated and `name` is all there
    // is.
    Labels interface{} `json:"labels"`
    // Display name of the rate bucket, in the operator's own language.
    Name string `json:"name"`
    // Sort position among this market's tax classes, ascending, default 0 — and
    // the tie-break that picks a class when none is flagged default.
    Position int `json:"position"`
    // Tax rate in PERCENT, 0–100 (default 0) — 20 means 20 %, not 0.2.
    // Whether a stored price already contains it is a separate question, answered
    // per market by `pricing.tax_basis` on the context.
    Rate float64 `json:"rate"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClassCreateRequest) New(data []byte) *MarketTaxClassCreateRequest {
    model.data = data
    return &model
}

func (model *MarketTaxClassCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}