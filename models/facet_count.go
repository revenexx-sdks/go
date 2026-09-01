package models

import (
    "encoding/json"
    "errors"
)

// FacetCount Facet values and their counts for one faceted field.
type FacetCount struct {
    // 
    Counts []interface{} `json:"counts"`
    // 
    FieldName string `json:"field_name"`

    // Used by Decode() method
    data []byte
}

func (model FacetCount) New(data []byte) *FacetCount {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FacetCount) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}