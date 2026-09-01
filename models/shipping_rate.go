package models

import (
    "encoding/json"
    "errors"
)

// ShippingRate One offerable shipping method with its computed price for this
// buyer context.
type ShippingRate struct {
    // The carrier CODE — unchanged for every caller that already reads it. The
    // method's carrier_id, else its `carrier` text, else the tenant's
    // default_carrier.
    Carrier string `json:"carrier"`
    // The carrier row's display name, or null when the code names no maintained
    // carrier.
    CarrierName string `json:"carrier_name"`
    // The class of service this rate is, from the carrier row — a code into the
    // tenant's service levels.
    CarrierServiceLevel string `json:"carrier_service_level"`
    // Which step of the chain answered: 'method' (carrier_id), 'method_code' (the
    // method's text matched a carrier), 'method_text' (it matched none),
    // 'tenant_default' / 'tenant_default_text' (the setting, matched or not).
    CarrierSource string `json:"carrier_source"`
    // Stable method code, unique per tenant (e.g. standard, express). What a
    // checkout and an order line store, so it is the value every integration
    // joins on.
    Code string `json:"code"`
    // ISO 4217 code (default EUR). Exactly three characters — the column says
    // so. Echoed into a rate, never converted: this app prices in the currency
    // the method carries.
    Currency string `json:"currency"`
    // The delivery window a checkout can print. Calendar days, cut-off evaluated
    // in UTC (send `at` to control the instant).
    Delivery ShippingDeliveryEstimate `json:"delivery"`
    // The sentence under the name in the checkout — the delivery promise in
    // words. Null when the name says enough.
    Description string `json:"description"`
    // Transit time upper bound in calendar days, as applied: the method's own,
    // else the carrier's.
    EtaDaysMax int `json:"eta_days_max"`
    // Transit time lower bound in calendar days, as applied: the method's own,
    // else the carrier's.
    EtaDaysMin int `json:"eta_days_min"`
    // Only when a free-above threshold applied. Names the compared value AND its
    // basis (net or gross), and says whether the threshold was the method's own
    // or shop-wide — the free-shipping promise is a common dispute and this is
    // the sentence that settles it.
    FreeReason string `json:"free_reason"`
    // Localized display names. A flat map keyed by locale — the Cockpit falls
    // back to `en`. Null means the row has no translations and every client shows
    // the untranslated column instead.
    Labels interface{} `json:"labels"`
    // Display name shown in the checkout.
    Name string `json:"name"`
    // Sort order in the checkout (default 0) — a rate answer is returned in
    // this order.
    Position int `json:"position"`
    // The shipping fee for this basket, in `currency`, rounded to two decimals
    // — 0 when a free-above threshold or a 'free' method applied. NULL when
    // `quote_required` is true: the price is unknown, not zero, and a checkout
    // must not add 0.00 for it.
    Price float64 `json:"price"`
    // Pricing model (default 'fixed'): 'fixed' is one price for every basket,
    // 'free' is no price at all, 'matrix' is a tiered price read off this
    // method's rate tiers. Only 'matrix' looks at matrix_basis, quote_above and
    // the tier table.
    PricingType string `json:"pricing_type"`
    // Only when quote_required — the measure and the threshold it exceeded, so
    // an operator pricing it by hand can see what triggered the referral.
    QuoteReason string `json:"quote_reason"`
    // True when the matrix measure is above the method's quote_above threshold:
    // the method is still offered, carries no price, and the storefront shows
    // 'shipping on request'. The order is placed without a computed shipping fee.
    QuoteRequired bool `json:"quote_required"`
    // The tax class this rate was taxed under, as a code in markets.tax_classes
    // — the method's own, the tenant's shipping_tax_class, or the market's
    // default, whichever answered. Null means unresolved, not untaxed.
    TaxClass string `json:"tax_class"`
    // The rate in percent from markets.tax_classes for this market and tax_class
    // — 19 means 19 %. Null means UNKNOWN, never 0: read `tax.resolved` before
    // treating a missing rate as tax-free.
    TaxRate float64 `json:"tax_rate"`
    // Which step of the chain supplied the rate: the method's own class, the
    // tenant's shipping_tax_class, the market default, or the tenant's
    // default_shipping_tax_rate. Null means unknown, NOT untaxed.
    TaxSource string `json:"tax_source"`

    // Used by Decode() method
    data []byte
}

func (model ShippingRate) New(data []byte) *ShippingRate {
    model.data = data
    return &model
}

func (model *ShippingRate) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}