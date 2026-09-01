package models

import (
    "encoding/json"
    "errors"
)

// OrderPlaceRequest The snapshot payload: items plus frozen
// buyer/addresses/payment/shipping. The order number is drawn from the order
// range, totals are computed from the items.
type OrderPlaceRequest struct {
    // The invoice address, FROZEN at place-time. Changing the customer's address
    // afterwards does not change what this order was billed to.
    BillingAddress interface{} `json:"billing_address"`
    // The ordering party as it was at place-time, FROZEN: a copy, not a
    // reference, so the order still reads correctly after the customer record is
    // renamed, merged or deleted. The caller decides what goes in; this app
    // stores it and reads nothing out of it.
    Buyer interface{} `json:"buyer"`
    // The cart this order was placed from, when a storefront handed one over. A
    // reference across an app boundary (the carts app), not a foreign key —
    // nothing here checks that it resolves. Null for an order an integration or
    // an operator created. The carts.order hand-over sets it.
    CartId string `json:"cart_id"`
    // The sales channel the order arrived through — webshop, app, phone desk,
    // EDI. Null when the caller named none.
    ChannelId string `json:"channel_id"`
    // The PERSON who ordered — a contact in the customers app. Resolved from
    // the acting principal whenever the caller carries one, and a body value that
    // disagrees is refused rather than silently overridden. Null for a guest
    // checkout. Ignored when the caller carries a principal — the RESOLVED
    // contact wins, and a body value that disagrees is a 400 rather than a silent
    // override.
    ContactId string `json:"contact_id"`
    // ISO 4217 code of EVERY amount on this order. Frozen at place-time from the
    // market's default_currency unless the caller named one. Nothing on this
    // order is ever converted, and the approval threshold is read in this
    // currency — which is why the threshold is a per-market setting. Defaults
    // to the market's default_currency setting.
    Currency string `json:"currency"`
    // The BUYER's own reference — their purchase-order number. Free text, not
    // unique, never generated here: it exists so the paperwork can carry the
    // number the buyer's accounts payable will look for. One of the few fields
    // PUT /orders/{id} may still change.
    CustomerOrderNumber string `json:"customer_order_number"`
    // Optional, and CHECKED rather than used: the order always computes its own
    // total from the positions, the shipping cost and the tax. Send it as a
    // checksum on that arithmetic — if it agrees the order is placed, and if it
    // disagrees the call is refused with 400 naming both numbers, yours and the
    // computed one. The comparison is at 2 decimal places (this app stores 4,
    // ERPs work to 2, so a difference below a cent is agreement). It is never
    // taken as the order value: the approval threshold and the revenue rollup
    // read the computed number, which is why a total that disagrees is an error
    // rather than an override.
    GrandTotal float64 `json:"grand_total"`
    // The order positions — at least one, and at most the tenant's
    // max_items_per_order (500 out of the box; a longer list is a 400 naming the
    // limit).
    Items []OrderItemCreateRequest `json:"items"`
    // Free-form data belonging to the INTEGRATION side — an ERP's own
    // bookkeeping about this order. Stored and returned untouched; nothing here
    // reads it.
    Metadata interface{} `json:"metadata"`
    // The COMPANY the order is booked on — an organization in the customers
    // app, and the B2B half of who ordered. This is what
    // orders.reports.customer-rollup aggregates by and what makes an order
    // visible to a buyer's colleagues. Null on a private or guest order, which
    // the rollup counts separately because it cannot attribute it. A principal's
    // own organization wins over this when it has one.
    OrganizationId string `json:"organization_id"`
    // The payment arrangement as it was chosen, FROZEN. This app reads exactly
    // two keys and stores the rest untouched: 'status' seeds payment_status at
    // place-time when it names one of the permitted values (anything else is
    // ignored and the order starts 'open'), and 'payment_id' is merged in by POST
    // /orders/{id}/payment-status. The method itself, its provider fields and any
    // redirect state belong to the payments app.
    Payment interface{} `json:"payment"`
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
    // no price, from the request's shipping_total. In `currency`. Only read when
    // the shipping snapshot carries no 'price'.
    ShippingTotal float64 `json:"shipping_total"`
    // Free-form data belonging to the ORDERING side — carried through from the
    // storefront or the cart and handed back untouched. One of the few fields PUT
    // /orders/{id} may still change.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderPlaceRequest) New(data []byte) *OrderPlaceRequest {
    model.data = data
    return &model
}

func (model *OrderPlaceRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}