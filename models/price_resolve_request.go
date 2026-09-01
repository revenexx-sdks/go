package models

import (
    "encoding/json"
    "errors"
)

// PriceResolveRequest Buyer context + items. Unpriceable items come back as
// on_request — a missing price is a first-class state, never 0.
type PriceResolveRequest struct {
    // The instant every validity window — list and entry — is evaluated at
    // (ISO 8601). Default now. This is how a promo price is previewed before it
    // starts, and it is echoed as `basis.evaluated_at`.
    At string `json:"at"`
    // Buyer context: the sales channel. Third scope — beats the open lists,
    // loses to contact and organization.
    ChannelId string `json:"channel_id"`
    // Buyer context: the contact this quote is for. The most specific scope — a
    // list naming this contact beats every other list, whatever their priority.
    // Sending it (or organization_id) is also what makes the buyer AUTHENTICATED
    // for `requires_auth` lists and for the tenant’s anonymous_resolve_allowed
    // setting.
    ContactId string `json:"contact_id"`
    // ISO 4217 code the quote is wanted in. ONLY lists in this currency are
    // candidates and nothing is ever converted, so a wrong value here is not a
    // rounding difference — it is no price at all. Omit to take the buyer
    // market’s currency, then the tenant’s default_currency;
    // `basis.currency_source` names which applied.
    Currency string `json:"currency"`
    // Items to price, at most 200 per call — a whole cart or a whole product
    // listing in one round trip. The answer holds one entry per item, in this
    // order.
    Items []PriceResolveItem `json:"items"`
    // Buyer context: the market, as a uuid pin for older callers. Prefer the
    // `X-Revenexx-Market` header, which carries a market CODE and is what scopes
    // the visible price lists. The market decides the tax rates AND which
    // per-market settings (rounding, tie-break, anonymous access) apply — with
    // several markets and no signal at all the answer says `tax.resolved: false`,
    // `reason: market_required` rather than quoting another market’s VAT.
    MarketId string `json:"market_id"`
    // Buyer context: the organization the buyer belongs to. Second most specific
    // scope; also counts as authenticated.
    OrganizationId string `json:"organization_id"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveRequest) New(data []byte) *PriceResolveRequest {
    model.data = data
    return &model
}

func (model *PriceResolveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}