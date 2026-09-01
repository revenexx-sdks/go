package models

import (
    "encoding/json"
    "errors"
)

// PriceResolveBasis The policy this answer was computed under — the tenant
// settings in force plus where the currency came from.
type PriceResolveBasis struct {
    // false ⇒ a buyer with no contact/organization is answered on_request for
    // everything.
    AnonymousResolveAllowed bool `json:"anonymous_resolve_allowed"`
    // Where `currency` came from: the request, the buyer market's own currency,
    // the tenant's default_currency setting, or the shipped fallback.
    CurrencySource string `json:"currency_source"`
    // The instant validity windows were evaluated at.
    EvaluatedAt string `json:"evaluated_at"`
    // Which list won where specificity and priority tied.
    PriceListPriorityTiebreak string `json:"price_list_priority_tiebreak"`
    // Decimals every DERIVED amount (net, gross, line totals) was rounded to.
    PricePrecision int `json:"price_precision"`
    // How those amounts landed on the last decimal.
    RoundingMode string `json:"rounding_mode"`
    // Tenant setting: the basis a price list that states none is read on.
    TaxInclusiveDefault string `json:"tax_inclusive_default"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveBasis) New(data []byte) *PriceResolveBasis {
    model.data = data
    return &model
}

func (model *PriceResolveBasis) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}