package models

import (
    "encoding/json"
    "errors"
)

// MeasurementFamiliesUpdateRequest Partial update — omitted fields keep
// their current value.
type MeasurementFamiliesUpdateRequest struct {
    // The measurement family's stable identifier. A `measure` attribute names one
    // and then offers that family's units.
    Code string `json:"code"`
    // What the measurement family is called, per language tag.
    Labels interface{} `json:"labels"`
    // The unit every value of this family is converted to before it is compared
    // or sorted — the unit each `convert_factor` is relative to.
    StandardUnit string `json:"standard_unit"`
    // The units this family offers. `convert_factor` multiplies a value into
    // `standard_unit`, so a gram is 0.001 kilograms; `symbol` is what a form
    // prints next to the number.
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