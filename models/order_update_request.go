package models

import (
    "encoding/json"
    "errors"
)

// OrderUpdateRequest Narrow modification — these six columns and no others.
// Anything else in the body is ignored, and a body with none of them at all
// is a 400 naming the allowed set. A whole key REPLACES the value it names;
// there is no merge into an existing snapshot. Nothing here moves the order:
// status, payment and fulfillment travel through the action routes.
type OrderUpdateRequest struct {
    // The invoice address, FROZEN at place-time. Changing the customer's address
    // afterwards does not change what this order was billed to. Replaced
    // wholesale — send the whole address, not a patch of it.
    BillingAddress interface{} `json:"billing_address"`
    // The ordering party as it was at place-time, FROZEN: a copy, not a
    // reference, so the order still reads correctly after the customer record is
    // renamed, merged or deleted. The caller decides what goes in; this app
    // stores it and reads nothing out of it. Replaced wholesale — send the
    // whole snapshot, not a patch of it.
    Buyer interface{} `json:"buyer"`
    // The BUYER's own reference — their purchase-order number. Free text, not
    // unique, never generated here: it exists so the paperwork can carry the
    // number the buyer's accounts payable will look for. One of the few fields
    // PUT /orders/{id} may still change.
    CustomerOrderNumber string `json:"customer_order_number"`
    // Free-form data belonging to the INTEGRATION side — an ERP's own
    // bookkeeping about this order. Stored and returned untouched; nothing here
    // reads it. Replaced wholesale.
    Metadata interface{} `json:"metadata"`
    // The delivery address, FROZEN at place-time — what goes on the label of
    // every shipment of this order. Null on an order that is never delivered (a
    // service, a digital item, a collection). Replaced wholesale. This is the one
    // correction that actually matters after placement: the label of every
    // shipment still to go out is printed from it.
    ShippingAddress interface{} `json:"shipping_address"`
    // Free-form data belonging to the ORDERING side — carried through from the
    // storefront or the cart and handed back untouched. One of the few fields PUT
    // /orders/{id} may still change. Replaced wholesale.
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderUpdateRequest) New(data []byte) *OrderUpdateRequest {
    model.data = data
    return &model
}

func (model *OrderUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}