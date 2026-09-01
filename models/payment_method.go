package models

import (
    "encoding/json"
    "errors"
)

// PaymentMethod model.
type PaymentMethod struct {
    // The machine name of the method, unique per tenant and lower case by
    // convention ('invoice', 'prepayment', 'card', 'paypal'). It is the string
    // the checkout asks for, the string every payment stores, and therefore the
    // one value here that cannot be changed freely: renaming it would leave the
    // ledger naming something that no longer exists, so it is refused with 409
    // for as long as any payment names it.
    Code string `json:"code"`
    // Allowed ISO 3166-1 alpha-2 country codes, compared upper-cased against the
    // buyer country. null or an empty list means unrestricted — the invoice
    // method this app seeds is restricted to DE, which is why an eligibility call
    // without a country sees it excluded.
    Countries []string `json:"countries"`
    // When this configuration was created.
    CreatedAt string `json:"created_at"`
    // One line explaining the method where it is offered — payment terms, what
    // happens after the order. Shown to the buyer, so it is the merchant's
    // wording rather than the app's.
    Description string `json:"description"`
    // A disabled method is never eligible and never reaches a checkout. This is
    // the switch an operator wants: deleting a method the ledger still names —
    // or renaming its `code` — is refused with 409.
    Enabled bool `json:"enabled"`
    // The surcharge this method costs the buyer, read as an amount or as a
    // percentage depending on `fee_type`. Never negative — a discount for
    // paying a certain way is not expressible here.
    FeeAmount float64 `json:"fee_amount"`
    // ISO 4217 code a fixed fee is expressed in. The database bounds the length
    // at three characters and nothing else, so lower case is stored as written.
    FeeCurrency string `json:"fee_currency"`
    // How `fee_amount` applies: 'none' (no surcharge), 'fixed' (that many units
    // of `fee_currency`) or 'percent' (that share of the order amount).
    FeeType string `json:"fee_type"`
    // Id of the configuration row. A payment names its method by `code`, never by
    // this — so an id is only ever used to address the configuration itself.
    Id string `json:"id"`
    // Who moves the money. 'self_managed' — invoice, prepayment — means the
    // merchant fulfils and reconciles it outside any PSP, and such a payment
    // authorizes the moment it is created. 'psp' means a configured provider
    // authorizes, captures and refunds it.
    Kind string `json:"kind"`
    // Buyer-facing names keyed by language tag — what a storefront shows
    // instead of the operator-facing `name`. Free jsonb: the database constrains
    // neither the tags nor the values, so a client reads the tag it wants and
    // falls back to `en`.
    Labels interface{} `json:"labels"`
    // Largest order amount this method may be used for — the usual credit-risk
    // cap on invoice and prepayment. null means no upper bound.
    MaxOrderValue float64 `json:"max_order_value"`
    // Free-form merchant data carried on the configuration. This app never reads
    // it — it is storage for the integrations that do (an ERP key for the
    // method, a ledger account, a display hint).
    Metadata interface{} `json:"metadata"`
    // Smallest order amount this method may be used for — the usual guard
    // against paying a €5 order by invoice. null means no lower bound.
    MinOrderValue float64 `json:"min_order_value"`
    // Operator-facing name, in the language the merchant administers in. What a
    // buyer sees comes from `labels`.
    Name string `json:"name"`
    // Sort order at checkout, ascending — the merchant's preferred payment
    // method first.
    Position int `json:"position"`
    // The PSP code this method transacts through, from GET
    // /payments/providers/catalog. Only meaningful for kind 'psp'; a PSP method
    // that names none falls back to the tenant's `default_provider` setting.
    Provider string `json:"provider"`
    // The provider's own payment-method id ('card', 'paypal', 'sepa_debit') —
    // what the driver is told to charge. Copied onto every payment created with
    // this method as `metadata.provider_method`.
    ProviderMethod string `json:"provider_method"`
    // The tenant the row belongs to — the same slug the request carried in
    // `X-Revenexx-Tenant`. Added by the platform rather than by this app, and
    // echoed so a caller that fans several tenants into one store can tell the
    // rows apart.
    TenantId string `json:"tenant_id"`
    // When it was last changed. The eligibility answer is computed live, so this
    // is the age of the configuration and not of any cached result.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model PaymentMethod) New(data []byte) *PaymentMethod {
    model.data = data
    return &model
}

func (model *PaymentMethod) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}