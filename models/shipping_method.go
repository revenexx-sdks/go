package models

import (
    "encoding/json"
    "errors"
)

// Model
type ShippingMethod struct {
    // 
    Carrier string `json:"carrier"`
    // 
    Code string `json:"code"`
    // 
    Countries interface{} `json:"countries"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Currency string `json:"currency"`
    // 
    Description string `json:"description"`
    // 
    Enabled bool `json:"enabled"`
    // 
    EtaDaysMax int `json:"eta_days_max"`
    // 
    EtaDaysMin int `json:"eta_days_min"`
    // 
    FreeAbove float64 `json:"free_above"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    MatrixAttribute string `json:"matrix_attribute"`
    // 
    MatrixBasis string `json:"matrix_basis"`
    // 
    Metadata interface{} `json:"metadata"`
    // 
    Name string `json:"name"`
    // 
    Position int `json:"position"`
    // 
    Price float64 `json:"price"`
    // 
    PricingType string `json:"pricing_type"`
    // 
    TaxClass string `json:"tax_class"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model ShippingMethod) New(data []byte) *ShippingMethod {
    model.data = data
    return &model
}

func (model *ShippingMethod) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}