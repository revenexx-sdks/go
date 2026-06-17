package models

import (
    "encoding/json"
    "errors"
)

// TheSnapshotPayloadItemsPlusFrozenBuyerAddressesPaymentShippingTheOrderNumberIsDrawnFromTheOrderRangeTotalsAreComputedFromTheItems
// Model
type OrderPlaceRequest struct {
    // Frozen billing address.
    BillingAddress interface{} `json:"billing_address"`
    // Frozen buyer snapshot (name, email, …).
    Buyer interface{} `json:"buyer"`
    // Source cart (the carts.order hand-over).
    CartId string `json:"cart_id"`
    // 
    ChannelId string `json:"channel_id"`
    // Ordering customer contact.
    ContactId string `json:"contact_id"`
    // ISO 4217 code (default EUR).
    Currency string `json:"currency"`
    // The buyer's own order/PO number.
    CustomerOrderNumber string `json:"customer_order_number"`
    // Override — computed as subtotal + shipping + tax when omitted.
    GrandTotal float64 `json:"grand_total"`
    // The order positions (at most 500).
    Items []OrderItemCreateRequest `json:"items"`
    // 
    MarketId string `json:"market_id"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // B2B organization.
    OrganizationId string `json:"organization_id"`
    // Frozen payment snapshot — a known 'payment.status' seeds payment_status
    // (otherwise 'open').
    Payment interface{} `json:"payment"`
    // Frozen shipping snapshot — 'shipping.price' seeds shipping_total.
    Shipping interface{} `json:"shipping"`
    // Frozen shipping address.
    ShippingAddress interface{} `json:"shipping_address"`
    // Shipping total (fallback when 'shipping.price' is absent).
    ShippingTotal float64 `json:"shipping_total"`
    // Free-form user data.
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