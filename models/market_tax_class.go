package models

import (
    "encoding/json"
    "errors"
)

// MarketTaxClass One rate bucket within a market — 'standard', 'reduced',
// 'zero' — and the source of record for that rate across the platform.
// Other apps point at it by CODE, with no foreign key behind it.
type MarketTaxClass struct {
    // Tax class code, unique per market — the rate bucket a product or a
    // shipping method is assigned to ('standard', 'reduced', 'zero'). Other apps
    // name a class by THIS and by nothing else: there is no foreign key behind it
    // and there cannot be (ADR-0055), which is why the delete route asks the
    // shipping app what still points at the code before removing it.
    Code string `json:"code"`
    // When the tax class was created on this market. Set by the database; never
    // writable.
    CreatedAt string `json:"created_at"`
    // Primary key of this tax class. The class is named by `code` everywhere
    // else, including by other apps.
    Id string `json:"id"`
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
    // The market this tax class belongs to. Filled from the route path on write
    // and never read out of the body; ON DELETE CASCADE, so deleting the market
    // deletes this row.
    MarketId string `json:"market_id"`
    // Display name of the rate bucket, in the operator's own language.
    Name string `json:"name"`
    // Sort position among this market's tax classes, ascending, default 0 — and
    // the tie-break that picks a class when none is flagged default.
    Position int `json:"position"`
    // Tax rate in PERCENT, 0–100 (default 0) — 20 means 20 %, not 0.2.
    // Whether a stored price already contains it is a separate question, answered
    // per market by `pricing.tax_basis` on the context.
    Rate float64 `json:"rate"`
    // When the tax class was last written. Set by the database on every update;
    // never writable.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model MarketTaxClass) New(data []byte) *MarketTaxClass {
    model.data = data
    return &model
}

func (model *MarketTaxClass) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}