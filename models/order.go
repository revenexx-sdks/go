package models

import (
    "encoding/json"
    "errors"
)

// Order An ORDER as it was placed: a snapshot. Buyer, addresses, payment and
// shipping are frozen copies, the totals were computed here, and three
// independent dimensions say where it stands — status (lifecycle),
// payment_status (fed from outside) and fulfillment_status (derived from the
// positions).
type Order struct {
    // When the fulfilling system took the order over. Written once. While it is
    // null the order can still be modified here; afterwards modification goes
    // through that system, unless the tenant sets
    // allow_modification_after_acknowledge.
    AcknowledgedAt string `json:"acknowledged_at"`
    // The invoice address, FROZEN at place-time. Changing the customer's address
    // afterwards does not change what this order was billed to.
    BillingAddress interface{} `json:"billing_address"`
    // The ordering party as it was at place-time, FROZEN: a copy, not a
    // reference, so the order still reads correctly after the customer record is
    // renamed, merged or deleted. The caller decides what goes in; this app
    // stores it and reads nothing out of it.
    Buyer interface{} `json:"buyer"`
    // When the order was cancelled, whether by a full cancel or by the last open
    // quantity being cancelled position by position. Null otherwise.
    CancelledAt string `json:"cancelled_at"`
    // The cart this order was placed from, when a storefront handed one over. A
    // reference across an app boundary (the carts app), not a foreign key —
    // nothing here checks that it resolves. Null for an order an integration or
    // an operator created.
    CartId string `json:"cart_id"`
    // The sales channel the order arrived through — webshop, app, phone desk,
    // EDI. Null when the caller named none.
    ChannelId string `json:"channel_id"`
    // When the order was closed — by a full shipment, by payment or by hand,
    // depending on the tenant's auto_complete_on. Null until then.
    CompletedAt string `json:"completed_at"`
    // The PERSON who ordered — a contact in the customers app. Resolved from
    // the acting principal whenever the caller carries one, and a body value that
    // disagrees is refused rather than silently overridden. Null for a guest
    // checkout.
    ContactId string `json:"contact_id"`
    // When the order row was written. For a placed order this is placed_at; for a
    // requested one it is when the request was submitted.
    CreatedAt string `json:"created_at"`
    // ISO 4217 code of EVERY amount on this order. Frozen at place-time from the
    // market's default_currency unless the caller named one. Nothing on this
    // order is ever converted, and the approval threshold is read in this
    // currency — which is why the threshold is a per-market setting.
    Currency string `json:"currency"`
    // The BUYER's own reference — their purchase-order number. Free text, not
    // unique, never generated here: it exists so the paperwork can carry the
    // number the buyer's accounts payable will look for. One of the few fields
    // PUT /orders/{id} may still change.
    CustomerOrderNumber string `json:"customer_order_number"`
    // The FULFILLING system's reference for this order, typically the ERP order
    // number. Written once by POST /orders/{id}/acknowledge and null until an
    // integration acknowledged it.
    ExternalRef string `json:"external_ref"`
    // Whether the order has SHIPPED, and the one dimension nobody writes: it is
    // DERIVED after every quantity change from the positions' own bookkeeping.
    // 'fulfilled' means shipped >= ordered − cancelled across all positions,
    // 'partial' means something went out. Sending it has no effect; ship, cancel
    // or return something and it moves.
    FulfillmentStatus string `json:"fulfillment_status"`
    // What the buyer owes: subtotal + shipping_total + tax_total, COMPUTED by
    // this app and NEVER taken from the caller — trusting a supplied total is
    // how inconsistent orders happened. This is the number the approval threshold
    // is compared against and the number the revenue rollup sums.
    GrandTotal float64 `json:"grand_total"`
    // Why the order is held, in the words the shipping guard quotes back. Null
    // when it is not held — releasing a hold clears it.
    HoldReason string `json:"hold_reason"`
    // Primary key of the order, and the id every other route takes. Not the order
    // number.
    Id string `json:"id"`
    // The summed ORDERED quantity over all positions, rounded to a whole number
    // — a headline figure for a list, computed once at place-time. It is
    // deliberately not reduced when something is cancelled or returned; the
    // positions carry that arithmetic.
    ItemCount int `json:"item_count"`
    // Free-form data belonging to the INTEGRATION side — an ERP's own
    // bookkeeping about this order. Stored and returned untouched; nothing here
    // reads it.
    Metadata interface{} `json:"metadata"`
    // The order number a human quotes — drawn from the tenant's order range at
    // place-time, unique per tenant and never reused. It is NOT the id: every
    // route addresses an order by uuid, and GET /orders?number=… is how a
    // number becomes one.
    Number string `json:"number"`
    // A business stop, ORTHOGONAL to status: a held order keeps its lifecycle
    // state and is refused at the guards. How far the hold reaches is the
    // tenant's call (on_hold_blocks: shipping only, shipping and cancellation, or
    // nothing at all).
    OnHold bool `json:"on_hold"`
    // The COMPANY the order is booked on — an organization in the customers
    // app, and the B2B half of who ordered. This is what
    // orders.reports.customer-rollup aggregates by and what makes an order
    // visible to a buyer's colleagues. Null on a private or guest order, which
    // the rollup counts separately because it cannot attribute it.
    OrganizationId string `json:"organization_id"`
    // The payment arrangement as it was chosen, FROZEN. This app reads exactly
    // two keys and stores the rest untouched: 'status' seeds payment_status at
    // place-time when it names one of the permitted values (anything else is
    // ignored and the order starts 'open'), and 'payment_id' is merged in by POST
    // /orders/{id}/payment-status. The method itself, its provider fields and any
    // redirect state belong to the payments app.
    Payment interface{} `json:"payment"`
    // Whether the order is PAID, and the dimension this app does not decide: it
    // is fed from outside through POST /orders/{id}/payment-status (the payments
    // app or an ERP), and only seeded at place-time from payment.status.
    // Orthogonal to the lifecycle — a completed order can still be open, and a
    // paid one can still be pending.
    PaymentStatus string `json:"payment_status"`
    // When the order was PLACED. Null while it is pending approval: an order
    // awaiting sign-off exists but was never placed, and that is exactly the
    // difference this field records.
    PlacedAt string `json:"placed_at"`
    // The shipping arrangement as it was chosen, FROZEN. Two keys are READ at
    // place-time and feed the totals: 'price' becomes shipping_total (the
    // shipping_total field is only the fallback when this is absent) and
    // 'tax_rate' is what shipping is taxed at, because shipping is a
    // Nebenleistung and is taxed too. Everything else — the carrier product,
    // the delivery window, the pickup point — is stored untouched and belongs
    // to the shipping app.
    Shipping interface{} `json:"shipping"`
    // The delivery address, FROZEN at place-time — what goes on the label of
    // every shipment of this order. Null on an order that is never delivered (a
    // service, a digital item, a collection).
    ShippingAddress interface{} `json:"shipping_address"`
    // NET shipping cost, taken from shipping.price or, when the snapshot carries
    // no price, from the request's shipping_total. In `currency`.
    ShippingTotal float64 `json:"shipping_total"`
    // Where the order stands in its LIFECYCLE, and one of three independent
    // status dimensions. 'pending' = created but not placed, an order waiting for
    // approval; 'placed' = accepted, nothing shipped; 'in_fulfillment' = part of
    // it has gone out, or all of it has and the tenant does not close on
    // shipment; 'completed' and 'cancelled' end it. Moved by the action routes
    // only — it is not writable through PUT /orders/{id}.
    Status string `json:"status"`
    // NET total of the positions (the sum of their line_total), COMPUTED here at
    // place-time. In `currency`, four decimal places. A caller cannot set it.
    Subtotal float64 `json:"subtotal"`
    // All tax on this order: the positions' tax_amount plus the tax on shipping
    // (shipping_total × shipping.tax_rate). COMPUTED here — a caller cannot
    // set it.
    TaxTotal float64 `json:"tax_total"`
    // When any column of the order last changed — every status move, every
    // re-derived fulfillment, every modification.
    UpdatedAt string `json:"updated_at"`
    // Free-form data belonging to the ORDERING side — carried through from the
    // storefront or the cart and handed back untouched. One of the few fields PUT
    // /orders/{id} may still change.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model Order) New(data []byte) *Order {
    model.data = data
    return &model
}

func (model *Order) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}