package models

import (
    "encoding/json"
    "errors"
)

// Organization A buying COMPANY — the unit a contract, a credit limit and a
// price list belong to. Its people are `contacts`, and it is mirrored into
// platform auth as a team.
type Organization struct {
    // Industry / line of business, in the merchant's own words. Free text: no
    // NACE code, no WZ number, no list to pick from — whatever somebody typed
    // on the company. Segment rules read it, and both `?branche=` and an `eq`
    // condition match it EXACTLY and case-sensitively, so 'Maschinenbau' and
    // 'maschinenbau' are two different industries. Indexed, so it stays cheap to
    // filter on.
    Branche string `json:"branche"`
    // When this company record was created in this app. Not when the customer
    // relationship began — an ERP import creates decade-old customers today.
    CreatedAt string `json:"created_at"`
    // Ceiling on open receivables in the market's currency, and one of the inputs
    // that decide whether an order is accepted at all. Null means NO limit —
    // not a limit of zero.
    CreditLimit float64 `json:"credit_limit"`
    // The number this company carries in the merchant's own ERP — the key an
    // ERP integration joins on, and what a service desk asks for on the phone.
    // Free text with NO enforced format (a letter prefix and a running number is
    // the common shape, but plain digits are just as valid), unique per tenant
    // while it is set, and one of the fields duplicate detection can be pointed
    // at. The real values come out of the merchant's ERP; nothing published here
    // can name one that exists.
    CustomerNumber string `json:"customer_number"`
    // True stops SHIPMENTS to this company while leaving login and ordering alone
    // — the "they may order, we are just not sending anything until this is
    // settled" state. Separate from `status` on purpose: blocking the login to
    // stop a delivery locks out the people who could settle it.
    DeliveryBlock bool `json:"delivery_block"`
    // Id of the platform TEAM this organization is mirrored as — what makes its
    // people a team for storefront auth (sessions, SSO, mobile SDKs). Written by
    // the mirror and ignored on every write a caller sends; null while the mirror
    // has not run yet.
    ExternalTeamId string `json:"external_team_id"`
    // Primary key of the company record. Stable for its whole life; contacts,
    // addresses, segment memberships and the metrics projection all point at it.
    Id string `json:"id"`
    // Where the company stands in the SALES PIPELINE, and a deliberately separate
    // axis from `status`: a prospect that may log in and a customer that may not
    // are both ordinary states, and one column cannot say that. One of the
    // tenant's own stages (GET /customers/lifecycle-stages) — a fresh install
    // starts with lead, prospect, customer, churned, and the merchant may add
    // their own. Nothing moves it automatically; a stage changes when a person or
    // an integration says so.
    LifecycleStage string `json:"lifecycle_stage"`
    // Legal or trading name of the COMPANY — never a person. Mirrored to the
    // platform team, so a rename here is a rename in storefront auth too.
    Name string `json:"name"`
    // When this company has to pay — one of the tenant's own terms (GET
    // /customers/payment-terms, seeded with prepayment, direct_debit,
    // net_7/14/30/60/90). Null means nothing was agreed and the order flow falls
    // back to the market's `default_payment_terms`. This is a commercial term,
    // not a payment method: HOW they pay is the payments app's business.
    PaymentTerms string `json:"payment_terms"`
    // Code of the price list this company buys on — plain text pointing into
    // the prices app. ADR-0055 forbids the cross-app foreign key, so nothing here
    // checks it: a code that names no list simply prices nothing. `standard` is
    // the list the prices app seeds on install.
    PriceList string `json:"price_list"`
    // Free-form per-organization settings, keyed by whatever the merchant's own
    // integrations agree on — this app never branches on a key in here. Segment
    // rules can address a TOP-LEVEL key as `setting:<key>`, which is the whole
    // reason the blob survives: a flag an ERP writes here selects a segment
    // without a schema change. Commercial terms are typed columns now
    // (payment_terms, credit_limit); writing them back in here leaves the
    // checkout reading the column and finding nothing.
    Settings interface{} `json:"settings"`
    // ACCESS, not pipeline: 'blocked' stops this company's people from logging in
    // and is where a rejected registration parks the company it founded. 'active'
    // is the default. For how far along a company is, read `lifecycle_stage` —
    // reading this one for that is how a won deal gets locked out.
    Status string `json:"status"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // When any column of this row last changed.
    UpdatedAt string `json:"updated_at"`
    // VAT identification number (USt-IdNr. in Germany) — the closest thing a
    // B2B buyer has to a legal identity. Validated against the EU VIES service
    // when the tenant's `organization_vat_id_required` setting is on, and stored
    // verbatim otherwise, including for buyers outside the EU.
    VatId string `json:"vat_id"`

    // Used by Decode() method
    data []byte
}

func (model Organization) New(data []byte) *Organization {
    model.data = data
    return &model
}

func (model *Organization) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}