package models

import (
    "encoding/json"
    "errors"
)

// MultiSearchResult model.
type MultiSearchResult struct {
    // One result per entry in `searches`, in the same order.
    Results []SearchResult `json:"results"`

    // Used by Decode() method
    data []byte
}

func (model MultiSearchResult) New(data []byte) *MultiSearchResult {
    model.data = data
    return &model
}

func (model *MultiSearchResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}