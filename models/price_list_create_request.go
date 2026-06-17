package models

import (
    "encoding/json"
    "errors"
)

// Model
type PriceListCreateRequest struct {
    // Scope: only this channel.
    ChannelId string `json:"channel_id"`
    // Unique list code per tenant.
    Code string `json:"code"`
    // Scope: only this contact — beats every other scope.
    ContactId string `json:"contact_id"`
    // ISO 4217 code (default EUR) — resolution only considers lists matching
    // the requested currency.
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // Default lists resolve last within their group.
    IsDefault bool `json:"is_default"`
    // Localised names ({de, en, …}).
    Labels interface{} `json:"labels"`
    // Scope: only this market.
    MarketId string `json:"market_id"`
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // Scope: only this organization.
    OrganizationId string `json:"organization_id"`
    // Tie-breaker within a specificity group (higher wins, default 0).
    Priority int `json:"priority"`
    // Default 'active' — only active lists resolve.
    Status string `json:"status"`
    // Gross (true) or net (false, default) prices.
    TaxIncluded bool `json:"tax_included"`
    // Validity window start.
    ValidFrom string `json:"valid_from"`
    // Validity window end.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceListCreateRequest) New(data []byte) *PriceListCreateRequest {
    model.data = data
    return &model
}

func (model *PriceListCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}