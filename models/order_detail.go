package models

import (
    "encoding/json"
    "errors"
)

// TheOrderAggregateEveryColumnOfTheOrderPlusItsItemsShipmentsWithPositionsReturnsAndCancellations
// Model
type OrderDetail struct {
    // 
    AcknowledgedAt string `json:"acknowledged_at"`
    // 
    BillingAddress interface{} `json:"billing_address"`
    // 
    Buyer interface{} `json:"buyer"`
    // 
    Cancellations []OrderCancellation `json:"cancellations"`
    // 
    CancelledAt string `json:"cancelled_at"`
    // 
    CartId string `json:"cart_id"`
    // 
    ChannelId string `json:"channel_id"`
    // 
    CompletedAt string `json:"completed_at"`
    // 
    ContactId string `json:"contact_id"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    CustomerOrderNumber string `json:"customer_order_number"`
    // 
    ExternalRef string `json:"external_ref"`
    // 
    FulfillmentStatus string `json:"fulfillment_status"`
    // 
    GrandTotal float64 `json:"grand_total"`
    // 
    HoldReason string `json:"hold_reason"`
    // 
    Id string `json:"id"`
    // 
    ItemCount int `json:"item_count"`
    // 
    Items []OrderItem `json:"items"`
    // 
    MarketId string `json:"market_id"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Number string `json:"number"`
    // 
    OnHold bool `json:"on_hold"`
    // 
    OrganizationId string `json:"organization_id"`
    // 
    Payment interface{} `json:"payment"`
    // 
    PaymentStatus string `json:"payment_status"`
    // 
    PlacedAt string `json:"placed_at"`
    // 
    Returns []OrderReturn `json:"returns"`
    // 
    Shipments []OrderShipment `json:"shipments"`
    // 
    Shipping interface{} `json:"shipping"`
    // 
    ShippingAddress interface{} `json:"shipping_address"`
    // 
    ShippingTotal float64 `json:"shipping_total"`
    // 
    Status string `json:"status"`
    // 
    Subtotal float64 `json:"subtotal"`
    // 
    TaxTotal float64 `json:"tax_total"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    UserData interface{} `json:"user_data"`

    // Used by Decode() method
    data []byte
}

func (model OrderDetail) New(data []byte) *OrderDetail {
    model.data = data
    return &model
}

func (model *OrderDetail) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}