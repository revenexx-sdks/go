package models

import (
    "encoding/json"
    "errors"
)

// BuyerContextItemsUnpriceableItemsComeBackAsOnRequestAMissingPriceIsAFirstClassStateNever0
// Model
type PriceResolveRequest struct {
    // Point in time for validity windows (ISO 8601 timestamp, default now).
    At string `json:"at"`
    // Buyer context: channel.
    ChannelId string `json:"channel_id"`
    // Buyer context: contact — most specific scope.
    ContactId string `json:"contact_id"`
    // ISO 4217 code (default EUR) — only lists in this currency resolve.
    Currency string `json:"currency"`
    // Items to price (at most 200 per call).
    Items []PriceResolveItem `json:"items"`
    // Buyer context: market.
    MarketId string `json:"market_id"`
    // Buyer context: organization.
    OrganizationId string `json:"organization_id"`

    // Used by Decode() method
    data []byte
}

func (model PriceResolveRequest) New(data []byte) *PriceResolveRequest {
    model.data = data
    return &model
}

func (model *PriceResolveRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}