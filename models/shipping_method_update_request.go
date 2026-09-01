package models

import (
    "encoding/json"
    "errors"
)

// ShippingMethodUpdateRequest Partial update — omitted fields keep their
// current value.
type ShippingMethodUpdateRequest struct {
    // Carrier CODE, kept from before shipping_carriers existed. Looked up in the
    // carrier table when carrier_id is not set, so an existing value keeps
    // working and gains a tracking template; a code nobody maintains is still
    // reported as a plain name.
    Carrier string `json:"carrier"`
    // The carrier this method ships with. Wins over `carrier` and supplies the
    // tracking template, pickup cut-off, handling time and transit days.
    CarrierId string `json:"carrier_id"`
    // Stable method code, unique per tenant (e.g. standard, express). What a
    // checkout and an order line store, so it is the value every integration
    // joins on.
    Code string `json:"code"`
    // The countries this method may be offered into. ISO 3166-1 alpha-2 codes;
    // null or an empty array means no restriction. Compared upper-cased, so a
    // lower-case entry still matches. Declared as an array rather than the bare
    // object a jsonb column derives to — this one is always a list. ANDed with
    // the carrier's own reach.
    Countries []string `json:"countries"`
    // ISO 4217 code (default EUR). Exactly three characters — the column says
    // so. Echoed into a rate, never converted: this app prices in the currency
    // the method carries.
    Currency string `json:"currency"`
    // The sentence under the name in the checkout — the delivery promise in
    // words. Null when the name says enough.
    Description string `json:"description"`
    // Only enabled methods are ever quoted (default false); a disabled one is
    // reported in `excluded` rather than hidden.
    Enabled bool `json:"enabled"`
    // Transit time upper bound in calendar days. Falls back to the carrier's when
    // null.
    EtaDaysMax int `json:"eta_days_max"`
    // Transit time lower bound in calendar days, for the checkout. Falls back to
    // the carrier's when null.
    EtaDaysMin int `json:"eta_days_min"`
    // Free shipping at or above this order value — wins over every pricing
    // model, including a matrix. Compared net or gross as the market's
    // free_above_compares setting declares. Null falls back to the tenant's
    // shop-wide free_shipping_threshold.
    FreeAbove float64 `json:"free_above"`
    // Localized display names. A flat map keyed by locale — the Cockpit falls
    // back to `en`. Null means the row has no translations and every client shows
    // the untranslated column instead.
    Labels interface{} `json:"labels"`
    // Attribute name for matrix_basis 'attribute' — the key the rate request's
    // `attributes` map is read at. Free text: the set of attributes is the
    // catalogue's, not this app's.
    MatrixAttribute string `json:"matrix_attribute"`
    // The measure a matrix method prices its tiers over: total basket weight (in
    // the market's weight unit), total item count, order value, or 'attribute'
    // — any number the rate request carries under matrix_attribute. Null falls
    // back to the tenant's matrix_basis_default. Ignored unless pricing_type is
    // 'matrix'.
    MatrixBasis string `json:"matrix_basis"`
    // Free-form jsonb the platform never reads or validates — whatever the
    // merchant or their integration needs to keep beside the row (a customer
    // number with the carrier, an ERP key, a label-printer id). The shape varies
    // BY INTEGRATION, not by anything this app knows, so no key is declared and
    // none is reserved; the example is one plausible instance rather than a
    // schema. A flat map of scalars is the convention, and nothing enforces it.
    Metadata interface{} `json:"metadata"`
    // Display name shown in the checkout.
    Name string `json:"name"`
    // Sort order in the checkout (default 0) — a rate answer is returned in
    // this order.
    Position int `json:"position"`
    // The fixed price (default 0), in `currency` — ignored for 'free' and
    // 'matrix'.
    Price float64 `json:"price"`
    // Pricing model (default 'fixed'): 'fixed' is one price for every basket,
    // 'free' is no price at all, 'matrix' is a tiered price read off this
    // method's rate tiers. Only 'matrix' looks at matrix_basis, quote_above and
    // the tier table.
    PricingType string `json:"pricing_type"`
    // Above this MATRIX MEASURE the method carries no automatic price: it is
    // still offered, flagged `quote_required` with a reason, and the storefront
    // shows 'shipping on request'. For bulky or overweight freight priced by
    // hand. Null = every measure is priced automatically.
    QuoteAbove float64 `json:"quote_above"`
    // This method's own tax class, as a CODE into the buyer market's tax classes
    // (markets.tax_classes) — never a rate. First step of the tax chain: unset
    // falls back to the tenant's shipping_tax_class setting, then the market
    // default. Not a foreign key and it could not be (ADR-0055); GET
    // /shipping/tax-classes/{code}/usage is the integrity question markets asks
    // in its place.
    TaxClass string `json:"tax_class"`

    // Used by Decode() method
    data []byte
}

func (model ShippingMethodUpdateRequest) New(data []byte) *ShippingMethodUpdateRequest {
    model.data = data
    return &model
}

func (model *ShippingMethodUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}