package models

import (
    "encoding/json"
    "errors"
)

// Model
type MeasurementFamilies struct {
    // 
    Code string `json:"code"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Id string `json:"id"`
    // 
    Labels interface{} `json:"labels"`
    // 
    StandardUnit string `json:"standard_unit"`
    // 
    Units interface{} `json:"units"`
    // 
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model MeasurementFamilies) New(data []byte) *MeasurementFamilies {
    model.data = data
    return &model
}

func (model *MeasurementFamilies) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}