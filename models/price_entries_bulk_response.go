package models

import (
    "encoding/json"
    "errors"
)

// PriceEntriesBulkResponse Counts, not rows: an import chunk of 5000 does not
// echo 5000 entries back.
type PriceEntriesBulkResponse struct {
    // Rows inserted — rungs this list did not have.
    Created int `json:"created"`
    // The mode actually applied — the request's, or the default `upsert`.
    Mode string `json:"mode"`
    // Existing rungs rewritten in place (always 0 in append mode).
    Updated int `json:"updated"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntriesBulkResponse) New(data []byte) *PriceEntriesBulkResponse {
    model.data = data
    return &model
}

func (model *PriceEntriesBulkResponse) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}