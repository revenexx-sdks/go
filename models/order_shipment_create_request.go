package models

import (
    "encoding/json"
    "errors"
)

// CreateAShipmentOmittedPositionsShipEverythingStillOpen Model
type OrderShipmentCreateRequest struct {
    // 
    Carrier string `json:"carrier"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // Delivery note number — drawn from the 'delivery' range when omitted.
    Number string `json:"number"`
    // Omitted = every position with open quantity, in full.
    Positions []OrderShipmentPosition `json:"positions"`
    // Defaults to now.
    ShippedAt string `json:"shipped_at"`
    // 
    TrackingCode string `json:"tracking_code"`
    // 
    TrackingUrl string `json:"tracking_url"`

    // Used by Decode() method
    data []byte
}

func (model OrderShipmentCreateRequest) New(data []byte) *OrderShipmentCreateRequest {
    model.data = data
    return &model
}

func (model *OrderShipmentCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}