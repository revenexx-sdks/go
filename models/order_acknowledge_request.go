package models

import (
    "encoding/json"
    "errors"
)

// Model
type OrderAcknowledgeRequest struct {
    // The fulfilling system's order reference (e.g. the ERP order number).
    ExternalRef string `json:"external_ref"`

    // Used by Decode() method
    data []byte
}

func (model OrderAcknowledgeRequest) New(data []byte) *OrderAcknowledgeRequest {
    model.data = data
    return &model
}

func (model *OrderAcknowledgeRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}