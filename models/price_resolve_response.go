package models

import (
    "encoding/json"
    "errors"
)

// PriceResolveResponse One answer per requested item, in request order, plus
// the currency, the tax context and the policy the numbers were computed
// under.
type PriceResolveResponse struct {
    // The policy this answer was computed under — the tenant settings in force
    // plus where the currency came from.
    Basis PriceResolveBasis `json:"basis"`
    // ISO 4217 currency the whole answer is quoted in, and the currency lists had
    // to match to be candidates at all. `basis.currency_source` says where it
    // came from: the request, the buyer market, the tenant setting, or the
    // shipped fallback.
    Currency string `json:"currency"`
    // One entry per requested item, in the order the items were sent. An item
    // that could not be priced is present and `on_request`, never missing.
    Prices []ResolvedPrice `json:"prices"`
    // Tax resolution status of this answer. resolved=false ⇒ tax_class/tax_rate
    // are unknown, NOT zero.
    Tax PriceTaxContext `json:"tax"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveResponse) New(data []byte) *PriceResolveResponse {
    model.data = data
    return &model
}

func (model *PriceResolveResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}