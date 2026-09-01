package models

import (
    "encoding/json"
    "errors"
)

// PriceList A price list: one currency, one tax basis, one validity window,
// one buyer scope — and the entries that price items in it. Which list wins
// for a given buyer is decided by scope first, then priority, then the
// default flag; see prices.resolve.
type PriceList struct {
    // Buyer scope: this list prices for this sales channel. Beats the open lists,
    // loses to contact and organization scope.
    ChannelId string `json:"channel_id"`
    // The unique per-tenant handle of the list — what an import, an ERP export
    // and every integration addresses it by, and what the
    // `default_price_list_code` setting names. It is never quietly reassigned: a
    // second list under a code that is taken answers 409.
    Code string `json:"code"`
    // Buyer scope: this list prices for this one contact. The most specific scope
    // there is — it beats organization, channel and every open list, whatever
    // their priority.
    ContactId string `json:"contact_id"`
    // When the list was created. Also the `newest` tie-break’s input when the
    // tenant settles genuine ties that way.
    CreatedAt string `json:"created_at"`
    // ISO 4217 currency of EVERY amount in this list — entries carry no
    // currency of their own, so this is the one that governs them. Resolution
    // only ever considers lists whose currency equals the currency of the call: a
    // list in another currency is not converted, it simply does not price the
    // item. This app never converts between currencies.
    Currency string `json:"currency"`
    // Free text for whoever maintains the list — why it exists and who it is
    // for. Never shown to a buyer.
    Description string `json:"description"`
    // The price list itself. Every sub-route addresses the list by this id, and a
    // resolve answer names the list that priced an item under `price_list.id`.
    Id string `json:"id"`
    // The fallback list. Within its group it deliberately sorts LAST, so a
    // default list wins only where nothing more specific priced the item. At most
    // one list per tenant holds the flag — `prices.lists.make-default` moves it
    // in one call.
    IsDefault bool `json:"is_default"`
    // Localised names, keyed by language tag: {"de": "Standardpreise", "en":
    // "Standard prices"}. Read the tag you need and fall back to `en`; `name` is
    // the untranslated original.
    Labels interface{} `json:"labels"`
    // Free-form bag, unvalidated and never read by this app: whatever JSON object
    // you write round-trips exactly. Its keys are the integration’s own — ERP
    // provenance is the usual content, e.g. {"source_system": "erp",
    // "erp_price_group": "A1"}.
    Metadata interface{} `json:"metadata"`
    // Operator-facing name, shown wherever a human picks a list. Not addressable
    // — integrations join on `code`.
    Name string `json:"name"`
    // Buyer scope: this list prices for buyers of this organization. Beats
    // channel-scoped and open lists, loses to a contact-scoped one.
    OrganizationId string `json:"organization_id"`
    // Tie-break WITHIN one specificity group, higher first. It never beats
    // specificity: an organization-scoped list at priority 0 still wins over an
    // open list at priority 100. Default 0.
    Priority int `json:"priority"`
    // Gate: when true the list resolves only for a buyer who has a contact or
    // organization context. An anonymous resolve never matches it, so a tenant
    // that prices only for logged-in customers flags its list and guests fall
    // through to price-on-request rather than to some other list’s number.
    RequiresAuth bool `json:"requires_auth"`
    // Whether the list takes part in resolution at all. Only `active` lists are
    // candidates; `inactive` retires a list without deleting the prices it holds.
    Status string `json:"status"`
    // Whether the amounts stored in this list are `net` (tax excluded) or `gross`
    // (tax included) — the one fact a price cannot be without. null inherits
    // the tenant’s `tax_inclusive_default` setting, and the resolve answer
    // names which of the two decided under `tax_basis_source`.
    TaxBasis string `json:"tax_basis"`
    // LEGACY mirror of `tax_basis`. `false` is the column default, so it is NOT
    // read as anybody having chosen net; only `true` is read as a statement
    // (gross), and only where `tax_basis` is null. Prefer `tax_basis`.
    TaxIncluded bool `json:"tax_included"`
    // When the row last changed. Written by the database, not by the caller.
    UpdatedAt string `json:"updated_at"`
    // Start of the validity window of the WHOLE list; null = open-ended. Outside
    // the window the list is not a candidate at all. The instant compared against
    // is the resolve call’s `at`, echoed as `basis.evaluated_at`.
    ValidFrom string `json:"valid_from"`
    // End of the validity window of the whole list; null = open-ended. Use it to
    // let a season expire on its own instead of deactivating a list by hand.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceList) New(data []byte) *PriceList {
    model.data = data
    return &model
}

func (model *PriceList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}