package models

import (
    "encoding/json"
    "errors"
)

// PriceListUpdateRequest Partial update — omitted fields keep their current
// value.
type PriceListUpdateRequest struct {
    // Scope: only this sales channel. Beats the open lists, loses to contact and
    // organization.
    ChannelId string `json:"channel_id"`
    // Unique list code per tenant — the handle every import and integration
    // addresses this list by. A code already in use answers 409.
    Code string `json:"code"`
    // Scope: only this contact. The most specific scope there is — it beats
    // organization, channel and every open list, whatever their priority.
    ContactId string `json:"contact_id"`
    // ISO 4217 code (default EUR) — the currency of EVERY amount in this list,
    // since entries carry none of their own. Resolution only considers lists
    // matching the currency of the call; nothing is ever converted.
    Currency string `json:"currency"`
    // Free text for whoever maintains the list — why it exists and who it is
    // for. Never shown to a buyer.
    Description string `json:"description"`
    // The fallback list. Within its group it sorts LAST, so it wins only where
    // nothing more specific priced the item. Use prices.lists.make-default to
    // move the flag rather than setting it here — two defaults leave a tie to
    // row order.
    IsDefault bool `json:"is_default"`
    // Localised names, keyed by language tag — {"de": "Händlerpreise", "en":
    // "Dealer prices"}. Omit to show `name` everywhere.
    Labels interface{} `json:"labels"`
    // Free-form bag: whatever JSON object you write round-trips exactly, and this
    // app never reads it. Its keys are yours — ERP provenance is the usual
    // content.
    Metadata interface{} `json:"metadata"`
    // Operator-facing name, shown wherever a human picks a list.
    Name string `json:"name"`
    // Scope: only buyers of this organization. Beats channel-scoped and open
    // lists.
    OrganizationId string `json:"organization_id"`
    // Tie-break WITHIN a specificity group (higher wins, default 0). It never
    // beats scope: an organization list at 0 still wins over an open list at 100.
    Priority int `json:"priority"`
    // Gate: when true the list resolves only for an authenticated buyer (contact
    // or organization context); anonymous resolve calls get on_request. Default
    // false (open to everyone).
    RequiresAuth bool `json:"requires_auth"`
    // Default 'active' — only active lists resolve. 'inactive' retires a list
    // without deleting its prices.
    Status string `json:"status"`
    // Whether the amounts in this list are net (tax excluded) or gross (tax
    // included) — the one fact a price cannot be without. Omit (null) to
    // inherit the tenant's tax_inclusive_default setting; the resolve answer
    // names which of the two decided under tax_basis_source.
    TaxBasis string `json:"tax_basis"`
    // LEGACY mirror of tax_basis. false is the column default and is NOT read as
    // a statement of intent; true is read as gross, and only where tax_basis is
    // null. Prefer tax_basis.
    TaxIncluded bool `json:"tax_included"`
    // Start of the validity window of the WHOLE list (ISO 8601); null =
    // open-ended. Outside it the list is not a candidate at all.
    ValidFrom string `json:"valid_from"`
    // End of the validity window of the whole list; null = open-ended. Lets a
    // season expire on its own instead of being deactivated by hand.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceListUpdateRequest) New(data []byte) *PriceListUpdateRequest {
    model.data = data
    return &model
}

func (model *PriceListUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}