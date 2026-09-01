package models

import (
    "encoding/json"
    "errors"
)

// FormKitStepMarker A Revenexx step marker. The storefront cuts the flat
// array at each marker and renders the nodes that follow it as one wizard
// step, then removes the marker before FormKit renders anything. A definition
// with no marker is a single-step form.
type FormKitStepMarker struct {
    // Stable id for the step, so a client can address it.
    Id string `json:"id"`
    // What the step is: 'fields' for a normal step, 'thankyou' for the
    // confirmation panel shown after a successful submit.
    Kind string `json:"kind"`
    // The step heading the visitor reads.
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model FormKitStepMarker) New(data []byte) *FormKitStepMarker {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormKitStepMarker) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}