package models

import (
    "encoding/json"
    "errors"
)

// RegisterAReturnAgainstTheShippedQuantitiesTheReturnNumberIsDrawnFromTheReturnRange
// Model
type OrderReturnCreateRequest struct {
    // Free-form metadata.
    Metadata interface{} `json:"metadata"`
    // 
    Positions []OrderReturnPosition `json:"positions"`
    // 
    Reason string `json:"reason"`

    // Used by Decode() method
    data []byte
}

func (model OrderReturnCreateRequest) New(data []byte) *OrderReturnCreateRequest {
    model.data = data
    return &model
}

func (model *OrderReturnCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}