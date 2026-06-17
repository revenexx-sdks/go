package models

import (
    "encoding/json"
    "errors"
)

// NarrowModificationOnlyTheseColumnsAreTouchableAndOnlyUntilTheOrderIsAcknowledgedStatusMovesThroughTheActionRoutes
// Model
type OrderUpdateRequest struct {
    // 
    BillingAddress interface{} `json:"billing_address"`
    // 
    Buyer interface{} `json:"buyer"`
    // 
    CustomerOrderNumber string `json:"customer_order_number"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // 
    ShippingAddress interface{} `json:"shipping_address"`
    // Free-form user data.
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