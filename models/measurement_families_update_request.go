package models

import (
    "encoding/json"
    "errors"
)

// PartialUpdateOmittedFieldsKeepTheirCurrentValue Model
type MeasurementFamiliesUpdateRequest struct {
    // 
    Code string `json:"code"`
    // 
    Labels interface{} `json:"labels"`
    // 
    StandardUnit string `json:"standard_unit"`
    // 
    Units interface{} `json:"units"`

    // Used by Decode() method
    data []byte
}

func (model MeasurementFamiliesUpdateRequest) New(data []byte) *MeasurementFamiliesUpdateRequest {
    model.data = data
    return &model
}

func (model *MeasurementFamiliesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}