package models

import (
    "encoding/json"
    "errors"
)

// ResolvedPrice What one item costs this buyer, and which list said so.
type ResolvedPrice struct {
    // ISO 4217 currency of every amount on this item. Always the winning list’s
    // currency, which always equals the call’s top-level `currency` —
    // resolution only considers lists that match it, so a list and its answer can
    // never disagree. null on an on-request item.
    Currency string `json:"currency"`
    // Present ONLY on an item that named neither `product_id` nor `sku`, and
    // always with this exact text. The call still answers 200 and the item comes
    // back on_request, because one malformed line must not cost a whole cart its
    // prices.
    Error string `json:"error"`
    // `unit_price × quantity`, on the SAME basis as `unit_price` (so net if the
    // list is net) and rounded to `basis.price_precision`. Not a tax-adjusted
    // total — a cart computes its own from the net/gross pair.
    LineTotal float64 `json:"line_total"`
    // true = no price for this buyer context — show "price on request", never
    // 0.
    OnRequest bool `json:"on_request"`
    // Why there is no price: nothing prices it, a list marks it on-request, the
    // tenant hides prices from anonymous buyers, or the item named neither
    // product_id nor sku.
    OnRequestReason string `json:"on_request_reason"`
    // The list that priced this item — null when nothing did. On an
    // `on_request_entry` answer it is the list that said "ask us".
    PriceList interface{} `json:"price_list"`
    // Echo of the requested `product_id` — null when the item was identified by
    // SKU.
    ProductId string `json:"product_id"`
    // The quantity this answer was computed for: what you sent, or 1 where you
    // sent nothing or a non-positive value. It selects the tier and multiplies
    // into `line_total`.
    Quantity float64 `json:"quantity"`
    // Echo of the requested `sku` — null when the item was identified by
    // product id.
    Sku string `json:"sku"`
    // Whether the stored amount is net or gross. THE fact a price cannot be
    // without.
    TaxBasis string `json:"tax_basis"`
    // Who decided it: the list's own tax_basis, a legacy tax_included=true on the
    // list, or the tenant's tax_inclusive_default setting.
    TaxBasisSource string `json:"tax_basis_source"`
    // The tax class code that produced `tax_rate`: the product’s own class
    // where the products app knows one, otherwise the buyer market’s default
    // class. The codes are the tenant’s, defined in `markets.tax_classes` —
    // conventionally `standard` and `reduced`. null when tax could not be
    // resolved.
    TaxClass string `json:"tax_class"`
    // Whether unit_price already contains tax. Never null on a priced item — it
    // is `tax_basis` as a boolean, kept for existing callers.
    TaxIncluded bool `json:"tax_included"`
    // Tax rate as a PERCENTAGE (19 means 19 %, not 0.19), read from
    // `markets.tax_classes` for this market and `tax_class`. null means UNKNOWN
    // — a checkout must be able to tell that apart from a genuine 0 %.
    TaxRate float64 `json:"tax_rate"`
    // The FULL quantity ladder the winning list holds for this item, ascending by
    // `quantity_min` — what a PDP renders as a tier table. Empty on an
    // on-request item.
    Tiers []PriceTier `json:"tiers"`
    // Price for ONE unit, in `currency` and on the basis `tax_basis` names — a
    // decimal amount in major units (19.90 EUR), never minor units/cents. It is
    // the stored rung exactly as a merchant typed it, unrounded. Do not display
    // it without reading `tax_basis`; prefer `unit_price_net`/`unit_price_gross`,
    // which are unambiguous.
    UnitPrice float64 `json:"unit_price"`
    // Unit price INCLUDING tax, in `currency`, rounded to `basis.price_precision`
    // under `basis.rounding_mode`. Derived from `unit_price` and `tax_rate` in
    // whichever direction `tax_basis` requires. Present only when `tax.resolved`
    // is true.
    UnitPriceGross float64 `json:"unit_price_gross"`
    // Unit price EXCLUDING tax, in `currency`, rounded to `basis.price_precision`
    // under `basis.rounding_mode`. Present only when `tax.resolved` is true —
    // null means the rate is unknown, not that there is no tax.
    UnitPriceNet float64 `json:"unit_price_net"`

    // Used by Decode() method
    data []byte
}

func (model ResolvedPrice) New(data []byte) *ResolvedPrice {
    model.data = data
    return &model
}

func (model *ResolvedPrice) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}