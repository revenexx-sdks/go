package models

import (
    "encoding/json"
    "errors"
)

// MultiSearchRequest Envelope for a federated search. Top-level search
// parameters outside `searches` are forwarded to Typesense unchanged and act
// as defaults for every entry.
type MultiSearchRequest struct {
    // The searches to run, in order. Must not be empty.
    Searches []MultiSearchEntry `json:"searches"`

    // Used by Decode() method
    data []byte
}

func (model MultiSearchRequest) New(data []byte) *MultiSearchRequest {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *MultiSearchRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}